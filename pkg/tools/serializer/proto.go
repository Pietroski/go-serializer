package go_serializer

import (
	"fmt"
	"gitlab.com/pietroski-software-company/devex/golang/serializer/models"
	"google.golang.org/protobuf/proto"
)

type (
	protoSerializer struct{}
)

func NewProtoSerializer() models.Serializer {
	return &protoSerializer{}
}

func (s *protoSerializer) Serialize(payload interface{}) ([]byte, error) {
	protoPayload, ok := payload.(proto.Message)
	if !ok {
		return []byte{}, fmt.Errorf(models.WrongPayloadTypeErrMsg, nil)
	}
	bs, err := proto.Marshal(protoPayload)
	if err != nil {
		return []byte{}, fmt.Errorf(models.EncodeErrMsg, err)
	}

	return bs, err
}

func (s *protoSerializer) Deserialize(payload []byte, target interface{}) error {
	protoTarget, ok := target.(proto.Message)
	if !ok {
		return fmt.Errorf(models.WrongTargetTypeErrMsg, nil)
	}
	if err := proto.Unmarshal(payload, protoTarget); err != nil {
		return fmt.Errorf(models.DecodeErrMsg, err)
	}

	return nil
}

func (s *protoSerializer) DataRebind(payload interface{}, target interface{}) error {
	bs, err := s.Serialize(payload)
	if err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	if err = s.Deserialize(bs, target); err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	return nil
}
