package data

import (
	"kratosent/ent"
	"kratosent/internal/conf"

	"github.com/go-kratos/kratos/contrib/polaris/v2"
	"github.com/go-kratos/kratos/v2/log"
	redis "github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	polarisV2 "github.com/polarismesh/polaris-go"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
	NewDepartmentRepo,
	NewRoleRepo,
	NewEntCli,
)

// Data .
type Data struct {
	EntClient *ent.Client
	// sdk       api.SDKContext
	pol      *polaris.Polaris
	redisCli *redis.Client
}

// NewData .
func NewData(entCli *ent.Client, data *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	sdk, err := polarisV2.NewSDKContextByAddress(data.Department.Addrs...)
	if err != nil {
		return nil, nil, err
	}

	pol := polaris.New(sdk,
		polaris.WithNamespace(data.Department.Namespace),
	)

	redisCli := redis.NewClient(&redis.Options{
		Addr:         data.Redis.Addr,
		ReadTimeout:  data.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: data.Redis.WriteTimeout.AsDuration(),
	})

	return &Data{
		EntClient: entCli,
		pol:       &pol,
		redisCli:  redisCli,
	}, cleanup, nil
}
