package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/vijayakumar-psg587/golang-tut/src/webservice-tut/controllers"
	"github.com/vijayakumar-psg587/golang-tut/src/webservice-tut/models"
)

func main() {
	val := "Test value"
	firstPointer := &val
	val = "Another val"
	fmt.Printf("Priting pointer and its value, %v, %v\n", firstPointer, *firstPointer)

	//sampleTut.SampleImp()

	// Shows sample conversion of int to string with reflect
	// sampleTut.SampleConv()

	// valString := *sampleTut.DateConverter(120345, 20210912)
	// fmt.Println(valString)

	userMap := map[uuid.UUID]models.User{}

	userMap[uuid.New()] = models.User{ID: 1, FirstName: "VV"}
	userMap[uuid.New()] = models.User{ID: 2, FirstName: "KK"}

	// I do not want to deal with error now hence using _ there
	usersCreated, _ := models.CreateUsers(userMap)
	for idC, userC := range usersCreated {
		fmt.Println(idC, userC)
	}
	// sampleTut.GetCurrentTimeString()

	models.GetUserById(2)

	//WebServiceSampleCreator()

}

func WebServiceSampleCreator() {

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
