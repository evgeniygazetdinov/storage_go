package main

import (
	"fmt"
	"reflect"
)

type myType struct{}

func (s myType) String() string {
	return ""
}

func main() {
	fmt.Println(isStringer("123"))
	fmt.Println(isStringer(1))
	d := myType{}
	fmt.Println(isStringer(d))
}

func isStringer(obj interface{}) bool {
	_, ok := obj.(fmt.Stringer)
	return ok
}

func isStringer2(obj interface{}, objForCOmpare interface{}) bool {
	stringerType := reflect.TypeOf((fmt.Stringer)(nil)).Elem()
	typeoF := reflect.TypeOf((*obj)(nil))
	fmt.Println(typeoF.Implements(stringerType))
	return true
}
