package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func Signup(w http.ResponseWriter, request *http.Request) {
	t := views.SignupPage()
	t.Execute(w, nil)
}
