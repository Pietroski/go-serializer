package main

import "fmt"

//func main() {
//	bs := make([]byte, 4)
//	fmt.Println(len(bs), cap(bs))
//	bbs := bs[: len(bs) : cap(bs)<<1]
//	fmt.Println(len(bbs), cap(bbs))
//}

//func main() {
//	bs := []byte{1, 2, 3, 4}
//	nbs := make([]byte, cap(bs)<<1)
//	fmt.Println(len(nbs), cap(nbs), nbs)
//	// nbs = bs[:]
//	copy(nbs, bs)
//	fmt.Println(len(nbs), cap(nbs), nbs)
//	bs = nbs
//	fmt.Println(len(bs), cap(bs), bs)
//}

func main() {
	fmt.Println(1 << 16)
}
