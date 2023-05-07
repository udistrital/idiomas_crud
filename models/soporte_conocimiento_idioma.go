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

type SoporteConocimientoIdioma struct {
	Id                 int                 `orm:"column(id);pk;auto"`
	Institucion        int                 `orm:"column(institucion)"`
	Documento          int                 `orm:"column(documento)"`
	Descripcion        string              `orm:"column(descripcion);null"`
	ConocimientoIdioma *ConocimientoIdioma `orm:"column(conocimiento_idioma);rel(fk)"`
	Activo             bool                `orm:"column(activo)"`
	FechaCreacion      string              `orm:"column(fecha_creacion);null"`
	FechaModificacion  string              `orm:"column(fecha_modificacion);null"`
}

type SoporteConocimientoIdiomaV2 struct {
	Id                   int                   `orm:"column(id);pk;auto"`
	Descripcion          string                `orm:"column(descripcion);null"`
	Activo               bool                  `orm:"column(activo)"`
	FechaCreacion        time.Time             `orm:"column(fecha_creacion);type(timestamp without time zone);null"`
	FechaModificacion    time.Time             `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	TercerosId           int                   `orm:"column(terceros_id)"`
	DocumentoId          int                   `orm:"column(documento_id)"`
	ConocimientoIdiomaId *ConocimientoIdiomaV2 `orm:"column(conocimiento_idioma_id);rel(fk)"`
}

func (t *SoporteConocimientoIdiomaV2) TableName() string {
	return "soporte_conocimiento_idioma"
}

func init() {
	orm.RegisterModel(new(SoporteConocimientoIdiomaV2))
}

// AddSoporteConocimientoIdioma insert a new SoporteConocimientoIdioma into database and returns
// last inserted Id on success.
func AddSoporteConocimientoIdioma(m *SoporteConocimientoIdiomaV2) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSoporteConocimientoIdiomaById retrieves SoporteConocimientoIdioma by Id. Returns error if
// Id doesn't exist
func GetSoporteConocimientoIdiomaById(id int) (v *SoporteConocimientoIdiomaV2, err error) {
	o := orm.NewOrm()
	v = &SoporteConocimientoIdiomaV2{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSoporteConocimientoIdioma retrieves all SoporteConocimientoIdioma matches certain condition. Returns empty list if
// no records exist
func GetAllSoporteConocimientoIdioma(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SoporteConocimientoIdiomaV2)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
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

	var l []SoporteConocimientoIdiomaV2
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

// UpdateSoporteConocimientoIdioma updates SoporteConocimientoIdioma by Id and returns error if
// the record to be updated doesn't exist
func UpdateSoporteConocimientoIdiomaById(m *SoporteConocimientoIdiomaV2) (err error) {
	o := orm.NewOrm()
	v := SoporteConocimientoIdiomaV2{Id: m.Id}
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

// DeleteSoporteConocimientoIdioma deletes SoporteConocimientoIdioma by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSoporteConocimientoIdioma(id int) (err error) {
	o := orm.NewOrm()
	v := SoporteConocimientoIdiomaV2{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SoporteConocimientoIdiomaV2{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
