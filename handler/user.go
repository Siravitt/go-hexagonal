package handler

import "net/http"

type UserHandler interface {
	GetUsers(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
}
