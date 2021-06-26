package sampleTut

import "fmt"

func SampleImp() {
	fmt.Printf("This is from sampleTut \n")
	// arraysUser
	// var arr [5]string
	arr := [5]string{}
	arr[1] = "rray 1 string"
	arr[0] = "rray 0 string"
	arr[4] = "This is the 4th value"
	slc := arr[2:5]
	fmt.Printf("Array and Slice is %v, %v \n", arr, slc)

	// slice is not a fixed size entity
	slice := []int{}
	slice = append(slice, 33)
	fmt.Printf("Slice which is not a fixed size entity:%v \n", len(slice))

	// Maps
	// Another way of creating maps -  make(map[key]val) -  Not needed an empty constructor
	maps := map[string]string{"key1": "lue1", "key2": "value2"}
	fmt.Printf("Map test, %v \n", maps["key1"])
}
