package internal

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/julienschmidt/httprouter"
	"github.com/rcrowley/go-metrics"
	"github.com/rcrowley/go-metrics/exp"
	"github.com/thoas/stats"
	"github.com/unrolled/logger"
)

type counters struct {
	r metrics.Registry
}

func newCounters() *counters {
	counters := &counters{
		r: metrics.NewRegistry(),
	}
	return counters
}

func (c *counters) Inc(name string) {
	metrics.GetOrRegisterCounter(name, c.r).Inc(1)
}

func (c *counters) Dec(name string) {
	metrics.GetOrRegisterCounter(name, c.r).Dec(1)
}

func (c *counters) IncBy(name string, n int64) {
	metrics.GetOrRegisterCounter(name, c.r).Inc(n)
}

func (c *counters) DecBy(name string, n int64) {
	metrics.GetOrRegisterCounter(name, c.r).Dec(n)
}

type server struct {
	// Config
	config *config

	// Mux / Router
	router *httprouter.Router

	// Logger
	logger *logger.Logger

	// Stats/Metrics
	counters *counters
	stats    *stats.Stats
}

func (s *server) Run() error {
	return http.ListenAndServe(
		s.config.bind,
		s.logger.Handler(
			s.stats.Handler(
				gziphandler.GzipHandler(
					s.router,
				),
			),
		),
	)
}

func (s *server) initRoutes() {
	s.router.Handler("GET", "/metrics", exp.ExpHandler(s.counters.r))
	s.router.GET("/stats", s.statsHandler())
	s.router.GET("/health", s.healthHandler())

	s.router.GET("/", s.indexHandler())
}

func NewServer(options ...Option) (*server, error) {
	server := &server{
		// Config
		config: newDefaultConfig(),

		// Mux / Router
		router: httprouter.New(),

		// Logger
		logger: logger.New(logger.Options{
			Prefix:               "petstore",
			RemoteAddressHeaders: []string{"X-Forwarded-For"},
		}),

		// Stats/Metrics
		counters: newCounters(),
		stats:    stats.New(),
	}

	for _, opt := range options {
		if err := opt(server.config); err != nil {
			return nil, err
		}
	}

	server.initRoutes()

	return server, nil
}
