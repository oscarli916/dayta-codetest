package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO
	//	Implement a grpc client connection and pass the HTTP request to grpc call
	e.POST("/", implementThisFunction)

	e.Logger.Fatal(e.Start(":8080"))
}

func implementThisFunction(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, "UNIMPLEMENTED")
}
