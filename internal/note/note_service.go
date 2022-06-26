package note

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("NService - Create - s.repo.Store: %w", err)
	}

	return note, nil
}
