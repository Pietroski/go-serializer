package file_writer

import (
	"encoding/binary"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	grpc_item "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/generated/go/pkg/item"
	item_models "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/models/item"
	go_serializer "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/serializer"
)

type Fields struct {
	Int int64
	//Str string
}

var (
	msg = &item_models.Item{
		Id:     "any-item",
		ItemId: 100,
		Number: 5_000_000_000,
		SubItem: &item_models.SubItem{
			Date:     time.Now().Unix(),
			Amount:   1_000_000_000,
			ItemCode: "code-status",
		},
	}

	protoMsg = &grpc_item.Item{
		Id:     "any-item",
		ItemId: 100,
		Number: 5_000_000_000,
		SubItem: &grpc_item.SubItem{
			Date:     time.Now().Unix(),
			Amount:   1_000_000_000,
			ItemCode: "code-status",
		},
	}

	fields = &Fields{
		Int: 5_000_000_000,
		//Str: "any-string",
	}
)

func BenchmarkBytesWrite(b *testing.B) {
	serializer := go_serializer.NewProtoSerializer()

	bs, err := serializer.Serialize(protoMsg)
	require.NoError(b, err)

	file, err := os.Create("data-bytes") // "playground/file-writer/docs/data-bytes"
	require.NoError(b, err)

	_, err = file.Write(bs)
	require.NoError(b, err)

	err = file.Close()
	require.NoError(b, err)

	file, err = os.Open("data-bytes")
	require.NoError(b, err)

	debuffer := []byte{}
	_, err = file.Read(debuffer)
	require.NoError(b, err)

	var target grpc_item.Item
	err = serializer.Deserialize(debuffer, &target)
	require.NoError(b, err)

	b.Log(target.Id)

	err = file.Close()
	require.NoError(b, err)
}

func BenchmarkBinaryWrite(b *testing.B) {
	file, err := os.Create("data-binary") // "playground/file-writer/docs/data-bytes"
	require.NoError(b, err)

	err = binary.Write(file, binary.BigEndian, fields)
	require.NoError(b, err)

	err = file.Close()
	require.NoError(b, err)

	file, err = os.Open("data-binary")
	require.NoError(b, err)

	var target Fields
	err = binary.Read(file, binary.BigEndian, &target)
	require.NoError(b, err)

	err = file.Close()
	require.NoError(b, err)

	b.Log(target.Int)
}

func BenchmarkBytesBinaryWrite(b *testing.B) {
	serializer := go_serializer.NewProtoSerializer()

	bs, err := serializer.Serialize(protoMsg)
	require.NoError(b, err)

	file, err := os.Create("data-binary") // "playground/file-writer/docs/data-bytes"
	require.NoError(b, err)

	err = binary.Write(file, binary.BigEndian, bs)
	require.NoError(b, err)

	err = file.Close()
	require.NoError(b, err)

	file, err = os.Open("data-binary")
	require.NoError(b, err)

	preBuffer := make([]byte, len(bs))
	err = binary.Read(file, binary.BigEndian, &preBuffer)
	require.NoError(b, err)

	err = file.Close()
	require.NoError(b, err)

	var target grpc_item.Item
	err = serializer.Deserialize(preBuffer, &target)
	require.NoError(b, err)

	//b.Log(target.Id)
}

func BenchmarkBytesNormalFileWrite(b *testing.B) {
	serializer := go_serializer.NewProtoSerializer()

	bs, err := serializer.Serialize(protoMsg)
	require.NoError(b, err)

	file, err := os.Create("data-normal-writer") // "playground/file-writer/docs/data-bytes"
	require.NoError(b, err)

	_, err = file.Write(bs)
	require.NoError(b, err)

	//b.Log(n, string(bs))

	err = file.Close()
	require.NoError(b, err)

	file, err = os.Open("data-normal-writer")
	require.NoError(b, err)

	preBuffer := make([]byte, len(bs))
	_, err = file.Read(preBuffer)
	require.NoError(b, err)

	//b.Log(n, string(preBuffer))

	err = file.Close()
	require.NoError(b, err)

	var target grpc_item.Item
	err = serializer.Deserialize(preBuffer, &target)
	require.NoError(b, err)

	//b.Log(target.Id)
}

func BenchmarkJsonBytesNormalFileWrite(b *testing.B) {
	serializer := go_serializer.NewJsonSerializer()

	bs, err := serializer.Serialize(msg)
	require.NoError(b, err)

	file, err := os.Create("data-normal-writer") // "playground/file-writer/docs/data-bytes"
	require.NoError(b, err)

	_, err = file.Write(bs)
	require.NoError(b, err)

	//b.Log(n, string(bs))

	err = file.Close()
	require.NoError(b, err)

	file, err = os.Open("data-normal-writer")
	require.NoError(b, err)

	preBuffer := make([]byte, len(bs))
	_, err = file.Read(preBuffer)
	require.NoError(b, err)

	//b.Log(n, string(preBuffer))

	err = file.Close()
	require.NoError(b, err)

	var target item_models.Item
	err = serializer.Deserialize(preBuffer, &target)
	require.NoError(b, err)

	//b.Log(target.Id)
}

func BenchmarkJsonBinaryFileWrite(b *testing.B) {
	serializer := go_serializer.NewJsonSerializer()

	bs, err := serializer.Serialize(msg)
	require.NoError(b, err)

	file, err := os.Create("data-normal-writer") // "playground/file-writer/docs/data-bytes"
	require.NoError(b, err)

	err = binary.Write(file, binary.BigEndian, bs)
	require.NoError(b, err)

	//_, err = file.Write(bs)
	//require.NoError(b, err)

	//b.Log(n, string(bs))

	err = file.Close()
	require.NoError(b, err)

	file, err = os.Open("data-normal-writer")
	require.NoError(b, err)

	preBuffer := make([]byte, len(bs))
	err = binary.Read(file, binary.BigEndian, &preBuffer)
	require.NoError(b, err)

	//_, err = file.Read(preBuffer)
	//require.NoError(b, err)

	//b.Log(n, string(preBuffer))

	err = file.Close()
	require.NoError(b, err)

	var target item_models.Item
	err = serializer.Deserialize(preBuffer, &target)
	require.NoError(b, err)

	//b.Log(target.Id)
}
