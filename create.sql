CREATE DATABASE IF NOT EXISTS `autodata`;

USE `autodata`;

CREATE TABLE IF NOT EXISTS `provincia` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `nombre` varchar(255)
) ENGINE=MyISAM;

CREATE TABLE IF NOT EXISTS `centro_examen` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `nombre` varchar(255),
  `id_provincia` int,
  
  FOREIGN KEY (`id_provincia`) REFERENCES `provincia` (`id`)
) ENGINE=MyISAM;

CREATE TABLE IF NOT EXISTS `autoescuela` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `codigo` varchar(255),
  `nombre` varchar(255)
) ENGINE=MyISAM;

CREATE TABLE IF NOT EXISTS `seccion` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `codigo` int,
  `id_autoescuela` int,

  FOREIGN KEY (`id_autoescuela`) REFERENCES `autoescuela` (`id`),

  UNIQUE INDEX `uniq_seccion` (`codigo`, `id_autoescuela`)
) ENGINE=MyISAM;

CREATE TABLE IF NOT EXISTS `permiso` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `nombre` varchar(255)
) ENGINE=MyISAM;

CREATE TABLE IF NOT EXISTS `tipo_examen` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `nombre` varchar(255)
) ENGINE=MyISAM;

CREATE TABLE IF NOT EXISTS `fecha` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `mes` int,
  `anyo` int
) ENGINE=MyISAM;

CREATE TABLE IF NOT EXISTS `report` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `id_seccion` int,
  `id_centro_examen` int,
  `id_permiso` int,
  `id_tipo_examen` int,
  `id_fecha` int,
  `total_aptos` int,
  `aptos_1` int,
  `aptos_2` int,
  `aptos_3_4` int,
  `aptos_5m` int,
  `no_aptos` int,

  FOREIGN KEY (`id_seccion`) REFERENCES `seccion` (`id`),
  FOREIGN KEY (`id_centro_examen`) REFERENCES `centro_examen` (`id`),
  FOREIGN KEY (`id_permiso`) REFERENCES `permiso` (`id`),
  FOREIGN KEY (`id_tipo_examen`) REFERENCES `tipo_examen` (`id`),
  FOREIGN KEY (`id_fecha`) REFERENCES `fecha` (`id`),

  UNIQUE INDEX `uniq_report` (`id_seccion`, `id_centro_examen`, `id_permiso`, `id_tipo_examen`, `id_fecha`)
) ENGINE=MyISAM;

CREATE TABLE IF NOT EXISTS `report_files` (
  `nombre` varchar(255),
  `leido` boolean default 0
) ENGINE=MyISAM;
