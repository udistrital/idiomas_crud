package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarParametroDeExamen_20210724_174732 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarParametroDeExamen_20210724_174732{}
	m.Created = "20210724_174732"

	migration.Register("AgregarParametroDeExamen_20210724_174732", m)
}

// Run the migrations
func (m *AgregarParametroDeExamen_20210724_174732) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE idiomas.conocimiento_idioma ADD COLUMN seleccion_examen boolean;")
}

// Reverse the migrations
func (m *AgregarParametroDeExamen_20210724_174732) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE idiomas.conocimiento_idioma DROP COLUMN seleccion_examen;")
}