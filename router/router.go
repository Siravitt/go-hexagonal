package router

import (
	"net/http"

	"github.com/Siravitt/go-hexagonal/handler"
	"github.com/gorilla/mux"
)

func InitRouter(userHdr handler.UserHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/usersDB", userHdr.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/userDB/{userId:[0-9]+}", userHdr.GetUser).Methods(http.MethodGet)

	router.HandleFunc("/usersMock", userHdr.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/userMock/{userId:[0-9]+}", userHdr.GetUser).Methods(http.MethodGet)
	// router.HandleFunc("")
}
