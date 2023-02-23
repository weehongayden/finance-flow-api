package main

import (
	"fmt"
	"net/http"
	"time"
	"weehongayden/finance-flow-app/internal/config"
)

func serve(app config.Config) {
	svr := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", app.Server.Address, app.Server.Port),
		Handler:      routes(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	svr.ListenAndServe()
}
