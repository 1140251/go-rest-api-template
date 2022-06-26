package repo

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/1140251/go-template/internal/note"
	"github.com/1140251/go-template/internal/note/entity"
	"github.com/1140251/go-template/pkg/postgres"
)

// NoteRepository -.
type NoteRepository struct {
	*postgres.Postgres
}

func (n NoteRepository) Store(ctx context.Context, e *entity.Note) error {
	sql, args, err := n.Builder.
		Insert("note").
		Columns("message, created_at, updated_at").
		Values(e.Message, time.Now(), time.Now()).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("NoteRepository - Store - r.Builder: %w", err)
	}

	_, err = n.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("NoteRepository - Store - r.Pool.Exec: %w", err)
	}

	return nil
}

// NewNoteRepo -.
func NewNoteRepo(pg *postgres.Postgres) note.Repo {
	return &NoteRepository{pg}
}
