package serializer

import (
	"encoding/json"

	error_builder "gitlab.com/pietroski-software-company/watcher/serializer/go-serializer/pkg/tools/builder/errors"
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
		return []byte{}, error_builder.Err(EncodeErrMsg, err)
	}

	return bs, err
}

func (s *jsonSerializer) Deserialize(payload []byte, target interface{}) error {
	if err := json.Unmarshal(payload, target); err != nil {
		return error_builder.Err(DecodeErrMsg, err)
	}

	return nil
}
