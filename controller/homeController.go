package controller

import "net/http"

type HomeController struct{}

func (c *HomeController) Home(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte("Hello world"))
}
