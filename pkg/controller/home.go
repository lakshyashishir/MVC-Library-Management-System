package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func Home(w http.ResponseWriter, request *http.Request) {
	t := views.StartPage()
	t.Execute(w, nil)
}
