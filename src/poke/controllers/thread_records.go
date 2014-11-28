package controllers

import (
	"github.com/astaxie/beego"
	"poke/base"
	"poke/models"
	"strings"
	"time"
)

type ThreadRecords struct {
	base.BaseController
}

func (this *ThreadRecords) Prepare() {
	this.BaseController.Prepare()
	this.Data[`subtitle`] = "回帖纪录"
	this.Data[`position`] = "record"
}

func (this *ThreadRecords) Records() {
	count, _ := models.GetThreadRecordsCounts()
	paginator := NewPaginator(this.Ctx.Request, 10, int(count))

	re, err := models.GetThreadRecordsRange(paginator.Offset(), paginator.PerPageNums)
	if err != nil {
		beego.Warn(`get page range error`, err)
		this.Abort("500")
	}

	succn, err := models.GetThreadRecordsCountSuccOrNot(true)
	if err != nil {
		beego.Warn(`get succ count error`, err)
	}
	failn, err := models.GetThreadRecordsCountSuccOrNot(false)
	if err != nil {
		beego.Warn(`get fail count error`, err)
	}
	this.Data[`paginator`] = paginator
	this.Data[`succ`] = succn
	this.Data[`fail`] = failn
	this.Data[`records`] = re

	this.TplNames = `threadrecords.html`
}

func (this *ThreadRecords) Search() {
	query_typ := []string{}
	query_val := []interface{}{}

	result, _ := this.GetInt("result")
	if result == 1 || result == -1 {
		query_typ = append(query_typ, `succ`)
		if result == 1 {
			query_val = append(query_val, true)
		} else {
			query_val = append(query_val, false)
		}

	}

	timestring := this.GetString("time")
	timestrings := strings.Split(timestring, "~")
	if len(timestrings) == 2 {
		var tf time.Time
		var tt time.Time
		tf, _ = time.Parse("2006-01-02", timestrings[0])
		tt, _ = time.Parse("2006-01-02", timestrings[1])

		query_typ = append(query_typ, `time__gte`)
		query_val = append(query_val, tf)

		query_typ = append(query_typ, `time__lte`)
		query_val = append(query_val, tt)
	}

	username := this.GetString("username")
	if u := strings.TrimSpace(username); u != `` {
		query_typ = append(query_typ, `username`)
		query_val = append(query_val, u)
	}

	place := this.GetString("place")
	if p := strings.TrimSpace(place); p != `` {
		query_typ = append(query_typ, `ret__contains`)
		query_val = append(query_val, p)
	}

	title := this.GetString(`title`)
	if t := strings.TrimSpace(title); t != `` {
		query_typ = append(query_typ, "title__contains")
		query_val = append(query_val, t)
	}

	re, err := models.SearchThreadRecord(query_typ, query_val)
	if err != nil {
		beego.Critical("search record fail:", err)
		this.Abort(`500`)
	}

	count := len(re)
	paginator := NewPaginator(this.Ctx.Request, 10, count)

	succn, failn := 0, 0
	for _, v := range re {
		if v.Succ {
			succn += 1
		} else {
			failn += 1
		}
	}

	this.Data[`paginator`] = paginator
	this.Data[`succ`] = succn
	this.Data[`fail`] = failn
	this.Data[`records`] = re[paginator.Offset():func() int {
		if paginator.Offset()+paginator.PerPageNums > len(re) {
			return len(re)
		} else {
			return paginator.Offset() + paginator.PerPageNums
		}
	}()]
	this.TplNames = `threadrecords.html`

	// this.Redirect(this.UrlFor("this.Records"), 302)
}
