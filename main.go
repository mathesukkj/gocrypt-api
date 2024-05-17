package main

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"

	_ "gocrypt-api/routers"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase(
		"default",
		"postgres",
		"user=mathesu password=password host=127.0.0.1 port=5432 dbname=crypto sslmode=disable",
	)

	orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
