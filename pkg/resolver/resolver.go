package resolver

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	graphql1 "github.com/Nerufa/go-blueprint/generated/graphql"
	"github.com/Nerufa/go-blueprint/pkg/db/domain"
	"github.com/Nerufa/go-blueprint/pkg/db/trx"
	gqErrs "github.com/Nerufa/go-blueprint/pkg/graphql/errors"
	"github.com/Nerufa/go-shared/logger"
	"github.com/Nerufa/go-shared/metric"
	"github.com/Nerufa/go-shared/provider"
	"github.com/Nerufa/go-shared/tracing"
)

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// Config custom graphql settings resolvers
type Config struct {
	Debug bool
}

// Resolver config graphql resolvers
type Resolver struct {
	ctx    context.Context
	log    logger.Logger
	tracer tracing.Tracer
	cfg    Config
	metric metric.Scope
	repo   Repo
	trx    *trx.Manager
}

// Mutation returns root graphql mutation resolver
func (r *Resolver) Mutation() graphql1.MutationResolver {
	return &mutationResolver{r}
}

// Query returns root graphql query resolver
func (r *Resolver) Query() graphql1.QueryResolver {
	return &queryResolver{r}
}

// AddErrorf is a convenience method for adding an error to the current response
func (r *Resolver) AddDebugErrorf(ctx context.Context, format string, args ...interface{}) {
	if r.cfg.Debug {
		graphql.AddError(ctx, gqErrs.WrapClientErr(fmt.Errorf(format, args...)))
	}
}

// Repo
type Repo struct {
	List domain.ListRepo
}

// New returns instance of config graphql resolvers
func New(ctx context.Context, set provider.AwareSet, appSet AppSet, cfg Config) graphql1.Config {
	c := graphql1.Config{
		Resolvers: &Resolver{
			ctx:    ctx,
			log:    set.Logger.WithFields(logger.Fields{"service": Prefix}),
			metric: set.Metric,
			tracer: set.Tracer,
			cfg:    cfg,
			repo:   appSet.Repo,
			trx:    appSet.Trx,
		},
	}
	return c
}
