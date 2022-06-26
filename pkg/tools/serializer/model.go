package go_serializer

const (
	EncodeErrMsg = "failed to encode payload - err: %v"
	DecodeErrMsg = "failed to decode payload to into target - err: %v"

	RebinderErrMsg = "failed to rebind data - err: %v"
)

type (
	Serializer interface {
		Serialize(payload interface{}) ([]byte, error)
		Deserialize(payload []byte, target interface{}) error

		DataRebind(payload interface{}, target interface{}) error
	}

	Beautifier interface {
		Serializer
		Beautify(payload interface{}, prefix string, indent string) ([]byte, error)
	}
)
