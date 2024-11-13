package main

import (
	"fmt"
	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/generated/go/pkg/item"
	"log"
	"time"

	"github.com/google/uuid"

	go_serializer "gitlab.com/pietroski-software-company/devex/golang/serializer/pkg/tools/serializer"
)

func main() {
	serializer := go_serializer.NewProtoSerializer()

	payload := &grpc_item.Item{
		Id:     "any-item",
		ItemId: 100,
		Number: 5_000_000_000,
		SubItem: &grpc_item.SubItem{
			Date:     time.Now().Unix(),
			Amount:   1_000_000_000,
			ItemCode: "code-status",
		},
	}
	bs, err := serializer.Serialize(payload)
	if err != nil {
		log.Fatalf("error serializing message - err: %v", err)
	}

	fmt.Printf("%v\n\n", string(bs))

	testStr := []byte("test-string")
	fmt.Println(testStr)
	fmt.Println(fmt.Sprintf("%s", testStr))

	newUUID, _ := uuid.NewUUID()
	//testStr := []byte("test-string")
	//fmt.Println(testStr)
	fmt.Println(fmt.Sprintf("%s", newUUID))
}
