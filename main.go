package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Siravitt/go-hexagonal/handler"
	"github.com/Siravitt/go-hexagonal/repository"
	"github.com/Siravitt/go-hexagonal/router"
	"github.com/Siravitt/go-hexagonal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	initTimeZone()

	db := initDatabase()

	// Create Repository
	repository := repository.NewRepository(db)
	// repository := repository.NewUserRepositoryMock(db)

	// Create Service
	service := service.NewService(repository)

	// Create Handler
	handler := handler.NewHandler(service)

	// Create Router
	router.InitRouter(handler)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDatabase() *sqlx.DB {
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

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
