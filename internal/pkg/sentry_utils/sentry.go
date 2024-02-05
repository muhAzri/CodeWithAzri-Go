package sentry_utils

import (
	"os"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
)

func Init() (*sentryhttp.Handler, error) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           os.Getenv("SENTRY_DSN"),
		EnableTracing: true,

		TracesSampleRate: 1.0,
	}); err != nil {
		return nil, err
	}

	sentryHandler := sentryhttp.New(sentryhttp.Options{})
	return sentryHandler, nil
}
