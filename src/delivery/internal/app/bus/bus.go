package bus

import (
	"log"

	"github.com/nats-io/nats.go"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Config *viper.Viper
	Logger *log.Logger
}

func New(params Params) *nats.Conn {
	nc, err := nats.Connect(params.Config.GetString("nats.url"))
	if err != nil {
		params.Logger.Fatalf("connecting nats: %v", err)
	}

	return nc
}
