package byteutils

import (
	"testing"
	"unsafe"
)

// UnsafeAppend appends src to dst using unsafe operations
// WARNING: dst must have sufficient capacity to hold src
// Returns the updated slice
func UnsafeAppend(dst, src []byte) []byte {
	if len(src) == 0 {
		return dst
	}

	dstLen := len(dst)
	srcLen := len(src)

	// Get pointer to the start of destination space
	dstPtr := unsafe.Pointer(unsafe.SliceData(dst))

	// Advance pointer to where we'll copy
	targetPtr := unsafe.Add(dstPtr, dstLen)

	// Get pointer to source data
	srcPtr := unsafe.Pointer(unsafe.SliceData(src))

	// Perform the copy
	copy(unsafe.Slice((*byte)(targetPtr), srcLen), unsafe.Slice((*byte)(srcPtr), srcLen))

	// Return slice with updated length
	return dst[:dstLen+srcLen]
}

func TestRawExample(t *testing.T) {
	// Create destination with correct capacity
	dst := make([]byte, 5, 10) // "Hello" with room for "World"
	copy(dst, []byte{5, 10, 15, 20, 25})

	src := []byte{45, 50, 55, 70, 75}

	// Safety check
	if cap(dst)-len(dst) < len(src) {
		panic("insufficient capacity")
	}

	result := UnsafeAppend(dst, src)
	// result now contains "HelloWorld"

	t.Log(result)
}

// Test function
func TestExample(t *testing.T) {
	dst := make([]byte, 5, 10)
	copy(dst, []byte("Hello"))

	src := []byte("World")

	result := UnsafeAppend(dst, src)

	expected := "HelloWorld"
	if string(result) != expected {
		panic("unexpected result")
	}
}

func unsafeCopy(dst, src []byte) {
	if len(src) == 0 {
		return
	}

	// Get pointer to dst data
	dstPtr := unsafe.Pointer(unsafe.SliceData(dst))

	// Get pointer to src data
	srcPtr := unsafe.Pointer(unsafe.SliceData(src))

	// Copy the data
	copy(unsafe.Slice((*byte)(dstPtr), len(src)), unsafe.Slice((*byte)(srcPtr), len(src)))
}

func TestUnsafeExample(t *testing.T) {
	// Create destination with correct capacity
	dst := make([]byte, 10) // "Hello" with room for "World"
	copy(dst, []byte{5, 10, 15, 20, 25})

	src := []byte{45, 50, 55, 70, 75, 80}

	unsafeCopy(dst, src)

	t.Log(dst)
}

type bytesWriter struct {
	data   []byte
	cursor int

	freeCap int // cap(data) - len(data)
}

func newBytesWriter(data []byte) *bytesWriter {
	capacity := cap(data)
	length := len(data)
	bbw := &bytesWriter{
		data:    data,
		freeCap: capacity - length,
	}
	if bbw.freeCap == 0 {
		bbw.freeCap = capacity
	}
	if length == 0 {
		bbw.data = bbw.data[:capacity]
	}

	return bbw
}

func (bbw *bytesWriter) put(b byte) {
	if 1 >= bbw.freeCap {
		newDataCap := cap(bbw.data) << 1
		newData := make([]byte, newDataCap)
		copy(newData, bbw.data)
		bbw.data = newData
		bbw.freeCap = newDataCap - bbw.cursor
	}

	bbw.data[bbw.cursor] = b
	bbw.cursor++
	bbw.freeCap--
}

func (bbw *bytesWriter) write(bs []byte) {
	bsLen := len(bs)
	if bsLen > bbw.freeCap {
		newCap := cap(bbw.data) << 1
		currentMaxSize := len(bbw.data) + bsLen - bbw.freeCap
		for currentMaxSize > newCap {
			newCap <<= 1
		}

		newData := make([]byte, newCap)
		copy(newData, bbw.data)
		bbw.data = newData
		bbw.freeCap = newCap - bbw.cursor
	}

	copy(bbw.data[bbw.cursor:], bs)
	bbw.cursor += bsLen
	bbw.freeCap -= bsLen
}

func (bbw *bytesWriter) bytes() []byte {
	return bbw.data[:bbw.cursor]
}

func (bbw *bytesWriter) unsafeCopy(dst, src []byte) {
	if len(src) == 0 {
		return
	}

	// Get pointer to dst data
	dstPtr := unsafe.Pointer(unsafe.SliceData(dst))

	// Get pointer to src data
	srcPtr := unsafe.Pointer(unsafe.SliceData(src))

	// Copy the data
	copy(unsafe.Slice((*byte)(dstPtr), len(src)), unsafe.Slice((*byte)(srcPtr), len(src)))
}

func (bbw *bytesWriter) unsafeCopyFrom(bs []byte) {
	dstPtr := unsafe.Add(unsafe.Pointer(unsafe.SliceData(bbw.data)), bbw.cursor)
	srcPtr := unsafe.Pointer(unsafe.SliceData(bs))
	copy(unsafe.Slice((*byte)(dstPtr), len(bs)), unsafe.Slice((*byte)(srcPtr), len(bs)))
}

func (bbw *bytesWriter) grow(n int) {
	newCap := (cap(bbw.data) << 1) + n
	newData := make([]byte, newCap)
	copy(newData, bbw.data)
	bbw.data = newData
	bbw.freeCap = newCap - bbw.cursor
}

func (bbw *bytesWriter) growAppend(n int) {
	bbw.data = append(bbw.data, make([]byte, n)...)
	bbw.freeCap = cap(bbw.data) + n - bbw.cursor
}

func (bbw *bytesWriter) yield() int {
	return bbw.cursor
}
