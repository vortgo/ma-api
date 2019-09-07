package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"ma-api/handler"
)

func main() {
	e := echo.New()
	e.Static("/static", "assets")
	handler.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
