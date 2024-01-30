package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"

	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Logger *log.Logger
}

func New(params Params) *viper.Viper {
	v := viper.New()
	v.AddRemoteProvider("etcd", os.Getenv("CONFIG_URL"), "/config/delivery.yml")
	v.SetConfigType("yml")
	if err := v.ReadRemoteConfig(); err != nil {
		log.Fatalf("cannot read remote config: %v", err)
	}

	return v
}
