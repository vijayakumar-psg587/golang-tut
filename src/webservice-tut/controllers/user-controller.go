package controllers

import (
	"net/http"
	"regexp"
)

type UserController struct {
	userPath *regexp.Regexp
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Custom", "New")
	w.Write([]byte("User controller"))
}

func newUserController() *UserController {
	return &UserController{
		userPath: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
