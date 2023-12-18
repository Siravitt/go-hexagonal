package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// type userHandler struct {
// 	userSrv service.UserService
// }

// func NewUserHandler(userSrv service.UserService) Handler {
// 	return userHandler{userSrv: userSrv}
// }

func (h handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.srv.GetAllUser()
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("content/type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])
	user, err := h.srv.GetUser(userId)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("content/type", "application/json")
	json.NewEncoder(w).Encode(user)
}
