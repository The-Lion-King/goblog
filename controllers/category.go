package controllers

import (
	"code/2/blog/models"
	"github.com/astaxie/beego"

)


type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get(){
	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return
	}

	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	this.Data["IsLogin"] = CheckLogin(this.Ctx)
	name,errCookie:=this.Ctx.Request.Cookie("name")
	if errCookie==nil{
		this.Data["Name"]=name.Value
	}
	var err error
	this.Data["Categories"] , err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

}
