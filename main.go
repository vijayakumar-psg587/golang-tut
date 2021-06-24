package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/vijayakumar-psg587/golang-tut/controllers"
	"github.com/vijayakumar-psg587/golang-tut/models"
)

func main() {
	val := "Test value"
	firstPointer := &val
	val = "Another val"
	fmt.Printf("Priting pointer and its value, %v, %v\n", firstPointer, *firstPointer)

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
	maps := map[string]string{"key1": "lue1", "key2": "value2"}
	fmt.Printf("Map test, %v \n", maps["key1"])

	// getting uuid
	//uId := uuid.New()

	// calling first service
	userMap := map[uuid.UUID]models.User{}
	userMap[uuid.New()] = models.User{
		ID:        1,
		FirstName: "Vijay",
	}
	userMap[uuid.New()] = models.User{
		ID:        2,
		FirstName: "Kumar",
	}
	// I do not want to deal with error now hence using _ there
	usersCreated, _ := models.CreateUsers(userMap)
	for idC, userC := range usersCreated {
		fmt.Println(idC, userC)
	}

	// Web service
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
