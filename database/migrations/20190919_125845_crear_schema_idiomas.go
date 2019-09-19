package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchemaIdiomas_20190919_125845 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchemaIdiomas_20190919_125845{}
	m.Created = "20190919_125845"

	migration.Register("CrearSchemaIdiomas_20190919_125845", m)
}

// Run the migrations
func (m *CrearSchemaIdiomas_20190919_125845) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA IF NOT EXISTS idiomas;")
	m.SQL("ALTER SCHEMA idiomas OWNER TO desarrollooas;")
	m.SQL("SET search_path TO pg_catalog,public,idiomas;")
}

// Reverse the migrations
func (m *CrearSchemaIdiomas_20190919_125845) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS idiomas");
}
