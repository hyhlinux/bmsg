package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bmsg/controllers:MessgeController"] = append(beego.GlobalControllerRouter["bmsg/controllers:MessgeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bmsg/controllers:MessgeController"] = append(beego.GlobalControllerRouter["bmsg/controllers:MessgeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bmsg/controllers:MessgeController"] = append(beego.GlobalControllerRouter["bmsg/controllers:MessgeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:Id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bmsg/controllers:MessgeController"] = append(beego.GlobalControllerRouter["bmsg/controllers:MessgeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bmsg/controllers:MessgeController"] = append(beego.GlobalControllerRouter["bmsg/controllers:MessgeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
