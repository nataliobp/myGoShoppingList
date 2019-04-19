package controller

import (
	"encoding/json"
	"github.com/nataliobp/myGoShoppingList/core"
	"github.com/nataliobp/myGoShoppingList/identityAccess"
	"net/http"
	"regexp"
)

type UserController struct {
	Container *core.Container
}

func (u *UserController) SignUp(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	userVM := UserVM{}
	e := decoder.Decode(&userVM)

	if e != nil {
		http.Error(writer, e.Error(), 500)
		return
	}

	singUpService := u.Container.Get("signUpService").(*identityAccess.SignUpService)
	singUpService.SignUp(userVM.Email, userVM.Password)

	writer.WriteHeader(http.StatusCreated)
}

func (u *UserController) GetUserByEmail(writer http.ResponseWriter, request *http.Request) {
	pattern, _ := regexp.Compile(`/users/email/([\w\.@]+)`)
	matches := pattern.FindStringSubmatch(request.URL.Path)
	email := matches[1]

	findUserService := u.Container.Get("findUserService").(*identityAccess.FindUserService)
	userByEmail := findUserService.FindByEmail(email)

	if userByEmail == nil {
		http.NotFound(writer, request)
		return
	}

	userVM := UserVM{Email: userByEmail.Email, Password: userByEmail.Password}

	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	e := encoder.Encode(userVM)

	if e != nil {
		http.Error(writer, e.Error(), 500)
	}
}
