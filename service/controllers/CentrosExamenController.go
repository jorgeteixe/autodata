package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jorgeteixe/autodata/service/models"
)

type CentrosExamenController struct{}

func NewCentrosExamenController() *CentrosExamenController {
	return &CentrosExamenController{}
}

func (ac CentrosExamenController) GetCentrosExamenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	centros, err := models.FetchCentrosExamen()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(centros)
}

func (pc CentrosExamenController) GetCentroExamenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	centro, err := models.FetchCentroExamen(ID)

	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(centro)
}
