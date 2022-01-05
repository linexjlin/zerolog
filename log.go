package zerolog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"time"
)

var consoleWriter io.Writer
var fileWriter io.Writer

func init() {
	if os.Getenv("debug")+os.Getenv("DEBUG")+os.Getenv("Debug") != "" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano}
	consoleWriter = output
	log.Logger = log.Output(output)
	log.Logger = log.With().Caller().Logger()
}

func LogToFile(path string) {
	if f, e := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660); e != nil {
		log.Warn().Err(e)
	} else {
		fileWriter = f
		multi := zerolog.MultiLevelWriter(consoleWriter, fileWriter)
		log.Logger = log.Output(multi)
		log.Logger = log.With().Caller().Logger()
	}
}
