package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	t := views.InternalServerErrorPage()
	t.Execute(w, nil)
}
