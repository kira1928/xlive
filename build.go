//go:build ignore

package main

import (
	"os"

	"github.com/kira1928/xlive/src/cmd/build"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   true,
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	os.Exit(build.RunCmd())
}
