/*
Copyright © 2024 Kai Blin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package web

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kblin.org/shortis/internal/model"
)

type appication struct {
	Model model.DbModel
	Mux   *http.ServeMux
}

func Run(address string, port int, model model.DbModel) {
	slog.Info("Starting web server", "address", address, "port", port)
	full_address := fmt.Sprintf("%s:%d", address, port)

	mux := http.NewServeMux()

	app := appication{
		Model: model,
		Mux:   mux,
	}

	app.setupRoutes()

	srv := &http.Server{
		Addr:         full_address,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	shutdownError := make(chan error)

	// Gracefully shut down
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit
		slog.Info("caught signal, shutting down", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		slog.Error(err.Error())
		panic(err.Error())
	}

	err = <-shutdownError
	if err != nil {
		slog.Error(err.Error())
		panic(err.Error())
	}

	slog.Info("stopped server", "address", address, "port", port)

}
