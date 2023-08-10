package api

import (
	"mvc/pkg/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home).Methods("GET")
	r.HandleFunc("/admin", controller.Admin).Methods("GET")
	r.HandleFunc("/user", controller.User).Methods("GET")
	r.HandleFunc("/requests", controller.Requests).Methods("GET")
	r.HandleFunc("/myRequests", controller.UserRequests).Methods("GET")
	r.HandleFunc("/getbook", controller.UserViewBook).Methods("GET")
	r.HandleFunc("/mybooks", controller.UserBooks).Methods("GET")
	r.HandleFunc("/reqAdmin", controller.RequestAdmin).Methods("GET")
	r.HandleFunc("/adminRequests", controller.AdminRequests).Methods("GET")
	r.HandleFunc("/issuedBooks", controller.IssuedBooks).Methods("GET")
	r.HandleFunc("/addBook", controller.AddBook).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("GET")
	r.HandleFunc("/signup", controller.Signup).Methods("GET")

	r.HandleFunc("/login", controller.LoginPost).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("POST")
	r.HandleFunc("/signup", controller.Signup).Methods("POST")
	r.HandleFunc("/addBook", controller.AddBook).Methods("POST")
	r.HandleFunc("/user/request", controller.UserRequests).Methods("POST")
	r.HandleFunc("/user/return", controller.UserReturnBook).Methods("POST")
	r.HandleFunc("/user/removeRequest", controller.UserRemoveRequestBook).Methods("POST")
	r.HandleFunc("/requests/approve", controller.ApproveBookRequest).Methods("POST")
	r.HandleFunc("/requests/reject", controller.RejectBookRequest).Methods("POST")
	r.HandleFunc("/adminRequests/approve", controller.ApproveAdmin).Methods("POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8000", r)
}
