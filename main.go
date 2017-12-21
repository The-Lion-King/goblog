package main

import (
	_ "code/2/blog/routers"
	"code/2/blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
)

func init(){
	models.RegisterDB()
}
func isExists(path string) bool{
	_,err:=os.Stat(path)

	return os.IsExist(err)
}
func main() {

	orm.RunSyncdb("default",false,true)
	if !isExists("attachment"){
		os.Mkdir("attachment",os.ModePerm)
	}
	beego.Run()

}

