package zerolog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"time"
)

var debug = true
var consoleWriter io.Writer
var fileWriter io.Writer

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
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

func DebugEnable(enable bool) {
	debug = enable
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
