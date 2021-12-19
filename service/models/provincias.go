package models

import (
	"github.com/jorgeteixe/autodata/service/config"
)

type Provincia struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre,omitempty"`
}

func FetchProvincias() ([]Provincia, error) {
	rows, err := config.DB.Query("SELECT `id`, `nombre` FROM `provincia`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	provincias := make([]Provincia, 0)

	for rows.Next() {
		prov := Provincia{}
		err := rows.Scan(&prov.ID, &prov.Nombre)

		if err != nil {
			return nil, err
		}

		provincias = append(provincias, prov)
	}
	return provincias, nil
}

func FetchProvincia(ID int) (*Provincia, error) {
	row := config.DB.QueryRow("SELECT `id`, `nombre` FROM `provincia` WHERE `id` = ?", ID)

	prov := Provincia{}

	err := row.Scan(&prov.ID, &prov.Nombre)

	if err != nil {
		return nil, err
	}

	return &prov, nil
}
