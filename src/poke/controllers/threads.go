package controllers

import (
	"fmt"
	"poke/base"
	"poke/models"
	"regexp"
	"strconv"
	"strings"
)

type ThreadController struct {
	base.BaseController
}

func (this *ThreadController) Single() {
	us, err := models.GetAllThreads()
	if err != nil {
		this.Data[`error`] = err
		this.TplNames = `error.html`
		return
	}
	this.Data[`threads`] = us
	this.Data[`position`] = "thread"
	this.Data[`subp`] = "thread-single"

	this.TplNames = `threads-single.html`
}

func (this *ThreadController) Bunch() {
	this.Data[`position`] = "thread"
	this.Data[`subp`] = "thread-bunch"
	this.TplNames = "threads-bunch.html"
}
func (this *ThreadController) BunchAdd() {
	this.Data[`position`] = "thread"
	this.Data[`subp`] = "thread-bunch"
	this.TplNames = "ret.html"

	bunchs := this.GetString(`threadsbatch`)
	if strings.TrimSpace(bunchs) == `` {
		this.Data[`ret`] = `输入数据不能为空`
		return
	}
	bunch := regexp.MustCompile(`(.*)\|(.+)`).FindAllStringSubmatch(bunchs, -1)
	ret := []string{}
	for _, v := range bunch {
		if len(v) < 3 {
			ret = append(ret, fmt.Sprintf("数据不正确,缺少标题或内容: %v", v))
			continue
		}
		err := models.CreateThread(v[1], v[2])
		if err != nil {
			ret = append(ret, fmt.Sprintf("添加 %v , %v 错误: %v", v[1], v[2], err))
			continue
		}
	}

	this.Data[`ret`] = ret

}

func (this *ThreadController) UpdateThread() {
	ids := this.Ctx.Input.Param(`:id`)
	id, _ := strconv.ParseInt(ids, 10, 0)
	title := this.GetString(`title`)
	text := this.GetString(`text`)
	err := models.UpdateThread(int(id), title, text)
	if err != nil {
		this.Data[`json`] = map[string]interface{}{`error`: err}
		this.ServeJson()
		return
	}
	this.Data[`json`] = map[string]interface{}{`succ`: `ok`}
	this.ServeJson()
	return

}
func (this *ThreadController) CreateThread() {
	title := this.GetString(`title`)
	text := this.GetString(`text`)
	err := models.CreateThread(title, text)
	if err != nil {
		this.Data[`json`] = map[string]interface{}{`error`: err}
		this.ServeJson()
		return
	}
	this.Data[`json`] = map[string]interface{}{`succ`: `ok`}
	this.ServeJson()
	return
}
func (this *ThreadController) DeleteThread() {
	ids := this.Ctx.Input.Param(`:id`)
	id, _ := strconv.ParseInt(ids, 10, 0)
	err := models.DeleteThread(int(id))
	if err != nil {
		this.Data[`json`] = map[string]interface{}{`error`: err}
		this.ServeJson()
		return
	}
	this.Data[`json`] = map[string]interface{}{`succ`: `ok`}
	this.ServeJson()
	return
}
