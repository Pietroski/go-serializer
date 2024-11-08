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
