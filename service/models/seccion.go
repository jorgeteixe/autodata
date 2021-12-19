package models

import "github.com/jorgeteixe/autodata/service/config"

type Seccion struct {
	ID          int          `json:"id"`
	Codigo      string       `json:"codigo,omitempty"`
	Autoescuela *Autoescuela `json:"autoescuela,omitempty"`
}

func FetchSecciones() ([]Seccion, error) {
	rows, err := config.DB.Query("SELECT `id`, `codigo`, `autoescuela` FROM `seccion`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	secciones := make([]Seccion, 0)

	for rows.Next() {
		seccion := Seccion{Autoescuela: &Autoescuela{}}

		err := rows.Scan(&seccion.ID, &seccion.Codigo, &seccion.Autoescuela.ID)

		if err != nil {
			return nil, err
		}

		secciones = append(secciones, seccion)
	}
	return secciones, nil
}

func FetchSeccion(ID int) (*Seccion, error) {
	row := config.DB.QueryRow("SELECT `id`, `codigo`, `autoescuela` FROM `seccion` WHERE `id` = ?", ID)

	seccion := Seccion{Autoescuela: &Autoescuela{}}

	err := row.Scan(&seccion.ID, &seccion.Codigo, &seccion.Autoescuela.ID)

	if err != nil {
		return nil, err
	}

	return &seccion, nil
}
