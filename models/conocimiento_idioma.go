package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type ConocimientoIdioma struct {
	Id                       int               `orm:"column(id);pk;auto"`
	Persona                  int               `orm:"column(persona)"`
	Idioma                   *Idioma           `orm:"column(idioma);rel(fk)"`
	NivelLee                 *ValorNivelIdioma `orm:"column(nivel_lee);rel(fk)"`
	NivelEscribe             *ValorNivelIdioma `orm:"column(nivel_escribe);rel(fk)"`
	NivelEscucha             *ValorNivelIdioma `orm:"column(nivel_escucha);rel(fk)"`
	NivelHabla               *ValorNivelIdioma `orm:"column(nivel_habla);rel(fk)"`
	Nativo                   bool              `orm:"column(nativo)"`
	ClasificacionNivelIdioma *Nivel            `orm:"column(clasificacion_nivel_idioma);rel(fk)"`
	Activo                   bool              `orm:"column(activo)"`
	FechaCreacion            string            `orm:"column(fecha_creacion);null"`
	FechaModificacion        string            `orm:"column(fecha_modificacion);null"`
}

type ConocimientoIdiomaV2 struct {
	Id                int               `orm:"column(id);pk;auto"`
	TercerosId        int               `orm:"column(terceros_id)"`
	IdiomaId          *Idioma           `orm:"column(idioma_id);rel(fk)"`
	Nativo            bool              `orm:"column(nativo)"`
	SeleccionExamen   bool              `orm:"column(seleccion_examen)"`
	NivelId           *Nivel            `orm:"column(nivel_id);rel(fk)"`
	NivelLeeId        *ValorNivelIdioma `orm:"column(nivel_lee_id);rel(fk)"`
	NivelEscribeId    *ValorNivelIdioma `orm:"column(nivel_escribe_id);rel(fk)"`
	NivelEscuchaId    *ValorNivelIdioma `orm:"column(nivel_escucha_id);rel(fk)"`
	Activo            bool              `orm:"column(activo)"`
	NivelHablaId      *ValorNivelIdioma `orm:"column(nivel_habla_id);rel(fk)"`
	FechaCreacion     time.Time         `orm:"column(fecha_creacion);type(timestamp without time zone);null"`
	FechaModificacion time.Time         `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
}

func (t *ConocimientoIdiomaV2) TableName() string {
	return "conocimiento_idioma"
}

func init() {
	orm.RegisterModel(new(ConocimientoIdiomaV2))
}

// AddConocimientoIdioma insert a new ConocimientoIdioma into database and returns
// last inserted Id on success.
func AddConocimientoIdioma(m *ConocimientoIdiomaV2) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetConocimientoIdiomaById retrieves ConocimientoIdioma by Id. Returns error if
// Id doesn't exist
func GetConocimientoIdiomaById(id int) (v *ConocimientoIdiomaV2, err error) {
	o := orm.NewOrm()
	v = &ConocimientoIdiomaV2{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllConocimientoIdioma retrieves all ConocimientoIdioma matches certain condition. Returns empty list if
// no records exist
func GetAllConocimientoIdioma(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ConocimientoIdiomaV2)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.HasSuffix(k, "__in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ConocimientoIdiomaV2
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateConocimientoIdioma updates ConocimientoIdioma by Id and returns error if
// the record to be updated doesn't exist
func UpdateConocimientoIdiomaById(m *ConocimientoIdiomaV2) (err error) {
	o := orm.NewOrm()
	v := ConocimientoIdiomaV2{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.FechaCreacion = v.FechaCreacion
		m.FechaModificacion = time_bogota.Tiempo_bogota()
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteConocimientoIdioma deletes ConocimientoIdioma by Id and returns error if
// the record to be deleted doesn't exist
func DeleteConocimientoIdioma(id int) (err error) {
	o := orm.NewOrm()
	v := ConocimientoIdiomaV2{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ConocimientoIdiomaV2{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
