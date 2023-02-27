package data

import (
	"context"
	"fmt"

	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"

	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/go-kratos/kratos/contrib/polaris/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	ogrpc "google.golang.org/grpc"
)

// func NewGRPCClient[T any](service string,sdk api.SDKContext, fc func (ogrpc.ClientConnInterface)T)(T,func(),error){
func NewGRPCClient[T any](service string,p *polaris.Polaris, fc func (ogrpc.ClientConnInterface)T)(T,func(),error){
	var empty T

	conn,err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(fmt.Sprintf("discovery:///%s",service)),
		grpc.WithDiscovery(p.Registry()),
		grpc.WithMiddleware(
			mmd.Client(),
		),
	)
	if err != nil {
		return empty,nil,err
	}

	return fc(conn),func(){
		conn.Close()
	},nil
}

func NewHTTPClient[T any](service string,p *polaris.Polaris, fc func(*http.Client)T)(T,func(),error){
	var empty T

	conn,err := http.NewClient(
		context.Background(),
		http.WithEndpoint(fmt.Sprintf("discovery:///%s",service)),
		http.WithDiscovery(p.Registry()),
		http.WithMiddleware(
			mmd.Client(),
		),
	)
	if err != nil {
		return empty,nil,err
	}

	return fc(conn),func(){
		conn.Close()
	},nil
}
