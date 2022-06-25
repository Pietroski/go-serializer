package serializer

import (
	"encoding/json"

	error_builder "gitlab.com/pietroski-software-company/watcher/serializer/go-serializer/pkg/tools/builder/errors"
)

type (
	jsonBeautifier struct {
		jsonSerializer Serializer
	}
)

func NewJsonBeautifier() Beautifier {
	return &jsonBeautifier{
		jsonSerializer: NewJsonSerializer(),
	}
}

func (b *jsonBeautifier) Beautify(payload interface{}, prefix string, indent string) ([]byte, error) {
	bs, err := json.MarshalIndent(payload, prefix, indent)
	if err != nil {
		return []byte{}, error_builder.Err(EncodeErrMsg, err)
	}

	return bs, err
}

func (b *jsonBeautifier) Serialize(payload interface{}) ([]byte, error) {
	return b.jsonSerializer.Serialize(payload)
}

func (b *jsonBeautifier) Deserialize(payload []byte, target interface{}) error {
	return b.jsonSerializer.Deserialize(payload, target)
}
