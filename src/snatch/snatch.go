package snatch

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"loginpool"
	"os"
	//"strings"
	"github.com/astaxie/beego/logs"
	"html"
	"regexp"
	"strconv"
	"time"
	. "tools"
	"uwants"
)

var waittime time.Duration = 10
var log *logs.BeeLogger
var cl *uwants.Uwants

func init() {
	log = logs.NewLogger(1024)
	log.SetLevel(logs.LevelInfo)
	log.SetLogger(`console`, ``)
	// log.Info(`use lookt and fneon123 as account`)
	uwants.Proxy = ``
	var err error
	cl, err = loginpool.GetUwants(`alaladu`, `asdfghj`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var root = `http://www.discuss.com.hk/`

func GetFid() (fids []string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	log.Trace(`获取首页中`)
	re, err := cl.Get(root + `index2.php`)
	E(`爬虫:获取首页失败`, err)
	defer re.Body.Close()

	doc, err := goquery.NewDocumentFromReader(cl.Decoder.NewReader(re.Body))
	E(`爬虫:分析首页失败`, err)
	as := doc.Find("a")

	content := as.Map(func(i int, q *goquery.Selection) string {
		s, _ := q.Attr("href")
		return s
	})
	reg := regexp.MustCompile(`forumdisplay\.php\?fid=(\d+)`)
	dumpmap := map[string]struct{}{}

	log.Trace(`内容去重中`)
	for _, v := range content {
		if t := reg.FindStringSubmatch(v); len(t) > 1 {
			dumpmap[t[1]] = struct{}{}
		}
	}

	for k := range dumpmap {
		fids = append(fids, k)
	}
	return

}

func ChkFidAv(fids []string) (fidsav []string, err error) {

	log.Trace(`检查板块页面可用性`)
	for _, fid := range fids {

		log.Debug("检查ForumID: %v 中", fid)
		re, err := cl.Get(root + `forumdisplay.php?fid=` + fid)
		if err != nil {
			continue
		}
		doc, err := goquery.NewDocumentFromResponse(re)
		if err != nil {
			continue
		}
		_, ok := doc.Find(`form#postform`).Attr(`action`)
		if ok {
			log.Info("检查ForumID: %v 成功", fid)
			fidsav = append(fidsav, fid)
		}
		re.Body.Close()
	}
	return
}

func GetTid(fid string) (tids []string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	log.Debug("获取板块Id %v 下的帖子", fid)
	re, err := cl.Get(root + `forumdisplay.php?fid=` + fid)
	E(fmt.Sprintf("爬虫:获取板块 %v 失败", fid), err)
	defer re.Body.Close()

	doc, err := goquery.NewDocumentFromReader(cl.Decoder.NewReader(re.Body))
	E(fmt.Sprintf("爬虫:分析板块 %v 失败", fid), err)

	log.Trace(`寻找板块最后页`)
	flast, ok := doc.Find(".pages>a.last").Attr(`href`)
	if !ok {
		log.Info("爬虫:分析板块 %v 最后页失败,尝试分析最后一页", fid)

		flast, ok = doc.Find(".pages>a").Not(".next").Last().Attr(`href`)
		if !ok {
			log.Info("爬虫:分析板块 %v 最后页失败,假定为只有一页", fid)
			flast = `forumdisplay.php?fid=` + fid + `&page=1`
		}

	}

	reg := regexp.MustCompile(`forumdisplay\.php\?fid=\d+&page=(\d+)`)
	//得到具体页数 整数 ex. 1 2 3
	slastpage := reg.FindStringSubmatch(html.UnescapeString(flast))

	if len(slastpage) < 2 {
		E(fmt.Sprintf("爬虫:分析板块 %v 最后页失败", fid), false)
	}

	lastpage, err := strconv.Atoi(slastpage[1])
	E(fmt.Sprintf("爬虫:分析板块 %v 最后页失败", fid), err)

	reg = regexp.MustCompile(`thread_(\d+)`)
	for i := 1; i < lastpage+1; i++ {

		time.Sleep(waittime * time.Second)

		log.Debug("解析fid: %v 的第%v页", fid, i)
		re, err := cl.Get(root + `forumdisplay.php?fid=` + fid + `&page=` + fmt.Sprint(i))
		if err != nil {
			log.Warn("获取板块%v的第%v页失败", fid, i)
			continue
		}
		doc, err := goquery.NewDocumentFromReader(cl.Decoder.NewReader(re.Body))
		if err != nil {
			log.Warn("解析板块%v的第%v页失败", fid, i)
			continue
		}
		doc.Find(`span.tsubject`).Each(
			func(i int, q *goquery.Selection) {
				ids, _ := q.Attr(`id`)

				th := reg.FindStringSubmatch(ids)
				if len(th) < 2 {
					//TODO error handle
					return
				}
				tids = append(tids, th[1])
			})
	}
	return

}

func ChkTidAv(tids []string) (tidsav []string, err error) {
	log.Trace(`检查帖子页面可用性`)
	for _, tid := range tids {

		time.Sleep(waittime * time.Second)
		log.Trace("检查tid %v中", tid)
		re, err := cl.Get(root + `viewthread.php?tid=` + tid)
		if err != nil {
			continue
		}
		doc, err := goquery.NewDocumentFromResponse(re)
		if err != nil {
			continue
		}
		_, ok := doc.Find(`form#postform`).Attr(`action`)
		if ok {
			log.Info("tid %v可用", tid)
			tidsav = append(tidsav, tid)
		} else {
			log.Warn("tid %v不可用", tid)
		}
		re.Body.Close()
	}
	return
}
