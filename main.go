package main

import (
	
	"fmt"
	"strconv"

	"DBOpration/common"
	"DBOpration/utils"
	"DBOpration/utils/cache"
	"DBOpration/utils/db"
	"DBOpration/utils/config"
	_ "DBOpration/routers"


	"github.com/astaxie/beego"
	clog "github.com/cihub/seelog"
	"github.com/go-ini/ini"

)

const(
	LINUX_PROFILE="/home/saisai/gosource/src/DBOpration/conf/profile.ini"
)

var(
	http_addr string
	beego_loglevel int
)

func main() {

	var cf *config.Config

	//可根据运行的服务器系统不同，加载不同的配置文件
	cf=config.NewConfig(LINUX_PROFILE)

	if cf == nil{
		clog.Critical("new config failed")
		return
	}

	config.MakeDefaultConfig(cf)
	cfg,err:=cf.Load()
	if err!=nil{
		clog.Critical(err)
		return
	}

	if err=setupRedis(cfg);err!=nil{
		clog.Critical(err)
		return
	}

	if err=setupHttp(cfg);err!=nil{
		clog.Critical(err)
		return
	}

	if err=setupDB(cfg);err!=nil{
		clog.Critical(err)
		return
	}	

	beego.Run()
}

func setupRedis(cfg *ini.File) (err error) {
	sec,err:=cfg.GetSection("redis")
	if err!=nil{
		return err
	}

	url:=sec.Key("url").String()
	max_idles:=sec.Key("max_idles").MustInt(3)
	idle_timeout:=sec.Key("idle_timeout").MustInt(240)

	cache.Init(url,max_idles,idle_timeout,utils.REDIS_DB_USER_SYSTEM)
	//test cache first
	err=cache.Ping()
	if err!=nil{
		return err
	}
	return
}

func setupDB(cfg *ini.File) (err error) {
	sec,err:=cfg.GetSection("db")
	utils.CheckErr(err,utils.CHECK_FLAG_EXIT)
	driver:=sec.Key("driver").String()
	url:=sec.Key("url").String()
	max_lt,_:=strconv.Atoi(sec.Key("max_life_time").String())
	max_oc,_:=strconv.Atoi(sec.Key("max_open_conns").String())
	max_ic,_:=strconv.Atoi(sec.Key("max_idle_conns").String())
	common.DBmysqlUser=db.OpenDB(driver,url,max_lt,max_oc,max_ic)

	return
}

func setupHttp(cfg *ini.File) (err error) {
	sec,err:=cfg.GetSection("http")
	if err!=nil{
		return err
	}

	addr:=sec.Key("addr").String()
	if addr == ""{
		err=fmt.Errorf("HTTP addr can't be empty")
		return err
	}

	http_addr=addr
	beego_loglevel=sec.Key("beego_loglevel").MustInt(beego.LevelInformational)
	fmt.Println("setup http finish")

	return
}
