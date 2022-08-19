package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/larlandis/shorturl/cmd"
	"github.com/larlandis/shorturl/internal/pkg/config"
	"github.com/larlandis/shorturl/internal/pkg/hash"
	"github.com/larlandis/shorturl/internal/pkg/rest/reader"
	"github.com/larlandis/shorturl/internal/platform/storage"
)

func main() {

	// parseArgs
	cmd.ParseArgs()

	// create services
	st := storage.New(cmd.StorageCluster)
	s := hash.New(st)

	// init libs
	config.Init(cmd.ConfigFile)

	// start server
	log.Printf("starting server on :%s", cmd.Port)
	if err := http.ListenAndServe(":"+cmd.Port, reader.New(s)); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error starting server: %s\n", err)
			return
		}
		log.Println("server stopped")
	}
}
