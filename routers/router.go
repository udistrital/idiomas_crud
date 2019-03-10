// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/idiomas_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/clasificacion_nivel_idioma",
			beego.NSInclude(
				&controllers.ClasificacionNivelIdiomaController{},
			),
		),

		beego.NSNamespace("/conocimiento_idioma",
			beego.NSInclude(
				&controllers.ConocimientoIdiomaController{},
			),
		),

		beego.NSNamespace("/idioma",
			beego.NSInclude(
				&controllers.IdiomaController{},
			),
		),

		beego.NSNamespace("/soporte_conocimiento_idioma",
			beego.NSInclude(
				&controllers.SoporteConocimientoIdiomaController{},
			),
		),

		beego.NSNamespace("/valor_nivel_idioma",
			beego.NSInclude(
				&controllers.ValorNivelIdiomaController{},
			),
		),

		beego.NSNamespace("/valor_nota",
			beego.NSInclude(
				&controllers.ValorNotaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
