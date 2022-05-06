package avltree

// Iterator holding the iterator's state
type Iterator[K, V any] struct {
	tree     *Tree[K, V]
	node     *Node[K, V]
	position position
}

type position byte

const (
	begin, between, end position = 0, 1, 2
)

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (tree *Tree[K, V]) Iterator() *Iterator[K, V] {
	return &Iterator[K, V]{tree: tree, node: nil, position: begin}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) Next() bool {
	switch iterator.position {
	case begin:
		iterator.position = between
		iterator.node = iterator.tree.Left()
	case between:
		iterator.node = iterator.node.Next()
	}

	if iterator.node == nil {
		iterator.position = end
		return false
	}
	return true
}

// Prev moves the iterator to the next element and returns true if there was a previous element in the container.
// If Prev() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Prev() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) Prev() bool {
	switch iterator.position {
	case end:
		iterator.position = between
		iterator.node = iterator.tree.Right()
	case between:
		iterator.node = iterator.node.Prev()
	}

	if iterator.node == nil {
		iterator.position = begin
		return false
	}
	return true
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator[K, V]) Value() (v V) {
	if iterator.node == nil {
		return v
	}
	return iterator.node.Value
}

// Key returns the current element's key.
// Does not modify the state of the iterator.
func (iterator *Iterator[K, V]) Key() (k K, ok bool) {
	if iterator.node == nil {
		return k, false
	}
	return iterator.node.Key, true
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator[K, V]) Begin() {
	iterator.node = nil
	iterator.position = begin
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator[K, V]) End() {
	iterator.node = nil
	iterator.position = end
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator
func (iterator *Iterator[K, V]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) Last() bool {
	iterator.End()
	return iterator.Prev()
}
