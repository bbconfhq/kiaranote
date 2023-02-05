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
    		n.id, h.parent_note_id, n.user_id, n.title, "" as content, n.is_private, n.create_dt, n.update_dt
		FROM
		    note n
		INNER JOIN
			note_hierarchy h ON n.id = h.note_id
		WHERE
			n.delete_dt is NULL
			AND n.is_private is FALSE
		ORDER BY
			h.order`,
	)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	privateRows, err := repo.Reader().Queryx(
		`SELECT
    		n.id, h.parent_note_id, n.user_id, n.title, "" as content, n.is_private, n.create_dt, n.update_dt
		FROM
		    note n
		INNER JOIN
		    note_hierarchy h on n.id = h.note_id
		WHERE
			n.delete_dt is NULL
		  	AND n.user_id = ? 
		    AND n.is_private is TRUE
		ORDER BY
			h.order`, userId,
	)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	for publicRows.Next() {
		var note GetNoteResponse
		err := publicRows.StructScan(&note)
		if err != nil {
			return common.Response{
				Code:  http.StatusInternalServerError,
				Error: constant.ErrInternal,
			}
		}
		notes.Public = append(notes.Public, note)
	}

	for privateRows.Next() {
		var note GetNoteResponse
		err := privateRows.StructScan(&note)
		if err != nil {
			return common.Response{
				Code:  http.StatusInternalServerError,
				Error: constant.ErrInternal,
			}
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
	Title        string `json:"title" validate:"required,lt=128"`
	Content      string `json:"content" validate:"required"`
	IsPrivate    bool   `json:"is_private" validate:"required,boolean"`
	Order        int64  `json:"order" validate:"required,number,gt=0"`
}

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
func V1PostNote(req *PostNoteRequest, c echo.Context) common.Response {
	sess, err := session.Get("session", c)
	if err != nil {
		panic(err)
	}

	userId := sess.Values["user_id"]
	if userId == nil {
		panic("cannot get user_id from session")
	}

	repo := dao.GetRepo()
	result, err := repo.Writer().Exec(
		`INSERT INTO note (user_id, title, content, is_private) VALUES (?, ?, ?, ?)`, userId, req.Title, req.Content, req.IsPrivate,
	)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	noteId, err := result.LastInsertId()
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	// If there is no request for parent_note_id, set note as root
	if req.ParentNoteId == 0 {
		req.ParentNoteId = noteId
	}
	_, err = repo.Writer().Exec(
		"INSERT INTO note_hierarchy (`note_id`, `parent_note_id`, `order`) VALUES (?, ?, ?)", noteId, req.ParentNoteId, req.Order,
	)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	return common.Response{
		Data: noteId,
		Code: http.StatusCreated,
	}
}

type GetNoteRequest struct {
	Id int64 `param:"note_id"`
}
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
func V1GetNote(req *GetNoteRequest, c echo.Context) common.Response {
	sess, err := session.Get("session", c)
	if err != nil {
		panic(err)
	}

	userId := sess.Values["user_id"]
	if userId == nil {
		panic("cannot get user_id from session")
	}

	repo := dao.GetRepo()
	var note GetNoteResponse
	if err := repo.Reader().QueryRowx(
		`SELECT
    		n.id, h.parent_note_id, n.user_id, n.title, n.content, n.is_private, n.create_dt, n.update_dt
		FROM
		    note n
		INNER JOIN
			note_hierarchy h ON n.id = h.note_id
		WHERE
		    (
			    n.id = ?
				AND n.delete_dt is NULL
			)
			AND
		    (
				n.is_private is FALSE
				OR n.user_id = ?
			)
		ORDER BY
			h.order`, req.Id, req.Id, userId).StructScan(&note); err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	return common.Response{
		Data: note,
		Code: http.StatusOK,
	}
}

type PutNoteRequest struct {
	Id           int64  `param:"note_id"`
	ParentNoteId int64  `json:"parent_note_id" validate:"required,number,gt=0"`
	Title        string `json:"title" validate:"required,lt=128"`
	Content      string `json:"content" validate:"required"`
	IsPrivate    bool   `json:"is_private" validate:"required,boolean"`
	Order        int64  `json:"order" validate:"required,number,gt=0"`
}

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
func V1PutNote(req *PutNoteRequest, c echo.Context) common.Response {
	sess, err := session.Get("session", c)
	if err != nil {
		panic(err)
	}

	userId := sess.Values["user_id"]
	if userId == nil {
		panic("cannot get user_id from session")
	}

	repo := dao.GetRepo()
	query := `UPDATE note SET title = ?, content = ?, is_private = ?
            WHERE (id = ? AND delete_dt is NULL) AND (is_private = FALSE OR user_id = ?)`
	result, err := repo.Writer().Exec(query, req.Title, req.Content, req.IsPrivate, req.Id)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	// 지워지지 않았고, Public이 아닐 경우에는 user_id 확인 후 다를 경우 BadRequest 에러
	affectedRows, err := result.RowsAffected()
	if affectedRows < 1 {
		return common.Response{
			Code:  http.StatusBadRequest,
			Error: constant.ErrBadRequest,
		}
	}

	query = "UPDATE note_hierarchy SET `parent_note_id` = ?, `order` = ? WHERE `note_id` = ?"
	_, err = repo.Writer().Exec(query, req.ParentNoteId, req.Order, req.Id)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	return common.Response{
		Data: req.Id,
		Code: http.StatusOK,
	}
}

type DeleteNoteRequest struct {
	Id int64 `param:"note_id"`
}

// V1DeleteNote   godoc
// @Summary      Delete note
// @Description  Delete note by note_id, role >= ADMIN
// @Tags         Note
// @Accept       json
// @Produce      json
// @Param        note_id	path		uint			true	"Note Id"
// @Success      200		{object}	int
// @Failure      400		{object}	nil
// @Failure      401		{object}	nil
// @Failure      500		{object}	nil
// @Router       /note/{note_id} [delete]
func V1DeleteNote(req *DeleteNoteRequest, c echo.Context) common.Response {
	sess, err := session.Get("session", c)
	if err != nil {
		panic(err)
	}

	userId := sess.Values["user_id"]
	if userId == nil {
		panic("cannot get user_id from session")
	}

	repo := dao.GetRepo()
	query := `UPDATE note SET delete_dt = ?
            WHERE (id = ? AND delete_dt is NULL) AND (is_private = FALSE OR user_id = ?)`
	result, err := repo.Writer().Exec(query, time.Now(), req.Id, userId)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	// 지워지지 않았고, Public이 아닐 경우에는 user_id 확인 후 다를 경우 BadRequest 에러
	affectedRows, err := result.RowsAffected()
	if affectedRows < 1 {
		return common.Response{
			Code:  http.StatusBadRequest,
			Error: constant.ErrBadRequest,
		}
	}

	// TODO: Cascade update delete_dt for note_id and parent_note_id
	// FIXME: Below code will delete note_hierarchy, not update delete_dt (will remain dirty rows)
	_, err = repo.Writer().Exec(
		`DELETE FROM note_hierarchy WHERE note_id = ?`, req.Id,
	)

	return common.Response{
		Data: nil,
		Code: http.StatusOK,
	}
}
