package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID        int
	FirstName string
}

var (
	users    []*User
	usersMap []map[uuid.UUID]User
)

// Creates all the users and returns list of createdUsers
func CreateUsers(inputMap map[uuid.UUID]User) ([]map[uuid.UUID]User, error) {
	// Just add the give array of users
	// get the map of users and put them in the list
	usrArr := []User{}
	usrMapArr := []map[uuid.UUID]User{}
	for id, user := range inputMap {
		// Issue here is "user" is actually an extracted value from map
		// so the same address space is being used inside the for of function which will be discaraded after the block
		// its a memory efficiancy that golang adopted.. but to overcome this, store it in a temp for the first iteration
		// so it becomes available during the upcoming iterations as well
		// fmt.Printf("this is id and userObj %v, %v %v %v\n", id, user, User(inputMap[id]), *&user)
		temp := User(inputMap[id])                     // simply user will do the trick
		fmt.Println("Working with temp:", temp, &temp) // !!! Def need the temp
		users = append(users, &temp)
		usrArr = append(usrArr, user)

	}

	usersMap = append(usersMap, inputMap)

	// This is JSON marshallling of data
	aJSON, _ := json.Marshal(users)
	fmt.Println("priting users from actual struct:", string(aJSON))

	return usrMapArr, nil
}

func GetAllUsers() []*User {

	usrArr := users
	fmt.Printf("Complete users: %v \n", usrArr)
	return users
}

func GetUserById(id int) (*User, error) {
	// means find user from the id value of user
	// !IMP since I am return pointers it is valid to return nil, if its object/struct refer below func
	if len(users) == 0 {
		return nil, errors.New("No users are available")
	}

	for index, val := range users {
		fmt.Println("Val and index will be printed here:", index, val)
	}
	return nil, nil
}

// func GetUserByUUId(id uuid.UUID) (User, error)
