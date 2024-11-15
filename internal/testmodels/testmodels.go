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

	Float32SliceTestData struct {
		Float32List []float32 `json:"float32_list,omitempty"`
	}

	Float64SliceTestData struct {
		Float64List []float64 `json:"float64_list,omitempty"`
	}

	Complex64SliceTestData struct {
		Complex64List []complex64 `json:"complex64_list,omitempty"`
	}

	Complex128SliceTestData struct {
		Complex128List []complex128 `json:"complex128_list,omitempty"`
	}

	ByteSliceTestData struct {
		ByteList []byte `json:"byte_list,omitempty"`
	}

	ByteByteSliceTestData struct {
		ByteByteList [][]byte `json:"byte_byte_list,omitempty"`
	}

	StringSliceTestData struct {
		StringList []string `json:"string_list,omitempty"`
	}

	StructSliceTestData struct {
		StructList []StructTestData `json:"struct_list,omitempty"`
	}

	StructPointerSliceTestData struct {
		StructPointerList []*StructTestData `json:"struct_pointer_list,omitempty"`
	}

	StructTestData struct {
		Bool   bool   `json:"bool,omitempty"`
		String string `json:"string,omitempty"`
		Int64  int64  `json:"int_64,omitempty"`
	}

	MapIntIntTestData struct {
		MapIntInt map[int]int `json:"map_int_int,omitempty"`
	}

	MapInt64Int64TestData struct {
		MapInt64Int64 map[int64]int64 `json:"map_int64_int64,omitempty"`
	}

	MapStringStringTestData struct {
		MapStringString map[string]string `json:"map_string_string,omitempty"`
	}

	MapIntStructTestData struct {
		MapIntStruct map[int]StructTestData `json:"map_int_struct_test_data,omitempty"`
	}

	MapIntStructPointerTestData struct {
		MapIntStructPointer map[int]*StructTestData `json:"map_int_struct_pointer_test_data,omitempty"`
	}

	MapInt64StructTestData struct {
		MapInt64Struct map[int64]StructTestData `json:"map_int64_struct_test_data,omitempty"`
	}

	MapInt64StructPointerTestData struct {
		MapInt64StructPointer map[int64]*StructTestData `json:"map_int64_struct_pointer_test_data,omitempty"`
	}

	MapStringStructTestData struct {
		MapStringStruct map[string]StructTestData `json:"map_string_struct_test_data,omitempty"`
	}

	MapStringStructPointerTestData struct {
		MapStringStructPointer map[string]*StructTestData `json:"map_string_struct_pointer_test_data,omitempty"`
	}

	SpecialStructTestData struct {
		Bool       bool       `json:"bool,omitempty"`
		String     string     `json:"string,omitempty"`
		Int        int        `json:"int,omitempty"`
		Int8       int8       `json:"int8,omitempty"`
		Int16      int16      `json:"int16,omitempty"`
		Int32      int32      `json:"int32,omitempty"`
		Int64      int64      `json:"int64,omitempty"`
		Uint       uint       `json:"uint,omitempty"`
		Uint8      uint8      `json:"uint8,omitempty"`
		Uint16     uint16     `json:"uint16,omitempty"`
		Uint32     uint32     `json:"uint32,omitempty"`
		Uint64     uint64     `json:"uint64,omitempty"`
		Float32    float32    `json:"float32,omitempty"`
		Float64    float64    `json:"float64,omitempty"`
		Complex    complex64  `json:"complex,omitempty"`
		Complex128 complex128 `json:"complex128,omitempty"`
		Byte       byte       `json:"byte,omitempty"`
		Bytes      []byte     `json:"bytes,omitempty"`
		BytesBytes [][]byte   `json:"bytes_bytes,omitempty"`
	}

	SimplifiedSpecialStructSliceTestData struct {
		SimplifiedSpecialStructSliceTestData []SimplifiedSpecialStructTestData `json:"simplified_special_struct_test_data,omitempty"`
	}

	SimplifiedSpecialStructPointerSliceTestData struct {
		SimplifiedSpecialStructPointerSliceTestData []*SimplifiedSpecialStructTestData `json:"simplified_special_struct_test_data,omitempty"`
	}

	SimplifiedSpecialStructTestData struct {
		Bool          bool     `json:"bool,omitempty"`
		String        string   `json:"string,omitempty"`
		Int32         int32    `json:"int32,omitempty"`
		Int64         int64    `json:"int64,omitempty"`
		Uint32        uint32   `json:"uint32,omitempty"`
		Uint64        uint64   `json:"uint64,omitempty"`
		Float32       float32  `json:"float32,omitempty"`
		Float64       float64  `json:"float64,omitempty"`
		Bytes         []byte   `json:"bytes,omitempty"`
		RepeatedBytes [][]byte `json:"repeated_bytes,omitempty"`
	}

	StringStruct struct {
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
		FieldStr   string `json:"field_str,omitempty"`
		FieldInt   int8   `json:"field_int,omitempty"`
		FieldBool  bool   `json:"field_bool,omitempty"`
		FieldBytes []byte `json:"field_bytes,omitempty"`

		FieldStrPtr   *string `json:"field_str_ptr,omitempty"`
		FieldIntPtr   *int    `json:"field_int_ptr,omitempty"`
		FieldBoolPtr  *bool   `json:"field_bool_ptr,omitempty"`
		FieldBytesPtr *[]byte `json:"field_bytes_ptr,omitempty"`

		SubTestData    SubTestData   `json:"sub_test_data,omitempty"`
		SubTestDataPtr *SubTestData  `json:"sub_test_data_ptr,omitempty"`
		SliceTestData  SliceTestData `json:"slice_test_data,omitempty"`
		MapTestData    MapTestData   `json:"map_test_data,omitempty"`
	}

	SubTestData struct {
		FieldStr   string `json:"field_str,omitempty"`
		FieldInt32 int32  `json:"field_int_32,omitempty"`
		FieldBool  bool   `json:"field_bool,omitempty"`
		FieldInt64 int64  `json:"field_int_64,omitempty"`
		FieldInt   int    `json:"field_int,omitempty"`

		FieldStrPtr   *string `json:"field_str_ptr,omitempty"`
		FieldInt32Ptr *int32  `json:"field_int_32_ptr,omitempty"`
		FieldBoolPtr  *bool   `json:"field_bool_ptr,omitempty"`
		FieldInt64Ptr *int64  `json:"field_int_64_ptr,omitempty"`
		FieldIntPtr   *int    `json:"field_int_ptr,omitempty"`
	}

	SliceTestData struct {
		IntList       []int     `json:"int_list,omitempty"`
		IntIntList    [][]int   `json:"int_int_list,omitempty"`
		ThreeDIntList [][][]int `json:"three_d_int_list,omitempty"`

		StrList    []string   `json:"str_list,omitempty"`
		StrStrList [][]string `json:"str_str_list,omitempty"`

		StructList       []SliceItem  `json:"struct_list,omitempty"`
		PtrStructList    []*SliceItem `json:"ptr_struct_list,omitempty"`
		PtrStructNilList []*SliceItem `json:"ptr_struct_nil_list,omitempty"`
	}

	MapTestData struct {
		Int64KeyMapInt64Value map[int64]int64
		StrKeyMapStrValue     map[string]string
	}

	SliceItem struct {
		Int  int    `json:"int,omitempty"`
		Str  string `json:"str,omitempty"`
		Bool bool   `json:"bool,omitempty"`
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
		FieldStr   string `json:"field_str,omitempty"`
		FieldInt   int8   `json:"field_int,omitempty"`
		FieldBool  bool   `json:"field_bool,omitempty"`
		FieldBytes []byte `json:"field_bytes,omitempty"`

		FieldStrPtr   *string `json:"field_str_ptr,omitempty"`
		FieldIntPtr   *int    `json:"field_int_ptr,omitempty"`
		FieldBoolPtr  *bool   `json:"field_bool_ptr,omitempty"`
		FieldBytesPtr *[]byte `json:"field_bytes_ptr,omitempty"`

		SubTestData    SubTestData   `json:"sub_test_data,omitempty"`
		SubTestDataPtr *SubTestData  `json:"sub_test_data_ptr,omitempty,omitempty"`
		SliceTestData  SliceTestData `json:"slice_test_data,omitempty"`
		MapTestData    MapTestData   `json:"map_test_data,omitempty"`
	}
)
