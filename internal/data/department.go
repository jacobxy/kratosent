package data

import (
	"context"
	"kratosent/ent/department"
	"kratosent/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type DepartmentRepo struct {
	data *Data
	log  *log.Helper
}

func NewDepartmentRepo(data *Data, logger log.Logger) biz.DepartmenterRepo {
	return &DepartmentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *DepartmentRepo) Create(ctx context.Context, d *biz.DepartmentInfo) (int64, error) {
	cli := r.data.EntClient
	res, err := cli.Department.Create().SetID(d.ID).
		SetName(d.Name).
		SetParentDepartmentID(d.ParentDepartmentID).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return res.ID, nil
}

func (r *DepartmentRepo) Update(ctx context.Context, d *biz.DepartmentInfo) (int64, error) {
	cli := r.data.EntClient
	res, err := cli.Department.UpdateOneID(d.ID).
		SetName(d.Name).
		SetParentDepartmentID(d.ParentDepartmentID).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return res.ID, nil
}

func (r *DepartmentRepo) Delete(ctx context.Context, departmentID int64) error {
	cli := r.data.EntClient
	return cli.Department.DeleteOneID(departmentID).Exec(ctx)
}

func (r *DepartmentRepo) List(ctx context.Context, departmentID []int64) ([]*biz.DepartmentInfo, error) {
	cli := r.data.EntClient
	query := cli.Department.Query()
	if len(departmentID) > 0 {
		query = query.Where(department.IDIn(departmentID...))
	}
	departments, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*biz.DepartmentInfo, 0, len(departments))
	for _, department := range departments {
		res = append(res, &biz.DepartmentInfo{
			ID:                 department.ID,
			Name:               department.Name,
			ParentDepartmentID: *department.ParentDepartmentID,
		})
	}
	return res, nil
}
