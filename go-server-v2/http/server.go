package http

import (
	"context"
	"net/http"
	"time"

	"go-server-v2/config"
	"go-server-v2/domain"
	"go-server-v2/http/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jesperkha/notifier"
)

type Server struct {
	mux     *chi.Mux
	config  *config.Config
	cleanup func()
}

type Dependencies struct {
	Logger domain.Logger
	Notif  *notifier.Notifier
	Config *config.Config
}

func Run(deps Dependencies) {
	mux := chi.NewMux()

	mux.Use(requestLogger(deps.Logger))
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Mount routes
	mux.Get("/", routes.StaticFile("web/index.html"))

	cleanup := func() {
		// Add cleanup here...
		// Runs after server shutdown and before exit
	}

	s := &Server{
		mux:     mux,
		config:  deps.Config,
		cleanup: cleanup,
	}

	s.ListenAndServe(deps.Notif, deps.Logger)
}

func (s *Server) ListenAndServe(notif *notifier.Notifier, logger domain.Logger) {
	done, finish := notif.Register()
	ctx := context.Background()

	server := &http.Server{
		Addr:    s.config.Port,
		Handler: s.mux,
	}

	go func() {
		<-done
		if err := server.Shutdown(ctx); err != nil {
			logger.Error(ctx, "unexpected shutdown", "error", err.Error())
		}

		logger.Info(ctx, "cleaning up")
		s.cleanup()
		finish()
	}()

	logger.Info(ctx, "server is running", "port", s.config.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error(ctx, "server closed with error", "error", err.Error())
	}
}

func requestLogger(logger domain.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := &responseWriter{ResponseWriter: w, status: 200}

			next.ServeHTTP(ww, r)

			duration := time.Since(start)
			logger.Info(r.Context(), "http request",
				"method", r.Method,
				"path", r.URL.Path,
				"status", ww.status,
				"duration", duration.String(),
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
			)
		})
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
