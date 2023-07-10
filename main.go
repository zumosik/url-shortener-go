package main

import (
	"flag"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"github.com/zumosik/url-shortener/config"
)

var ConfigPath string

func init() {
	flag.StringVar(&ConfigPath, "confpath", "./config/config.yaml", "Path to config (.yaml)")
}

func main() {
	flag.Parse()

	config, err := config.ReadConfig(ConfigPath)
	if err != nil {
		panic(err.Error())
	}

	if err := ConfigureLogger(config.Server.LogLevel); err != nil {
		errors.Wrap(err, "Failed to configure Logger")
	}

}

func ConfigureLogger(logLevel string) error {
	logLvl, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(logLvl)
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg% \n",
	})

	return nil

}
