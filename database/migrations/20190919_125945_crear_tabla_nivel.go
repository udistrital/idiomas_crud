package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNivel_20190919_125945 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNivel_20190919_125945{}
	m.Created = "20190919_125945"

	migration.Register("CrearTablaNivel_20190919_125945", m)
}

// Run the migrations
func (m *CrearTablaNivel_20190919_125945) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS idiomas.nivel( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP, fecha_modificacion TIMESTAMP, CONSTRAINT pk_nivel PRIMARY KEY (id), CONSTRAINT unique_nombre_nivel UNIQUE (nombre) );")
	m.SQL("ALTER TABLE idiomas.nivel OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE idiomas.nivel IS 'Tabla que parametriza los los niveles de idiomas existentes en la clasificación del Marco Común Europeo de Referencia para las lenguas (MCERL).';")
	m.SQL("COMMENT ON COLUMN idiomas.nivel.id IS 'Identificador unico de la tabla nivel.';")
	m.SQL("COMMENT ON COLUMN idiomas.nivel.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN idiomas.nivel.descripcion IS 'Campo en el que se puede registrar una descripcion de la información de nivel idioma.';")
	m.SQL("COMMENT ON COLUMN idiomas.nivel.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN idiomas.nivel.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN idiomas.nivel.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN idiomas.nivel.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN idiomas.nivel.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaNivel_20190919_125945) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS idiomas.nivel")
}
