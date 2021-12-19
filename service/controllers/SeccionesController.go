package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jorgeteixe/autodata/service/models"
)

type SeccionesController struct{}

func NewSeccionesController() *SeccionesController {
	return &SeccionesController{}
}

func (ac SeccionesController) GetSeccionesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	secciones, err := models.FetchSecciones()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(secciones)
}

func (pc SeccionesController) GetSeccionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	seccion, err := models.FetchSeccion(ID)

	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(seccion)
}
