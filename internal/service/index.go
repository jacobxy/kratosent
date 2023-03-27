package service

import (
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "kratosent/api/helloworld/v1"
	"kratosent/api/role"

	department "kratosent/api/department/v1"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewGRPCServiceSlice,
	NewHTTPServiceSlice,

	NewGreeterService,
	NewDepartmentService,
	NewRoleService,
)

type RegisterGRPCFunc func(*grpc.Server)
type RegisterHTTPFunc func(*http.Server)

func NewGRPCServiceSlice(
	greeter *GreeterService, // Gretter 服务
	dep *DepartmentService, //dep 服务
	r *RoleService,
) []RegisterGRPCFunc {
	res := make([]RegisterGRPCFunc, 0, 5)
	res = append(res, func(s *grpc.Server) {
		v1.RegisterGreeterServer(s, greeter)
	})
	res = append(res, func(s *grpc.Server) {
		department.RegisterDepartmentServer(s, dep)
	})

	res = append(res, func(s *grpc.Server) {
		role.RegisterRoleServer(s, r)
	})
	return res
}

func NewHTTPServiceSlice(
	greeter *GreeterService, // Gretter
	dep *DepartmentService,
	r *RoleService,
) []RegisterHTTPFunc {
	res := make([]RegisterHTTPFunc, 0, 5)
	res = append(res, func(s *http.Server) {
		v1.RegisterGreeterHTTPServer(s, greeter)
	})
	res = append(res, func(s *http.Server) {
		department.RegisterDepartmentHTTPServer(s, dep)
	})

	res = append(res, func(s *http.Server) {
		role.RegisterRoleHTTPServer(s, r)
	})
	return res
}
