package handlers

import (
	"library/models"
	"library/services"
	"net/http"
)

type UserHandler struct {
	Service *services.UserService
}


func () Register() {

//pana and joshua 

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

	
}