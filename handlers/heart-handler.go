package handlers

import (
	"fmt"
	"net/http"
)

func HearthHandlerfunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w,"Método não suportado", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Serviço está saudável!")
}