package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaConocimientoIdioma_20190827_094755 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaConocimientoIdioma_20190827_094755{}
	m.Created = "20190827_094755"

	migration.Register("CrearTablaConocimientoIdioma_20190827_094755", m)
}

// Run the migrations
func (m *CrearTablaConocimientoIdioma_20190827_094755) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS idiomas.conocimiento_idioma( id serial NOT NULL, terceros_id integer NOT NULL, idioma_id integer NOT NULL, nativo boolean NOT NULL, nivel_id integer, nivel_lee_id integer, nivel_escribe_id integer, nivel_escucha_id integer, nivel_habla_id integer, fecha_creacion TIMESTAMP, fecha_modificacion TIMESTAMP, CONSTRAINT pk_conocimiento_idioma PRIMARY KEY (id), CONSTRAINT fk_idioma_conocimiento_idioma FOREIGN KEY (idioma_id) REFERENCES idiomas.idioma(id), CONSTRAINT fk_nivel_conocimiento_idioma FOREIGN KEY (nivel_id) REFERENCES idiomas.nivel(id), CONSTRAINT fk_valor_nivel_idioma_lee FOREIGN KEY (nivel_lee_id) REFERENCES idiomas.valor_nivel_idioma(id), CONSTRAINT fk_valor_nivel_idioma_escribe FOREIGN KEY (nivel_escribe_id) REFERENCES idiomas.valor_nivel_idioma(id), CONSTRAINT fk_valor_nivel_idioma_escucha FOREIGN KEY (nivel_escucha_id) REFERENCES idiomas.valor_nivel_idioma(id), CONSTRAINT fk_valor_nivel_idioma_habla FOREIGN KEY (nivel_habla_id) REFERENCES idiomas.valor_nivel_idioma(id) );")
	m.SQL("ALTER TABLE idiomas.conocimiento_idioma OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE idiomas.conocimiento_idioma IS 'Tabla que almacena los conocimientos en idiomas que tiene una Persona (Terceros).';")
	m.SQL("COMMENT ON COLUMN idiomas.conocimiento_idioma.id IS 'Identificador unico de la tabla conocimiento_idioma.';")
	m.SQL("COMMENT ON COLUMN idiomas.conocimiento_idioma.terceros_id IS 'Persona a la que se le relaciona un nivel de idioma.';")
	m.SQL("	COMMENT ON COLUMN idiomas.conocimiento_idioma.idioma_id IS 'Identificador unico de la tabla idioma, Idioma al cual se le relaciona el nivel y clasificación.';")
	m.SQL("COMMENT ON COLUMN idiomas.conocimiento_idioma.nativo IS 'Valor booleano para indicar si el idioma es nativo de la Persona.';")
	m.SQL("COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_id IS 'Identificador unico de la tabla nivel que parametriza el nivel de idioma de la Persona en la clasificación (MCERL).';")
	m.SQL("COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_lee_id IS 'Nivel de lectura que tiene la persona sobre el idioma.';")
	m.SQL("COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_escribe_id IS 'Nivel de escritura que tiene la persona sobre el idioma.';")
	m.SQL("COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_escucha_id IS 'Nivel de escucha que tiene la persona sobre el idioma.';")
	m.SQL("COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_habla_id IS 'Nivel de habla que tiene la persona sobre el idioma.';")
	m.SQL("COMMENT ON COLUMN idiomas.conocimiento_idioma.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN idiomas.conocimiento_idioma.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaConocimientoIdioma_20190827_094755) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS idiomas.conocimiento_idioma")
}
