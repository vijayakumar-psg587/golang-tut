package controllers

import "net/http"

func RegisterController() {
	userC := newUserController()

	http.Handle("/users", *userC)
	http.Handle("/users/", *userC)

}
