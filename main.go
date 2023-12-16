package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Siravitt/go-hexagonal/handler"
	"github.com/Siravitt/go-hexagonal/repository"
	"github.com/Siravitt/go-hexagonal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
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

	log.Printf("User service started at port %v", viper.GetInt("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
