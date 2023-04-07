package main

import (
	"gdt-api/config"
	jbData "gdt-api/features/job/data"
	jbHdl "gdt-api/features/job/handler"
	jbSrv "gdt-api/features/job/services"
	usrData "gdt-api/features/user/data"
	usrHdl "gdt-api/features/user/handler"
	usrSrv "gdt-api/features/user/services"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	config.Migrate(db)
	uData := usrData.New(db)
	uSrv := usrSrv.New(uData)
	uHdl := usrHdl.New(uSrv)
	jData := jbData.New(db)
	jSrv := jbSrv.New(jData)
	jHdl := jbHdl.New(jSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	// user Routes
	e.POST("/register", uHdl.Register())
	e.POST("/login", uHdl.Login())
	e.GET("/job-list", jHdl.GetJobList())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
