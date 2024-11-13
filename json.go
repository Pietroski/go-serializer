package serializer

import (
	"encoding/json"
	"fmt"

	"gitlab.com/pietroski-software-company/devex/golang/serializer/models"
)

type jsonSerializer struct{}

func NewJsonSerializer() models.Serializer {
	return &jsonSerializer{}
}

func (s *jsonSerializer) Serialize(payload interface{}) ([]byte, error) {
	bs, err := json.Marshal(payload)
	if err != nil {
		return []byte{}, fmt.Errorf(models.EncodeErrMsg, err)
	}

	return bs, err
}

func (s *jsonSerializer) Deserialize(payload []byte, target interface{}) error {
	if err := json.Unmarshal(payload, target); err != nil {
		return fmt.Errorf(models.DecodeErrMsg, err)
	}

	return nil
}

func (s *jsonSerializer) DataRebind(payload interface{}, target interface{}) error {
	bs, err := s.Serialize(payload)
	if err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	if err = s.Deserialize(bs, target); err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	return nil
}
