package routers

import (
	"github.com/LINQQ1212/mygoshop/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
