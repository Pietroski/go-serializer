package go_serializer

import (
	"bytes"
	"encoding/gob"
	"io"

	error_builder "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/builder/errors"
)

type gobSerializer struct {
	encoder func(w io.Writer) *gob.Encoder
	decoder func(r io.Reader) *gob.Decoder
}

func NewGobSerializer() Serializer {
	return &gobSerializer{
		encoder: gob.NewEncoder,
		decoder: gob.NewDecoder,
	}
}

func (s *gobSerializer) Serialize(payload interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(payload); err != nil {
		return []byte{}, error_builder.Err(EncodeErrMsg, err)
	}

	return buf.Bytes(), nil
}

func (s *gobSerializer) Deserialize(payload []byte, target interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(payload))
	if err := decoder.Decode(target); err != nil {
		return error_builder.Err(DecodeErrMsg, err)
	}

	return nil
}

func (s *gobSerializer) DataRebind(payload interface{}, target interface{}) error {
	bs, err := s.Serialize(payload)
	if err != nil {
		return error_builder.Err(RebinderErrMsg, err)
	}

	if err = s.Deserialize(bs, target); err != nil {
		return error_builder.Err(RebinderErrMsg, err)
	}

	return nil
}
