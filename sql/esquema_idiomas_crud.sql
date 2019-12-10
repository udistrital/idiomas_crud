--Migracion:20190827_091707_crear_schema_idiomas
CREATE SCHEMA idiomas;

--idioma >> idioma
--clasificación_nivel_idioma >> nivel
--soporte_conocimiento_idioma >> soporte_conocimiento_idioma
--conocimiento_idioma >> conocimiento_idioma
--valor_nivel_idioma >> valor_nivel_idioma

--######################################################################################################################################################
--######################################################################################################################################################
--Creación de secuencia y tabla idioma:
--Migracion:20190827_091952_crear_tabla_idioma
CREATE SEQUENCE idiomas.idioma_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE idiomas.idioma(
	id integer NOT NULL DEFAULT nextval('idiomas.idioma_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	CONSTRAINT pk_idioma PRIMARY KEY (id),
    CONSTRAINT unique_nombre_idioma UNIQUE (nombre)
);
COMMENT ON TABLE idiomas.idioma IS 'Tabla que parametriza los Idiomas.';
COMMENT ON COLUMN idiomas.idioma.id IS 'Identificador unico de la tabla idioma.';
COMMENT ON COLUMN idiomas.idioma.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN idiomas.idioma.descripcion IS 'Campo en el que se puede registrar una descripcion de la información de idioma.';
COMMENT ON COLUMN idiomas.idioma.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN idiomas.idioma.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN idiomas.idioma.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN idiomas.idioma.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN idiomas.idioma.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
--######################################################################################################################################################
--######################################################################################################################################################
--Creación de secuencia y tabla nivel:
--Migracion:20190827_092656_crear_tabla_nivel
CREATE SEQUENCE idiomas.nivel_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE idiomas.nivel(
	id integer NOT NULL DEFAULT nextval('idiomas.nivel_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
    descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	CONSTRAINT pk_nivel PRIMARY KEY (id),
    CONSTRAINT unique_nombre_nivel UNIQUE (nombre)
);
COMMENT ON TABLE idiomas.nivel IS 'Tabla que parametriza los los niveles de idiomas existentes en la clasificación del Marco Común Europeo de Referencia para las lenguas (MCERL).';
COMMENT ON COLUMN idiomas.nivel.id IS 'Identificador unico de la tabla nivel.';
COMMENT ON COLUMN idiomas.nivel.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN idiomas.nivel.descripcion IS 'Campo en el que se puede registrar una descripcion de la información de nivel idioma.';
COMMENT ON COLUMN idiomas.nivel.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN idiomas.nivel.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN idiomas.nivel.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN idiomas.nivel.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN idiomas.nivel.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
--######################################################################################################################################################
--######################################################################################################################################################
--Creación de secuencia y tabla valor_nivel_idioma:
--Migracion:20190827_093911_crear_tabla_valor_nivel_idioma
CREATE SEQUENCE idiomas.valor_nivel_idioma_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE idiomas.valor_nivel_idioma(
	id integer NOT NULL DEFAULT nextval('idiomas.valor_nivel_idioma_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
    descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	CONSTRAINT pk_valor_nivel_idioma PRIMARY KEY (id),
    CONSTRAINT unique_nombre_valor_nivel_idioma UNIQUE (nombre)
);
COMMENT ON TABLE idiomas.valor_nivel_idioma IS 'Tabla paramétrica que almacena los posibles valores para los niveles de lectura, escritura, escucha y habla de un idioma..';
COMMENT ON COLUMN idiomas.valor_nivel_idioma.id IS 'Identificador unico de la tabla valor_nivel_idioma.';
COMMENT ON COLUMN idiomas.valor_nivel_idioma.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN idiomas.valor_nivel_idioma.descripcion IS 'Campo en el que se puede registrar una descripcion de la información del valor de lectura, escritura, escucha y habla.';
COMMENT ON COLUMN idiomas.valor_nivel_idioma.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN idiomas.valor_nivel_idioma.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN idiomas.valor_nivel_idioma.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN idiomas.valor_nivel_idioma.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN idiomas.valor_nivel_idioma.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
--######################################################################################################################################################
--######################################################################################################################################################
--Creación de secuencia y tabla conocimiento_idioma:
--Migracion:20190827_094755_crear_tabla_conocimiento_idioma
CREATE SEQUENCE idiomas.conocimiento_idioma_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE idiomas.conocimiento_idioma(
	id integer NOT NULL DEFAULT nextval('idiomas.conocimiento_idioma_id_seq'::regclass),
    terceros_id integer NOT NULL,
    idioma_id integer NOT NULL,
    nativo boolean NOT NULL,
    nivel_id integer,
    nivel_lee_id integer,
    nivel_escribe_id integer,
    nivel_escucha_id integer,
    nivel_habla_id integer,
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	CONSTRAINT pk_conocimiento_idioma PRIMARY KEY (id),
    CONSTRAINT fk_idioma_conocimiento_idioma FOREIGN KEY (idioma_id) REFERENCES idiomas.idioma(id),
    CONSTRAINT fk_nivel_conocimiento_idioma FOREIGN KEY (nivel_id) REFERENCES idiomas.nivel(id),
    CONSTRAINT fk_valor_nivel_idioma_lee FOREIGN KEY (nivel_lee_id) REFERENCES idiomas.valor_nivel_idioma(id),
    CONSTRAINT fk_valor_nivel_idioma_escribe FOREIGN KEY (nivel_escribe_id) REFERENCES idiomas.valor_nivel_idioma(id),
    CONSTRAINT fk_valor_nivel_idioma_escucha FOREIGN KEY (nivel_escucha_id) REFERENCES idiomas.valor_nivel_idioma(id),
    CONSTRAINT fk_valor_nivel_idioma_habla FOREIGN KEY (nivel_habla_id) REFERENCES idiomas.valor_nivel_idioma(id)
);

COMMENT ON TABLE idiomas.conocimiento_idioma IS 'Tabla que almacena los conocimientos en idiomas que tiene una Persona (Terceros).';
COMMENT ON COLUMN idiomas.conocimiento_idioma.id IS 'Identificador unico de la tabla conocimiento_idioma.';
COMMENT ON COLUMN idiomas.conocimiento_idioma.terceros_id IS 'Persona a la que se le relaciona un nivel de idioma.';
COMMENT ON COLUMN idiomas.conocimiento_idioma.idioma_id IS 'Identificador unico de la tabla idioma, Idioma al cual se le relaciona el nivel y clasificación.';
COMMENT ON COLUMN idiomas.conocimiento_idioma.nativo IS 'Valor booleano para indicar si el idioma es nativo de la Persona.';
COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_id IS 'Identificador unico de la tabla nivel que parametriza el nivel de idioma de la Persona en la clasificación (MCERL).';
COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_lee_id IS 'Nivel de lectura que tiene la persona sobre el idioma.';
COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_escribe_id IS 'Nivel de escritura que tiene la persona sobre el idioma.';
COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_escucha_id IS 'Nivel de escucha que tiene la persona sobre el idioma.';
COMMENT ON COLUMN idiomas.conocimiento_idioma.nivel_habla_id IS 'Nivel de habla que tiene la persona sobre el idioma.';
COMMENT ON COLUMN idiomas.conocimiento_idioma.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN idiomas.conocimiento_idioma.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
--######################################################################################################################################################
--######################################################################################################################################################
--Creación de secuencia y tabla soporte_conocimiento_idioma:
--Migracion:20190827_095153_crear_tabla_soporte_conocimiento_idioma
CREATE SEQUENCE idiomas.soporte_conocimiento_idioma_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE idiomas.soporte_conocimiento_idioma(
	id integer NOT NULL DEFAULT nextval('idiomas.soporte_conocimiento_idioma_id_seq'::regclass),
    descripcion character varying(100),    
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
    terceros_id integer NOT NULL,
    documento_id integer NOT NULL,
    conocimiento_idioma_id integer NOT NULL,
	CONSTRAINT pk_soporte_conocimiento_idioma PRIMARY KEY (id),
    CONSTRAINT fk_conocimiento_idioma_soporte_conocimiento_idioma FOREIGN KEY (conocimiento_idioma_id) REFERENCES idiomas.conocimiento_idioma(id)
);
COMMENT ON TABLE idiomas.soporte_conocimiento_idioma IS 'Tabla que almacena los documentos que certifican el nivel que tiene la persona sobre el idioma.';
COMMENT ON COLUMN idiomas.soporte_conocimiento_idioma.id IS 'Identificador unico de la tabla soporte_conocimiento_idioma.';
COMMENT ON COLUMN idiomas.soporte_conocimiento_idioma.descripcion IS 'Campo en el que se puede registrar una descripcion acerca del certificado.';
COMMENT ON COLUMN idiomas.soporte_conocimiento_idioma.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN idiomas.soporte_conocimiento_idioma.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
COMMENT ON COLUMN idiomas.soporte_conocimiento_idioma.terceros_id IS 'Institución que certifica el nivel del idioma.';
COMMENT ON COLUMN idiomas.soporte_conocimiento_idioma.documento_id IS 'Identificador de la tabla documento.';
COMMENT ON COLUMN idiomas.soporte_conocimiento_idioma.conocimiento_idioma_id IS 'Identificador de la tabla conocimiento_idioma.';
--######################################################################################################################################################
--######################################################################################################################################################

--D10 - Se hacen migraciones para insertar data del Repositorio de Planestic
