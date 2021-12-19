package models

import "github.com/jorgeteixe/autodata/service/config"

type Permiso struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre,omitempty"`
}

func FetchPermisos() ([]Permiso, error) {
	rows, err := config.DB.Query("SELECT `id`, `nombre` FROM `permiso`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	permisos := make([]Permiso, 0)

	for rows.Next() {
		perm := Permiso{}
		err := rows.Scan(&perm.ID, &perm.Nombre)

		if err != nil {
			return nil, err
		}

		permisos = append(permisos, perm)
	}
	return permisos, nil
}

func FetchPermiso(ID int) (*Permiso, error) {
	row := config.DB.QueryRow("SELECT `id`, `nombre` FROM `permiso` WHERE `id` = ?", ID)

	perm := Permiso{}

	err := row.Scan(&perm.ID, &perm.Nombre)

	if err != nil {
		return nil, err
	}

	return &perm, nil
}
