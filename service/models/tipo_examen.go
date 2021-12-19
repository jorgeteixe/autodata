package models

import "github.com/jorgeteixe/autodata/service/config"

type TipoExamen struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre,omitempty"`
}

func FetchTiposExamen() ([]TipoExamen, error) {
	rows, err := config.DB.Query("SELECT `id`, `nombre` FROM `tipo_examen`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tipos := make([]TipoExamen, 0)

	for rows.Next() {
		tipo := TipoExamen{}
		err := rows.Scan(&tipo.ID, &tipo.Nombre)

		if err != nil {
			return nil, err
		}

		tipos = append(tipos, tipo)
	}
	return tipos, nil
}

func FetchTipoExamen(ID int) (*TipoExamen, error) {
	row := config.DB.QueryRow("SELECT `id`, `nombre` FROM `tipo_examen` WHERE `id` = ?", ID)

	tipo := TipoExamen{}

	err := row.Scan(&tipo.ID, &tipo.Nombre)

	if err != nil {
		return nil, err
	}

	return &tipo, nil
}
