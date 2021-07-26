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

// ConocimientoIdiomaController operations for ConocimientoIdioma
type ConocimientoIdiomaController struct {
	beego.Controller
}

// URLMapping ...
func (c *ConocimientoIdiomaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ConocimientoIdioma
// @Param	body		body 	models.ConocimientoIdioma	true		"body for ConocimientoIdioma content"
// @Success 201 {int} models.ConocimientoIdioma
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *ConocimientoIdiomaController) Post() {
	var v models.ConocimientoIdioma
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		//-------------- Temporal: Cambio por transición ------- //
		i := &models.Idioma{
			Id: v.Idioma.Id,
		}

		ni := &models.Nivel {
			Id: v.ClasificacionNivelIdioma.Id,
		}

		nli:= &models.ValorNivelIdioma {
			Id: v.NivelLee.Id,
		}

		nei:= &models.ValorNivelIdioma {
			Id: v.NivelEscribe.Id,
		}

		nescui:= &models.ValorNivelIdioma {
			Id: v.NivelEscucha.Id,
		}

		nhi:= &models.ValorNivelIdioma {
			Id: v.NivelHabla.Id,
		}


		
		temp:= models.ConocimientoIdiomaV2 {

			Id : v.Id,          
			TercerosId  : v.Persona,     
			Idioma  : i,        
			Nativo   : v.Nativo,        
			ClasificacionNivelIdioma  : ni        ,
			NivelLee : nli      ,
			NivelEscribe : nei   ,
			NivelEscucha  : nescui  ,
			NivelHabla   : nhi   ,
			FechaCreacion     : time.Now(),
			FechaModificacion  :time.Now(),
		}


		if _, err := models.AddConocimientoIdioma(&temp); err == nil {
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
// @Description get ConocimientoIdioma by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ConocimientoIdioma
// @Failure 404 not found resource
// @router /:id [get]
func (c *ConocimientoIdiomaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetConocimientoIdiomaById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {

		i := &models.Idioma{
			Id: v.Idioma.Id,
		    
		}

		ni := &models.Nivel {
			Id: v.ClasificacionNivelIdioma.Id,
		}

		nli:= &models.ValorNivelIdioma {
			Id: v.NivelLee.Id,
		}

		nei:= &models.ValorNivelIdioma {
			Id: v.NivelEscribe.Id,
		}

		nescui:= &models.ValorNivelIdioma {
			Id: v.NivelEscucha.Id,
		}

		nhi:= &models.ValorNivelIdioma {
			Id: v.NivelHabla.Id,
		}

		f_c := v.FechaCreacion.String()
		f_m := v.FechaModificacion.String()
		
		temp := models.ConocimientoIdioma {
			Id    : v.Id,            
			Persona  : v.TercerosId,                
			Idioma  : i,                
			NivelLee  :    nli,           
			NivelEscribe  : nei,  
			NivelEscucha  : nescui,           
			NivelHabla  : nhi,             
			Nativo : v.Nativo,              
			ClasificacionNivelIdioma : ni,
			FechaCreacion   : f_c,
			FechaModificacion : f_m,

		}

		c.Data["json"] = temp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ConocimientoIdioma
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ConocimientoIdioma
// @Failure 404 not found resource
// @router / [get]
func (c *ConocimientoIdiomaController) GetAll() {
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

	l, err := models.GetAllConocimientoIdioma(query, fields, sortby, order, offset, limit)
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
			var temp []models.ConocimientoIdioma
			for _, i := range l {
				field, _ := i.(models.ConocimientoIdiomaV2)

				i := &models.Idioma{
					Id: field.Idioma.Id,
					Nombre: field.Idioma.Nombre,  
					Descripcion :  field.Idioma.Descripcion  , 
					CodigoAbreviacion :field.Idioma.CodigoAbreviacion ,
					Activo            :field.Idioma.Activo,
					NumeroOrden :field.Idioma.NumeroOrden     , 
					FechaCreacion  : field.Idioma.FechaCreacion ,
					FechaModificacion : field.Idioma.FechaModificacion,
				}
		
				ni := &models.Nivel {
					Id: field.ClasificacionNivelIdioma.Id,
					Nombre: field.ClasificacionNivelIdioma.Nombre,  
					Descripcion :  field.ClasificacionNivelIdioma.Descripcion  , 
					CodigoAbreviacion :field.ClasificacionNivelIdioma.CodigoAbreviacion ,
					Activo            :field.ClasificacionNivelIdioma.Activo,
					NumeroOrden :field.ClasificacionNivelIdioma.NumeroOrden     , 
					FechaCreacion  : field.ClasificacionNivelIdioma.FechaCreacion ,
					FechaModificacion : field.ClasificacionNivelIdioma.FechaModificacion,

				}
		
				nli:= &models.ValorNivelIdioma {
					Id: field.NivelLee.Id,
					Nombre: field.NivelLee.Nombre,  
					Descripcion :  field.NivelLee.Descripcion  , 
					CodigoAbreviacion :field.NivelLee.CodigoAbreviacion ,
					Activo            :field.NivelLee.Activo,
					NumeroOrden :field.NivelLee.NumeroOrden     , 
					FechaCreacion  : field.NivelLee.FechaCreacion ,
					FechaModificacion : field.NivelLee.FechaModificacion,
				}
		
				nei:= &models.ValorNivelIdioma {
					Id: field.NivelEscribe.Id,
					Nombre: field.NivelEscribe.Nombre,  
					Descripcion :  field.NivelEscribe.Descripcion  , 
					CodigoAbreviacion :field.NivelEscribe.CodigoAbreviacion ,
					Activo            :field.NivelEscribe.Activo,
					NumeroOrden :field.NivelEscribe.NumeroOrden     , 
					FechaCreacion  : field.NivelEscribe.FechaCreacion ,
					FechaModificacion : field.NivelEscribe.FechaModificacion,
				}
		
				nescui:= &models.ValorNivelIdioma {
					Id: field.NivelEscucha.Id,
					Nombre: field.NivelEscucha.Nombre,  
					Descripcion :  field.NivelEscucha.Descripcion  , 
					CodigoAbreviacion :field.NivelEscucha.CodigoAbreviacion ,
					Activo            :field.NivelEscucha.Activo,
					NumeroOrden :field.NivelEscucha.NumeroOrden     , 
					FechaCreacion  : field.NivelEscucha.FechaCreacion ,
					FechaModificacion : field.NivelEscucha.FechaModificacion,
				}
		
				nhi:= &models.ValorNivelIdioma {
					Id: field.NivelHabla.Id,
					Nombre: field.NivelHabla.Nombre,  
					Descripcion :  field.NivelHabla.Descripcion  , 
					CodigoAbreviacion :field.NivelHabla.CodigoAbreviacion ,
					Activo            :field.NivelHabla.Activo,
					NumeroOrden :field.NivelHabla.NumeroOrden     , 
					FechaCreacion  : field.NivelHabla.FechaCreacion ,
					FechaModificacion : field.NivelHabla.FechaModificacion,
				}
		
				f_c := field.FechaCreacion.String()
				f_m := field.FechaModificacion.String()

				x := models.ConocimientoIdioma {
					Id    : field.Id,            
					Persona  : field.TercerosId,                
					Idioma  : i,                
					NivelLee  :    nli,           
					NivelEscribe  : nei,  
					NivelEscucha  : nescui,           
					NivelHabla  : nhi,             
					Nativo : field.Nativo,              
					ClasificacionNivelIdioma : ni,
					FechaCreacion   : f_c,
					FechaModificacion : f_m,
		
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
// @Description update the ConocimientoIdioma
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ConocimientoIdioma	true		"body for ConocimientoIdioma content"
// @Success 200 {object} models.ConocimientoIdioma
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *ConocimientoIdiomaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ConocimientoIdioma{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		//-------------- Temporal: Cambio por transición ------- //
		i := &models.Idioma{
			Id: v.Idioma.Id,
		}

		ni := &models.Nivel {
			Id: v.ClasificacionNivelIdioma.Id,
		}

		nli:= &models.ValorNivelIdioma {
			Id: v.NivelLee.Id,
		}

		nei:= &models.ValorNivelIdioma {
			Id: v.NivelEscribe.Id,
		}

		nescui:= &models.ValorNivelIdioma {
			Id: v.NivelEscucha.Id,
		}

		nhi:= &models.ValorNivelIdioma {
			Id: v.NivelHabla.Id,
		}


		layout := "2006-01-02"  //Cambiar si la hora que viene en string tiene otro formato
		f_c, _ := time.Parse(layout, v.FechaCreacion)
		
		v2:= models.ConocimientoIdiomaV2 {

			Id : v.Id,          
			TercerosId  : v.Persona,     
			Idioma  : i,        
			Nativo   : v.Nativo,        
			ClasificacionNivelIdioma  : ni        ,
			NivelLee : nli      ,
			NivelEscribe : nei   ,
			NivelEscucha  : nescui  ,
			NivelHabla   : nhi   ,
			FechaCreacion     : f_c,
			FechaModificacion  :time.Now(),
		}


		if err := models.UpdateConocimientoIdiomaById(&v2); err == nil {
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
// @Description delete the ConocimientoIdioma
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *ConocimientoIdiomaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteConocimientoIdioma(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
