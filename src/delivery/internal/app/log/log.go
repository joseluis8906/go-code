package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type (
	Deps struct {
		fx.In

		Config *viper.Viper
	}

	Logger struct {
		stdOut io.Writer
		conn   *fluent.Fluent
		tag    string
		env    string
	}
)

func (l *Logger) Write(data []byte) (int, error) {
	if _, err := l.stdOut.Write(data); err != nil {
		return 0, err
	}

	structured := make(map[string]string, 6)
	values := strings.Split(string(data), " ")

	structured["env"] = l.env
	structured["app"] = values[0]
	structured["datetime"] = fmt.Sprintf("%v %v", values[1], values[2])
	structured["line"] = strings.TrimSuffix(values[3], ":")
	structured["level"] = values[4]
	structured["message"] = strings.Join(values[5:], " ")

	return len(data), l.conn.Post(l.tag, structured)
}

func New(deps Deps) *log.Logger {
	fluentd, err := fluent.New(fluent.Config{
		FluentHost:    deps.Config.GetString("fluentd.host"),
		FluentPort:    deps.Config.GetInt("fluentd.port"),
		FluentNetwork: "tcp",
		MarshalAsJSON: true,
		Async:         true,
	})
	if err != nil {
		log.Fatalf("connecting fluentd: %v", err)
	}

	logger := Logger{
		stdOut: os.Stderr,
		conn:   fluentd,
		tag:    deps.Config.GetString("fluentd.tag"),
		env:    deps.Config.GetString("env"),
	}

	l := log.New(&logger, "delivery ", log.LstdFlags)
	l.SetFlags(log.LstdFlags | log.Llongfile)

	return l
}

func Info(message string) string {
	return fmt.Sprintf("INFO %v", message)
}

func Error(message string) string {
	return fmt.Sprintf("ERROR %v", message)
}
