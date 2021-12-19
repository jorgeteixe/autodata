package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jorgeteixe/autodata/service/models"
)

type AutoescuelasController struct{}

func NewAutoescuelasController() *AutoescuelasController {
	return &AutoescuelasController{}
}

func (ac AutoescuelasController) GetAutoescuelasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	autoescuelas, err := models.FetchAutoescuelas()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(autoescuelas)
}

func (pc AutoescuelasController) GetAutoescuelaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	auto, err := models.FetchAutoescuela(ID)

	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(auto)
}
