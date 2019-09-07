package graphql

import (
	"context"
	"github.com/Nerufa/go-shared/provider"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/gqlerror"
	"net/http"
	"strings"
	"time"

	gqlgen "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/Nerufa/go-blueprint/generated/graphql"
	"github.com/Nerufa/go-shared/logger"
	"github.com/Nerufa/go-shared/metric"
	"github.com/Nerufa/go-shared/tracing"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"

	gqErrs "github.com/Nerufa/go-blueprint/pkg/graphql/errors"
)

var errInternalServer = errors.New("internal server error")

// GraphQL
type GraphQL struct {
	ctx      context.Context
	resolver *graphql.Config
	log      logger.Logger
	cfg      Config
	metric   metric.Scope
	tracing  tracing.Tracer
}

// Use
func (g *GraphQL) Use(router *chi.Mux) {
	router.Use(g.cfg.Middleware...)
}

// Routers
func (g *GraphQL) Routers(router *chi.Mux) {

	upgrader := websocket.Upgrader{}

	options := []handler.Option{
		handler.WebsocketUpgrader(upgrader),
		handler.IntrospectionEnabled(g.cfg.Introspection),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			if e, ok := err.(error); ok {
				return gqErrs.WrapPanicErr(e)
			}
			g.log.Alert("unhandled panic, err: %v", logger.Args(err))
			return nil
		}),
		handler.ErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
			// panic error
			if _, ok := e.(*gqErrs.PanicErr); ok {
				g.log.Alert("recover on middleware, err: %v", logger.Args(e))
				goto done
			}
			// client error
			g.log.Error("internal server error, err: %v", logger.Args(e))
			if _, ok := e.(*gqErrs.ClientErr); !ok {
				e = errInternalServer
			}
		done:
			return gqlgen.DefaultErrorPresenter(ctx, e)
		}),
	}

	if g.cfg.Debug {
		router.Handle(g.cfg.Playground.Route, handler.Playground(g.cfg.Playground.Name, g.cfg.Playground.Endpoint))
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
		options = append(options, handler.RequestMiddleware(func(ctx context.Context, next func(ctx context.Context) []byte) []byte {
			startTime := time.Now()
			rc := gqlgen.GetRequestContext(ctx)
			resp := next(ctx)
			e := strings.ReplaceAll(rc.Errors.Error(), "\n", " ")
			g.log.Debug("\nVARS:\n%+v\nQUERY:\n%v\nRESPONSE:\n%v\nERROR:\n%v\n",
				logger.Args(rc.Variables, strings.TrimRight(rc.RawQuery, "\n"), string(resp), e),
				logger.WithFields(logger.Fields{
					"time": time.Since(startTime).String(),
				}),
			)
			return resp
		}))
	}

	router.Handle(g.cfg.Route,
		handler.GraphQL(
			graphql.NewExecutableSchema(*g.resolver),
			options...,
		),
	)
}

type PlaygroundCfg struct {
	Route    string
	Name     string
	Endpoint string
}

// Config
type Config struct {
	Debug         bool
	Introspection bool
	Middleware    []func(http.Handler) http.Handler
	Playground    PlaygroundCfg
	Route         string
}

// New
func New(ctx context.Context, resolver graphql.Config, set provider.AwareSet, cfg Config) *GraphQL {
	return &GraphQL{
		ctx:      ctx,
		resolver: &resolver,
		cfg:      cfg,
		metric:   set.Metric,
		tracing:  set.Tracer,
		log:      set.Logger.WithFields(logger.Fields{"service": Prefix}),
	}
}
