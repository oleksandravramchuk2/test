package pool

import (
	"math/rand"
	"sync"
	"time"

	"github.com/test/metric"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyz"
	poolCap     = 50
	reqLen      = 2
	frequency   = 200
)

// ActiveRequestsPool struct for storing 50 records.
type ActiveRequestsPool struct {
	sync.RWMutex
	items []string
}

func randOfN(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func RandStringBytes() []byte {
	b := make([]byte, reqLen)
	for i := range b {
		b[i] = letterBytes[randOfN(len(letterBytes))]
	}
	return b
}

func NewPool(m *metric.CountersMap) *ActiveRequestsPool {
	pool := make([]string, 50)
	for i := range pool {
		req := string(RandStringBytes())
		pool[i] = string(req)
		m.Store(req)
	}
	return &ActiveRequestsPool{items: pool}
}

// changePool Given the task, every 200 ms we delete and add one record, without using ready-made solutions.
// My solution is to use a slice, but without deleting and adding each record, and changing a random existing record to another.
// Thus, we have a slice of 50 elements in which one element changes every 200 ms.
// Additionally, there is a registration of a new element and deactivation of the old element in the metrics.
func (p *ActiveRequestsPool) changePool(m *metric.CountersMap) {
	var canceled, newValue string
	var active bool

	i := randOfN(poolCap)
	// we start random find new request combination, if we will find the same existing with active status(it presents in pool at this moment),
	// try to generate another one
	for {
		newValue = string(RandStringBytes())
		active = m.CheckActive(newValue)
		if active {
			break
		}
	}

	m.Store(newValue)
	p.Lock()
	canceled, p.items[i] = p.items[i], newValue
	p.Unlock()
	m.Cancel(canceled)
}

// Changer iterator for changing one element in loop
func (p *ActiveRequestsPool) Changer(m *metric.CountersMap) {
	for {
		p.changePool(m)
		time.Sleep(frequency * time.Millisecond)
	}
}

// GetValue returns a random value for the request/
func (p *ActiveRequestsPool) GetValue(m *metric.CountersMap) (s string) {
	i := randOfN(poolCap)
	s = p.items[i]
	m.Inc(s)
	return s
}
