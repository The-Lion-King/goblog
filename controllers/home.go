package controllers

import (
	"github.com/astaxie/beego"

	"code/2/blog/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "home.html"
	name,err:=this.Ctx.Request.Cookie("name")
	if err==nil{
		this.Data["Name"]=name.Value
	}
	cate:=this.Input().Get("cate")
	topics,err:=models.GetAllTopic(cate,this.Input().Get("label"))
	if err!=nil{
		beego.Error(err)
	}
	this.Data["Topic"]=topics
	for _,v:=range topics{
		_,err:=models.GetCommentCount(v.Id)
		if err!=nil{
			beego.Error(err)
		}
	}
	this.Data["Categories"] , err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	this.Data["IsHome"]=true
	this.Data["IsLogin"]=CheckLogin(this.Ctx)
}
