package serializer

import (
	"testing"
	"time"
)

type (
	BenchStruct struct {
		StringField string
		IntField    int
		TimeField   time.Time
		StructField NestedStructLevel1
	}

	NestedStructLevel1 struct {
		NestedStringFieldLevel1 string
		NestedIntFieldLevel1    int
		NestedTimeFieldLevel1   time.Time
	}
)

var (
	gobS     = NewGobSerializer()
	jsonS    = NewJsonSerializer()
	msgpackS = NewMsgPackSerializer()
)

func BenchmarkSerializers(b *testing.B) {
	ts := BenchStruct{
		StringField: "anything",
		IntField:    10,
		TimeField:   time.Now(),
		StructField: NestedStructLevel1{
			NestedStringFieldLevel1: "anything-nested",
			NestedIntFieldLevel1:    50,
			NestedTimeFieldLevel1:   time.Now(),
		},
	}

	var gobBS []byte
	b.Run("benchmark glob serializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gobBS, _ = gobS.Serialize(ts)
		}
	})

	var gobPayload BenchStruct
	b.Run("benchmark glob deserializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = gobS.Deserialize(gobBS, &gobPayload)
		}
	})

	var jsonBS []byte
	b.Run("benchmark json serializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			jsonBS, _ = jsonS.Serialize(ts)
		}
	})

	var jsonPayload BenchStruct
	b.Run("benchmark json deserializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = jsonS.Deserialize(jsonBS, &jsonPayload)
		}
	})

	var msgpackBS []byte
	b.Run("benchmark masgpack serializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msgpackBS, _ = msgpackS.Serialize(ts)
		}
	})

	var msgpackPayload BenchStruct
	b.Run("benchmark masgpack deserializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = msgpackS.Deserialize(msgpackBS, &msgpackPayload)
		}
	})
}

func BenchmarkSerializers2(b *testing.B) {
	type A struct {
		Name     string
		BirthDay time.Time
		Phone    string
		Siblings int
		Spouse   bool
		Money    float64
	}
	ts := A{
		Name:     "any-name",
		BirthDay: time.Now(),
		Phone:    "any-phone",
		Siblings: 50,
		Spouse:   true,
		Money:    5_000_000,
	}

	var gobBS []byte
	b.Run("benchmark glob serializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gobBS, _ = gobS.Serialize(ts)
		}
	})

	var gobPayload BenchStruct
	b.Run("benchmark glob deserializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = gobS.Deserialize(gobBS, &gobPayload)
		}
	})

	var jsonBS []byte
	b.Run("benchmark json serializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			jsonBS, _ = jsonS.Serialize(ts)
		}
	})

	var jsonPayload BenchStruct
	b.Run("benchmark json deserializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = jsonS.Deserialize(jsonBS, &jsonPayload)
		}
	})

	var msgpackBS []byte
	b.Run("benchmark masgpack serializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msgpackBS, _ = msgpackS.Serialize(ts)
		}
	})

	var msgpackPayload BenchStruct
	b.Run("benchmark masgpack deserializer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = msgpackS.Deserialize(msgpackBS, &msgpackPayload)
		}
	})
}
