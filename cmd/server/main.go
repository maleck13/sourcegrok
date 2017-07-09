package main

import (
	"flag"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
	"github.com/maleck13/sourcegrok/pkg/logger"
	"github.com/maleck13/sourcegrok/pkg/web"
)

var logLevel string

func setupLogger() *logrus.Logger {
	switch logLevel {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.ErrorLevel)
	}
	return logrus.StandardLogger()
}

func main() {
	//flags defined here
	{
		flag.StringVar(&logLevel, "log-level", "info", "use this to set log level: error, info, debug")
		flag.Parse()
	}

	var (
		httpHandler  http.Handler
		router       *chi.Mux
		logrusLogger = setupLogger()
		log          = logger.New(logrusLogger)
	)

	//router
	{
		router = chi.NewRouter()
		httpHandler = web.BuildHTTPHandler(router)
	}
	//sys route
	{
		web.SysRoute(router, &web.SysHandler{Logger: log})
	}
	//http handler
	{
		port := ":3010"
		logrusLogger.Info("starting sourcegrok 0.0.1 on  port " + port)
		if err := http.ListenAndServe(port, httpHandler); err != nil {
			logrusLogger.Fatal(err)
		}
	}

}
