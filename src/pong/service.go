package pong

import (
	"go.uber.org/fx"

	"github.com/joseluis8906/go-code/idl/pongpb"
)

var Module = fx.Provide(NewPongService)

type (
	PongService struct {
		pongpb.UnimplementedPingServer
	}
)

func NewPongService() *PongService {
	return &PongService{}
}
