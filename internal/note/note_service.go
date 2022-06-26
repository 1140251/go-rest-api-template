package note

import (
	"context"

	"github.com/pkg/errors"

	"github.com/1140251/go-template/internal/note/entity"
)

// NService -.
type NService struct {
	repo Repo
}

// NewNoteService -.
func NewNoteService(r Repo) Service {
	return &NService{
		repo: r,
	}
}

func (s *NService) Create(ctx context.Context, note *entity.Note) (*entity.Note, error) {
	err := s.repo.Store(ctx, note)
	if err != nil {
		return nil, errors.Wrap(err, "NService - Create - s.repo.Store")
	}

	return note, nil
}
