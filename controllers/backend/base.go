package backend

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

// ModuleName 后台模块名 和 模板目录
const ModuleName = "backend"

//  github.com/allegro/bigcache

type baseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
	//cache          *util.LruCache
}

func (base *baseController) Prepare() {
	controllerName, actionName := base.GetControllerAndAction()
	base.moduleName = ModuleName
	base.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	base.actionName = strings.ToLower(actionName)
	base.auth()
	base.checkPermission()
}

//登录状态验证
func (base *baseController) auth() {

}

//权限验证
func (base *baseController) checkPermission() {

}

//渲染模版
func (base *baseController) display(tpl ...string) {
	fmt.Println(base.Ctx.Request.URL.String())
	var tplname string
	if len(tpl) == 1 {
		tplname = base.moduleName + "/" + tpl[0] + ".html"
	} else {
		tplname = base.moduleName + "/" + base.controllerName + "/" + base.actionName + ".html"
	}

	fmt.Println(tpl)
	//this.Layout = this.moduleName + "/layout.html"
	base.TplName = tplname
}

//显示错误提示
func (base *baseController) showmsg(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, base.Ctx.Request.Referer())
	}
	base.Data["msg"] = msg[0]
	base.Data["redirect"] = msg[1]
	//this.Layout = this.moduleName + "/layout.html"
	base.TplName = base.moduleName + "/" + "showmsg.html"
	base.Render()
	base.StopRun()
}

//获取用户IP地址
func (base *baseController) getClientIP() string {
	s := strings.Split(base.Ctx.Request.RemoteAddr, ":")

	return s[0]
}

/*
func (this *baseController) getTime() time.Time {

	timezone, _ := strconv.ParseFloat(option.Get("timezone"), 64)
	add := timezone * float64(time.Hour)
	return time.Now().UTC().Add(time.Duration(add))
}
*/
