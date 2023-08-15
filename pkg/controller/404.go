package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	t := views.NotFoundPage()
	t.Execute(w, nil)
}
