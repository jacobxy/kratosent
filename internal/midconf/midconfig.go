package midconf

import (
	"kratosent/internal/conf"
	"log"
	"sync/atomic"

	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"
	"gopkg.in/yaml.v2"
)

type MidConfig struct {
	WhiteList []string `yaml:"white_list"`
}

var midconfig atomic.Value

func NewMidConfig(config *conf.MidConfig) error {
	sdk, err := polaris.NewSDKContextByAddress(config.Addrs...)
	if err != nil {
		return err
	}
	configAPI := polaris.NewConfigAPIByContext(sdk)
	if err != nil {
		log.Println("fail to start example", err)
		return err
	}

	configFile, err := configAPI.GetConfigFile(config.Namespace, config.FileGroup, config.FileName)
	if err != nil {
		return err
	}

	content := configFile.GetContent()

	mid := MidConfig{}
	yaml.UnmarshalStrict([]byte(content), &mid)
	midconfig.Store(mid)
	log.Println("configFile:", configFile.GetContent())
	configFile.AddChangeListener(changeLister)
	return nil
}

func changeLister(event model.ConfigFileChangeEvent) {
	content := event.NewValue
	mid := MidConfig{}
	yaml.UnmarshalStrict([]byte(content), &mid)
	log.Println("midconf change", mid)
	midconfig.Store(mid)
}
