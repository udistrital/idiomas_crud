package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterAddCamposActivo_20221028_113406 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterAddCamposActivo_20221028_113406{}
	m.Created = "20221028_113406"

	migration.Register("AlterAddCamposActivo_20221028_113406", m)
}

// Run the migrations
func (m *AlterAddCamposActivo_20221028_113406) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE IF EXISTS idiomas.conocimiento_idioma ADD COLUMN activo boolean NOT NULL;")
	m.SQL("COMMENT ON COLUMN idiomas.conocimiento_idioma.activo IS 'Indica el estado del registro';")

	m.SQL("ALTER TABLE IF EXISTS idiomas.soporte_conocimiento_idioma ADD COLUMN activo boolean NOT NULL;")
	m.SQL("COMMENT ON COLUMN idiomas.soporte_conocimiento_idioma.activo IS 'Indica el estado del registro';")

}

// Reverse the migrations
func (m *AlterAddCamposActivo_20221028_113406) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE IF EXISTS idiomas.conocimiento_idioma DROP COLUMN IF EXISTS activo;")

	m.SQL("ALTER TABLE IF EXISTS idiomas.soporte_conocimiento_idioma DROP COLUMN IF EXISTS activo;")

}
