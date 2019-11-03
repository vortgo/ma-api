package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"ma-api/handler"
	"ma-api/templates"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "assets")
	handler.RegisterRoutes(e)
	templates.TemplateRegistry(e)
	e.Logger.Fatal(e.Start(":1323"))

}
