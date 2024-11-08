package main

import (
	"reflect"
	"testing"
	"unsafe"
)

// FastSet provides the fastest way to set a value using reflection
// Warning: This uses unsafe operations - use with caution
func FastSet(field reflect.Value, value interface{}) {
	// Get the pointer to the field's underlying data
	ptr := unsafe.Pointer(field.UnsafeAddr())

	// Create a new reflect.Value from the pointer
	newVal := reflect.NewAt(field.Type(), ptr).Elem()

	// Set the value
	newVal.Set(reflect.ValueOf(value))
}

// SafeSet provides a safe but slower alternative
func SafeSet(field reflect.Value, value interface{}) {
	if field.CanSet() {
		field.Set(reflect.ValueOf(value))
	}
}

// Example usage and benchmark:

type Example struct {
	Field string
}

func BenchmarkReflectionSet(b *testing.B) {
	example := &Example{}
	v := reflect.ValueOf(example).Elem()
	field := v.FieldByName("Field")

	// Fast but unsafe
	b.Run("Fast but unsafe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FastSet(field, "new value")
		}
	})

	// Safe but slower
	b.Run("Safe but slower", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SafeSet(field, "new value")
		}
	})
}

// Direct field access (for comparison)
func DirectSet(e *Example) {
	e.Field = "new value"
}
