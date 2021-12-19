package controllers

import (
	"fmt"
	"net/http"
)

type GeneralController struct{}

func NewGeneralController() *GeneralController {
	return &GeneralController{}
}

func (gc GeneralController) ReadyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok\n")
}
