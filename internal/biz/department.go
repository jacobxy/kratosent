package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type DepartmentInfo struct {
	ID                 int64  `copier:"DepartmentId"`
	Name               string `copier:"DepartmentName"`
	ParentDepartmentID int64  `copier:"ParentDepartmentId"`
}

type DepartmenterRepo interface {
	Create(ctx context.Context, d *DepartmentInfo) (int64, error)
	Update(ctx context.Context, d *DepartmentInfo) (int64, error)
	Delete(ctx context.Context, departmentID int64) error
	List(ctx context.Context, departmentID []int64) ([]*DepartmentInfo, error)
}

type DepartmentUsecase struct {
	DepartmenterRepo
	log *log.Helper
}

func NewDepartmentUsecase(repo DepartmenterRepo, logger log.Logger) *DepartmentUsecase {
	return &DepartmentUsecase{
		DepartmenterRepo: repo,
		log:              log.NewHelper(logger),
	}
}

func (uc *DepartmentUsecase) Create(ctx context.Context, d *DepartmentInfo) (int64, error) {
	uc.log.WithContext(ctx).Infof("CreateDepartment: %v", d.Name)
	return uc.DepartmenterRepo.Create(ctx, d)
}

func (uc *DepartmentUsecase) Update(ctx context.Context, d *DepartmentInfo) (int64, error) {
	uc.log.WithContext(ctx).Infof("UpdateDepartment: %v", d.Name)
	return uc.DepartmenterRepo.Update(ctx, d)
}

func (uc *DepartmentUsecase) Delete(ctx context.Context, departmentID int64) error {
	uc.log.WithContext(ctx).Infof("DeleteDepartment: %v", departmentID)
	return uc.DepartmenterRepo.Delete(ctx, departmentID)
}

func (uc *DepartmentUsecase) List(ctx context.Context, departmentID []int64) ([]*DepartmentInfo, error) {
	uc.log.WithContext(ctx).Infof("ListDepartment: %v", departmentID)
	return uc.DepartmenterRepo.List(ctx, departmentID)
}
