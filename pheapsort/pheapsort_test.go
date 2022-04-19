package pheapsort_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/covrom/galg/pheapsort"
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

func TestParallelHeapSort(t *testing.T) {
	rand.Seed(0)

	ints := fillInts()

	h1 := make([]int, 10000)
	h2 := make([]intVal, 10000)
	h3 := make([]intVal, 10000)

	copy(h1, ints)

	for i := range h2 {
		h2[i] = intVal(ints[i])
	}
	for i := range h3 {
		h3[i] = intVal(ints[i])
	}

	sort.Ints(h1)
	h2 = pheapsort.ParallelHeapSort(3, h2)
	h3 = pheapsort.NormalHeapSort(h3)

	for i, v := range h1 {
		if intVal(v) != h2[i] || intVal(v) != h3[i] {
			t.Errorf("idx=%v, v1=%v, v2=%v, v3=%v", i, v, h2[i], h3[i])
		}
	}
}
