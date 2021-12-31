package main

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/dayta-ai/dayta-se-test3/proto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
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
	conn, err := grpc.Dial("192.168.1.43:3030", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDecodeTokenServiceClient(conn)

	header := c.Request().Header

	r, err := client.DecodeToken(context.Background(), &pb.DecodeTokenRequest{Token: header.Get("Authorization")})
	if err != nil {
		fmt.Printf("Could not make Token: %v", err)
	}
	fmt.Println(r)
	return c.JSON(http.StatusOK, r)
}
