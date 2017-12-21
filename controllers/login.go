package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
	exitLogin:=this.Input().Get("exit")
	if exitLogin=="true" {
		this.Ctx.SetCookie("name","",-1,"/")
		this.Ctx.SetCookie("pwd","",-1,"/")
	}

}

func (this *LoginController) Post(){
	uname:=this.Input().Get("uname")
	pwd:=this.Input().Get("pwd")
	autoLogin:=this.Input().Get("autoLogin")=="on"
	if uname==beego.AppConfig.String("adminName")&&
		pwd==beego.AppConfig.String("adminPass"){
			maxAge:=0
			if autoLogin{
				maxAge=1<<31-1
			}
		this.Ctx.SetCookie("name",uname,maxAge,"/")
		this.Ctx.SetCookie("pwd",pwd,maxAge,"/")
		this.Redirect("/",301)
		return
	}

}

func CheckLogin(ctx *context.Context) bool {
ck,err:=ctx.Request.Cookie("name")
if err!=nil{
	return false
}
name:=ck.Value
cb,err2:=ctx.Request.Cookie("pwd")
	if err2!=nil{
		return false
	}
	pwd:=cb.Value
	return name==beego.AppConfig.String("adminName") && pwd==beego.AppConfig.String("adminPass")
}