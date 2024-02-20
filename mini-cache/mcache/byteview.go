package mcache

// ByteView a ByteView holds  an Immutable view of bytes.
type ByteView struct {
	b []byte
}

// Len return the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(originalBytes []byte) []byte {
	c := make([]byte, len(originalBytes))
	copy(c, originalBytes)
	return c
}
