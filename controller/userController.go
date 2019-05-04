package controller

import (
	"encoding/json"
	"github.com/nataliobp/myGoShoppingList/core"
	"github.com/nataliobp/myGoShoppingList/identityAccess"
	"net/http"
	"regexp"
	"strconv"
)

type UserController struct {
	Container *core.Container
}

func (u *UserController) ManageUser(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		u.createUser(request, writer)
	} else if request.Method == "GET" {
		u.getUserById(request, writer)
	}
}

func (u *UserController) createUser(request *http.Request, writer http.ResponseWriter) {
	decoder := json.NewDecoder(request.Body)
	signUpVM := SignUpVM{}
	e := decoder.Decode(&signUpVM)

	if e != nil {
		http.Error(writer, e.Error(), 500)
		return
	}

	singUpService := u.Container.Get("signUpService").(*identityAccess.SignUpService)
	id := singUpService.SignUp(signUpVM.Name, signUpVM.Email, signUpVM.Password)

	writer.WriteHeader(http.StatusCreated)
	_, _ = writer.Write([]byte("/users/" + strconv.FormatInt(id, 10)))
}

func (u *UserController) getUserById(request *http.Request, writer http.ResponseWriter) {
	pattern, _ := regexp.Compile(`/users/(\d+)`)
	matches := pattern.FindStringSubmatch(request.URL.Path)
	id, _ := strconv.ParseInt(matches[1], 10, 64)

	findUserService := u.Container.Get("findUserService").(*identityAccess.FindUserService)
	userById, err := findUserService.FindById(id)

	if err != nil {
		http.NotFound(writer, request)
		return
	}
	userVM := UserVM{Id: userById.Id, Name: userById.Name, Email: userById.Email}
	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	e := encoder.Encode(userVM)

	if e != nil {
		http.Error(writer, e.Error(), 500)
	}
}

func (u *UserController) GetUserByEmail(writer http.ResponseWriter, request *http.Request) {
	pattern, _ := regexp.Compile(`/users/email/([\w.@]+)`)
	matches := pattern.FindStringSubmatch(request.URL.Path)
	email := matches[1]

	findUserService := u.Container.Get("findUserService").(*identityAccess.FindUserService)
	userByEmail, err := findUserService.FindByEmail(email)

	if err != nil {
		http.NotFound(writer, request)
		return
	}

	userVM := UserVM{Id: userByEmail.Id, Email: userByEmail.Email}
	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	e := encoder.Encode(userVM)

	if e != nil {
		http.Error(writer, e.Error(), 500)
	}
}
