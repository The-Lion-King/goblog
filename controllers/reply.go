package controllers

import (
	"code/2/blog/models"
	"github.com/astaxie/beego")

type ReplyController struct {
	beego.Controller
}




func (this *ReplyController) Add(){
	nickname:=this.Input().Get("nickname")
	content:=this.Input().Get("content")
	tid:=this.Input().Get("tid")
	err:=models.AddComment(nickname,tid,content)
	if err!=nil{
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)

}

func (this *ReplyController) Delete(){
	if !CheckLogin(this.Ctx){
		return
	}
	tid:=this.Input().Get("tid")
	id:=this.Input().Get("rid")
	err:=models.DeleteComment(tid,id)
	if err!=nil{
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}