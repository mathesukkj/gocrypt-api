package main

import (
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	aescrypt "gocrypt-api/crypto"
	_ "gocrypt-api/routers"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	orm.RegisterDataBase(
		"default",
		"postgres",
		connString,
	)

	orm.RunSyncdb("default", false, true)

	aescrypt.Init()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
