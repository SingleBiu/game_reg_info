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

func GetUserByNumber(number string) (user *models.UserTable, err error) {
	o := orm.NewOrm()
	user = &models.UserTable{Number: number}
	err = o.Read(user)
	return user, err
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
	o := orm.NewOrm()

	user := models.UserTable{Name: Name, Number: Number}
	flag1 := 0
	exist := o.Raw("SELECT * FROM user_table WHERE Name = ? AND Number = ?", user.Name, user.Number).QueryRow(&user)
	if exist == nil {
		flag1 = 1
	}
	if flag1 == 1 {
		c.TplName = "User.html"
		c.Data["qqName"] = user.Name
		c.Data["qqNumber"] = user.Number
		c.Data["ysUID"] = user.UID_YUAN
		c.Data["btUID"] = user.UID_BENG
		c.Data["zzzUID"] = user.UID_JUE

		//全部展示 不分页
		var users []models.UserTable
		_, err := o.QueryTable("user_table").All(&users)
		if err != nil {
			c.Abort("500")
			return
		}
		c.Data["users"] = users

	} else {
		c.TplName = "Login.html"
	}
}

func (c *LoginController) ShowRegister() {
	c.TplName = "Check.html"
}
