package handler

import (
	"github.com/bbconfhq/kiaranote/internal/common"
	"github.com/bbconfhq/kiaranote/internal/constant"
	"github.com/bbconfhq/kiaranote/internal/dao"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type GetNotesRequest struct{}

type GetNotesResponse struct {
	Public  []GetNoteResponse `json:"public"`
	Private []GetNoteResponse `json:"private"`
}

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
func V1GetNotes(_ *GetNotesRequest, c echo.Context) common.Response {
	sess, err := session.Get("session", c)
	if err != nil {
		panic(err)
	}

	userId := sess.Values["user_id"]
	if userId == nil {
		panic("cannot get user_id from session")
	}

	notes := GetNotesResponse{
		Public:  make([]GetNoteResponse, 0),
		Private: make([]GetNoteResponse, 0),
	}

	repo := dao.GetRepo()
	publicRows, err := repo.Reader().Queryx(
		`SELECT
    		n.id, h.parent_note_id, n.user_id, n.title, n.content, n.is_private, n.create_dt, n.update_dt
		FROM
		    note n
		INNER JOIN
			note_hierarchy h ON n.id = h.note_id
		WHERE
			delete_dt is NULL
			AND is_private is FALSE
		ORDER BY
			h.order`,
	)
	if err != nil {
		panic(err)
	}

	privateRows, err := repo.Reader().Queryx(
		`SELECT
    		n.id, h.parent_note_id, n.user_id, n.title, n.content, n.is_private, n.create_dt, n.update_dt
		FROM
		    note n
		INNER JOIN
		    note_hierarchy h on n.id = h.note_id
		WHERE
			delete_dt is NULL
		  	AND user_id = ? 
		    AND is_private is TRUE
		ORDER BY
			h.order`, userId,
	)
	if err != nil {
		panic(err)
	}

	for publicRows.Next() {
		var note GetNoteResponse
		err := publicRows.StructScan(&note)
		if err != nil {
			panic(err)
		}
		notes.Public = append(notes.Public, note)
	}

	for privateRows.Next() {
		var note GetNoteResponse
		err := privateRows.StructScan(&note)
		if err != nil {
			panic(err)
		}
		notes.Private = append(notes.Private, note)
	}

	return common.Response{
		Data: notes,
		Code: http.StatusOK,
	}
}

type PostNoteRequest struct {
	ParentNoteId int64  `json:"parent_note_id" validate:"-"`
	Title        string `json:"title" validate:"required"`
	Content      string `json:"content" validate:"required"`
	IsPrivate    bool   `json:"is_private" validate:"required"`
}

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
type GetNoteResponse struct {
	Id           int64     `json:"id"`
	ParentNoteId int64     `db:"parent_note_id" json:"parent_note_id"`
	UserId       int64     `db:"user_id" json:"user_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	IsPrivate    bool      `db:"is_private" json:"is_private"`
	CreateDt     time.Time `db:"create_dt" json:"create_dt"`
	UpdateDt     time.Time `db:"update_dt" json:"update_dt"`
}

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
func V1PutNote(_ *PutNoteRequest, _ echo.Context) common.Response {
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
