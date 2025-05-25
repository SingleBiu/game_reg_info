/*
 * @Author: SingleBiu
 * @Date: 2025-05-23 14:18:57
 * @LastEditors: SingleBiu
 * @LastEditTime: 2025-05-25 15:28:44
 * @Description: file content
 */
package routers

import (
	"regInfo/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.UsrController{})
	beego.Router("/Check", &controllers.UsrController{}, "get:Get;post:HandleUsrPost")
	beego.Router("/login", &controllers.LoginController{}, "get:Get;post:HandleLoginPost")
	beego.Router("/Success", &controllers.LoginController{}, "get:ShowRegister")
}
