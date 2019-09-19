package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaValorNivelIdioma_20190827_093911 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaValorNivelIdioma_20190827_093911{}
	m.Created = "20190827_093911"

	migration.Register("CrearTablaValorNivelIdioma_20190827_093911", m)
}

// Run the migrations
func (m *CrearTablaValorNivelIdioma_20190827_093911) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS idiomas.valor_nivel_idioma( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP, fecha_modificacion TIMESTAMP, CONSTRAINT pk_valor_nivel_idioma PRIMARY KEY (id), CONSTRAINT unique_nombre_valor_nivel_idioma UNIQUE (nombre) );")
	m.SQL("ALTER TABLE idiomas.valor_nivel_idioma OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE idiomas.valor_nivel_idioma IS 'Tabla paramétrica que almacena los posibles valores para los niveles de lectura, escritura, escucha y habla de un idioma..';")
	m.SQL("COMMENT ON COLUMN idiomas.valor_nivel_idioma.id IS 'Identificador unico de la tabla valor_nivel_idioma.';")
	m.SQL("COMMENT ON COLUMN idiomas.valor_nivel_idioma.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN idiomas.valor_nivel_idioma.descripcion IS 'Campo en el que se puede registrar una descripcion de la información del valor de lectura, escritura, escucha y habla.';")
	m.SQL("COMMENT ON COLUMN idiomas.valor_nivel_idioma.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN idiomas.valor_nivel_idioma.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN idiomas.valor_nivel_idioma.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN idiomas.valor_nivel_idioma.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN idiomas.valor_nivel_idioma.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaValorNivelIdioma_20190827_093911) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS idiomas.valor_nivel_idioma")
}
