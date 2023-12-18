package router

import (
	"fmt"
	"net/http"

	"github.com/Siravitt/go-hexagonal/handler"
	"github.com/Siravitt/go-hexagonal/logs"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func InitRouter(userHdr handler.UserHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/usersDB", userHdr.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/userDB/{userId:[0-9]+}", userHdr.GetUser).Methods(http.MethodGet)

	router.HandleFunc("/usersMock", userHdr.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/userMock/{userId:[0-9]+}", userHdr.GetUser).Methods(http.MethodGet)

	logs.Info("User service started at port " + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}
