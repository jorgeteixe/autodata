package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jorgeteixe/autodata/service/models"
)

type TiposExamenController struct{}

func NewTiposExamenController() *TiposExamenController {
	return &TiposExamenController{}
}

func (pc TiposExamenController) GetTiposExamenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tipos, err := models.FetchTiposExamen()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(tipos)
}

func (pc TiposExamenController) GetTipoExamenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	tipo, err := models.FetchTipoExamen(ID)

	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(tipo)
}
