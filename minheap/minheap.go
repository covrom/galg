package minheap

type MinHeapComparable[T any] interface {
	LessThan(T) bool
}

func InitMinHeap[T MinHeapComparable[T]](h *[]T) {
	n := len(*h)
	for i := n/2 - 1; i >= 0; i-- {
		downHeap(h, i, n)
	}
}

// heapLimit - size limit, if zero then unlimited
func PushMinHeap[T MinHeapComparable[T]](heapLimit int, h *[]T, x T) {
	n := len(*h)
	if heapLimit <= 0 || n < heapLimit {
		*h = append(*h, x)
		upHeap(h, n)
	} else if !x.LessThan((*h)[0]) {
		// collect only max N
		(*h)[0] = x
		downHeap(h, 0, n)
	}
}

func PopMinHeap[T MinHeapComparable[T]](h *[]T) T {
	n := len(*h) - 1
	(*h)[0], (*h)[n] = (*h)[n], (*h)[0]
	downHeap(h, 0, n)
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func RemoveMinHeap[T MinHeapComparable[T]](h *[]T, i int) T {
	n := len(*h) - 1
	if n != i {
		(*h)[i], (*h)[n] = (*h)[n], (*h)[i]
		if !downHeap(h, i, n) {
			upHeap(h, i)
		}
	}
	x := (*h)[n]
	*h = (*h)[0 : n-1]
	return x
}

func FixMinHeap[T MinHeapComparable[T]](h *[]T, i int) {
	if !downHeap(h, i, len(*h)) {
		upHeap(h, i)
	}
}

func upHeap[T MinHeapComparable[T]](h *[]T, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !(*h)[j].LessThan((*h)[i]) {
			break
		}
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		j = i
	}
}

func downHeap[T MinHeapComparable[T]](h *[]T, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && (*h)[j2].LessThan((*h)[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !((*h)[j]).LessThan((*h)[i]) {
			break
		}
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		i = j
	}
	return i > i0
}
