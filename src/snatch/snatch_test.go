package snatch

import (
	"html"
	"testing"
)

func Test_GetFid(t *testing.T) {
	t.SkipNow()
	t.Log(GetFid())
}

func Test_url(t *testing.T) {
	t.SkipNow()
	t.Log(html.UnescapeString(`&amp;`))

}

func Test_GetTid(t *testing.T) {
	t.SkipNow()
	t.Log(GetTid(`forumdisplay.php?fid=1830`))

}

func Test_ChkTidSav(t *testing.T) {
	tid, err := GetTid(`forumdisplay.php?fid=1830`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ChkTidAv(tid))

}
