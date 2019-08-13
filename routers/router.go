package routers

import (
	"github.com/LINQQ1212/mygoshop/controllers"
	"github.com/LINQQ1212/mygoshop/controllers/backend"
	"github.com/LINQQ1212/mygoshop/controllers/frontend"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	var FilterUser = func(ctx *context.Context) {

	}
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

	// error
	beego.ErrorController(&controllers.ErrorController{})

	//前端
	beego.Router("/", &frontend.IndexController{}, "*:Index")

	//后端
	admin := backend.ModuleName

	beego.NewNamespace("/"+admin,
		beego.NSRouter("/test", &backend.IndexController{}, "*:Test"),
	)

	beego.Router("/"+admin, &backend.IndexController{}, "*:Index")
	beego.Router("/"+admin, &backend.IndexController{}, "*:Index")

}
