package clients

import (
	"context"

	"github.com/auenkr/go-templates/server/gen/db"
	"github.com/auenkr/go-templates/server/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func NewDBQueries(cfg *config.Config, lc fx.Lifecycle) (*db.Queries, error) {
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseConnectionString)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			pool.Close()
			return nil
		},
	})

	return db.New(pool), nil
}
