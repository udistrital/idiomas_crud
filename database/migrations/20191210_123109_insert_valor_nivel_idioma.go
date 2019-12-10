package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertValorNivelIdioma_20191210_123109 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertValorNivelIdioma_20191210_123109{}
	m.Created = "20191210_123109"

	migration.Register("InsertValorNivelIdioma_20191210_123109", m)
}

// Run the migrations
func (m *InsertValorNivelIdioma_20191210_123109) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO idiomas.valor_nivel_idioma(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (1,'Regular','Valor regular','Regular', true,1, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO idiomas.valor_nivel_idioma(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (2,'Bien','Valor bien','Bien', true,2, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO idiomas.valor_nivel_idioma(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (3,'Muy bien','Valor muy bien','Muy bien', true,3, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertValorNivelIdioma_20191210_123109) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM idiomas.valor_nivel_idioma WHERE id BETWEEN 1 AND 3;")
}
