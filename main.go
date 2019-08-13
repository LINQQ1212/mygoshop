package main

import (
	_ "github.com/LINQQ1212/mygoshop/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.SetLogger("file", `{"filename":"logs/test.log"}`)
		beego.BeeLogger.DelLogger("console")
	}
	//输出文件名和行号
	beego.SetLogFuncCall(true)
	//beego.BConfig.EnableGzip = true
	//beego.SetStaticPath("/download2","down2")
	beego.Run()

}

/*

日志等级
LevelEmergency
LevelAlert
LevelCritical
LevelError
LevelWarning
LevelNotice
LevelInformational
LevelDebug


EnableXSRF
启用XSRF

beego.BConfig.WebConfig.EnableXSRF = false

XSRFKEY

设置XSRF密钥。默认情况下这是beegoxsrf。

beego.BConfig.WebConfig.XSRFKEY = "beegoxsrf"

XSRFExpire

设置XSRF过期时间。默认设置为0。

beego.BConfig.WebConfig.XSRFExpire = 0

*/
