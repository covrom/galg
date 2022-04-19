package minheap_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/covrom/galg/minheap"
)

func fillInts() []int {
	rand.Seed(time.Now().Unix())

	ints := make([]int, 1000000)

	for i := range ints {
		ints[i] = rand.Int()
	}
	return ints
}

type intVal int

func (x intVal) LessThan(y intVal) bool {
	return int(x) < int(y)
}

func TestMinHeap(t *testing.T) {
	rand.Seed(0)

	ints := fillInts()

	h1 := make([]int, len(ints))
	h2 := make([]intVal, 0, 10000)

	copy(h1, ints)

	// first 10k maximum elements (ascending order)
	sort.Sort(sort.Reverse(sort.IntSlice(h1)))
	h1 = h1[:10000]
	sort.Ints(h1)

	// min-heap algo (first 10k maximum elements)
	for _, v := range ints {
		minheap.PushMinHeap(10000, &h2, intVal(v))
	}

	for i, v := range h1 {
		v2 := minheap.PopMinHeap(&h2)
		if int(v2) != v {
			t.Errorf("idx=%v, v=%v, v2=%v", i, v, v2)
		}
	}
}

func TestMinHeapOrdered(t *testing.T) {
	rand.Seed(0)

	ints := fillInts()

	h1 := make([]int, len(ints))
	h2 := make([]int, 0, 10000)

	copy(h1, ints)

	// first 10k maximum elements (ascending order)
	sort.Sort(sort.Reverse(sort.IntSlice(h1)))
	h1 = h1[:10000]
	sort.Ints(h1)

	// min-heap algo (first 10k maximum elements)
	for _, v := range ints {
		minheap.PushMinHeapOrdered(10000, &h2, v)
	}

	for i, v := range h1 {
		v2 := minheap.PopMinHeapOrdered(&h2)
		if int(v2) != v {
			t.Errorf("idx=%v, v=%v, v2=%v", i, v, v2)
		}
	}
}
