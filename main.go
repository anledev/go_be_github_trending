package main

import (
	"github.com/labstack/echo/v4"
	"go_be_github_trending/handler"
)

func main() {
	e := echo.New()

	e.GET("/", handler.Welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}
