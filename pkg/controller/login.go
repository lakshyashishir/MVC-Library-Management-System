package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func Login(w http.ResponseWriter, request *http.Request) {
	t := views.LoginPage()
	t.Execute(w, nil)
}
