package main

import (
	"pizzatime/version"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	logger         = logrus.WithField("component", "pizzatime")
	globalLogLevel = "warn"
	root           = &cobra.Command{
		Use:     "pizzatime",
		Version: version.Version.String(),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			logrus.SetLevel(logrus.InfoLevel)
			logrus.SetFormatter(&logrus.JSONFormatter{})
			if l, err := logrus.ParseLevel(globalLogLevel); err != nil {
				logger.WithError(err).
					Warn("got error when parsing log level, defaulting to INFO")
			} else {
				logrus.SetLevel(l)
			}
		},
	}
)

func init() {
	root.PersistentFlags().StringVar(&globalLogLevel, "log-level", globalLogLevel, "minimum level of logs to print to STDERR")
	root.AddCommand(
		cmdVersion,
	)
}

func main() {
	root.Execute()
}
