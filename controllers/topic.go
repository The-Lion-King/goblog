package controllers

import (
	"github.com/astaxie/beego"
	"code/2/blog/models"

	"strings"
	"path"
)

type TopicController struct {
	beego.Controller
}


func (this *TopicController) Get(){
	this.Data["IsTopic"] = true
	this.TplName="topic.html"
	this.Data["IsLogin"]=CheckLogin(this.Ctx)
	this.Data["Name"]=beego.AppConfig.String("adminName")
	topics,err:=models.GetAllTopic("","")
	if err!=nil{
		beego.Error(err)
	}
	this.Data["Topics"]=topics


}
func (this *TopicController)Post(){
	if !CheckLogin(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	// 解析表单
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category:=this.Input().Get("category")
	labels:=this.Input().Get("labels")
	//attachment:=this.Input().Get("attachment")
	var err error
	_,fn,err:=this.GetFile("attachment")

	if err != nil {
		beego.Error(err)
	}
	var attachment string
	if fn!=nil{
		attachment=fn.Filename
		err=this.SaveToFile("attachment",path.Join("attachment",attachment))
		if err!=nil{
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err = models.AddTopic(title, content,category,labels,attachment)
	} else {
		err = models.ModifyTopic(tid, title, content,category,labels,attachment)
	}

	this.Redirect("/topic", 302)
}
func (this *TopicController)Add(){
	err:=CheckLogin(this.Ctx)
	if !err{
		this.Redirect("/login",302)
		return
	}
	this.TplName="topic_add.html"
}


func (this *TopicController)View(){

	this.TplName="topic_view.html"
	tid:=this.Ctx.Input.Param("0")
	topic,err:=models.GetOneTopic(tid)
	replies,err:=models.GetAllComments(tid)
	if err!=nil{
		beego.Error(err)
	}
	this.Data["Labels"] = strings.Split(topic.Lables, " ")
	this.Data["Name"]=beego.AppConfig.String("adminName")
	this.Data["Replies"]=replies
	this.Data["IsLogin"]=CheckLogin(this.Ctx)
	this.Data["Topic"]=topic
}

func (this *TopicController) Delete(){
	tid:=this.Input().Get("tid")
	topic,err:=models.GetOneTopic(tid)
	err=models.DeleteTopic(tid,topic.Category)

	if err!=nil{
		beego.Error(err)
	}
	this.Redirect("/topic",302)

}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetOneTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}