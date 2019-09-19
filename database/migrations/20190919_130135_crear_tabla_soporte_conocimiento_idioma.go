package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaSoporteConocimientoIdioma_20190919_130135 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaSoporteConocimientoIdioma_20190919_130135{}
	m.Created = "20190919_130135"

	migration.Register("CrearTablaSoporteConocimientoIdioma_20190919_130135", m)
}

// Run the migrations
func (m *CrearTablaSoporteConocimientoIdioma_20190919_130135) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS idiomas.soporte_conocimiento_idioma( id serial NOT NULL, descripcion character varying(100), fecha_creacion TIMESTAMP, fecha_modificacion TIMESTAMP, terceros_id integer NOT NULL, documento_id integer NOT NULL, conocimiento_idioma_id integer NOT NULL, CONSTRAINT pk_soporte_conocimiento_idioma PRIMARY KEY (id), CONSTRAINT fk_conocimiento_idioma_soporte_conocimiento_idioma FOREIGN KEY (conocimiento_idioma_id) REFERENCES idiomas.conocimiento_idioma(id) );")
	m.SQL("ALTER TABLE idiomas.soporte_conocimiento_idioma OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaSoporteConocimientoIdioma_20190919_130135) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS idiomas.soporte_conocimiento_idioma")
}
