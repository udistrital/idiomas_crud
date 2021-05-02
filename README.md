# :speech_balloon: IDIOMAS CRUD


API CRUD para la gestión del conocimiento en idiomas con el que cuentan las personas relacionadas a la Universidad Distrital. El proyecto está escrito en el lenguaje GO, generado mediante el framework beego.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
IDIOMAS_CRUD_PGDB=[nombre de la base de datos]
IDIOMAS_CRUD_PGPASS=[password del usuario]
IDIOMAS_CRUD_PGHOST=[direccion de la base de datos]
IDIOMAS_CRUD_PGPORT=[Puerto de conexión con la base de datos]
IDIOMAS_CRUD_PGUSER=[usuario con acceso a la base de datos]
IDIOMAS_CRUD_PGSCHEMA=[esquema donde se ubican las tablas]
IDIOMAS_CRUD_HTTP_PORT=[puerto de ejecucion] bee run
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con IDIOMAS_CRUD_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/idiomas_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/idiomas_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
IDIOMAS_CRUD_HTTP_PORT=8080 IDIOMAS_CRUD_PGHOST=127.0.0.1:27017 IDIOMAS_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=idiomas_crud . --no-cache
# docker run -p 80:80 idiomas_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/idiomas_crud

#2. Moverse a la carpeta del repositorio
cd idiomas_crud

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```
## Estado CI

| Develop | Release | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/idiomas_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/idiomas_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/idiomas_crud/status.svg?ref=refs/heads/release)](https://hubci.portaloas.udistrital.edu.co/udistrital/idiomas_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/idiomas_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/idiomas_crud/) |




## Modelo de datos
![Modelo de datos idiomas](https://github.com/udistrital/idiomas_crud/blob/dev/sql/idiomas_crud.png?raw=true)


  
## Licencia

This file is part of idiomas_crud

idiomas_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

idiomas_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with idiomas_crud. If not, see https://www.gnu.org/licenses/.
