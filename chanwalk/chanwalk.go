package chanwalk

// walk over chan with buffered processing
func WalkChan[T any](chin chan T, bufsize int, f func([]T)) {
	buf := make([]T, 0, bufsize)
	for {
		v, ok := <-chin
		if ok {
			buf = append(buf, v)
		}
		if !ok || len(buf) >= bufsize {
			if len(buf) == 0 {
				return
			}
			f(buf)
			buf = buf[:0]
		}
	}
}
