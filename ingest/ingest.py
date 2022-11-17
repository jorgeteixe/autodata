import os
import pandas as pd

import db

DATA_DIR = './data'


def main():
    print(f'starting ingest')
    for filename in os.listdir(DATA_DIR):
        if db.get_read_file(filename) is None:
            print(f'started to ingest file: {filename}')
            import_to_db(os.path.join(DATA_DIR, filename))
            db.set_read_file(filename)
        else:
            print(f'skipped file: {filename}')

    db.close_cnx()
    print(f'finished ingest')


def import_to_db(file):
    df = get_dataframe(file)
    provincias = sorted(df['DESC_PROVINCIA'].unique())
    db.update_provincias(provincias)
    centros = df.groupby(by='CENTRO_EXAMEN', as_index=False)['DESC_PROVINCIA'].apply(set)
    db.update_centros(centros)
    autoescuelas = df.groupby(by='CODIGO_AUTOESCUELA', as_index=False)['NOMBRE_AUTOESCUELA'].apply(set)
    #autoescuelas['CODIGO_AUTOESCUELA'] = autoescuelas['CODIGO_AUTOESCUELA'].str.replace(" ", "")
    db.update_autoescuelas(autoescuelas)
    secciones = df.groupby(by=['CODIGO_AUTOESCUELA', 'CODIGO_SECCION'], as_index=False)['CODIGO_SECCION'].apply(set)
    db.update_secciones(secciones)
    tipos_examen = df['TIPO_EXAMEN'].unique()
    db.update_tipos_examen(tipos_examen)
    permisos = df['NOMBRE_PERMISO'].unique()
    db.update_permisos(permisos)
    fechas = df.groupby(by='MES', as_index=False)['ANYO'].apply(set)
    db.update_fecha(fechas)
    reports = df
    db.create_reports(reports)


def get_dataframe(file):
    return pd.read_csv(file, sep=';', header=0, encoding='ISO-8859-1')


if __name__ == "__main__":
    main()
