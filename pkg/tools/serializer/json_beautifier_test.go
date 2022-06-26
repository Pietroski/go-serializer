package serializer

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewJsonBeautifier(t *testing.T) {
	tests := []struct {
		name string
		want Beautifier
	}{
		{
			name: "constructor call",
			want: NewJsonBeautifier(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJsonBeautifier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJsonBeautifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonBeautifier(t *testing.T) {
	t.Run("success - string", func(t *testing.T) {
		var target string
		payload := "anything"
		s := NewJsonBeautifier()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, &target)
		require.NoError(t, err)
		require.Equal(t, payload, target)
	})

	t.Run("success - struct", func(t *testing.T) {
		type TestStruct struct {
			Str string
			Num int
		}
		target := &TestStruct{}
		payload := &TestStruct{
			Str: "anything",
			Num: 10,
		}
		s := NewJsonBeautifier()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, &target)
		require.NoError(t, err)
		require.Equal(t, payload, target)
	})

	t.Run("success - stringified struct", func(t *testing.T) {
		type TestStruct struct {
			Str string `json:"str"`
			Num int    `json:"num"`
		}
		target := &TestStruct{}
		bs := []byte(`
			{
				"str": "anything",
				"num": 10
			}
		`)
		s := NewJsonBeautifier()
		err := s.Deserialize(bs, &target)
		require.NoError(t, err)
		require.Equal(t, target.Str, "anything")
		require.Equal(t, target.Num, 10)
	})

	t.Run("success - number", func(t *testing.T) {
		var target int
		payload := 100
		s := NewJsonBeautifier()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, &target)
		require.NoError(t, err)
		require.Equal(t, payload, target)
	})

	t.Run("failure - decode error - struct", func(t *testing.T) {
		type TestStruct struct {
			Str string
			Num int
		}
		target := &TestStruct{}
		payload := []byte(`
			{
				"str": "anything",
				"num": 10
			}
		`)
		s := NewJsonBeautifier()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, &target)
		require.Error(t, err)
	})

	t.Run("failure - encode error - struct", func(t *testing.T) {
		payload := make(chan string)
		s := NewJsonBeautifier()
		bs, err := s.Serialize(payload)
		require.Error(t, err)
		require.NotNil(t, bs)
		require.Empty(t, bs)
	})

	t.Run("success - map with time", func(t *testing.T) {
		target := &map[string]interface{}{}
		payload := &map[string]interface{}{
			"str":     "anything",
			"num":     10,
			"timeNow": time.Now().String(),
			"isValid": true,
			"nesting": map[string]interface{}{
				"trial": "another-anything",
			},
		}
		s := NewJsonBeautifier()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, target)
		require.NoError(t, err)
		require.NotNil(t, target)
		require.NotEmpty(t, target)
		t.Log(target)
	})

	t.Run("success - nested struct", func(t *testing.T) {
		type (
			NestedStructLevel1 struct {
				StringFieldLevel1 string
				IntFieldLevel1    int
				TimeFieldLevel1   time.Time
			}

			TestStruct struct {
				StringField string
				IntField    int
				TimeField   time.Time
				StructField NestedStructLevel1
			}
		)
		target := &TestStruct{}
		payload := &TestStruct{
			StringField: "anything",
			IntField:    10,
			TimeField:   time.Now(),
			StructField: NestedStructLevel1{
				StringFieldLevel1: "anything-nested",
				IntFieldLevel1:    50,
				TimeFieldLevel1:   time.Now(),
			},
		}
		s := NewJsonBeautifier()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, target)
		require.NoError(t, err)
		require.Equal(t, payload.StringField, target.StringField)
		require.Equal(t, payload.IntField, target.IntField)
		require.True(t, payload.TimeField.Equal(target.TimeField))
		require.Equal(t, payload.StructField.StringFieldLevel1, target.StructField.StringFieldLevel1)
		require.Equal(t, payload.StructField.IntFieldLevel1, target.StructField.IntFieldLevel1)
		require.True(t, payload.StructField.TimeFieldLevel1.Equal(target.StructField.TimeFieldLevel1))
	})

	t.Run(
		"test beautifier",
		func(t *testing.T) {
			type (
				NestedStructLevel1 struct {
					StringFieldLevel1 string
					IntFieldLevel1    int
					TimeFieldLevel1   time.Time
				}

				TestStruct struct {
					StringField string
					IntField    int
					TimeField   time.Time
					StructField NestedStructLevel1
				}
			)
			target := &TestStruct{}
			payload := &TestStruct{
				StringField: "anything",
				IntField:    10,
				TimeField:   time.Now(),
				StructField: NestedStructLevel1{
					StringFieldLevel1: "anything-nested",
					IntFieldLevel1:    50,
					TimeFieldLevel1:   time.Now(),
				},
			}
			s := NewJsonBeautifier()
			bs, err := s.Beautify(payload, "", "  ")
			require.NoError(t, err)

			err = s.Deserialize(bs, target)
			require.NoError(t, err)
			require.Equal(t, payload.StringField, target.StringField)
			require.Equal(t, payload.IntField, target.IntField)
			require.True(t, payload.TimeField.Equal(target.TimeField))
			require.Equal(t, payload.StructField.StringFieldLevel1, target.StructField.StringFieldLevel1)
			require.Equal(t, payload.StructField.IntFieldLevel1, target.StructField.IntFieldLevel1)
			require.True(t, payload.StructField.TimeFieldLevel1.Equal(target.StructField.TimeFieldLevel1))
		},
	)

	t.Run("failure - encode error - struct", func(t *testing.T) {
		payload := make(chan string)
		s := NewJsonBeautifier()
		bs, err := s.Beautify(payload, "", "  ")
		require.Error(t, err)
		require.NotNil(t, bs)
		require.Empty(t, bs)
	})

	t.Run("success rebind - stringified struct", func(t *testing.T) {
		type TestStructPayload struct {
			Str string `json:"str"`
			Num int    `json:"num"`
		}

		type TestStructDestination struct {
			Str string `json:"str" validate:"required"`
			Num int    `json:"num" validate:"required"`
		}
		payload := &TestStructPayload{
			Str: "any-string",
			Num: 10,
		}
		var target TestStructDestination

		s := NewJsonBeautifier()
		err := s.DataRebind(payload, &target)
		require.NoError(t, err)
		require.Equal(t, target.Str, "any-string")
		require.Equal(t, target.Num, 10)
	})

	t.Run("fail rebind serialization - stringified struct", func(t *testing.T) {
		type TestStructPayload struct {
			Str  string `json:"str"`
			Num  int    `json:"num"`
			Chan chan int
		}

		type TestStructDestination struct {
			Str  string   `json:"str" validate:"required"`
			Num  int      `json:"num" validate:"required"`
			Chan chan int `json:"my-lil-channel"`
		}
		payload := &TestStructPayload{
			Str:  "any-string",
			Num:  10,
			Chan: make(chan int),
		}
		var target TestStructDestination

		s := NewJsonBeautifier()
		err := s.DataRebind(payload, &target)
		require.Error(t, err)
	})

	t.Run("fail rebind serialization - stringified struct", func(t *testing.T) {
		type TestStructPayload struct {
			Str string `json:"str"`
			Num int    `json:"num"`
		}

		type TestStructDestination int8
		payload := &TestStructPayload{
			Str: "any-string",
			Num: 10,
		}
		var target TestStructDestination

		s := NewJsonBeautifier()
		err := s.DataRebind(payload, &target)
		require.Error(t, err)
	})
}
