package serializer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewGobSerializer(t *testing.T) {
	t.Run("success - string", func(t *testing.T) {
		var target string
		payload := "anything"
		s := NewGobSerializer()
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
		s := NewGobSerializer()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, target)
		require.NoError(t, err)
		require.Equal(t, payload, target)
	})

	t.Run("success - number", func(t *testing.T) {
		var target int
		payload := 100
		s := NewGobSerializer()
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
		s := NewGobSerializer()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, &target)
		require.Error(t, err)
	})

	t.Run("failure - encode error - struct", func(t *testing.T) {
		payload := make(chan string)
		s := NewGobSerializer()
		bs, err := s.Serialize(payload)
		require.Error(t, err)
		require.NotNil(t, bs)
		require.Empty(t, bs)
	})

	t.Run("success - struct with time", func(t *testing.T) {
		type TestStruct struct {
			Str        string
			Num        int
			TimeNow    time.Time
			PtrTimeNow *time.Time
		}
		timeNow := time.Now()
		target := &TestStruct{}
		payload := &TestStruct{
			Str:        "anything",
			Num:        10,
			TimeNow:    timeNow,
			PtrTimeNow: &timeNow,
		}
		s := NewGobSerializer()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, target)
		require.NoError(t, err)
		require.Equal(t, payload.Str, target.Str)
		require.Equal(t, payload.Num, target.Num)
		require.True(t, payload.TimeNow.Equal(target.TimeNow))
		require.True(t, (*payload).PtrTimeNow.Equal(*target.PtrTimeNow))
	})

	t.Run("failure - map with time and nesting", func(t *testing.T) {
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
		s := NewGobSerializer()
		bs, err := s.Serialize(payload)
		require.Error(t, err)
		require.NotNil(t, bs)
		require.Empty(t, bs)

		//err = s.Deserialize(bs, target)
		//require.NoError(t, err)
		//require.NotNil(t, target)
		//require.NotEmpty(t, target)
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
		s := NewGobSerializer()
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
}
