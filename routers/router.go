package routers

import (
	"adConsole/controllers"

	"adConsole/controllers/api"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/init", &api.InitController{})
}
