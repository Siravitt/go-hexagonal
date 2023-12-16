package main

import (
	"net/http"

	"github.com/Siravitt/go-hexagonal/handler"
	"github.com/Siravitt/go-hexagonal/repository"
	"github.com/Siravitt/go-hexagonal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/todoist")
	if err != nil {
		panic(err)
	}

	userRepositoryDB := repository.NewUserRepositoryDB(db)
	userRepositoryMock := repository.NewUserRepositoryMock()

	userServiceDB := service.NewUserService(userRepositoryDB)
	userServiceMock := service.NewUserService(userRepositoryMock)

	userHandlerDB := handler.NewUserHandler(userServiceDB)
	userHandlerMock := handler.NewUserHandler(userServiceMock)

	router := mux.NewRouter()

	router.HandleFunc("/usersDB", userHandlerDB.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/userDB/{userId:[0-9]+}", userHandlerDB.GetUser).Methods(http.MethodGet)

	router.HandleFunc("/usersMock", userHandlerMock.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/userMock/{userId:[0-9]+}", userHandlerMock.GetUser).Methods(http.MethodGet)

	http.ListenAndServe(":8000", router)
}
