package go_serializer

import (
	"fmt"

	"github.com/vmihailenco/msgpack/v5"

	"gitlab.com/pietroski-software-company/devex/golang/serializer/models"
)

type msgpackSerializer struct{}

func NewMsgPackSerializer() Serializer {
	return &msgpackSerializer{}
}

func (s *msgpackSerializer) Serialize(payload interface{}) ([]byte, error) {
	bs, err := msgpack.Marshal(payload)
	if err != nil {
		return []byte{}, fmt.Errorf(models.EncodeErrMsg, err)
	}

	return bs, nil
}

func (s *msgpackSerializer) Deserialize(payload []byte, target interface{}) error {
	if err := msgpack.Unmarshal(payload, target); err != nil {
		return fmt.Errorf(models.DecodeErrMsg, err)
	}

	return nil
}

func (s *msgpackSerializer) DataRebind(payload interface{}, target interface{}) error {
	bs, err := s.Serialize(payload)
	if err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	if err = s.Deserialize(bs, target); err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	return nil
}
