package pheapsort

import "github.com/covrom/galg/minheap"

func ParallelHeapSort[T minheap.MinHeapComparable[T]](heap_workers int, ints []T) []T {
	chres := make(chan T, heap_workers*3)
	workSize := len(ints) / heap_workers
	for i := 0; i < len(ints); i += workSize {
		right := i + workSize
		if right >= len(ints) {
			right = len(ints)
		}
		go heapSort(ints[i:right], chres)
	}
	h := make([]T, 0, len(ints))
	for v := range chres {
		minheap.PushMinHeap(0, &h, v)
		if len(h) == len(ints) {
			break
		}
	}
	res := make([]T, len(h))
	for i := range res {
		res[i] = minheap.PopMinHeap(&h)
	}
	return res
}

func heapSort[T minheap.MinHeapComparable[T]](ints []T, chout chan T) {
	h := make([]T, 0, len(ints))
	for _, v := range ints {
		minheap.PushMinHeap(0, &h, v)
	}
	for len(h) > 0 {
		chout <- minheap.PopMinHeap(&h)
	}
}

func NormalHeapSort[T minheap.MinHeapComparable[T]](ints []T) []T {
	h := make([]T, 0, len(ints))
	for _, v := range ints {
		minheap.PushMinHeap(0, &h, v)
	}
	res := make([]T, 0, len(h))
	for len(h) > 0 {
		res = append(res, minheap.PopMinHeap(&h))
	}
	return res
}
