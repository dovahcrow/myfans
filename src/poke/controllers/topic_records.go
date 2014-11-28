package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"poke/base"
	"poke/models"
	"strings"
	"time"
)

type TopicRecords struct {
	base.BaseController
}

func (this *TopicRecords) Prepare() {
	this.BaseController.Prepare()
	this.Data[`subtitle`] = "发帖纪录"
	this.Data[`position`] = "record"
}

func (this *TopicRecords) Records() {
	count, _ := models.GetRecordsCounts()
	paginator := NewPaginator(this.Ctx.Request, 10, int(count))

	re, err := models.GetRecordsRange(paginator.Offset(), paginator.PerPageNums)
	if err != nil {
		beego.Warn(`get page range error`, err)
	}

	succn, err := models.GetRecordsCountSuccOrNot(true)
	if err != nil {
		beego.Warn(`get succ count error`, err)
	}
	failn, err := models.GetRecordsCountSuccOrNot(false)
	if err != nil {
		beego.Warn(`get fail count error`, err)
	}
	this.Data[`paginator`] = paginator
	this.Data[`succ`] = succn
	this.Data[`fail`] = failn
	this.Data[`records`] = re
	this.TplNames = `records.html`
}

func (this *TopicRecords) Search() {
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
	fmt.Println("timestring:", timestring)
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
	fmt.Println("username", username)
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

	re, err := models.SearchRecord(query_typ, query_val)
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
	this.TplNames = `records.html`

	// this.Redirect(this.UrlFor("this.Records"), 302)
}
