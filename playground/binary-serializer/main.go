package main

import (
	"fmt"
	go_serializer "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/serializer"
	"math"
)

type (
	TestData struct {
		FieldStr  string `binary:"1"`
		FieldInt  int8   `binary:"2"`
		FieldBool bool   `binary:"3"`

		FieldStrPtr *string `binary:"4"`
		//FieldIntPtr *int `binary:"5"`
		//FieldBoolPtr *bool   `binary:"6"`
		SubTestData    SubTestData
		SubTestDataPtr *SubTestData
	}

	SubTestData struct {
		FieldStr   string
		FieldInt32 int32
		FieldBool  bool
		FieldInt64 int64
		FieldInt   int
	}
)

func main() {
	serializer := go_serializer.NewBinarySerializer()

	strPtr := "test-str-ptr"
	testData := TestData{
		FieldStr:  "test-data",
		FieldInt:  127,
		FieldBool: true,

		FieldStrPtr: &strPtr,

		SubTestData: SubTestData{
			FieldStr:   "test-sub-data",
			FieldInt32: 127567,
			FieldBool:  false,
			FieldInt64: math.MaxInt64,
			FieldInt:   0,
		},

		SubTestDataPtr: &SubTestData{
			FieldStr:   "test-sub-data-ptr",
			FieldInt32: 765432,
			FieldBool:  true,
			FieldInt64: math.MaxInt16,
			FieldInt:   5432,
		},
	}

	bs, err := serializer.Marshal(&testData)
	if err != nil {
		panic(err)
	}

	//t.Log(string(bs), bs)

	var td TestData
	err = serializer.Unmarshal(bs, &td)
	if err != nil {
		panic(err)
	}

	fmt.Println(td)
	fmt.Println(*td.FieldStrPtr)
	fmt.Println(*td.SubTestDataPtr)
}
