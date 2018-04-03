// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"bmsg/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1/",
		beego.NSNamespace("/api/msg",
			//CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
			beego.NSNamespace("/create",
				// by FromUserId 发送一封邮件
				beego.NSRouter("/", &controllers.MessgeController{}, "post:CreateMessge"),
			),
			beego.NSNamespace("/read",
				beego.NSRouter("/", &controllers.MessgeController{}, "post:ReadMessge"),
			),
			beego.NSNamespace("/update",
				beego.NSRouter("/", &controllers.MessgeController{}, "post:UpdateMessge"),
			),
			beego.NSNamespace("/delete",
				//by messge id 删除已读信息
				beego.NSRouter("/", &controllers.MessgeController{}, "post:DeleteMessge"),
			),
		),
	)
	beego.AddNamespace(ns)
}
