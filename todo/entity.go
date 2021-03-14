package todo

import (
	"context"
	"time"
)

// Entity db model
type Entity struct {
	tableName   struct{} `pg:"todos"`
	ID          int      `pg:"id,pk"`
	Title       string   `pg:"title,notnull"`
	Description string
	Completed   bool      `sql:"completed,notnull"`
	CreatedAt   time.Time `pg:"created_at,default:now()"`
	UpdatedAt   time.Time `pg:"updated_at,default:now()"`
}

// BeforeUpdate update updated_at timestamp
func (t *Entity) BeforeUpdate(ctx context.Context) (context.Context, error) {
	t.UpdatedAt = time.Now()
	return ctx, nil
}

// ToModel conversion
func (t *Entity) ToModel() Todo {
	return Todo{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
		CreatedAt:   t.CreatedAt.Format(time.Stamp),
		UpdatedAt:   t.UpdatedAt.Format(time.Stamp),
	}
}
