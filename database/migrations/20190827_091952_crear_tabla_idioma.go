package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaIdioma_20190827_091952 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaIdioma_20190827_091952{}
	m.Created = "20190827_091952"

	migration.Register("CrearTablaIdioma_20190827_091952", m)
}

// Run the migrations
func (m *CrearTablaIdioma_20190827_091952) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS idiomas.idioma( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP, fecha_modificacion TIMESTAMP, CONSTRAINT pk_idioma PRIMARY KEY (id), CONSTRAINT unique_nombre_idioma UNIQUE (nombre));")
	m.SQL("ALTER TABLE idiomas.idioma OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE idiomas.idioma IS 'Tabla que parametriza los Idiomas.';")
	m.SQL("COMMENT ON COLUMN idiomas.idioma.id IS 'Identificador unico de la tabla idioma.';")
	m.SQL("COMMENT ON COLUMN idiomas.idioma.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN idiomas.idioma.descripcion IS 'Campo en el que se puede registrar una descripcion de la información de idioma.';")
	m.SQL("COMMENT ON COLUMN idiomas.idioma.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN idiomas.idioma.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN idiomas.idioma.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN idiomas.idioma.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN idiomas.idioma.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaIdioma_20190827_091952) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS idiomas.idioma")
}
