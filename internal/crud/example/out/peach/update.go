package peach_repository

import (
	"context"

	peach "github.com/not-for-prod/speedrun/internal/crud/example/in"
	"github.com/not-for-prod/speedrun/internal/crud/example/out/peach/sql"
	"go.opentelemetry.io/otel"
)

func (r *PeachRepository) Update(ctx context.Context, peach peach.Peach) error {
	ctx, span := otel.Tracer("").Start(ctx, "/repository/peach/update.go")
	defer span.End()

	_, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(
		ctx,
		sql.Update,
		peach.Id,
		peach.Size,
		peach.Juice,
	)
	if err != nil {

	}

	return nil
}
