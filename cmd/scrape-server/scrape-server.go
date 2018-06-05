package main

import (
	"flag"

	"github.com/daregod/amazon-co-uk-scraper/server"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
)

// Config used by flags parse
type config struct {
	addr string
}

var (
	cfg config
)

func init() {
	flag.StringVar(&cfg.addr, "listen", ":8007", "addr:port to listen on")
}

func main() {
	flag.Parse()
	sServ := server.NewServer()
	cLog := console.New(true)
	cLog.SetTimestampFormat(`2006-01-02 15:04:05.0000:`)
	log.AddHandler(cLog, log.AllLevels...)
	r := gin.Default()
	sServ.MountRoutes(r)
	err := r.Run(cfg.addr)
	if err != nil {
		log.WithError(err).WithField("listen address", cfg.addr).Error("Cannot run server") //nolint: lll
	}
}
