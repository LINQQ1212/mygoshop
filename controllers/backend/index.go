package backend

// IndexController 后端类
type IndexController struct {
	baseController
}

// Index 后端首页
func (c *IndexController) Index() {
	c.Data["Title"] = "c.userid"
	c.Data["Body"] = "hello 后端首页"
	c.TplName = c.moduleName + "/index/index.html"

}

func (c *IndexController) Test() {
	c.Ctx.Output.Body([]byte("456"))
}

func (c *IndexController) Hello() {
	c.Ctx.Output.Body([]byte("456"))
}
