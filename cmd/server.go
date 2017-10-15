package cmd

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"net/http"
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

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
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

	logger.Info("Starting Server",
		zap.String("Bind", bind),
	)

	logger.Error("Http Server",
		zap.Error(http.ListenAndServe(bind, r)),
	)
}
