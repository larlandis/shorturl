package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/larlandis/shorturl/cmd"
	"github.com/larlandis/shorturl/internal/pkg/config"
	"github.com/larlandis/shorturl/internal/pkg/hash"
	"github.com/larlandis/shorturl/internal/pkg/rest/writer"
	"github.com/larlandis/shorturl/internal/platform/storage"
)

func main() {

	// parseArgs
	cmd.ParseArgs()

	// init config
	config.Init(cmd.ConfigFile)

	// create services
	st := storage.NewLocal()
	s := hash.New(st)

	// start server
	log.Printf("starting server on :%s", cmd.Port)
	if err := http.ListenAndServe(":"+cmd.Port, writer.New(s)); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error starting server: %s\n", err)
			return
		}
		log.Println("server stopped")
	}
}
