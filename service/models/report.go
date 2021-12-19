package models

import "github.com/jorgeteixe/autodata/service/config"

type Report struct {
	ID           int           `json:"id"`
	Seccion      *Seccion      `json:"seccion"`
	CentroExamen *CentroExamen `json:"centroExamen"`
	Permiso      *Permiso      `json:"permiso"`
	TipoExamen   *TipoExamen   `json:"tipoExamen"`
	Mes          int           `json:"mes"`
	Anyo         int           `json:"anyo"`
	TotalAptos   int           `json:"totalAptos"`
	Aptos1       int           `json:"aptos1"`
	Aptos2       int           `json:"aptos2"`
	Aptos34      int           `json:"aptos34"`
	Aptos5M      int           `json:"aptos5m"`
	NoAptos      int           `json:"noAptos"`
}

func FetchReports() ([]Report, error) {
	rows, err := config.DB.Query("SELECT `id`, `seccion`, `centro_examen`, `permiso`, `tipo_examen`, `mes`, `anyo`, `total_aptos`, `aptos_1`, `aptos_2`, `aptos_3_4`, `aptos_5m`, `no_aptos` FROM `report`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	reports := make([]Report, 0)

	for rows.Next() {
		report := Report{Seccion: &Seccion{}, CentroExamen: &CentroExamen{}, Permiso: &Permiso{}, TipoExamen: &TipoExamen{}}
		err := rows.Scan(&report.ID, &report.Seccion.ID, &report.CentroExamen.ID, &report.Permiso.ID, &report.TipoExamen.ID,
			&report.Mes, &report.Anyo, &report.TotalAptos, &report.Aptos1, &report.Aptos2, &report.Aptos34, &report.Aptos5M, &report.NoAptos)

		if err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}
	return reports, nil
}

func FetchReport(ID int) (*Report, error) {
	row := config.DB.QueryRow("SELECT `id`, `seccion`, `centro_examen`, `permiso`, `tipo_examen`, `mes`, `anyo`, `total_aptos`, `aptos_1`, `aptos_2`, `aptos_3_4`, `aptos_5m`, `no_aptos` FROM `report` WHERE `id` = ?", ID)

	report := Report{Seccion: &Seccion{}, CentroExamen: &CentroExamen{}, Permiso: &Permiso{}, TipoExamen: &TipoExamen{}}

	err := row.Scan(&report.ID, &report.Seccion.ID, &report.CentroExamen.ID, &report.Permiso.ID, &report.TipoExamen.ID,
		&report.Mes, &report.Anyo, &report.TotalAptos, &report.Aptos1, &report.Aptos2, &report.Aptos34, &report.Aptos5M, &report.NoAptos)

	if err != nil {
		return nil, err
	}

	return &report, nil
}
