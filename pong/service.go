package pong

import (
	"go.uber.org/fx"

	"github.com/joseluis8906/yummies/go-code/protobuf/pong/pongpb"
)

var Module = fx.Provide(NewPongService)

type (
	PingService struct {
		pongpb.UnimplementedPingServer
	}
)

func NewPongService() *PingService {
	return &PingService{}
}
