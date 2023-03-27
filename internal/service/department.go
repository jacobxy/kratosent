package service

import (
	"context"
	"fmt"
	"log"

	pb "kratosent/api/department/v1"
	"kratosent/internal/biz"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/jinzhu/copier"
)

type DepartmentService struct {
	pb.UnimplementedDepartmentServer
	uc *biz.DepartmentUsecase
}

func NewDepartmentService(repo *biz.DepartmentUsecase) *DepartmentService {
	return &DepartmentService{
		uc: repo,
	}
}

func (s *DepartmentService) CreateDepartment(ctx context.Context, req *pb.CreateDepartmentRequest) (*pb.CreateDepartmentReply, error) {
	var extra string
	if md, ok := metadata.FromServerContext(ctx); ok {
		// if md, ok := metadata.FromClientContext(ctx); ok {
		extra = md.Get("x-md-global-extra")
		log.Println("extra:", extra, req)
	}

	id, err := s.uc.Create(ctx, &biz.DepartmentInfo{
		ParentDepartmentID: req.ParentDepartmentId,
		Name:               req.GetDepartmentName(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateDepartmentReply{
		DepartmentId: id,
	}, nil
}
func (s *DepartmentService) UpdateDepartment(ctx context.Context, req *pb.UpdateDepartmentRequest) (*pb.UpdateDepartmentReply, error) {
	id, err := s.uc.Update(ctx, &biz.DepartmentInfo{
		ID:   req.DepartmentId,
		Name: req.GetDepartmentName(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateDepartmentReply{
		DeparmentId: id,
	}, nil
}
func (s *DepartmentService) DeleteDepartment(ctx context.Context, req *pb.DeleteDepartmentRequest) (*pb.DeleteDepartmentReply, error) {
	err := s.uc.Delete(ctx, req.DepartmentId)
	return &pb.DeleteDepartmentReply{}, err
}
func (s *DepartmentService) GetDepartment(ctx context.Context, req *pb.GetDepartmentRequest) (*pb.GetDepartmentReply, error) {
	res, err := s.uc.List(ctx, []int64{req.DepartmentId})
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("department %d not found", req.DepartmentId)
	}
	return &pb.GetDepartmentReply{
		Department: &pb.DepartmentInfo{
			DepartmentId:       res[0].ID,
			DepartmentName:     res[0].Name,
			ParentDepartmentId: res[0].ParentDepartmentID,
		},
	}, nil
}

func (s *DepartmentService) ListDepartment(ctx context.Context, req *pb.ListDepartmentRequest) (*pb.ListDepartmentReply, error) {
	departments, err := s.uc.List(ctx, req.DepartmentIds)
	if err != nil {
		return nil, err
	}
	res := &pb.ListDepartmentReply{}
	err = copier.Copy(&res.Departments, departments)
	if err != nil {
		return nil, err
	}
	return res, nil
}
