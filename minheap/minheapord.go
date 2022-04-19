package minheap

import "golang.org/x/exp/constraints"

func InitMinHeapOrdered[T constraints.Ordered](h *[]T) {
	n := len(*h)
	for i := n/2 - 1; i >= 0; i-- {
		downHeapOrdered(h, i, n)
	}
}

// heapLimit - size limit, if zero then unlimited
func PushMinHeapOrdered[T constraints.Ordered](heapLimit int, h *[]T, x T) {
	n := len(*h)
	if heapLimit <= 0 || n < heapLimit {
		*h = append(*h, x)
		upHeapOrdered(h, n)
	} else if !(x < (*h)[0]) {
		// collect only max N
		(*h)[0] = x
		downHeapOrdered(h, 0, n)
	}
}

func PopMinHeapOrdered[T constraints.Ordered](h *[]T) T {
	n := len(*h) - 1
	(*h)[0], (*h)[n] = (*h)[n], (*h)[0]
	downHeapOrdered(h, 0, n)
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func RemoveMinHeapOrdered[T constraints.Ordered](h *[]T, i int) T {
	n := len(*h) - 1
	if n != i {
		(*h)[i], (*h)[n] = (*h)[n], (*h)[i]
		if !downHeapOrdered(h, i, n) {
			upHeapOrdered(h, i)
		}
	}
	x := (*h)[n]
	*h = (*h)[0 : n-1]
	return x
}

func FixMinHeapOrdered[T constraints.Ordered](h *[]T, i int) {
	if !downHeapOrdered(h, i, len(*h)) {
		upHeapOrdered(h, i)
	}
}

func upHeapOrdered[T constraints.Ordered](h *[]T, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !((*h)[j] < (*h)[i]) {
			break
		}
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		j = i
	}
}

func downHeapOrdered[T constraints.Ordered](h *[]T, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && ((*h)[j2] < (*h)[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !((*h)[j] < (*h)[i]) {
			break
		}
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		i = j
	}
	return i > i0
}
