package serializer

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"gitlab.com/pietroski-software-company/devex/golang/serializer/models"
)

type GobSerializer struct {
	encoder func(w io.Writer) *gob.Encoder
	decoder func(r io.Reader) *gob.Decoder
}

func NewGobSerializer() *GobSerializer {
	return &GobSerializer{
		encoder: gob.NewEncoder,
		decoder: gob.NewDecoder,
	}
}

func (s *GobSerializer) Serialize(payload interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(payload); err != nil {
		return []byte{}, fmt.Errorf(models.EncodeErrMsg, err)
	}

	return buf.Bytes(), nil
}

func (s *GobSerializer) Deserialize(payload []byte, target interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(payload))
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf(models.DecodeErrMsg, err)
	}

	return nil
}

func (s *GobSerializer) DataRebind(payload interface{}, target interface{}) error {
	bs, err := s.Serialize(payload)
	if err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	if err = s.Deserialize(bs, target); err != nil {
		return fmt.Errorf(models.RebinderErrMsg, err)
	}

	return nil
}
