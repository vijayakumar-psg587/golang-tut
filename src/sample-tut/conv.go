package sampleTut

import (
	"fmt"
	"reflect"
	"strconv"
)

func SampleConv() {
	num1 := 10
	fmt.Println(num1, reflect.TypeOf(num1))
	str := strconv.Itoa(num1)
	fmt.Println(str, reflect.TypeOf(str))
}
