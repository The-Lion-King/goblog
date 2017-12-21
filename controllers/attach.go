package controllers

import (
	"github.com/astaxie/beego"
	"net/url"
	"os"
	"io"
)

type AttachController struct {
	beego.Controller
}

func (this *AttachController)Get(){
	url,err:=url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
	if err!=nil{
		beego.Error(err)
	}
	f,err:=os.Open(url)
	if err!=nil{
		beego.Error(err)
	}
	_,err=io.Copy(this.Ctx.ResponseWriter,f)
	if err!=nil{
		beego.Error(err)
	}
	defer f.Close()

}
