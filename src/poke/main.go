package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/session/redis"
	_ "poke/controllers"
	_ "poke/routers"

	"uwants"
)

func main() {
	cfg, err := config.NewConfig(`ini`, `config.conf`)
	if err != nil {
		fmt.Println(`open config error`)
		return
	}
	proxy := cfg.String(`proxy`)
	beego.Trace("set proxy as", proxy)
	uwants.Proxy = proxy
	beego.AddFuncMap("stringfy", func(i interface{}) string { return fmt.Sprint(i) })
	beego.Run()
}
