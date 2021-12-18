import os
import mysql.connector
from dotenv import load_dotenv
from mysql.connector.errors import DatabaseError

load_dotenv()
MYSQL_HOST = os.getenv('MYSQL_HOST')
MYSQL_USER = os.getenv('MYSQL_USER')
MYSQL_PASSWORD = os.getenv('MYSQL_PASSWORD')
MYSQL_DATABASE = os.getenv('MYSQL_DATABASE')

cnx = mysql.connector.connect(
    host=MYSQL_HOST,
    user=MYSQL_USER,
    password=MYSQL_PASSWORD,
    database=MYSQL_DATABASE
)


def close_cnx():
    cnx.close()


def clean_param(param):
    if type(param) is str:
        param = param.strip()
    return param


def clean_params(params):
    return tuple(map(clean_param, params))


def find_one(query, params):
    params = clean_params(params)
    cursor = cnx.cursor()
    cursor.execute(query, params)
    result = cursor.fetchone()
    cursor.close()
    return result


def add_one(query, params):
    created_id = None
    try:
        params = clean_params(params)
        cursor = cnx.cursor()
        cursor.execute(query, params)
        created_id = cursor.lastrowid
        cnx.commit()
        cursor.close()
    except DatabaseError:
        print(f'WARNING: could not add {params}')
    return created_id


def add_provincia(nombre):
    query = 'INSERT INTO `provincia` (`nombre`) VALUES (%s);'
    params = (nombre,)
    return add_one(query, params)


def find_provincia(nombre):
    query = 'SELECT `id`, `nombre` FROM `provincia` WHERE `nombre` = %s;'
    params = (nombre,)
    return find_one(query, params)


def update_provincias(provincias):
    for provincia in provincias:
        if find_provincia(provincia) is None:
            add_provincia(provincia)


def add_centro(nombre, id_provincia):
    query = 'INSERT INTO `centro_examen` (`nombre`, `provincia`) VALUES (%s, %s);'
    params = (nombre, id_provincia)
    return add_one(query, params)


def find_centro(nombre, id_provincia):
    query = 'SELECT `id`, `nombre`, `provincia` FROM `centro_examen` WHERE `nombre` = %s AND `provincia` = %s;'
    params = (nombre, id_provincia)
    return find_one(query, params)


def update_centros(centros):
    for _, (centro, provincia) in centros.iterrows():
        id_provincia, _ = find_provincia(provincia)
        if find_centro(centro, id_provincia) is None:
            add_centro(centro, id_provincia)


def add_autoescuela(codigo, nombre):
    query = 'INSERT INTO `autoescuela` (`codigo`, `nombre`) VALUES (%s, %s);'
    params = (codigo, nombre)
    return add_one(query, params)


def find_autoescuela(codigo):
    query = 'SELECT `id`, `codigo`, `nombre` FROM `autoescuela` WHERE `codigo` = %s;'
    params = (codigo,)
    return find_one(query, params)


def update_autoescuelas(autoescuelas):
    for _, (codigo, nombre) in autoescuelas.iterrows():
        if find_autoescuela(codigo) is None:
            add_autoescuela(codigo, nombre)


def add_seccion(codigo, id_autoescuela):
    query = 'INSERT INTO `seccion` (`codigo`, `autoescuela`) VALUES (%s, %s);'
    params = (codigo, id_autoescuela)
    return add_one(query, params)


def find_seccion(codigo, id_autoescuela):
    query = 'SELECT `id`, `codigo`, `autoescuela` ' \
            'FROM `seccion` WHERE `codigo` = %s AND `autoescuela` = %s;'
    params = (codigo, id_autoescuela)
    return find_one(query, params)


def update_secciones(secciones):
    for _, (autoescuela, codigo) in secciones.iterrows():
        id_autoescuela, _, _ = find_autoescuela(autoescuela)
        if find_seccion(codigo, id_autoescuela) is None:
            add_seccion(codigo, id_autoescuela)


def add_tipo_examen(nombre):
    query = 'INSERT INTO `tipo_examen` (`nombre`) VALUES (%s);'
    params = (nombre,)
    return add_one(query, params)


def find_tipo_examen(nombre):
    query = 'SELECT `id`, `nombre` FROM `tipo_examen` WHERE `nombre` = %s'
    params = (nombre,)
    return find_one(query, params)


def update_tipos_examen(tipos_examen):
    for tipo in tipos_examen:
        if find_tipo_examen(tipo) is None:
            add_tipo_examen(tipo)


def add_permiso(nombre):
    query = 'INSERT INTO `permiso` (`nombre`) VALUES (%s);'
    params = (nombre,)
    return add_one(query, params)


def find_permiso(nombre):
    query = 'SELECT `id`, `nombre` FROM `permiso` WHERE `nombre` = %s'
    params = (nombre,)
    return find_one(query, params)


def update_permisos(permisos):
    for permiso in permisos:
        if find_permiso(permiso) is None:
            add_permiso(permiso)


def add_report(secc, centro, per, tipo, mes, anyo, a, a1, a2, a34, a5m, na):
    query = "INSERT INTO `report` (`seccion`, `centro_examen`, `permiso`, `tipo_examen`, `mes`, `anyo`, " \
            "`total_aptos`, `aptos_1`, `aptos_2`, `aptos_3_4`, `aptos_5m`, `no_aptos`) " \
            "VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"
    params = (secc, centro, per, tipo, mes, anyo, a, a1, a2, a34, a5m, na)
    return add_one(query, params)


def create_reports(records):
    for _, (prov, centro, auto, _, secc, mes, anyo, tipo, per, a, a1, a2, a34, a5m, na) in records.iterrows():
        id_provincia = find_provincia(prov)[0]
        id_centro = find_centro(centro, id_provincia)[0]
        id_autoescuela = find_autoescuela(auto)[0]
        id_seccion = find_seccion(secc, id_autoescuela)[0]
        id_permiso = find_permiso(per)[0]
        id_tipo = find_tipo_examen(tipo)[0]
        add_report(id_seccion, id_centro, id_permiso, id_tipo, mes, anyo, a, a1, a2, a34, a5m, na)
