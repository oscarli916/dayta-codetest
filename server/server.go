package main

import (
	"errors"
	"fmt"
	"os"
	"net"
	"context"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	pb "github.com/dayta-ai/dayta-se-test3/proto"
)

var port = ":3030"

var secretkey = os.Getenv("SECRET_KEY")

type DecodeTokenServiceServer struct {
	pb.UnimplementedDecodeTokenServiceServer
}

func (s *DecodeTokenServiceServer) DecodeToken(ctx context.Context, input *pb.DecodeTokenRequest) (*pb.DecodeTokenResponse, error){
	fmt.Printf("Received user input: %v\n", input.GetToken())
	m, err := verifyToken(input.GetToken())
	if err != nil{
		fmt.Printf("Error: %v", err)
	}
	fmt.Println("Title: ", m["title"].(string))
	fmt.Println("Message: ", m["message"].(string))
	return &pb.DecodeTokenResponse{Title: m["title"].(string), Message: m["message"].(string)}, nil
}
	

func main() {
	/*
		TODO:
		Implement a GRPC server that accepts a rpc call Decode Token.
		function verifyToken() is given below.
	*/
	lis, err := net.Listen("tcp", port)
	if err != nil{
		fmt.Println("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDecodeTokenServiceServer(s, &DecodeTokenServiceServer{})
	fmt.Println("GRPC server running on port", port)
	if err := s.Serve(lis); err != nil{
		fmt.Println("Failed to serve.")
	}
}

func verifyToken(token string) (map[string]interface{}, error) {
	parsedToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretkey, nil
	})
	mappedClaims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Token claims map error")
	}
	return mappedClaims, nil
}
