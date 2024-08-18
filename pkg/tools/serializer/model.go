package go_serializer

const (
	WrongPayloadTypeErrMsg = "wrong payload type"
	WrongTargetTypeErrMsg  = "wrong target type"
	EncodeErrMsg           = "failed to encode payload - err: %v"
	DecodeErrMsg           = "failed to decode payload to into target - err: %v"

	RebinderErrMsg = "failed to rebind data - err: %v"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate -o ../../../fakes/fake_serializer.go . Serializer
//counterfeiter:generate -o ../../../fakes/fake_beautifier.go . Beautifier
//go:generate mockgen -package mocks -destination ../../../mocks/mocked_serializer.go . Serializer
//go:generate mockgen -package mocks -destination ../../../mocks/mocked_beautifier.go . Beautifier

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
