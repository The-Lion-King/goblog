package models

import (
	"time"
	"os"
	"path"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"

	//"strings"

)

const(
	_DB_NAME="data/blog.db"
	_SQLITE3_DRIVER="sqlite3"
)

type Category struct{
	Id  int64
	Title string
	Created time.Time 	`orm:"index"`
	Views int64 		`orm:"index"`
	TopicTime time.Time `orm:"index"`
	TopicCount int64
	TopicLastUserId int64
}

type Topic struct {
	Id int64
	Uid int64
	Title string
	Lables string
	Category   string
	Content string 		`orm:"size(5000)"`
	Attachment string
	Createed time.Time 	`orm:"index"`
	Updated time.Time 	`orm:"index"`
	Views int64 		`orm:"index"`
	Author string
	ReplayTime time.Time `orm:"index"`
	ReplyCount int64
	ReplayLastUserId int64

}
type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

func isExists(path string) bool{
	_,err:=os.Stat(path)
	if err==nil{
		return true
	}
	return  os.IsExist(err)
}

func RegisterDB(){
	if !isExists(_DB_NAME){
		os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category),new(Topic),new(Comment))
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)
}

func AddCategory(name string) error {
	o:=orm.NewOrm()

	cate:=&Category{
		Title: name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}
	qs:=o.QueryTable("category")
	err:=qs.Filter("title",name).One(cate)
	if err==nil{
		return err
	}
	_,err=o.Insert(cate)
	qs=o.QueryTable("category")
	cateOne:=new(Category)
	err=qs.Filter("title",name).One(cateOne)
	qs1:=o.QueryTable("topic")
	count,err:=qs1.Filter("category",name).Count()
	if err==nil{
		cateOne.TopicCount=count
		o.Update(cateOne)
	}
	if err!=nil{
		return err
	}
	return nil
}
func DeleteCategory(id string)error{
	cid,err:=strconv.ParseInt(id,10,64)
	if err!=nil{
		return err
	}
	o:=orm.NewOrm()
	cate:=Category{Id:cid}
	_,err=o.Delete(&cate)
	return err
}
func GetAllCategories()([]*Category,error){
	o:=orm.NewOrm()
	cates:=make([]*Category,0)
	qs:=o.QueryTable("category")


	_,err:=qs.All(&cates)
	return cates,err

}

func changeStringToInt(str string) (int64,error){
	num,err:=strconv.ParseInt(str,10,64)
	return num,err
}
func AddTopic(title,content,category,lables,attachment string) error {
	//lables = "$" + strings.Join(strings.Split(lables, " "), "#$") + "#"
	o := orm.NewOrm()

	topic := &Topic{
		Title:   title,
		Content: content,
		Lables:lables,
		Attachment:attachment,
		Createed: time.Now(),
		Category:category,
		Updated: time.Now(),
		ReplayTime:time.Now(),
	}
	_, err := o.Insert(topic)
	qs:=o.QueryTable("category")
	cate:=new(Category)
	err=qs.Filter("title",category).One(cate)
	qs1:=o.QueryTable("topic")
	count,err:=qs1.Filter("category",category).Count()
	if err==nil{
		cate.TopicCount=count
		o.Update(cate)
	}
	return err
}

func AddComment(nickname,tid,content string) error {
	tidNum,err:=changeStringToInt(tid)
	if err!=nil{
		return err
	}
	o:=orm.NewOrm()
	comment:=&Comment{
		Tid:tidNum,
		Name:nickname,
		Content:content,
		Created: time.Now(),
	}
	_,err=o.Insert(comment)

	qs:=o.QueryTable("comment")
	topic := &Topic{Id: tidNum}
	count,_:=qs.Filter("tid",tidNum).Count()
	if o.Read(topic) == nil {
		topic.ReplyCount=count
		o.Update(topic)
	}
		return err

}
func GetAllComments(tid string) ([]*Comment,error){
	tidNum,err:=changeStringToInt(tid)
	o:=orm.NewOrm()
	comments:=make([]*Comment,0)
	qs:=o.QueryTable("comment")
	_,err=qs.Filter("tid",tidNum).All(&comments)
	return comments,err
}

func DeleteComment(tid,id string)error{
	idNum,err:=changeStringToInt(id)
	tidNum,err1:=changeStringToInt(tid)
	if err1!=nil{
		return err1
	}
	if err!=nil{
		return err
	}
	o:=orm.NewOrm()
	comment:=&Comment{
		Tid:tidNum,
		Id:idNum,
	}

		_,err=o.Delete(comment)
	qs:=o.QueryTable("comment")
	topic := &Topic{Id: tidNum}
	count,_:=qs.Filter("tid",tidNum).Count()
	if o.Read(topic) == nil {
		topic.ReplyCount=count
		o.Update(topic)
	}
	return err
}
func GetOneTopic(tid string) (*Topic,error){
	tidNum,err:=changeStringToInt(tid)
	if err!=nil{
		return nil,err
	}
	topic:=new(Topic)
	o:=orm.NewOrm()
	qs:=o.QueryTable("topic")
	err=qs.Filter("id",tidNum).One(topic)
	topic.Views++
	_,err=o.Update(topic)
	return topic,err

}

func ModifyTopic(tid,title,content,category,labels,attachment string)error{
	//labels = "$" + strings.Join(strings.Split(labels, " "), "#$") + "#"
	o:=orm.NewOrm()

	tidNum,err:=changeStringToInt(tid)
	if err!=nil{
		return err
	}
	topic := &Topic{Id: tidNum}
	var oldAttach string
	if o.Read(topic) == nil {
		oldAttach=topic.Attachment
		topic.Lables=labels
		topic.Attachment=attachment
		topic.Title = title
		topic.Content = content
		topic.Category=category
		topic.Updated = time.Now()
		o.Update(topic)
	}
	qs:=o.QueryTable("category")
	cate:=new(Category)
	err=qs.Filter("title",category).One(cate)
	qs1:=o.QueryTable("topic")
	count,err:=qs1.Filter("category",category).Count()
	if err==nil{
		cate.TopicCount=count
		o.Update(cate)
	}
	if len(oldAttach)>0{
		os.Remove(path.Join("attachment",oldAttach))
	}
		return err
}

func DeleteTopic(tid,category string)error{
	o:=orm.NewOrm()
	tidNum,err:=changeStringToInt(tid)
	if err!=nil{
		return err
	}
	qs:=o.QueryTable("Topic")
	_,err=qs.Filter("id",tidNum).Delete()

	qs=o.QueryTable("category")
	cate:=new(Category)
	err=qs.Filter("title",category).One(cate)
	qs1:=o.QueryTable("topic")
	count,err:=qs1.Filter("category",category).Count()
	if err==nil{
		cate.TopicCount=count
		o.Update(cate)
	}
	return err

}
func GetCommentCount(tid int64)(int64,error){
	tidNum:=tid
	o:=orm.NewOrm()
	qs:=o.QueryTable("comment")
	topic := &Topic{Id: tidNum}
	count,err:=qs.Filter("tid",tidNum).Count()
	if o.Read(topic) == nil {
		topic.ReplyCount=count
		o.Update(topic)
	}
	return count,err

}
func GetAllTopic(cate,lable string)([]*Topic,error){
	o:=orm.NewOrm()
	topics:=make([]*Topic,0)
	qs:=o.QueryTable("topic")
	if len(cate)>0{
		qs=qs.Filter("category",cate)

	}
	if len(lable) > 0 {
		qs = qs.Filter("lables__contains", lable)
	}
	_,err:=qs.All(&topics)
	return topics,err
}

