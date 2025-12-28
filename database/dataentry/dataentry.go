package dataentry

import (
	"context"
	"database/sql"

	"github.com/Darkness4/auth-htmx/database"
)

type Repository struct {
	*database.Queries
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Queries: database.New(db)}
}

func (r *Repository) Create(ctx context.Context, datapointID int64, typ string, textValue sql.NullString, intValue sql.NullInt64) (database.Dataentry, error) {
	return r.CreateDataentry(ctx, database.CreateDataentryParams{
		DatapointID: datapointID,
		Type:        typ,
		TextValue:   textValue,
		IntValue:    intValue,
	})
}

func (r *Repository) Get(ctx context.Context, id int64) (database.Dataentry, error) {
	return r.GetDataentry(ctx, id)
}

func (r *Repository) ListByDatapoint(ctx context.Context, datapointID int64) ([]database.Dataentry, error) {
	return r.ListDataentriesByDatapoint(ctx, datapointID)
}

func (r *Repository) Update(ctx context.Context, id int64, typ string, textValue sql.NullString, intValue sql.NullInt64) (database.Dataentry, error) {
	return r.UpdateDataentry(ctx, database.UpdateDataentryParams{
		Type:      typ,
		TextValue: textValue,
		IntValue:  intValue,
		ID:        id,
	})
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	return r.DeleteDataentry(ctx, id)
}
