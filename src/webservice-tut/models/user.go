package models

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID        int
	FirstName string
}

var (
	users []*User
)

func CreateUsers(usersMap map[uuid.UUID]User) ([]map[uuid.UUID]User, error) {
	// Just add the give array of users
	// get the map of users and put them in the list
	usrArr := []User{}
	usrMapArr := []map[uuid.UUID]User{}
	for id, user := range usersMap {
		fmt.Printf("this is id and userObj %v, %v \n", id, user)
		users = append(users, &user)
		usrArr = append(usrArr, user)

	}

	usrMapArr = append(usrMapArr, usersMap)

	fmt.Println(users)
	return usrMapArr, nil
}
