package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) showVer() {
	c.Data["version"] = beego.AppConfig.String("version")
	c.Data["appname"] = beego.BConfig.AppName
}

func (c *ErrorController) Error404() {

	c.Data["title"] = "404 Not Found"
	c.Data["content"] = "page not found"
	c.showVer()
	c.TplName = "error/default.html"
}

func (c *ErrorController) Error501() {
	c.Data["title"] = "501 Server Error"
	c.Data["content"] = "server error"
	c.showVer()
	c.TplName = "error/default.html"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "default.html"
}
