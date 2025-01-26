
package peach_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	peach "github.com/not-for-prod/speedrun/internal/crud/example/in"
	sql "github.com/not-for-prod/speedrun/internal/crud/example/out/peach/sql"
	
)

func (r *PeachRepository) Delete(ctx context.Context, id int) error {
	ctx, span := otel.Tracer("").Start(ctx, "/repository/peach/get.go")
	defer span.End()

	_, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(
		ctx,
		sql.Delete,
		id,
	)	
	if err != nil {
		
	}

	return nil
}
