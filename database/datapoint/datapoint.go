package datapoint

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

func (r *Repository) Create(ctx context.Context, userID []byte, name string) (database.Datapoint, error) {
	return r.CreateDatapoint(ctx, database.CreateDatapointParams{
		UserID: userID,
		Name:   name,
	})
}

func (r *Repository) Get(ctx context.Context, id int64) (database.Datapoint, error) {
	return r.GetDatapoint(ctx, id)
}

func (r *Repository) ListByUser(ctx context.Context, userID []byte) ([]database.Datapoint, error) {
	return r.ListDatapointsByUser(ctx, userID)
}

func (r *Repository) UpdateName(ctx context.Context, id int64, name string) (database.Datapoint, error) {
	return r.UpdateDatapointName(ctx, database.UpdateDatapointNameParams{
		Name: name,
		ID:   id,
	})
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	return r.DeleteDatapoint(ctx, id)
}
