package config

import (
	"github.com/getsentry/sentry-go"
	"os"
	"strconv"
	"time"
)

func Sentry() error {
	sentryIsActive, _ := strconv.ParseBool(os.Getenv("SENTRY_ENABLED"))

	if sentryIsActive {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: os.Getenv("SENTRY_DSN"),
		})

		defer sentry.Flush(2 * time.Second)
		return err
	}
	return nil
}
