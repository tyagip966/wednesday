package main

import (
	"github.com/labstack/echo"
	"strconv"
	"wednesday/container"
	"wednesday/router"
)


func main() {
	e := echo.New()
	controller := container.Container{Profile: "local"}
	controller.TriggerDI()
    router.InitRouter(e,&controller)
	port := strconv.FormatInt(controller.GetYamlConfig().Server.Port, 10)
	e.Logger.Fatal(e.Start(":"+port))
}
