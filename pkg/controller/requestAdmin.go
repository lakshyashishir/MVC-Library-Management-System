package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func RequestAdmin(w http.ResponseWriter, request *http.Request) {
	t := views.RequestAdminPage()
	t.Execute(w, nil)
}
