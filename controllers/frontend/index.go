package frontend

import "github.com/astaxie/beego"

// IndexController 前端类
type IndexController struct {
	beego.Controller
}

// Index 前端首页
func (c *IndexController) Index() {
	c.Ctx.WriteString("hello frontend")
}
