package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/timmbarton/utils/tracing"
)

func New(pg *sqlx.DB) *Repository {
	return &Repository{
		pg: pg,
	}
}

type Repository struct {
	pg *sqlx.DB
}

func (r *Repository) GetFooBar(ctx context.Context) (res []int64, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer span.End()

	err = r.pg.SelectContext(ctx, &res, queryGetFooBar)
	if err != nil {
		return res, err
	}

	return
}
