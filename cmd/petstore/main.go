package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/prologic/petstore/internal"
)

var (
	bind    string
	debug   bool
	dbpath  string
	version bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <path>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVarP(&version, "version", "v", false, "display version information")
	flag.BoolVarP(&debug, "debug", "d", false, "enable debug logging")

	flag.StringVarP(&bind, "bind", "b", internal.DefaultBind, "interface and port to bind to")
	flag.StringVarP(&dbpath, "dbpath", "", internal.DefaultDBPath, "storage path for the database")
}

func main() {
	flag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if version {
		fmt.Printf("petstore version %s", internal.FullVersion())
		os.Exit(0)
	}

	server, err := internal.NewServer(
		internal.WithBind(bind),
		internal.WithDBPath(dbpath),
	)
	if err != nil {
		log.WithError(err).Error("error creating server")
		os.Exit(2)
	}

	log.WithField("bind", bind).
		WithField("version", internal.FullVersion()).
		Infof("petstore v%s listening on %s", internal.FullVersion(), bind)

	if err := server.Run(); err != nil {
		log.WithError(err).Fatal("error running server")
	}
}
