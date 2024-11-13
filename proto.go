package serializer

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"gitlab.com/pietroski-software-company/devex/golang/serializer/models"
)

type ProtoSerializer struct{}

func NewProtoSerializer() *ProtoSerializer {
	return &ProtoSerializer{}
}

func (s *ProtoSerializer) Serialize(payload interface{}) ([]byte, error) {
	protoPayload, ok := payload.(proto.Message)
	if !ok {
		return []byte{}, fmt.Errorf(models.WrongPayloadTypeErrMsg)
	}
	bs, err := proto.Marshal(protoPayload)
	if err != nil {
		return []byte{}, fmt.Errorf(models.EncodeErrMsg, err)
	}

	return bs, err
}

func (s *ProtoSerializer) Deserialize(payload []byte, target interface{}) error {
	protoTarget, ok := target.(proto.Message)
	if !ok {
		return fmt.Errorf(models.WrongTargetTypeErrMsg)
	}
	if err := proto.Unmarshal(payload, protoTarget); err != nil {
		return fmt.Errorf(models.DecodeErrMsg, err)
	}

	return nil
}

func (s *ProtoSerializer) DataRebind(payload interface{}, target interface{}) error {
	bs, err := s.Serialize(payload)
	if err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	if err = s.Deserialize(bs, target); err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	return nil
}
