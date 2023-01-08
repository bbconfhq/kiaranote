package handler

import (
	"github.com/bbconfhq/kiaranote/internal/common"
	"github.com/bbconfhq/kiaranote/internal/constant"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetNotesRequest struct{}
type GetNotesResponse struct{}

// V1GetNotes   godoc
// @Summary      Get notes
// @Description  Get list of notes (exclude content), divided by public and private, role >= USER
// @Tags         Note
// @Accept       json
// @Produce      json
// @Success      200		{object}	[]GetNotesResponse
// @Failure      400		{object}	nil
// @Failure      401		{object}	nil
// @Failure      500		{object}	nil
// @Router       /note [get]
func V1GetNotes(_ *GetNotesRequest, _ echo.Context) common.Response {
	return common.Response{
		Data: nil,
		Code: http.StatusOK,
	}
}

type PostNoteRequest struct{}
type PostNoteResponse struct{}

// V1PostNote   godoc
// @Summary      Post note
// @Description  Register new note, role >= ADMIN
// @Tags         Note
// @Accept       json
// @Produce      json
// @Param        req		body		PostNoteRequest	true	"Note content, parent id and public/private"
// @Success      201		{object}	nil
// @Failure      400		{object}	nil
// @Failure      401		{object}	nil
// @Failure      500		{object}	nil
// @Router       /note [post]
func V1PostNote(_ *PostNoteRequest, _ echo.Context) common.Response {
	return common.Response{
		Code: http.StatusCreated,
	}
}

type GetNoteRequest struct{}
type GetNoteResponse struct{}

// V1GetNote   godoc
// @Summary      Get note
// @Description  Get note content, role >= USER
// @Tags         Note
// @Accept       json
// @Produce      json
// @Param        note_id	path		uint			true	"Note Id"
// @Success      200		{object}	GetNoteResponse
// @Failure      400		{object}	nil
// @Failure      401		{object}	nil
// @Failure      500		{object}	nil
// @Router       /note/{note_id} [get]
func V1GetNote(_ *GetNoteRequest, _ echo.Context) common.Response {
	return common.Response{
		Code:  http.StatusInternalServerError,
		Error: constant.ErrInternal,
	}
}

type PutNoteRequest struct{}
type PutNoteResponse struct{}

// V1PutNote   godoc
// @Summary      Put note
// @Description  Put note by note_id, role >= ADMIN
// @Tags         Note
// @Accept       json
// @Produce      json
// @Param        note_id	path		uint			true	"Note Id"
// @Success      200		{object}	[]GetNotesResponse
// @Failure      400		{object}	nil
// @Failure      401		{object}	nil
// @Failure      500		{object}	nil
// @Router       /note/{note_id} [put]
func V1PutNote(req *PutNoteRequest, _ echo.Context) common.Response {
	return common.Response{
		Data: nil,
		Code: http.StatusOK,
	}
}

type DeleteNoteRequest struct{}
type DeleteNoteResponse struct{}

// V1DeleteNote   godoc
// @Summary      Delete note
// @Description  Delete note by note_id, role >= ADMIN
// @Tags         Note
// @Accept       json
// @Produce      json
// @Param        note_id	path		uint			true	"Note Id"
// @Success      200		{object}	nil
// @Failure      400		{object}	nil
// @Failure      401		{object}	nil
// @Failure      500		{object}	nil
// @Router       /note/{note_id} [delete]
func V1DeleteNote(_ *DeleteNoteRequest, _ echo.Context) common.Response {
	return common.Response{
		Data: nil,
		Code: http.StatusOK,
	}
}
