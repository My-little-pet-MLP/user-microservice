package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	service "github.com/my-little-pet/user-microservice/services"
)

func GetByIdUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w,"Método não suportado", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/users/id=")
	user,err := service.GetByIdUser(id);

	if err != nil {
		http.Error(w, "Erro ao buscar o usuário: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
}