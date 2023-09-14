package go_proto_mock

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

type ProtoMessage proto.Message

type ProtoReflectProtoMessage protoreflect.ProtoMessage

type ProtoReflectMessage protoreflect.Message

type ProtoReflectStruct struct{}

func (p *ProtoReflectStruct) ProtoMethods() *protoiface.Methods {
	return &protoiface.Methods{
		Flags: 0,
		Size: func(input protoiface.SizeInput) protoiface.SizeOutput {
			return protoiface.SizeOutput{}
		},
		Marshal: func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
			return protoiface.MarshalOutput{}, nil
		},
		Unmarshal: func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
			return protoiface.UnmarshalOutput{}, nil
		},
		Merge: func(input protoiface.MergeInput) protoiface.MergeOutput {
			return protoiface.MergeOutput{}
		},
		CheckInitialized: func(input protoiface.CheckInitializedInput) (protoiface.CheckInitializedOutput, error) {
			return protoiface.CheckInitializedOutput{}, nil
		},
	}
}

//func test() {
//	nilPRMsg := protoreflect.Message(nil)
//	prs := ProtoReflectStruct{}
//	nilPRMsg.ProtoMethods = func() *interface{} {
//		pm := prs.ProtoMethods()
//		i := interface{}(pm)
//		return &i
//	}
//}
