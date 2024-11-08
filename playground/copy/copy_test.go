package main

import (
	"bytes"
	"testing"
)

// Method 1: Direct slice manipulation
func WriteBytesDirect(b []byte, data []byte) {
	copy(b, data)
}

// Method 2: Using bytes.Buffer
func WriteBytesBuffer(data []byte) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, len(data)))
	buf.Write(data)
	return buf.Bytes()
}

// Method 3: Pre-allocated slice with append
func WriteBytesAppend(b []byte, data []byte) []byte {
	return append(b[:0], data...)
}

// Benchmark to compare methods
func BenchmarkByteWriting(b *testing.B) {
	data := []byte("Hello, World! This is some test data.")
	dst := make([]byte, len(data))

	b.Run("Direct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			WriteBytesDirect(dst, data)
		}
	})

	b.Run("Buffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			WriteBytesBuffer(data)
		}
	})

	b.Run("Append", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			WriteBytesAppend(dst, data)
		}
	})
}

func BenchmarkSliceCopy(b *testing.B) {
	data := []byte("Hello, World! This is some test data for benchmarking slice operations.")

	b.Run("PreAllocatedLength", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Pre-allocate with length
			dst := make([]byte, len(data))
			copy(dst, data)
		}
	})

	b.Run("PreAllocatedLength", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Pre-allocate with length
			dst := make([]byte, len(data))
			dst = append(dst, data...)
		}
	})

	b.Run("PreAllocatedCapacity", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Pre-allocate with capacity only
			dst := make([]byte, 0, len(data))
			dst = append(dst, data...)
		}
	})

	b.Run("PreAllocatedCapacityWithCopy", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Pre-allocate with capacity only, but use copy
			dst := make([]byte, 0, len(data))
			dst = dst[:len(data)]
			copy(dst, data)
		}
	})

	// To demonstrate growth penalty
	b.Run("GrowingSlice", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Start with small capacity
			dst := make([]byte, 0, 1)
			dst = append(dst, data...)
		}
	})
}

// Method 1: Grow with exact final size
func growWithPrealloc(size int) []int {
	return make([]int, 0, size)
}

// Method 2: Grow with append one by one
func growWithAppendOne(size int) []int {
	s := make([]int, 0)
	for i := 0; i < size; i++ {
		s = append(s, i)
	}
	return s
}

// Method 3: Grow with batch append
func growWithBatchAppend(size int) []int {
	s := make([]int, 0)
	batch := make([]int, 100)
	for i := 0; i < size; i += 100 {
		if i+100 > size {
			s = append(s, batch[:size-i]...)
		} else {
			s = append(s, batch...)
		}
	}
	return s
}

// Method 4: Grow with slice resize
func growWithResize(size int) []int {
	s := make([]int, 0)
	newCap := size
	if size > 1024 {
		newCap = int(float64(size) * 1.25) // Go's growth factor
	}
	s = append(s[:cap(s)], make([]int, newCap-cap(s))...)
	return s[:size]
}

// Method 5: Copy to larger slice
func growWithCopy(size int) []int {
	s := make([]int, 1<<6)

	sLen := len(s)
	if sLen < size {
		newCap := sLen << 1
		for newCap < size {
			newCap <<= 1
		}
		newS := make([]int, newCap)
		copy(newS, s)
	}

	return s
}

func BenchmarkSliceGrowth(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run("Prealloc-"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = growWithPrealloc(size)
			}
		})

		b.Run("AppendOne-"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = growWithAppendOne(size)
			}
		})

		b.Run("BatchAppend-"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = growWithBatchAppend(size)
			}
		})

		b.Run("Resize-"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = growWithResize(size)
			}
		})

		b.Run("Copy-"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = growWithCopy(size)
			}
		})
	}
}
