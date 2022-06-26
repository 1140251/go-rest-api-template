package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/1140251/go-template/internal/note"
	"github.com/1140251/go-template/internal/note/entity"
	"github.com/1140251/go-template/pkg/logger"
)

type noteRoutes struct {
	noteService note.Service
	l           logger.Interface
}

func newNoteRoutes(handler *gin.RouterGroup, s note.Service, l logger.Interface) {
	r := &noteRoutes{s, l}

	h := handler.Group("/note")
	{
		h.POST("", r.create)
	}
}

type ListResponse struct {
	Notes []entity.Note `json:"history"`
}

type createNoteRequest struct {
	Message string `json:"message" binding:"required"  example:"example message"`
}

// @Summary     Create note
// @Description Create a new note
// @ID          notes
// @Accept      json
// @Produce     json
// @Param       request body createNoteRequest true "Set up translation"
// @Success     200 {object} entity.Note
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /note [post]
func (r *noteRoutes) create(ctx *gin.Context) {
	var request createNoteRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - create")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")

		return
	}

	createdNote, err := r.noteService.Create(ctx, entity.NewNote(request.Message))
	if err != nil {
		r.l.Error(err, "http - v1 - create")
		errorResponse(ctx, http.StatusInternalServerError, "note service problems")

		return
	}

	ctx.JSON(http.StatusOK, createdNote.ID)
}
