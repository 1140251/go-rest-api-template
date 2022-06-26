// Package usecase implements application business logic. Each logic group in own file.
package note

import (
	"context"

	"github.com/1140251/go-template/internal/note/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=note_test

type (
	// Service -.
	Service interface {
		Create(context.Context, *entity.Note) (*entity.Note, error)
	}

	// Repo -.
	Repo interface {
		Store(context.Context, *entity.Note) error
	}
)
