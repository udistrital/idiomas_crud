package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ConocimientoIdiomaV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:IdiomaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:IdiomaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:IdiomaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:IdiomaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:IdiomaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:NivelController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:NivelController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:NivelController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:NivelController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:NivelController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:NivelController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:NivelController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:NivelController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:NivelController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:NivelController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:SoporteConocimientoIdiomaV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ValorNivelIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ValorNivelIdiomaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ValorNivelIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ValorNivelIdiomaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ValorNivelIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ValorNivelIdiomaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ValorNivelIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ValorNivelIdiomaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ValorNivelIdiomaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/idiomas_crud/controllers:ValorNivelIdiomaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
