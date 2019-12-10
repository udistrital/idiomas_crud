package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertNivel_20191210_124030 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertNivel_20191210_124030{}
	m.Created = "20191210_124030"

	migration.Register("InsertNivel_20191210_124030", m)
}

// Run the migrations
func (m *InsertNivel_20191210_124030) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO idiomas.nivel(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (1,'A1','Nivel A1','A1', true,1, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO idiomas.nivel(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (2,'A2','Nivel A2','A2', true,2, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO idiomas.nivel(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (3,'B1','Nivel B1','B1', true,3, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO idiomas.nivel(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (4,'B2','Nivel B2','B2', true,4, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO idiomas.nivel(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (5,'C1','Nivel C1','C1', true,5, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO idiomas.nivel(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (6,'C2','Nivel C2','C2', true,6, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	
}

// Reverse the migrations
func (m *InsertNivel_20191210_124030) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM idiomas.nivel WHERE id BETWEEN 1 AND 6;")
}
