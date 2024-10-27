package main

import (
	"fmt"
	"reflect"
)

const (
	interfaceInterfaceMapType = "map[interface{}]interface{}{}"
)

func main() {
	fmt.Println(reflect.ValueOf(map[interface{}]interface{}{}).Type().String())

	//fmt.Println(reflect.TypeOf(interfaceInterfaceMapType).)

	limit := 4
	//tmpMap := make(map[int8]int8, limit)
	var tmpMap map[int8]int8
	field := reflect.New(reflect.TypeOf(tmpMap)).Elem()
	fmt.Println(field.Type())
	//field.Set(reflect.MakeMapWithSize(field.Type(), limit))
	tmdt := map[interface{}]interface{}{}
	for i := 0; i < limit; i++ {
		//fmt.Printf("%v", field.M)

		//field.SetMapIndex(reflect.ValueOf(int8(i)), reflect.ValueOf(int8(i)))
		//var i8 int8
		//mk := field.MapKeys()[i].Type()
		//mv := field.MapIndex(field.MapKeys()[i]).Type()
		//field.SetMapIndex(
		//	reflect.ValueOf(i).Convert(reflect.TypeOf(mk)),
		//	reflect.ValueOf(i).Convert(reflect.TypeOf(mv)),
		//)
		tmdt[int8(i)] = int8(i)

	}
	mr := reflect.ValueOf(tmdt).MapRange()
	mr.Next()
	fmt.Println(mr.Key().Type())
	fmt.Println(mr.Value().Type())

	fmt.Println(reflect.ValueOf(tmdt).CanConvert(field.Type()))
	fmt.Println(field.CanConvert(reflect.ValueOf(tmdt).Type()))

	fmt.Println(field.Interface())

	fmt.Println(reflect.ValueOf("interfaceInterfaceMapType"))
}
