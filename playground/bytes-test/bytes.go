package main

import (
	"fmt"
)

var file_item_item_messages_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x22, 0x51, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x6f, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x69,
	0x74, 0x65, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x49, 0x74, 0x65, 0x6d, 0x22, 0x58, 0x0a, 0x0e, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53,
	0x75, 0x62, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x7d,
	0x0a, 0x0b, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x75, 0x62,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x73, 0x75, 0x62, 0x49, 0x74, 0x65, 0x6d, 0x22, 0xda, 0x01,
	0x0a, 0x0d, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x54, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74,
	0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x5f, 0x6e, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x10, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0e, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x09, 0x53, 0x6c,
	0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x22,
	0xe1, 0x02, 0x0a, 0x0b, 0x4d, 0x61, 0x70, 0x54, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x66, 0x0a, 0x19, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6d, 0x61, 0x70,
	0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x4d, 0x61, 0x70, 0x54, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4b, 0x65, 0x79, 0x4d, 0x61,
	0x70, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x15, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x5a, 0x0a, 0x15, 0x73, 0x74, 0x72, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x4d, 0x61,
	0x70, 0x54, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x72, 0x4b, 0x65, 0x79,
	0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x11, 0x73, 0x74, 0x72, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x48, 0x0a, 0x1a, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4b, 0x65, 0x79, 0x4d,
	0x61, 0x70, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x44, 0x0a,
	0x16, 0x53, 0x74, 0x72, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x75, 0x5a, 0x73, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x69, 0x65, 0x74, 0x72, 0x6f, 0x73, 0x6b, 0x69, 0x2d, 0x73, 0x6f, 0x66, 0x74,
	0x77, 0x61, 0x72, 0x65, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x74, 0x65, 0x6d,
	0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_item_error_item_messages_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x74, 0x65, 0x6d, 0x2d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x65, 0x0a, 0x09,
	0x49, 0x74, 0x65, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x42, 0x81, 0x01, 0x5a, 0x7f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x65, 0x74, 0x72, 0x6f, 0x73, 0x6b, 0x69, 0x2d, 0x73, 0x6f, 0x66,
	0x74, 0x77, 0x61, 0x72, 0x65, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x74, 0x65,
	0x6d, 0x2d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

func main() {
	fmt.Printf("%s", file_item_item_messages_proto_rawDesc)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Printf("%s", file_item_error_item_messages_proto_rawDesc)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	v := 256
	fmt.Println(v)
	fmt.Println(v < 1<<7)
	fmt.Println(v < 1<<14)
	fmt.Println(v < 1<<21)
	fmt.Println(v < 1<<28)
	fmt.Println((v>>0)&0x7f | 0x80)
	fmt.Println((v>>7)&0x7f | 0x80)
	fmt.Println(v >> 14)
}
