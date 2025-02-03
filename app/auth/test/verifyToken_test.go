package test

import (
	auths "Toktik-E-commerce/app/auth/idl/authV1"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestVerifyTokenTrue(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	client := auths.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	deliverResp, err := client.DeliverTokenByRPC(ctx, &auths.DeliverTokenReq{UserId: 996948441})
	assert.Equal(t, nil, err, "deliverToken失败")

	verifyResp, err := client.VerifyTokenByRPC(ctx, &auths.VerifyTokenReq{Token: deliverResp.Token})
	assert.Equal(t, nil, err, err)
	assert.Equal(t, true, verifyResp.Res, "验证错误")
}

func TestVerifyTokenTimeout(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	client := auths.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	deliverResp, err := client.DeliverTokenByRPC(ctx, &auths.DeliverTokenReq{UserId: 996948441})
	assert.Equal(t, nil, err, "deliverToken失败")

	verifyResp, err := client.VerifyTokenByRPC(ctx, &auths.VerifyTokenReq{Token: deliverResp.Token + "1"})
	assert.Equal(t, nil, err, err)
	assert.Equal(t, false, verifyResp.Res, "验证错误")
}
