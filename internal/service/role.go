package service

import (
	"context"
	"fmt"
	"log"

	pb "kratosent/api/role"
	"kratosent/internal/biz"

	"github.com/jinzhu/copier"
)

type RoleService struct {
	pb.UnimplementedRoleServer
	u *biz.RoleUsecase
}

func NewRoleService(u *biz.RoleUsecase) *RoleService {
	return &RoleService{
		u: u,
	}
}

func (s *RoleService) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleReply, error) {
	role, err := s.u.Create(ctx, &biz.RoleInfo{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	res := &pb.CreateRoleReply{}
	return res, copier.Copy(res, role)
}

func (s *RoleService) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleReply, error) {
	newId, err := s.u.UpdateRole(ctx, &biz.RoleInfo{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Name,
	})
	return &pb.UpdateRoleReply{
		Id: newId,
	}, err
}
func (s *RoleService) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleReply, error) {
	newID, err := s.u.DeleteRole(ctx, req.Id)
	return &pb.DeleteRoleReply{
		Id: newID,
	}, err
}
func (s *RoleService) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleReply, error) {
	roleInfos, err := s.u.ListRole(ctx, []int64{req.Id})
	if err != nil {
		return nil, err
	}
	if len(roleInfos) == 0 {
		return nil, fmt.Errorf("role not found")
	}
	return &pb.GetRoleReply{
		Role: &pb.RoleInfo{
			Id:          roleInfos[0].ID,
			Name:        roleInfos[0].Name,
			Description: roleInfos[0].Description,
		},
	}, nil
}
func (s *RoleService) ListRole(ctx context.Context, req *pb.ListRoleRequest) (*pb.ListRoleReply, error) {
	log.Println("ListRole ", req.Ids)
	roleInfos, err := s.u.ListRole(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	res := &pb.ListRoleReply{
		Roles: make([]*pb.RoleInfo, 0, len(roleInfos)),
	}
	err = copier.Copy(&res.Roles, roleInfos)
	return res, err
}
