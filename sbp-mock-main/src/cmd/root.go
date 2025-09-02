package cmd

import (
	"crypto/tls"
	"espp-mock/configs"
	"espp-mock/espp"
	"espp-mock/metrics"
	"espp-mock/server"
	sbpHandler "espp-mock/server/sbp"
	sbpUsecase "espp-mock/usecase/sbp"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "espp-mock",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configs.New(cfgFile)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		log.Debug().Msgf("%+v", config)

		// ctx, cancel := context.WithCancel(context.Background())
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient := &http.Client{Transport: tr}
		ipsAdapter := espp.New(config, httpClient)
		sbpUsecase := sbpUsecase.New(config, ipsAdapter)
		sbpHandler := sbpHandler.NewHandler(sbpUsecase)
		router := &server.RouterBuilder{
			SBPHandler: sbpHandler,
		}

		srv := &http.Server{
			Addr:    config.ListenAddress,
			Handler: router.Build(),
		}

		metrics.InitInflux(config.Influx)

		go func() {
			log.Info().Str("address", config.ListenAddress).Msg("server start")
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal().Err(err).Msg("listen server")
			}
		}()

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		signal := <-stop
		log.Info().Str("signal", signal.String()).Msg("received os signal")

		log.Info().Msg("server stopped")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./espp-mock.yaml)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("failed to start espp-mock")
	}
}
