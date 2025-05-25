package controllers

import (
	"regInfo/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "Login.html"
}

func (c *LoginController) HandleLoginPost() {
	//1.拿到数据
	Name := c.GetString("Name")
	Number := c.GetString("Number")
	println(Name, Number)
	//2.对数据进行校验
	if Name == "" || Number == "" {
		c.Ctx.WriteString("未填写完全")
		c.Ctx.WriteString("请补全信息")
		c.Redirect("/Login", 302)
		return
	}
	//3.插入数据库
	o := orm.NewOrm()

	user := models.UserTable{}
	user.Name = Name
	user.Number = Number
	flag1 := 0
	flag2 := 0
	exist := o.Read(&user, "number")
	if exist == nil {
		flag1 = 1
	}
	exist1 := o.Read(&user, "name")
	if exist1 == nil {
		flag2 = 1
	}
	if flag1 == 1 && flag2 == 1 {
		c.TplName = "User.html"
	} else {
		c.TplName = "Login.html"
	}
}

func (c *LoginController) ShowRegister() {
	c.TplName = "Check.html"
}
