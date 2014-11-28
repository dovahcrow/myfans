package controllers

import (
	"fmt"
	"poke/base"
	"poke/models"
	"regexp"
	"strconv"
	"strings"
)

type UserController struct {
	base.BaseController
}

func (this *UserController) Single() {
	us, err := models.GetAllUsers()
	if err != nil {
		this.Data[`error`] = err
		this.TplNames = `error.html`
		return
	}

	this.Data[`users`] = us
	this.Data[`position`] = "user"
	this.Data[`subp`] = "user-single"
	this.TplNames = `users-single.html`
}
func (this *UserController) Bunch() {
	this.Data[`position`] = "user"
	this.Data[`subp`] = "user-bunch"
	this.TplNames = `users-bunch.html`
}

func (this *UserController) BunchAdd() {
	this.Data[`position`] = "user"
	this.Data[`subp`] = "user-bunch"
	this.TplNames = "ret.html"

	bunchs := this.GetString(`usersbatch`)
	if strings.TrimSpace(bunchs) == `` {
		this.Data[`ret`] = `输入数据不能为空`
		return
	}
	bunch := regexp.MustCompile(`(.*)\|(.+)`).FindAllStringSubmatch(bunchs, -1)

	ret := []string{}
	for _, v := range bunch {
		if len(v) < 3 {
			ret = append(ret, fmt.Sprintf("数据不正确,缺少用户名或密码: %v", v))
			continue
		}
		err := models.CreateUser(v[1], v[2])
		if err != nil {
			ret = append(ret, fmt.Sprintf("添加 %v , %v 错误: %v", v[1], v[2], err))
			continue
		}
	}
	this.Data[`ret`] = ret
}

func (this *UserController) UpdateUser() {
	defer this.ServeJson()
	ids := this.Ctx.Input.Param(`:id`)
	id, _ := strconv.ParseInt(ids, 10, 0)
	usnm := this.GetString(`username`)
	pswd := this.GetString(`password`)
	fmt.Println(pswd, usnm)
	if strings.TrimSpace(usnm) == `` || strings.TrimSpace(pswd) == `` {
		this.Data[`ret`] = `输入数据不能为空`
		return
	}

	err := models.UpdateUser(int(id), usnm, pswd)
	if err != nil {
		this.Data[`json`] = map[string]interface{}{`error`: err}
		this.ServeJson()
		return
	}
	this.Data[`json`] = map[string]interface{}{`succ`: `ok`}

	return

}
func (this *UserController) CreateUser() {
	usnm := this.GetString(`username`)
	pswd := this.GetString(`password`)
	err := models.CreateUser(usnm, pswd)
	if err != nil {
		this.Data[`json`] = map[string]interface{}{`error`: err}
		this.ServeJson()
		return
	}
	this.Data[`json`] = map[string]interface{}{`succ`: `ok`}
	this.ServeJson()
	return
}
func (this *UserController) DeleteUser() {
	ids := this.Ctx.Input.Param(`:id`)
	id, _ := strconv.ParseInt(ids, 10, 0)
	err := models.DeleteUser(int(id))
	if err != nil {
		this.Data[`json`] = map[string]interface{}{`error`: err}
		this.ServeJson()
		return
	}
	this.Data[`json`] = map[string]interface{}{`succ`: `ok`}
	this.ServeJson()
	return
}
