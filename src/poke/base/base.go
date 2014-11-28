package base

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Layout = `layout.html`
	this.Data[`position`] = ""
	this.Data[`subp`] = ""
	if user := this.GetSession(`login`); user == nil {
		this.Redirect(this.UrlFor(`LoginController.Login`), 302)
		return
	}
}
