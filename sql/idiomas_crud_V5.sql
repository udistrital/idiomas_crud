-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler version: 0.9.4
-- PostgreSQL version: 13.0
-- Project Site: pgmodeler.io
-- Model Author: ---
-- object: desarrollooas | type: ROLE --
-- DROP ROLE IF EXISTS desarrollooas;
CREATE ROLE desarrollooas WITH 
	SUPERUSER
	CREATEDB
	CREATEROLE
	INHERIT
	LOGIN
	REPLICATION
	ENCRYPTED PASSWORD '********';
-- ddl-end --
COMMENT ON ROLE desarrollooas IS E'Usuario pruebas udistrital';
-- ddl-end --


-- Database creation must be performed outside a multi lined SQL file. 
-- These commands were put in this file only as a convenience.
-- 
-- object: idiomas | type: DATABASE --
-- DROP DATABASE IF EXISTS idiomas;
CREATE DATABASE idiomas
	ENCODING = 'UTF8'
	LC_COLLATE = 'es_CO.UTF-8'
	LC_CTYPE = 'es_CO.UTF-8'
	TABLESPACE = pg_default
	OWNER = desarrollooas;
-- ddl-end --


-- object: idiomas | type: SCHEMA --
-- DROP SCHEMA IF EXISTS idiomas CASCADE;
CREATE SCHEMA idiomas;
-- ddl-end --
ALTER SCHEMA idiomas OWNER TO desarrollooas;
-- ddl-end --

SET search_path TO pg_catalog,public,idiomas;
-- ddl-end --

