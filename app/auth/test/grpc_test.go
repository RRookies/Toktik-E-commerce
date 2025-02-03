package test

import (
	auths "Toktik-E-commerce/app/auth/idl/authV1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func TestDial(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := auths.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.DeliverTokenByRPC(ctx, &auths.DeliverTokenReq{UserId: 996948441})
	if err != nil {

	}
	fmt.Println(resp)
	time.Sleep(time.Second)
	respp, err := client.RefreshTokenByRPC(ctx, &auths.RefreshTokenReq{Token: resp.Token})
	if err != nil {
	}
	fmt.Println(respp)
	resppp, err := client.VerifyTokenByRPC(ctx, &auths.VerifyTokenReq{Token: respp.Token})
	fmt.Println(resppp)
	resppp, err = client.VerifyTokenByRPC(ctx, &auths.VerifyTokenReq{Token: resp.Token})
	fmt.Println(resppp)
	respppp, err := client.LogoutTokenByRPC(ctx, &auths.LogoutTokenReq{Token: resp.Token})
	fmt.Println(respppp)
	resppp, err = client.VerifyTokenByRPC(ctx, &auths.VerifyTokenReq{Token: resp.Token})
	fmt.Println(resppp)
}
