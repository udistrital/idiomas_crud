package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
	"github.com/udistrital/idiomas_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// SoporteConocimientoIdiomaController operations for SoporteConocimientoIdioma
type SoporteConocimientoIdiomaController struct {
	beego.Controller
}

// URLMapping ...
func (c *SoporteConocimientoIdiomaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create SoporteConocimientoIdioma
// @Param	body		body 	models.SoporteConocimientoIdioma	true		"body for SoporteConocimientoIdioma content"
// @Success 201 {int} models.SoporteConocimientoIdioma
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *SoporteConocimientoIdiomaController) Post() {
	var v models.SoporteConocimientoIdioma

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//-------------- Temporal: Cambio por transición ------- //
		ci := &models.ConocimientoIdiomaV2 {
			Id: v.ConocimientoIdioma.Id,
		}
		
		layout := "2006-01-02"  //Cambiar si la hora que viene en string tiene otro formato
		f_c, _ := time.Parse(layout, v.FechaCreacion)

		temp := models.SoporteConocimientoIdiomaV2 {
			Id  :   v.Id,          
			Descripcion:  v.Descripcion,    
			FechaCreacion :  f_c,   
			FechaModificacion :  time.Now(),	
			TercerosId: v.Institucion, 	
			DocumentoId: v.Documento,
			ConocimientoIdiomaId: ci,
		}

		if _, err := models.AddSoporteConocimientoIdioma(&temp); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get SoporteConocimientoIdioma by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SoporteConocimientoIdioma
// @Failure 404 not found resource
// @router /:id [get]
func (c *SoporteConocimientoIdiomaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSoporteConocimientoIdiomaById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {

		ci := &models.ConocimientoIdioma {
			Id: v.ConocimientoIdiomaId.Id,
			//NIVELES

		}

		f_c := v.FechaCreacion.String()
		f_m := v.FechaModificacion.String()

		temp := models.SoporteConocimientoIdioma {
			Id  :   v.Id,          
			Descripcion:  v.Descripcion,    
			FechaCreacion : f_c,   
			FechaModificacion :  f_m,	
			Institucion: v.TercerosId, 	
			Documento: v.DocumentoId,
			ConocimientoIdioma: ci,

		}
	
		c.Data["json"] = temp
		//-------------- Temporal: Cambio por transición ------- //
		//c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get SoporteConocimientoIdioma
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SoporteConocimientoIdioma
// @Failure 404 not found resource
// @router / [get]
func (c *SoporteConocimientoIdiomaController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllSoporteConocimientoIdioma(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
			c.Data["json"] = l
		}else{
			//-------------- Temporal: Cambio por transición ------- //
			var temp []models.SoporteConocimientoIdioma
			for _, i := range l {
				field, _ := i.(models.SoporteConocimientoIdiomaV2)

				i := &models.Idioma{
					Id: field.ConocimientoIdiomaId.IdiomaId.Id,
					Nombre: field.ConocimientoIdiomaId.IdiomaId.Nombre,  
					Descripcion :  field.ConocimientoIdiomaId.IdiomaId.Descripcion  , 
					CodigoAbreviacion :field.ConocimientoIdiomaId.IdiomaId.CodigoAbreviacion ,
					Activo            :field.ConocimientoIdiomaId.IdiomaId.Activo,
					NumeroOrden :field.ConocimientoIdiomaId.IdiomaId.NumeroOrden     , 
					FechaCreacion  : field.ConocimientoIdiomaId.IdiomaId.FechaCreacion ,
					FechaModificacion : field.ConocimientoIdiomaId.IdiomaId.FechaModificacion,
				}
		
				ni := &models.Nivel {
					Id: field.ConocimientoIdiomaId.NivelId.Id,
					Nombre: field.ConocimientoIdiomaId.NivelId.Nombre,  
					Descripcion :  field.ConocimientoIdiomaId.NivelId.Descripcion  , 
					CodigoAbreviacion :field.ConocimientoIdiomaId.NivelId.CodigoAbreviacion ,
					Activo            :field.ConocimientoIdiomaId.NivelId.Activo,
					NumeroOrden :field.ConocimientoIdiomaId.NivelId.NumeroOrden     , 
					FechaCreacion  : field.ConocimientoIdiomaId.NivelId.FechaCreacion ,
					FechaModificacion : field.ConocimientoIdiomaId.NivelId.FechaModificacion,

				}
		
				nli:= &models.ValorNivelIdioma {
					Id: field.ConocimientoIdiomaId.NivelLeeId.Id,
					Nombre: field.ConocimientoIdiomaId.NivelLeeId.Nombre,  
					Descripcion :  field.ConocimientoIdiomaId.NivelLeeId.Descripcion  , 
					CodigoAbreviacion :field.ConocimientoIdiomaId.NivelLeeId.CodigoAbreviacion ,
					Activo            :field.ConocimientoIdiomaId.NivelLeeId.Activo,
					NumeroOrden :field.ConocimientoIdiomaId.NivelLeeId.NumeroOrden     , 
					FechaCreacion  : field.ConocimientoIdiomaId.NivelLeeId.FechaCreacion ,
					FechaModificacion : field.ConocimientoIdiomaId.NivelLeeId.FechaModificacion,
				}
		
				nei:= &models.ValorNivelIdioma {
					Id: field.ConocimientoIdiomaId.NivelEscribeId.Id,
					Nombre: field.ConocimientoIdiomaId.NivelEscribeId.Nombre,  
					Descripcion :  field.ConocimientoIdiomaId.NivelEscribeId.Descripcion  , 
					CodigoAbreviacion :field.ConocimientoIdiomaId.NivelEscribeId.CodigoAbreviacion ,
					Activo            :field.ConocimientoIdiomaId.NivelEscribeId.Activo,
					NumeroOrden :field.ConocimientoIdiomaId.NivelEscribeId.NumeroOrden     , 
					FechaCreacion  : field.ConocimientoIdiomaId.NivelEscribeId.FechaCreacion ,
					FechaModificacion : field.ConocimientoIdiomaId.NivelEscribeId.FechaModificacion,
				}
		
				nescui:= &models.ValorNivelIdioma {
					Id: field.ConocimientoIdiomaId.NivelEscuchaId.Id,
					Nombre: field.ConocimientoIdiomaId.NivelEscuchaId.Nombre,  
					Descripcion :  field.ConocimientoIdiomaId.NivelEscuchaId.Descripcion  , 
					CodigoAbreviacion :field.ConocimientoIdiomaId.NivelEscuchaId.CodigoAbreviacion ,
					Activo            :field.ConocimientoIdiomaId.NivelEscuchaId.Activo,
					NumeroOrden :field.ConocimientoIdiomaId.NivelEscuchaId.NumeroOrden     , 
					FechaCreacion  : field.ConocimientoIdiomaId.NivelEscuchaId.FechaCreacion ,
					FechaModificacion : field.ConocimientoIdiomaId.NivelEscuchaId.FechaModificacion,
				}
		
				nhi:= &models.ValorNivelIdioma {
					Id: field.ConocimientoIdiomaId.NivelHablaId.Id,
					Nombre: field.ConocimientoIdiomaId.NivelHablaId.Nombre,  
					Descripcion :  field.ConocimientoIdiomaId.NivelHablaId.Descripcion  , 
					CodigoAbreviacion :field.ConocimientoIdiomaId.NivelHablaId.CodigoAbreviacion ,
					Activo            :field.ConocimientoIdiomaId.NivelHablaId.Activo,
					NumeroOrden :field.ConocimientoIdiomaId.NivelHablaId.NumeroOrden     , 
					FechaCreacion  : field.ConocimientoIdiomaId.NivelHablaId.FechaCreacion ,
					FechaModificacion : field.ConocimientoIdiomaId.NivelHablaId.FechaModificacion,
				}

				f_ci := field.ConocimientoIdiomaId.FechaCreacion.String()
				f_mi := field.ConocimientoIdiomaId.FechaModificacion.String()

				ci := &models.ConocimientoIdioma {
					Id    : field.ConocimientoIdiomaId.Id,            
					Persona  : field.ConocimientoIdiomaId.TercerosId,                
					Idioma  : i,                
					NivelLee  :    nli,           
					NivelEscribe  : nei,  
					NivelEscucha  : nescui,           
					NivelHabla  : nhi,             
					Nativo : field.ConocimientoIdiomaId.Nativo,              
					ClasificacionNivelIdioma : ni,
					FechaCreacion   : f_ci,
					FechaModificacion : f_mi,
		
				}
		
				f_c := field.FechaCreacion.String()
				f_m := field.FechaModificacion.String()

				x := models.SoporteConocimientoIdioma {
					Id  :   field.Id,          
					Descripcion:  field.Descripcion,    
					FechaCreacion : f_c,   
					FechaModificacion :  f_m,	
					Institucion: field.TercerosId, 	
					Documento: field.DocumentoId,
					ConocimientoIdioma: ci,
		
				}

				temp = append(temp,x)
			}
			c.Data["json"] = temp
		}
	
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the SoporteConocimientoIdioma
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SoporteConocimientoIdioma	true		"body for SoporteConocimientoIdioma content"
// @Success 200 {object} models.SoporteConocimientoIdioma
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *SoporteConocimientoIdiomaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SoporteConocimientoIdioma{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		ci := &models.ConocimientoIdiomaV2 {
			Id: v.ConocimientoIdioma.Id,
		}
		
		layout := "2006-01-02"  //Cambiar si la hora que viene en string tiene otro formato
		f_c, _ := time.Parse(layout, v.FechaCreacion)

		v2 := models.SoporteConocimientoIdiomaV2 {
			Id  :   v.Id,          
			Descripcion:  v.Descripcion,    
			FechaCreacion :  f_c,   
			FechaModificacion :  time.Now(),	
			TercerosId: v.Institucion, 	
			DocumentoId: v.Documento,
			ConocimientoIdiomaId: ci,
		}

		if err := models.UpdateSoporteConocimientoIdiomaById(&v2); err == nil {
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the SoporteConocimientoIdioma
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *SoporteConocimientoIdiomaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSoporteConocimientoIdioma(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
