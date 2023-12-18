package handler

import (
	"fmt"
	"net/http"

	"github.com/Siravitt/go-hexagonal/errs"
	"github.com/Siravitt/go-hexagonal/service"
)

type Handler interface {
	GetUsers(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	AddTodo(http.ResponseWriter, *http.Request)
	GetTodos(http.ResponseWriter, *http.Request)
}

type handler struct {
	srv service.Service
}

func NewHandler(srv service.Service) handler {
	return handler{srv: srv}
}

func handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case errs.AppError:
		w.WriteHeader(e.Code)
		fmt.Fprintln(w, err)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, e)
	}
}
