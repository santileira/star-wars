package server

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"rings/domain/skywalker/handler"
	"rings/domain/skywalker/service"
	"rings/http"
	"time"
)

type Runnable struct{}

func NewRunnable() *Runnable {
	return &Runnable{}
}

func (r *Runnable) Cmd() *cobra.Command {
	options := &Options{}

	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "Runs an API",
		Long:  `Runs an API`,
	}

	cmd.Flags().StringVar(&options.LogLevel, "log-level", defaultLogLevel, "log leve to use")
	cmd.Flags().IntVar(&options.TimeoutInMilliseconds, "timeout-in-milliseconds", defaultTimeoutInMilliseconds,
		"timeout of the api calls")

	cmd.Run = func(_ *cobra.Command, _ []string) {
		server := r.Run(options)
		server.Start()
	}
	return cmd
}

func (r *Runnable) Run(options *Options) *Server {
	r.configureLog(options.LogLevel)

	httpClient := http.NewClientImpl(time.Duration(options.TimeoutInMilliseconds) * time.Millisecond)
	service := service.NewServiceImpl(httpClient)
	handler := handler.NewHandler(service)

	return NewServer(handler)
}

func (r *Runnable) configureLog(logLevel string) {
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Warnf("Error passing the log level: %s", logLevel)
		return
	}

	logrus.Infof("Setting log level: %s", lvl.String())
	logrus.SetLevel(lvl)
}
