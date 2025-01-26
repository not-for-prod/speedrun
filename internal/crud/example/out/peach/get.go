
package peach_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	peach "github.com/not-for-prod/speedrun/internal/crud/example/in"
	sql "github.com/not-for-prod/speedrun/internal/crud/example/out/peach/sql"
	
)

func (r *PeachRepository) Get(ctx context.Context, id int) (.Peach, error) {
	ctx, span := otel.Tracer("").Start(ctx, "/repository/peach/get.go")
	defer span.End()

	var dbPeach dbPeach

	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx, 
		&dbPeach,
		sql.Get,
		id,
	)	
	if err != nil {

	}

	return dbPeach.toEntity(), nil
}
