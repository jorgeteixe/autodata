package models

import "github.com/jorgeteixe/autodata/service/config"

type Autoescuela struct {
	ID     int    `json:"id"`
	Codigo string `json:"codigo,omitempty"`
	Nombre string `json:"nombre,omitempty"`
}

func FetchAutoescuelas() ([]Autoescuela, error) {
	rows, err := config.DB.Query("SELECT `id`, `codigo`, `nombre` FROM `autoescuela`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	autoescuelas := make([]Autoescuela, 0)

	for rows.Next() {
		auto := Autoescuela{}
		err := rows.Scan(&auto.ID, &auto.Codigo, &auto.Nombre)

		if err != nil {
			return nil, err
		}

		autoescuelas = append(autoescuelas, auto)
	}
	return autoescuelas, nil
}

func FetchAutoescuela(ID int) (*Autoescuela, error) {
	row := config.DB.QueryRow("SELECT `id`, `codigo`, `nombre` FROM `autoescuela` WHERE `id` = ?", ID)

	auto := Autoescuela{}

	err := row.Scan(&auto.ID, &auto.Codigo, &auto.Nombre)

	if err != nil {
		return nil, err
	}

	return &auto, nil
}
