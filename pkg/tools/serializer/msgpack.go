package go_serializer

import (
	"github.com/vmihailenco/msgpack/v5"

	error_builder "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/builder/errors"
)

type msgpackSerializer struct{}

func NewMsgPackSerializer() Serializer {
	return &msgpackSerializer{}
}

func (s *msgpackSerializer) Serialize(payload interface{}) ([]byte, error) {
	bs, err := msgpack.Marshal(payload)
	if err != nil {
		return []byte{}, error_builder.Err(EncodeErrMsg, err)
	}

	return bs, nil
}

func (s *msgpackSerializer) Deserialize(payload []byte, target interface{}) error {
	if err := msgpack.Unmarshal(payload, target); err != nil {
		return error_builder.Err(DecodeErrMsg, err)
	}

	return nil
}

func (s *msgpackSerializer) DataRebind(payload interface{}, target interface{}) error {
	bs, err := s.Serialize(payload)
	if err != nil {
		return error_builder.Err(RebinderErrMsg, err)
	}

	if err = s.Deserialize(bs, target); err != nil {
		return error_builder.Err(RebinderErrMsg, err)
	}

	return nil
}
