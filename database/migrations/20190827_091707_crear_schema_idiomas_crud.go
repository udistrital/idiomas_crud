package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchemaIdiomasCrud_20190827_091707 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchemaIdiomasCrud_20190827_091707{}
	m.Created = "20190827_091707"

	migration.Register("CrearSchemaIdiomasCrud_20190827_091707", m)
}

// Run the migrations
func (m *CrearSchemaIdiomasCrud_20190827_091707) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA IF NOT EXISTS idiomas_crud;")
	m.SQL("ALTER SCHEMA idiomas_crud OWNER TO desarrollooas;")
	m.SQL("SET search_path TO pg_catalog,public,idiomas_crud;")
	
}

// Reverse the migrations
func (m *CrearSchemaIdiomasCrud_20190827_091707) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS idiomas_crud");
}