-- object: idiomas.idioma_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS idiomas.idioma_id_seq CASCADE;
CREATE SEQUENCE idiomas.idioma_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE idiomas.idioma_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: idiomas.idioma | type: TABLE --
-- DROP TABLE IF EXISTS idiomas.idioma CASCADE;
CREATE TABLE idiomas.idioma (
	id integer NOT NULL DEFAULT nextval('idiomas.idioma_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_idioma PRIMARY KEY (id),
	CONSTRAINT unique_nombre_idioma UNIQUE (nombre)
);
-- ddl-end --
COMMENT ON TABLE idiomas.idioma IS E'Tabla que parametriza los Idiomas.';
-- ddl-end --
COMMENT ON COLUMN idiomas.idioma.id IS E'Identificador unico de la tabla idioma.';
-- ddl-end --
COMMENT ON COLUMN idiomas.idioma.nombre IS E'Campo obligatorio de la tabla que indica el nombre del parámetro.';
-- ddl-end --
COMMENT ON COLUMN idiomas.idioma.descripcion IS E'Campo en el que se puede registrar una descripcion de la información de idioma.';
-- ddl-end --
COMMENT ON COLUMN idiomas.idioma.codigo_abreviacion IS E'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
-- ddl-end --
COMMENT ON COLUMN idiomas.idioma.activo IS E'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN idiomas.idioma.numero_orden IS E'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON COLUMN idiomas.idioma.fecha_creacion IS E'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN idiomas.idioma.fecha_modificacion IS E'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
ALTER TABLE idiomas.idioma OWNER TO desarrollooas;
-- ddl-end --

-- object: idiomas.nivel_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS idiomas.nivel_id_seq CASCADE;
CREATE SEQUENCE idiomas.nivel_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE idiomas.nivel_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: idiomas.nivel | type: TABLE --
-- DROP TABLE IF EXISTS idiomas.nivel CASCADE;
CREATE TABLE idiomas.nivel (
	id integer NOT NULL DEFAULT nextval('idiomas.nivel_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_nivel PRIMARY KEY (id),
	CONSTRAINT unique_nombre_nivel UNIQUE (nombre)
);
-- ddl-end --
COMMENT ON TABLE idiomas.nivel IS E'Tabla que parametriza los los niveles de idiomas existentes en la clasificación del Marco Común Europeo de Referencia para las lenguas (MCERL).';
-- ddl-end --
COMMENT ON COLUMN idiomas.nivel.id IS E'Identificador unico de la tabla nivel.';
-- ddl-end --
COMMENT ON COLUMN idiomas.nivel.nombre IS E'Campo obligatorio de la tabla que indica el nombre del parámetro.';
-- ddl-end --
COMMENT ON COLUMN idiomas.nivel.descripcion IS E'Campo en el que se puede registrar una descripcion de la información de nivel idioma.';
-- ddl-end --
COMMENT ON COLUMN idiomas.nivel.codigo_abreviacion IS E'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
-- ddl-end --
COMMENT ON COLUMN idiomas.nivel.activo IS E'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN idiomas.nivel.numero_orden IS E'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON COLUMN idiomas.nivel.fecha_creacion IS E'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN idiomas.nivel.fecha_modificacion IS E'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
ALTER TABLE idiomas.nivel OWNER TO desarrollooas;
-- ddl-end --

-- object: idiomas.valor_nivel_idioma_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS idiomas.valor_nivel_idioma_id_seq CASCADE;
CREATE SEQUENCE idiomas.valor_nivel_idioma_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE idiomas.valor_nivel_idioma_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: idiomas.valor_nivel_idioma | type: TABLE --
-- DROP TABLE IF EXISTS idiomas.valor_nivel_idioma CASCADE;
CREATE TABLE idiomas.valor_nivel_idioma (
	id integer NOT NULL DEFAULT nextval('idiomas.valor_nivel_idioma_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_valor_nivel_idioma PRIMARY KEY (id),
	CONSTRAINT unique_nombre_valor_nivel_idioma UNIQUE (nombre)
);
-- ddl-end --
COMMENT ON TABLE idiomas.valor_nivel_idioma IS E'Tabla paramétrica que almacena los posibles valores para los niveles de lectura, escritura, escucha y habla de un idioma..';
-- ddl-end --
COMMENT ON COLUMN idiomas.valor_nivel_idioma.id IS E'Identificador unico de la tabla valor_nivel_idioma.';
-- ddl-end --
COMMENT ON COLUMN idiomas.valor_nivel_idioma.nombre IS E'Campo obligatorio de la tabla que indica el nombre del parámetro.';
-- ddl-end --
COMMENT ON COLUMN idiomas.valor_nivel_idioma.descripcion IS E'Campo en el que se puede registrar una descripcion de la información del valor de lectura, escritura, escucha y habla.';
-- ddl-end --
COMMENT ON COLUMN idiomas.valor_nivel_idioma.codigo_abreviacion IS E'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
-- ddl-end --
COMMENT ON COLUMN idiomas.valor_nivel_idioma.activo IS E'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN idiomas.valor_nivel_idioma.numero_orden IS E'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON COLUMN idiomas.valor_nivel_idioma.fecha_creacion IS E'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN idiomas.valor_nivel_idioma.fecha_modificacion IS E'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
ALTER TABLE idiomas.valor_nivel_idioma OWNER TO desarrollooas;
-- ddl-end --

-- object: idiomas.conocimiento_idioma_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS idiomas.conocimiento_idioma_id_seq CASCADE;
CREATE SEQUENCE idiomas.conocimiento_idioma_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE idiomas.conocimiento_idioma_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: idiomas.conocimiento_idioma | type: TABLE --
-- DROP TABLE IF EXISTS idiomas.conocimiento_idioma CASCADE;
CREATE TABLE idiomas.conocimiento_idioma (
	id integer NOT NULL DEFAULT nextval('idiomas.conocimiento_idioma_id_seq'::regclass),
	terceros_id integer NOT NULL,
	idioma_id integer NOT NULL,
	nativo boolean NOT NULL,
	nivel_id integer,
	nivel_lee_id integer,
	nivel_escribe_id integer,
	nivel_escucha_id integer,
	nivel_habla_id integer,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	activo boolean NOT NULL,
	CONSTRAINT pk_conocimiento_idioma PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE idiomas.conocimiento_idioma IS E'Tabla que almacena los conocimientos en idiomas que tiene una Persona (Terceros).';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.id IS E'Identificador unico de la tabla conocimiento_idioma.';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.terceros_id IS E'Persona a la que se le relaciona un nivel de idioma.';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.idioma_id IS E'Identificador unico de la tabla idioma, Idioma al cual se le relaciona el nivel y clasificación.';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.nativo IS E'Valor booleano para indicar si el idioma es nativo de la Persona.';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_id IS E'Identificador unico de la tabla nivel que parametriza el nivel de idioma de la Persona en la clasificación (MCERL).';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_lee_id IS E'Nivel de lectura que tiene la persona sobre el idioma.';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_escribe_id IS E'Nivel de escritura que tiene la persona sobre el idioma.';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_escucha_id IS E'Nivel de escucha que tiene la persona sobre el idioma.';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_habla_id IS E'Nivel de habla que tiene la persona sobre el idioma.';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.fecha_creacion IS E'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.fecha_modificacion IS E'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN idiomas.conocimiento_idioma.activo IS E'Indica el estado del registro';
-- ddl-end --
ALTER TABLE idiomas.conocimiento_idioma OWNER TO desarrollooas;
-- ddl-end --

-- object: idiomas.soporte_conocimiento_idioma_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS idiomas.soporte_conocimiento_idioma_id_seq CASCADE;
CREATE SEQUENCE idiomas.soporte_conocimiento_idioma_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE idiomas.soporte_conocimiento_idioma_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: idiomas.soporte_conocimiento_idioma | type: TABLE --
-- DROP TABLE IF EXISTS idiomas.soporte_conocimiento_idioma CASCADE;
CREATE TABLE idiomas.soporte_conocimiento_idioma (
	id integer NOT NULL DEFAULT nextval('idiomas.soporte_conocimiento_idioma_id_seq'::regclass),
	descripcion character varying(100),
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	terceros_id integer NOT NULL,
	documento_id integer NOT NULL,
	conocimiento_idioma_id integer NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_soporte_conocimiento_idioma PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON COLUMN idiomas.soporte_conocimiento_idioma.activo IS E'Indica el estado del registro';
-- ddl-end --
ALTER TABLE idiomas.soporte_conocimiento_idioma OWNER TO desarrollooas;
-- ddl-end --

-- object: fk_idioma_conocimiento_idioma | type: CONSTRAINT --
-- ALTER TABLE idiomas.conocimiento_idioma DROP CONSTRAINT IF EXISTS fk_idioma_conocimiento_idioma CASCADE;
ALTER TABLE idiomas.conocimiento_idioma ADD CONSTRAINT fk_idioma_conocimiento_idioma FOREIGN KEY (idioma_id)
REFERENCES idiomas.idioma (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_nivel_conocimiento_idioma | type: CONSTRAINT --
-- ALTER TABLE idiomas.conocimiento_idioma DROP CONSTRAINT IF EXISTS fk_nivel_conocimiento_idioma CASCADE;
ALTER TABLE idiomas.conocimiento_idioma ADD CONSTRAINT fk_nivel_conocimiento_idioma FOREIGN KEY (nivel_id)
REFERENCES idiomas.nivel (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_valor_nivel_idioma_lee | type: CONSTRAINT --
-- ALTER TABLE idiomas.conocimiento_idioma DROP CONSTRAINT IF EXISTS fk_valor_nivel_idioma_lee CASCADE;
ALTER TABLE idiomas.conocimiento_idioma ADD CONSTRAINT fk_valor_nivel_idioma_lee FOREIGN KEY (nivel_lee_id)
REFERENCES idiomas.valor_nivel_idioma (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_valor_nivel_idioma_escribe | type: CONSTRAINT --
-- ALTER TABLE idiomas.conocimiento_idioma DROP CONSTRAINT IF EXISTS fk_valor_nivel_idioma_escribe CASCADE;
ALTER TABLE idiomas.conocimiento_idioma ADD CONSTRAINT fk_valor_nivel_idioma_escribe FOREIGN KEY (nivel_escribe_id)
REFERENCES idiomas.valor_nivel_idioma (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_valor_nivel_idioma_escucha | type: CONSTRAINT --
-- ALTER TABLE idiomas.conocimiento_idioma DROP CONSTRAINT IF EXISTS fk_valor_nivel_idioma_escucha CASCADE;
ALTER TABLE idiomas.conocimiento_idioma ADD CONSTRAINT fk_valor_nivel_idioma_escucha FOREIGN KEY (nivel_escucha_id)
REFERENCES idiomas.valor_nivel_idioma (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_valor_nivel_idioma_habla | type: CONSTRAINT --
-- ALTER TABLE idiomas.conocimiento_idioma DROP CONSTRAINT IF EXISTS fk_valor_nivel_idioma_habla CASCADE;
ALTER TABLE idiomas.conocimiento_idioma ADD CONSTRAINT fk_valor_nivel_idioma_habla FOREIGN KEY (nivel_habla_id)
REFERENCES idiomas.valor_nivel_idioma (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_conocimiento_idioma_soporte_conocimiento_idioma | type: CONSTRAINT --
-- ALTER TABLE idiomas.soporte_conocimiento_idioma DROP CONSTRAINT IF EXISTS fk_conocimiento_idioma_soporte_conocimiento_idioma CASCADE;
ALTER TABLE idiomas.soporte_conocimiento_idioma ADD CONSTRAINT fk_conocimiento_idioma_soporte_conocimiento_idioma FOREIGN KEY (conocimiento_idioma_id)
REFERENCES idiomas.conocimiento_idioma (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


