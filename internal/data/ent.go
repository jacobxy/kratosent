package data

import (
	"context"
	"kratosent/ent"
	"kratosent/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

func NewEntCli(c *conf.Data, logger log.Logger) (*ent.Client, func(), error) {
	entClient, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		logger.Log(log.LevelError, "failed opening mysql")
		return nil, nil, err
	}

	err = entClient.Schema.Create(context.Background())
	logger.Log(log.LevelInfo, "create schema", err)

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		entClient.Close()
	}
	return entClient, cleanup, nil
}
