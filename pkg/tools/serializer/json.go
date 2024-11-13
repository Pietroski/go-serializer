package go_serializer

import (
	"encoding/json"
	"fmt"
)

type (
	jsonSerializer struct{}
)

func NewJsonSerializer() Serializer {
	return &jsonSerializer{}
}

func (s *jsonSerializer) Serialize(payload interface{}) ([]byte, error) {
	bs, err := json.Marshal(payload)
	if err != nil {
		return []byte{}, fmt.Errorf(EncodeErrMsg, err)
	}

	return bs, err
}

func (s *jsonSerializer) Deserialize(payload []byte, target interface{}) error {
	if err := json.Unmarshal(payload, target); err != nil {
		return fmt.Errorf(DecodeErrMsg, err)
	}

	return nil
}

func (s *jsonSerializer) DataRebind(payload interface{}, target interface{}) error {
	bs, err := s.Serialize(payload)
	if err != nil {
		return fmt.Errorf(RebinderErrMsg, err)
	}

	if err = s.Deserialize(bs, target); err != nil {
		return fmt.Errorf(RebinderErrMsg, err)
	}

	return nil
}
