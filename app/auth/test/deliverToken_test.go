package test

import (
	auths "Toktik-E-commerce/app/auth/idl/authV1"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestDeliverToken(t *testing.T) {
	tests := []struct {
		name string
		req  *auths.DeliverTokenReq
		err  error
	}{
		{
			name: "正常使用",
			req: &auths.DeliverTokenReq{
				UserId: 123,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
			if err != nil {
				t.Errorf("did not connect: %v", err)
			}
			defer conn.Close()
			client := auths.NewAuthServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_, err = client.DeliverTokenByRPC(ctx, test.req)
			assert.Equal(t, test.err, err, "正常使用失败")
		})
	}
}
