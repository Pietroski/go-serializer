package main

import (
	"fmt"
	item_models "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/models/item"
	"math"
	"time"
)
import "unsafe"

func main() {
	var x struct {
		a int64
		b bool
		c string
	}
	const M, N = unsafe.Sizeof(x.c), unsafe.Sizeof(x)
	fmt.Println(M, N) // 16 32

	fmt.Println(unsafe.Sizeof(x.a), unsafe.Sizeof(x.b), unsafe.Sizeof(x.c))
	fmt.Println(unsafe.Sizeof(x)) // 16 32

	y := struct {
		a int64
		b bool
		c string
	}{
		a: math.MinInt64,
		b: true,
		c: "abcdefghijklmnopqrstuvwxyz",
	}
	//x.a = math.MaxInt64
	//x.b = true
	//x.c = "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(unsafe.Sizeof(y.a), unsafe.Sizeof(y.b), unsafe.Sizeof(y.c))
	fmt.Println(unsafe.Sizeof(y))
	fmt.Println(unsafe.Sizeof(&y.a), unsafe.Sizeof(y.b), unsafe.Sizeof(y.c))

	//fmt.Println(unsafe.Alignof(x.a)) // 8
	//fmt.Println(unsafe.Alignof(x.b)) // 1
	//fmt.Println(unsafe.Alignof(x.c)) // 8
	//
	//fmt.Println(unsafe.Offsetof(x.a)) // 0
	//fmt.Println(unsafe.Offsetof(x.b)) // 8
	//fmt.Println(unsafe.Offsetof(x.c)) // 16

	msg := &item_models.Item{
		Id:     "any-item",
		ItemId: 100,
		Number: 5_000_000_000,
		SubItem: &item_models.SubItem{
			Date:     time.Now().Unix(),
			Amount:   1_000_000_000,
			ItemCode: "code-status",
		},
	}
	fmt.Println(unsafe.Sizeof(*msg))
	fmt.Println(
		unsafe.Sizeof(item_models.Item{}),
		unsafe.Sizeof(item_models.Item{}.Id),
		unsafe.Sizeof(item_models.Item{}.ItemId),
		unsafe.Sizeof(item_models.Item{}.Number),
		unsafe.Sizeof(item_models.Item{}.SubItem),
		unsafe.Sizeof(item_models.SubItem{}),
	)

	var bs [8]byte
	//i64 := int64(100)
	//*(*int64)(unsafe.Pointer(&bs)) = i64
	//fmt.Println(*(*int64)(unsafe.Pointer(&bs[0])))

	*(*uint64)(unsafe.Pointer(&bs)) = math.MaxUint64
	fmt.Println(*(*uint64)(unsafe.Pointer(&bs[0])))
}
