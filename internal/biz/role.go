package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type RoleInfo struct {
	ID          int64 `json:"id"`
	Name        string
	Description string
}

type RoleRepo interface {
	Create(ctx context.Context, role *RoleInfo) (int64, error)
	Update(ctx context.Context, role *RoleInfo) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
	List(ctx context.Context, ids []int64) ([]*RoleInfo, error)
}

type RoleUsecase struct {
	RoleRepo
	log *log.Helper
}

func NewRoleUsecase(repo RoleRepo, logger log.Logger) *RoleUsecase {
	return &RoleUsecase{
		RoleRepo: repo,
		log:      log.NewHelper(logger),
	}
}

func (uc *RoleUsecase) CreateRole(ctx context.Context, role *RoleInfo) (int64, error) {
	uc.log.WithContext(ctx).Infof("CreateRole: %v", role.Name)
	return uc.RoleRepo.Create(ctx, role)
}

func (uc *RoleUsecase) UpdateRole(ctx context.Context, role *RoleInfo) (int64, error) {
	uc.log.WithContext(ctx).Infof("UpdateRole: %v", role.Name)
	return uc.RoleRepo.Update(ctx, role)
}

func (uc *RoleUsecase) DeleteRole(ctx context.Context, id int64) (int64, error) {
	uc.log.WithContext(ctx).Infof("DeleteRole: %v", id)
	return uc.RoleRepo.Delete(ctx, id)
}

func (uc *RoleUsecase) ListRole(ctx context.Context, ids []int64) ([]*RoleInfo, error) {
	uc.log.WithContext(ctx).Infof("ListRole: %v", ids)
	return uc.RoleRepo.List(ctx, ids)
}
