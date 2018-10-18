# idiomas_crud
API de idiomas, Integración con CI
idiomas_crud master/develop
 ## Requirements
Go version >= 1.8.
 ## Preparation:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/idiomas_crud
 ## Run
 Definir los valores de las siguientes variables de entorno:
  - `IDIOMAS_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `IDIOMAS_CRUD__PGUSER`: Usuario de la base de datos
 - `IDIOMAS_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `IDIOMAS_CRUD__PGURLS`: Host de conexión
 - `IDIOMAS_CRUD__PGDB`: Nombre de la base de datos
 - `IDIOMAS_CRUD__SCHEMA`: Esquema a utilizar en la base de datos
 ## Example:
IDIOMAS_HTTP_PORT=8095 IDIOMAS_CRUD__PGUSER=postgres IDIOMAS_CRUD__PGPASS=password IDIOMAS_CRUD__PGURLS=localhost IDIOMAS_CRUD__PGDB=local IDIOMAS_CRUD__SCHEMA=core_new bee run
 ## Model DB
![image](https://github.com/udistrital/idiomas_crud/blob/develop/modelo_idiomas_crud.png).
