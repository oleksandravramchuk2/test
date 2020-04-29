package metric

import (
	"fmt"
	"log"
	"sync"
)

// CountersMap concurrent-safe struct for storing metrics about count of requests and their calls
type CountersMap struct {
	mx sync.Mutex
	m  map[string]Status
}

// Status struct for each request record counting calls and active or canceled state
type Status struct {
	requests int
	active   bool
}

// String convenient representative view for status
func (s Status) String() string {
	if s.active {
		return fmt.Sprintf("%d, active", s.requests)
	}
	return fmt.Sprintf("%d, canceled", s.requests)
}

// NewCounters returns new CountersMap instance
func NewCounters() *CountersMap {
	return &CountersMap{
		m: make(map[string]Status),
	}
}

func (c *CountersMap) Store(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	// for first generation there is a chance of doublicating key,
	// so, commented peace of code that prints "doubled" warning, cause for first 50 records may be ~47-50 records in metrics.
	// To check this, u can disable changer and count requests /admin/requests
	//if _, ok := c.m[key]; ok {
	//		fmt.Println("doubled", key)
	//	}

	// Due to the fact that duplicates are possible, due to the limited total number of query options (26 * 26 = 676)
	// and the constant generation of new random records.
	// I decided to keep a request counter, simply turn the canceled status into active.
	s, ok := c.m[key]
	if ok {
		s.active = true
		c.m[key] = s
		return
	}
	c.m[key] = Status{0, true}
}

// Inc increases requests count for record with input key
func (c *CountersMap) Inc(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	s, ok := c.m[key]
	if !ok {
		log.Println("Inc warning: There is not record with key:", key)
		return
	}
	s.requests++
	c.m[key] = s
}

// Check checks if request is active at this moment
func (c *CountersMap) CheckActive(key string) bool {
	s, ok := c.m[key]
	if !ok {
		return false
	}
	return s.active
}


// Cancel set active flag to false for record with input key
func (c *CountersMap) Cancel(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	s, ok := c.m[key]
	if !ok {
		log.Println("Cancel warning: There is not record with key:", key)
		return
	}
	s.active = false
	c.m[key] = s
}

// Range returns string for response about all records metrics
func (c *CountersMap) Range() (resp string) {
	resp += fmt.Sprintf("Count: %v\n", len(c.m))
	c.mx.Lock()
	defer c.mx.Unlock()
	for k, v := range c.m {
		resp += fmt.Sprintf("%s - %v\n", k, v)
	}
	return
}
