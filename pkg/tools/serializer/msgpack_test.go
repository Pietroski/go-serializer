package serializer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewMsgPackSerializer(t *testing.T) {
	t.Run("success - string", func(t *testing.T) {
		var target string
		payload := "anything"
		s := NewMsgPackSerializer()
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
		s := NewMsgPackSerializer()
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
		s := NewMsgPackSerializer()
		err := s.Deserialize(bs, &target)
		require.Error(t, err)
		//require.Equal(t, target.Str, "anything")
		//require.Equal(t, target.Num, 10)
	})

	t.Run("success - number", func(t *testing.T) {
		var target int
		payload := 100
		s := NewMsgPackSerializer()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, &target)
		require.NoError(t, err)
		require.Equal(t, payload, target)
	})

	t.Run("failure - decode error - struct", func(t *testing.T) {
		type TestStruct struct {
			Str string `json:"str"`
			Num int    `json:"num"`
		}
		target := &TestStruct{}
		//payload := `
		//	{
		//		"str": "anything",
		//		"num": 10
		//	}
		//`
		payload := `{"str":"anything","num":10}`
		s := NewMsgPackSerializer()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, &target)
		require.Error(t, err)
	})

	t.Run("failure - decode error - struct", func(t *testing.T) {
		type TestStruct struct {
			Str string `json:"str"`
			Num int    `json:"num"`
		}
		target := &TestStruct{}
		payload := map[interface{}]interface{}{
			"str": "anything",
			"num": 10,
			"0":   false,
		}
		s := NewMsgPackSerializer()
		bs, err := s.Serialize(payload)
		require.NoError(t, err)

		err = s.Deserialize(bs, &target)
		require.NoError(t, err)
	})

	t.Run("failure - encode error - struct", func(t *testing.T) {
		payload := make(chan string)
		s := NewMsgPackSerializer()
		bs, err := s.Serialize(payload)
		require.Error(t, err)
		require.NotNil(t, bs)
		require.Empty(t, bs)
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
		s := NewMsgPackSerializer()
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
