package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jorgeteixe/autodata/service/models"
)

type ProvinciasController struct{}

func NewProvinciasController() *ProvinciasController {
	return &ProvinciasController{}
}

func (pc ProvinciasController) GetProvinciasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	provincias, err := models.FetchProvincias()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(provincias)
}

func (pc ProvinciasController) GetProvinciaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	provincia, err := models.FetchProvincia(ID)

	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(provincia)
}
