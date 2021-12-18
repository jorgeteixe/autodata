<div id="top"></div>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/jorgeteixe/autodata">
    <img src="images/blue_car.png" alt="Logo" width="160" height="160">
  </a>

  <h3 align="center">AUTODATA</h3>

  <p align="center">
    Find out what is the best Driving School in Spain with real data.
    <br />
    <br />
    <a href="https://github.com/jorgeteixe/autodata">View online</a>
    ·
    <a href="https://github.com/jorgeteixe/autodata/issues">Report Bug</a>
    ·
    <a href="https://github.com/jorgeteixe/autodata/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#how-it-works">How it works</a>
      <ul>
        <li><a href="#data-collection">Data collection</a></li>
        <li><a href="#data-ingestion">Data ingestion</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

There are many driving schools out there but no one knows how to choose one, which is the better or any data. This page
helps you visualize the main metrics for all the Spanish Driving Schools out there.

### Built With

To build this app is used:

* [Python 3](https://python.org/)
* [Docker](https://www.docker.com/)
* [MySQL](https://www.mysql.com/)
* more soon...

<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

To run the app you will need:
* docker and docker-compose
  * to install visit [the docs](https://docs.docker.com/get-docker/)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/jorgeteixe/autodata.git
   cd autodata
   ```
2. Set the environment variables
   ```sh
   cp .env.dist .env
   ```
3. Run the app with
   ```sh
   docker-compose up -d
   ```

## How it works

### Data collection
The data that the app uses is gathered from the public statistics of the [DGT](https://www.dgt.es/inicio/). From their
page you can download a ZIP containing a `.txt` file for each month. The file has the following format:

```csv
DESC_PROVINCIA;CENTRO_EXAMEN;CODIGO_AUTOESCUELA;NOMBRE_AUTOESCUELA;CODIGO_SECCION;MES;ANYO;TIPO_EXAMEN;NOMBRE_PERMISO;NUM_APTOS;NUM_APTOS_1conv;NUM_APTOS_2conv;NUM_APTOS_3o4conv;NUM_APTOS_5_o_mas_conv;NUM_NO_APTOS
Albacete;Albacete;XX0234;AUTOESCUELA DE PRUEBA;01;11;2021;PRUEBA CONDUCCIÓN Y CIRCULACIÓN;B  ;3;1;2;0;0;1
# ...
```

To get this data, you need to download it manually from their page: [Portal estadístico DGT](https://sedeapl.dgt.gob.es/WEB_IEST_CONSULTA)

### Data ingestion
The process of converting the data into a manageable one is done in the `ingest/` folder and (with container name _autodata_ingest_).
It transforms the data and inserts it into a database with the following shema:

![Database model][db-model]

<!-- ROADMAP -->
## Roadmap

- [x] Create ingestion mechanism
- [x] Container infrastructure
- [ ] Create API access layer
- [ ] Create frontend
- [ ] Configure CI/CD

See the [open issues](https://github.com/jorgeteixe/autodata/issues) for a full list of proposed features (and known issues).


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.


<!-- CONTACT -->
## Contact

Jorge Teixeira Crespo - [webpage](https://jorgeteixeira.es) - hi@jorgeteixeira.es

Project Link: [https://github.com/jorgeteixe/autodata](https://github.com/jorgeteixe/autodata)



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Portal estadístico DGT](https://sedeapl.dgt.gob.es/WEB_IEST_CONSULTA)


<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/jorgeteixe/autodata.svg?style=for-the-badge
[contributors-url]: https://github.com/jorgeteixe/autodata/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/jorgeteixe/autodata.svg?style=for-the-badge
[forks-url]: https://github.com/jorgeteixe/autodata/network/members
[stars-shield]: https://img.shields.io/github/stars/jorgeteixe/autodata.svg?style=for-the-badge
[stars-url]: https://github.com/jorgeteixe/autodata/stargazers
[issues-shield]: https://img.shields.io/github/issues/jorgeteixe/autodata.svg?style=for-the-badge
[issues-url]: https://github.com/jorgeteixe/autodata/issues
[license-shield]: https://img.shields.io/github/license/jorgeteixe/autodata.svg?style=for-the-badge
[license-url]: https://github.com/jorgeteixe/autodata/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/jorge-teixeira-crespo
[db-model]: images/db.png
