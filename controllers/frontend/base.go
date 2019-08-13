package frontend

import (
	"os"
	"strings"

	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
	options        map[string]string
}

// Prepare beego用户扩展
func (base *baseController) Prepare() {
	controllerName, actionName := base.GetControllerAndAction()
	base.moduleName = "frontend"
	base.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	base.actionName = strings.ToLower(actionName)
	base.Data["options"] = base.options
}

func (base *baseController) display(tpl string) {
	var theme string
	if v, ok := base.options["theme"]; ok && v != "" {
		theme = v
	} else {
		theme = "default"
	}
	if _, err := os.Stat(beego.BConfig.WebConfig.ViewsPath + "/" + theme + "/layout.html"); err == nil {
		base.Layout = theme + "/layout.html"
	}
	base.TplName = theme + "/" + tpl + ".html"
}
