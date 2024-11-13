package testmodels

type (
	IntSliceTestData struct {
		IntList []int `json:"int_list,omitempty"`
	}

	Int8SliceTestData struct {
		Int8List []int8 `json:"int8_list,omitempty"`
	}

	Int16SliceTestData struct {
		Int16List []int16 `json:"int16_list,omitempty"`
	}

	Int32SliceTestData struct {
		Int32List []int32 `json:"int32_list,omitempty"`
	}

	Int64SliceTestData struct {
		Int64List []int64 `json:"int64_list,omitempty"`
	}

	UintSliceTestData struct {
		UintList []uint `json:"uint_list,omitempty"`
	}

	Uint8SliceTestData struct {
		Uint8List []uint8 `json:"uint8_list,omitempty"`
	}

	Uint16SliceTestData struct {
		Uint16List []uint16 `json:"uint16_list,omitempty"`
	}

	Uint32SliceTestData struct {
		Uint32List []uint32 `json:"uint32_list,omitempty"`
	}

	Uint64SliceTestData struct {
		Uint64List []uint64 `json:"uint64_list,omitempty"`
	}

	ByteSliceTestData struct {
		ByteList []byte `json:"byte_list,omitempty"`
	}

	StringSliceTestData struct {
		StringList []string `json:"string_list,omitempty"`
	}

	StringsStruct struct {
		FirstString  string `json:"first_string,omitempty"`
		SecondString string `json:"second_string,omitempty"`
		ThirdString  string `json:"third_string,omitempty"`
		FourthString string `json:"fourth_string,omitempty"`
		FifthString  string `json:"fifth_string,omitempty"`
	}

	Int64Struct struct {
		FirstInt64  int64 `json:"first_int64,omitempty"`
		SecondInt64 int64 `json:"second_int64,omitempty"`
		ThirdInt64  int64 `json:"third_int64,omitempty"`
		FourthInt64 int64 `json:"fourth_int64,omitempty"`
		FifthInt64  int64 `json:"fifth_int64,omitempty"`
		SixthInt64  int64 `json:"sixth_int64,omitempty"`
	}

	Item struct {
		Id      string   `json:"id,omitempty"`
		ItemId  uint64   `json:"itemId,omitempty"`
		Number  int64    `json:"number,omitempty"`
		SubItem *SubItem `json:"subItem,omitempty"`
	}

	SubItem struct {
		Date     int64  `json:"date,omitempty"`
		Amount   int64  `json:"amount,omitempty"`
		ItemCode string `json:"itemCode,omitempty"`
	}

	TestData struct {
		FieldStr   string
		FieldInt   int8
		FieldBool  bool
		FieldBytes []byte

		FieldStrPtr   *string
		FieldIntPtr   *int
		FieldBoolPtr  *bool
		FieldBytesPtr *[]byte

		SubTestData    SubTestData
		SubTestDataPtr *SubTestData
		SliceTestData  SliceTestData
		MapTestData    MapTestData
	}

	SubTestData struct {
		FieldStr   string
		FieldInt32 int32
		FieldBool  bool
		FieldInt64 int64
		FieldInt   int

		FieldStrPtr   *string
		FieldInt32Ptr *int32
		FieldBoolPtr  *bool
		FieldInt64Ptr *int64
		FieldIntPtr   *int
	}

	SliceTestData struct {
		IntList       []int
		IntIntList    [][]int
		ThreeDIntList [][][]int

		StrList    []string
		StrStrList [][]string

		StructList       []SliceItem
		PtrStructList    []*SliceItem
		PtrStructNilList []*SliceItem
	}

	MapTestData struct {
		Int64KeyMapInt64Value map[int64]int64
		StrKeyMapStrValue     map[string]string
	}

	SliceItem struct {
		Int  int
		Str  string
		Bool bool
	}

	ProtoTypeSliceTestData struct {
		IntList        []int64      `json:"int_list,omitempty"`
		UintList       []uint64     `json:"uint_list,omitempty"`
		StrList        []string     `json:"str_list,omitempty"`
		StructList     []SliceItem  `json:"struct_list,omitempty"`
		PtrStructList  []*SliceItem `json:"ptr_struct_list,omitempty"`
		BytesBytesList [][]byte     `json:"bytes_bytes_list,omitempty"`
		BytesList      []byte       `json:"bytes_list,omitempty"`
	}

	ProtoEquivalentTestData struct {
		FieldStr   string
		FieldInt   int8
		FieldBool  bool
		FieldBytes []byte

		FieldStrPtr   *string
		FieldIntPtr   *int
		FieldBoolPtr  *bool
		FieldBytesPtr *[]byte

		SubTestData    SubTestData
		SubTestDataPtr *SubTestData
		SliceTestData  SliceTestData
		MapTestData    MapTestData
	}
)
