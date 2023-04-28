package data

import (
	"context"
	v1 "kratosent/api/department/v1"
	"kratosent/ent/role"
	"kratosent/internal/biz"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/jinzhu/copier"
)

var _ biz.RoleRepo = (*RoleRepo)(nil)

type RoleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &RoleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RoleRepo) Create(ctx context.Context, role *biz.RoleInfo) (int64, error) {
	cli := r.data.EntClient
	roleInfo, err := cli.Role.Create().SetName(role.Name).
		SetDescription(role.Description).Save(ctx)
	if err != nil {
		return 0, err
	}

	// departmentCli, clearUp, err := NewDepartmentClient(r.data.sdk)
	// departmentCli, clearUp, err := NewGRPCClient[v1.DepartmentClient]("kratosent", r.data.pol, v1.NewDepartmentClient)
	departmentCli, clearUp, err := NewHTTPClient[v1.DepartmentHTTPClient]("kratosent", r.data.pol, v1.NewDepartmentHTTPClient)
	if err != nil {
		return 0, err
	}
	defer clearUp()
	clientCtx := context.Background()
	clientCtx = metadata.AppendToClientContext(clientCtx, "x-md-global-extra", "lok")
	r.log.Info(
		departmentCli.CreateDepartment(clientCtx, &v1.CreateDepartmentRequest{
			DepartmentName:     role.Name,
			ParentDepartmentId: -1,
		}),
	)

	return roleInfo.ID, nil
}

func (r *RoleRepo) Update(ctx context.Context, roleInfo *biz.RoleInfo) (int64, error) {
	cli := r.data.EntClient
	query := cli.Role.UpdateOneID(roleInfo.ID)
	if roleInfo.Name != "" {
		query = query.SetName(roleInfo.Name)
	}
	if roleInfo.Description != "" {
		query = query.SetDescription(roleInfo.Description)
	}
	role, err := query.Save(ctx)
	if err != nil {
		return 0, err
	}

	return role.ID, nil
}

func (r *RoleRepo) Delete(ctx context.Context, roleID int64) (int64, error) {
	cli := r.data.EntClient
	err := cli.Role.DeleteOneID(roleID).Exec(ctx)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (r *RoleRepo) List(ctx context.Context, roleIDs []int64) ([]*biz.RoleInfo, error) {
	cli := r.data.EntClient
	query := cli.Role.Query()
	roleList := make([]string, 0, len(roleIDs))
	for _, v := range roleIDs {
		roleList = append(roleList, strconv.Itoa(int(v)))
	}
	list, err := r.data.redisCli.HMGet("role", roleList...).Result()
	if err != nil && len(list) != 0 {
		res := make([]*biz.RoleInfo, 0, len(list))
		err = copier.Copy(&res, &list)
		return res, err
	}
	switch len(roleIDs) {
	case 0:
	case 1:
		query = query.Where(role.IDEQ(roleIDs[0]))
	default:
		query = query.Where(role.IDIn(roleIDs...))
	}
	roles, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*biz.RoleInfo, 0, len(roles))
	err = copier.Copy(&res, &roles)

	datas := make(map[string]interface{}, len(res))
	for _, v := range res {
		datas[strconv.Itoa(int(v.ID))] = v
	}
	r.data.redisCli.HMSet("role", datas)
	return res, err
}
