package handlers

import (
	"encoding/json"
	"library/models"
	"library/services"
	"net/http"
)

type UserHandler struct {
	Service *services.UserService
}


func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	//the collection of request details
	var signUp models.User
	err := json.NewDecoder(r.Body).Decode(&signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//calling the service layer
	err = h.Service.RegisterUser(&signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	//response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signUp)


}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}