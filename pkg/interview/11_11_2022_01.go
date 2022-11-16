package main

import (
	"fmt"
	"sync"
)

// singleton variables, note the solution also allows for partial hierarchical tags.
// which means multiple entries could have multiple sized tags, and insert the first n tag values instead of all of them.
var ids map[int][]string = make(map[int][]string)
var tkr *tracker = newTracker()
var tkrMutex *sync.RWMutex = &sync.RWMutex{}

type tracker struct {
	count int
	tree  map[string]*tracker
}

func StartTracking(id int, tags []string) error {
	tkrMutex.Lock()
	defer tkrMutex.Unlock()
	ids[id] = tags
	return startTracking(tkr, tags)
}

func EndTracking(id int) error {
	tkrMutex.Lock()
	defer tkrMutex.Unlock()
	tags, ok := ids[id]
	if !ok {
		return fmt.Errorf("id %d not found", id)
	}
	delete(ids, id)
	return endTracking(tkr, tags)
}

func GetCounts(tags ...string) int {
	tkrMutex.RLock()
	defer tkrMutex.RUnlock()
	return getCounts(tkr, tags...)
}

func startTracking(t *tracker, tags []string) error {
	t.count++
	if len(tags) == 0 {
		return nil
	}
	if _, ok := t.tree[tags[0]]; !ok {
		t.tree[tags[0]] = newTracker()
	}
	next := t.tree[tags[0]]
	return startTracking(next, tags[1:])
}

func endTracking(t *tracker, tags []string) error {
	t.count--
	if len(tags) == 0 {
		return nil
	}
	next := t.tree[tags[0]]
	return endTracking(next, tags[1:])
}

func getCounts(t *tracker, tags ...string) int {
	if t == nil {
		return 0
	}
	if len(tags) == 0 {
		return t.count
	}
	next, _ := t.tree[tags[0]]
	return getCounts(next, tags[1:]...)
}

func newTracker() *tracker {
	return &tracker{
		count: 0,
		tree:  make(map[string]*tracker),
	}
}

func main() {
	StartTracking(1112, []string{"UPI", "Karnataka", "Bangalore"})
	StartTracking(2451, []string{"UPI", "Karnataka", "Mysore"})
	StartTracking(3421, []string{"UPI", "Rajasthan", "Jaipur"})
	StartTracking(1221, []string{"Wallet", "Karnataka", "Bangalore"})

	fmt.Println(GetCounts("UPI"))
	fmt.Println(GetCounts("UPI", "Karnataka"))
	fmt.Println(GetCounts("UPI", "Karnataka", "Bangalore"))
	fmt.Println(GetCounts("Bangalore"))

	StartTracking(4221, []string{"Wallet", "Karnataka", "Bangalore"})
	EndTracking(1112)
	EndTracking(2451)

	fmt.Println(GetCounts("UPI"))
	fmt.Println(GetCounts("Wallet"))
	fmt.Println(GetCounts("UPI", "Karnataka", "Bangalore"))
}
