package main

import (
	"github.com/labstack/echo/v4"
	"go_be_github_trending/db"
	"go_be_github_trending/handler"
)

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "123456",
		DbName:   "github",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()

	e.GET("/", handler.Welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}
