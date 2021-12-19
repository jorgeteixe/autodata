package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jorgeteixe/autodata/service/models"
)

type PermisosController struct{}

func NewPermisosController() *PermisosController {
	return &PermisosController{}
}

func (pc PermisosController) GetPermisosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	permisos, err := models.FetchPermisos()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(permisos)
}

func (pc PermisosController) GetPermisoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	permiso, err := models.FetchPermiso(ID)

	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(permiso)
}
