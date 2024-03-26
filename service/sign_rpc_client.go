package service

import (
	"context"
	signPb "reserve/proto/sign"
	"sync"
	"time"
)

var (
	signRpcClient signPb.SignClient
	once          sync.Once
)

func InitSignRpcClient(client signPb.SignClient) {
	once.Do(func() {
		signRpcClient = client
	})
}

func NewSignRpcService() ISignRpc {
	return &SignRpcService{
		signRpcClient: signRpcClient,
	}
}

type ISignRpc interface {
	VerifyToken(token string) (*signPb.VerifyTokenRespond, error)
	GetUserInfo(id int64) (*signPb.UserInfoRespond, error)
}

type SignRpcService struct {
	signRpcClient signPb.SignClient
}

func (s *SignRpcService) GetUserInfo(id int64) (*signPb.UserInfoRespond, error) {
	ctx, f := context.WithTimeout(context.Background(), time.Duration(3*int(time.Second)))
	defer f()
	in := &signPb.UserInfoRequest{
		UserId: id,
	}
	return s.signRpcClient.UserInfo(ctx, in)
}

func (s *SignRpcService) VerifyToken(token string) (*signPb.VerifyTokenRespond, error) {
	ctx, f := context.WithTimeout(context.Background(), time.Duration(3*int(time.Second)))
	defer f()
	in := &signPb.VerifyTokenRequest{
		Token: token,
	}
	_, err := s.signRpcClient.VerifyToken(ctx, in)
	if err != nil {
		return nil, err
	}
	resp := &signPb.VerifyTokenRespond{
		//UserId:   data.UserId,
		//UserName: data.UserName,
	}
	return resp, nil
}
