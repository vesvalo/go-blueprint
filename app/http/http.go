package http

import (
	"context"
	"github.com/Nerufa/go-shared/logger"
	"github.com/Nerufa/go-shared/metric"
	"github.com/Nerufa/go-shared/provider"
	"github.com/Nerufa/go-shared/tracing"
	"github.com/go-chi/chi"
	"net/http"
)

// Http
type Http struct {
	ctx     context.Context
	log     logger.Logger
	cfg     Config
	metric  metric.Scope
	tracing tracing.Tracer
	mux     *chi.Mux
}

// ListenAndServe
func (m *Http) ListenAndServe(bind ...string) (err error) {

	bindAdrr := m.cfg.Bind

	if len(bind) > 0 && len(bind[0]) > 0 {
		bindAdrr = bind[0]
	}

	server := &http.Server{
		Addr:    bindAdrr,
		Handler: m.mux,
	}

	m.log.Info("start listen and serve http at %v", logger.Args(bindAdrr))

	go func() {
		<-m.ctx.Done()
		m.log.Info("context cancelled, shutdown is raised")
		if e := server.Shutdown(context.Background()); e != nil {
			m.log.Emergency("graceful shutdown error, %v", logger.Args(e))
		}
	}()

	if err = server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			m.log.Emergency("server is shutdown with error, %v", logger.Args(err))
		} else {
			err = nil
		}
	}
	return
}

type Cors struct {
	Allowed []string
}

// Config
type Config struct {
	Debug bool
	Bind  string
	Cors  Cors
}

// New
func New(ctx context.Context, mux *chi.Mux, set provider.AwareSet, cfg Config) *Http {
	return &Http{
		ctx:     ctx,
		cfg:     cfg,
		metric:  set.Metric,
		tracing: set.Tracer,
		mux:     mux,
		log:     set.Logger.WithFields(logger.Fields{"service": prefix}),
	}
}
