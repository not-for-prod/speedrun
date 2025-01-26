package peach_repository

import (
	"context"

	peach "github.com/not-for-prod/speedrun/cmd/crud/example/in"
	"go.opentelemetry.io/otel"

	sql "github.com/not-for-prod/speedrun/cmd/crud/example/out/peach/sql"
)

func (r *PeachRepository) Create(ctx context.Context, peach peach.Peach) (int, error) {
	ctx, span := otel.Tracer("").Start(ctx, "/repository/peach/create.go")
	defer span.End()

	var dbID int

	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx,
		&dbID,
		sql.Insert,
		peach.Id,
		peach.Size,
		peach.Juice,
	)
	if err != nil {

	}

	return dbID, nil
}
