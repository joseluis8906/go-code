package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joseluis8906/go-code/src/delivery/internal/app"

	"github.com/spf13/viper"
)

func main() {
	conf := viper.New()
	conf.SetConfigName("delivery")
	conf.AddConfigPath("/usr/local/etc")
	conf.SetConfigType("yml")
	err := conf.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(conf.GetString("logging.path"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := log.New(f, "[DELIVERY] ", log.LstdFlags)
	logger.SetFlags(log.LstdFlags | log.Llongfile)

	app := app.GRPCServer{
		Args: os.Args,
		Log:  logger,
		Conf: conf,
	}

	err = app.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
