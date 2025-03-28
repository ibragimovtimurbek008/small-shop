package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/ibragimovtimurbek008/small-shop/internal/handlers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	authCmd = &cobra.Command{
		Use:   "auth",
		Short: "Runs auth service",
		Run: func(cmd *cobra.Command, args []string) {
			initViper()

			var (
				router = mux.NewRouter()
				web    = http.Server{
					Addr:    viper.GetString("APPLICATION_ADDR"),
					Handler: router,
				}

				serverContext = context.Background()
			)

			router.HandleFunc("/auth/v1/login", handlers.LoginHandler).
				Methods(http.MethodPost, http.MethodOptions)

			router.HandleFunc("/auth/v1/verify", handlers.VerifyTokenHandler).
				Methods(http.MethodPost, http.MethodOptions)

			router.HandleFunc("/v1/_health", handlers.HealthHandler)

			go func() {
				slog.Debug("starting web server", "addr", viper.GetString("APPLICATION_ADDR"))

				if err := web.ListenAndServe(); err != http.ErrServerClosed {
					panic(err)
				}
			}()

			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

			if s, ok := <-signalChan; ok {
				slog.Info("signal found", "signal", s.String())
			} else {
				slog.Info("signal channel closed")
			}

			slog.Info("terminating web server...")

			if err := web.Shutdown(serverContext); err != nil {
				slog.With("err", err).Error("web server shutdown error")
			}

			slog.Info("web server terminated")
		},
	}
)

func initViper() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	viper.SetConfigFile("local.yaml")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config file: %w", err))
	}
}
