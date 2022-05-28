package server

import (
	"Atlantis/config"
	"Atlantis/routes"
	"Atlantis/services/es"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func Init() {

	config := config.Get()

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.SentryDSN,
		Debug:            true,
		Environment:      config.AppEnv,
		TracesSampleRate: float64(config.SentrySamplingRate),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
	es.Init()

	r := routes.NewRouter()
	r.Run(":" + "4000")
}
