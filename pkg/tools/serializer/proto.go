package go_serializer

import (
	"google.golang.org/protobuf/proto"

	error_builder "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/builder/errors"
)

type (
	protoSerializer struct{}
)

func NewProtoSerializer() Serializer {
	return &protoSerializer{}
}

func (s *protoSerializer) Serialize(payload interface{}) ([]byte, error) {
	protoPayload, ok := payload.(proto.Message)
	if !ok {
		return []byte{}, error_builder.Err(WrongPayloadTypeErrMsg, nil)
	}
	bs, err := proto.Marshal(protoPayload)
	if err != nil {
		return []byte{}, error_builder.Err(EncodeErrMsg, err)
	}

	return bs, err
}

func (s *protoSerializer) Deserialize(payload []byte, target interface{}) error {
	protoTarget, ok := target.(proto.Message)
	if !ok {
		return error_builder.Err(WrongTargetTypeErrMsg, nil)
	}
	if err := proto.Unmarshal(payload, protoTarget); err != nil {
		return error_builder.Err(DecodeErrMsg, err)
	}

	return nil
}

func (s *protoSerializer) DataRebind(payload interface{}, target interface{}) error {
	bs, err := s.Serialize(payload)
	if err != nil {
		return error_builder.Err(RebinderErrMsg, err)
	}

	if err = s.Deserialize(bs, target); err != nil {
		return error_builder.Err(RebinderErrMsg, err)
	}

	return nil
}
