package beautifier

import (
	"encoding/json"
	"fmt"
	"gitlab.com/pietroski-software-company/devex/golang/serializer"
	"gitlab.com/pietroski-software-company/devex/golang/serializer/models"
)

type (
	jsonBeautifier struct {
		jsonSerializer models.Serializer
	}
)

func NewJsonBeautifier() models.Beautifier {
	return &jsonBeautifier{
		jsonSerializer: serializer.NewJsonSerializer(),
	}
}

func (b *jsonBeautifier) Beautify(payload interface{}, prefix string, indent string) ([]byte, error) {
	bs, err := json.MarshalIndent(payload, prefix, indent)
	if err != nil {
		return []byte{}, fmt.Errorf(models.EncodeErrMsg, err)
	}

	return bs, err
}

func (b *jsonBeautifier) Serialize(payload interface{}) ([]byte, error) {
	return b.jsonSerializer.Serialize(payload)
}

func (b *jsonBeautifier) Deserialize(payload []byte, target interface{}) error {
	return b.jsonSerializer.Deserialize(payload, target)
}

func (b *jsonBeautifier) DataRebind(payload interface{}, target interface{}) error {
	bs, err := b.Serialize(payload)
	if err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	if err = b.Deserialize(bs, target); err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	return nil
}
