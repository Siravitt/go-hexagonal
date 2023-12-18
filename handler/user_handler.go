package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Siravitt/go-hexagonal/service"
	"github.com/gorilla/mux"
)

type userHandler struct {
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userSrv.GetAllUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("content/type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])
	user, err := h.userSrv.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("content/type", "application/json")
	json.NewEncoder(w).Encode(user)
}
