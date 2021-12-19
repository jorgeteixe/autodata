package models

import "github.com/jorgeteixe/autodata/service/config"

type CentroExamen struct {
	ID        int        `json:"id"`
	Nombre    string     `json:"nombre,omitempty"`
	Provincia *Provincia `json:"provincia,omitempty"`
}

func FetchCentrosExamen() ([]CentroExamen, error) {
	rows, err := config.DB.Query("SELECT `id`, `nombre`, `provincia` FROM `centro_examen`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	centros := make([]CentroExamen, 0)

	for rows.Next() {
		centro := CentroExamen{Provincia: &Provincia{}}

		err := rows.Scan(&centro.ID, &centro.Nombre, &centro.Provincia.ID)

		if err != nil {
			return nil, err
		}

		centros = append(centros, centro)
	}
	return centros, nil
}

func FetchCentroExamen(ID int) (*CentroExamen, error) {
	row := config.DB.QueryRow("SELECT `id`, `nombre`, `provincia` FROM `centro_examen` WHERE `id` = ?", ID)

	centro := CentroExamen{Provincia: &Provincia{}}

	err := row.Scan(&centro.ID, &centro.Nombre, &centro.Provincia.ID)

	if err != nil {
		return nil, err
	}

	return &centro, nil
}
