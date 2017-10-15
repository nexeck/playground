package cmd

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/valve"
	mw "github.com/nexeck/playground/pkg/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var bind string

func init() {
	cmdServer := cmdServer()
	cmdServer.Flags().StringVarP(&bind, "bind", "b", ":8080", "Bind")
	RootCmd.AddCommand(cmdServer)
}

// NewCmdVersion adds version command
func cmdServer() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start server",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
}

func startServer() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	valv := valve.New()
	baseCtx := valv.Context()

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(mw.NewZapLogger(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/heartbeat"))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Mount("/debug", middleware.Profiler())

	r.Mount("/metrics", promhttp.Handler())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("oops")
	})

	srv := http.Server{
		Addr:    bind,
		Handler: chi.ServerBaseContext(baseCtx, r),
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			fmt.Println("shutting down..")

			// first valv
			valv.Shutdown(20 * time.Second)

			// create context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			// start http shutdown
			srv.Shutdown(ctx)

			// verify, in worst case call cancel via defer
			select {
			case <-time.After(21 * time.Second):
				fmt.Println("not all connections done")
			case <-ctx.Done():

			}
		}
	}()

	logger.Error("Http Server",
		zap.Error(srv.ListenAndServe()),
	)
}
