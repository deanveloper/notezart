package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/deanveloper/notezart/api"
)

// Start configures and starts the webserver.
// Blocks until webserver shuts down.
func Start(cfg api.Config) {
	startServer(cfg)
}

func startServer(cfg api.Config) {
	var err error
	if cfg.Web.CertFile != "" && cfg.Web.KeyFile != "" {
		err = http.ListenAndServeTLS(
			cfg.Web.Addr,
			cfg.Web.CertFile,
			cfg.Web.KeyFile,
			nil,
		)
	} else {
		err = http.ListenAndServe(cfg.Web.Addr, nil)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occured in webserver: %s", err)
	}
}
