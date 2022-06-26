package note_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	internalNote "github.com/1140251/go-template/internal/note"
	"github.com/1140251/go-template/internal/note/entity"
)

type test struct {
	name string
	mock func()
	res  *entity.Note
	err  error
	send interface{}
}

var errInternalServErr = errors.New("internal server error")

func note(t *testing.T) (internalNote.Service, *MockRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockRepo(mockCtl)

	noteService := internalNote.NewNoteService(repo)

	return noteService, repo
}

func TestNoteService_Create(t *testing.T) {
	t.Parallel()

	notes, repo := note(t)

	tests := []test{
		{
			name: "empty result",
			send: entity.Note{Message: "test"},
			mock: func() {
				repo.EXPECT().Store(context.Background(), &entity.Note{Message: "test"}).Return(nil)
			},
			res: &entity.Note{Message: "test"},
			err: nil,
		},
		{
			name: "repo error",
			send: entity.Note{},
			mock: func() {
				repo.EXPECT().Store(context.Background(), &entity.Note{}).Return(errInternalServErr)
			},
			res: nil,
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			noteS, ok := tc.send.(entity.Note)
			if !ok {
				t.Error("send is not note")
			}

			res, err := notes.Create(context.Background(), &noteS)

			if !errors.Is(err, tc.err) {
				t.Errorf("GetByPath() error = %v, wantErr %v", err, tc.err)
			}

			if !reflect.DeepEqual(res, tc.res) {
				t.Errorf("GetByPath() got = %s, want %s", res, tc.res)
			}
		})
	}
}
