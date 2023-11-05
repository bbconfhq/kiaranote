package handler

import (
	"net/http"
	"time"

	"github.com/bbconfhq/kiaranote/internal/common"
	"github.com/bbconfhq/kiaranote/internal/constant"
	"github.com/bbconfhq/kiaranote/internal/dao"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type GetNotesRequest struct{}
type GetNotesResponse struct {
	Public  []*NoteDto `json:"public"`
	Private []*NoteDto `json:"private"`
}
type NoteDto struct {
	Id         int64      `json:"id"`
	UserId     int64      `json:"user_id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	IsPrivate  bool       `json:"is_private"`
	CreateDt   time.Time  `json:"create_dt"`
	UpdateDt   time.Time  `json:"update_dt"`
	ChildNotes []*NoteDto `json:"child_notes"`
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

	repo := dao.GetRepo()
	rows, err := repo.Reader().Queryx(
		`SELECT
    		n.id, h.parent_note_id, n.user_id, n.title, "" as content, n.is_private, n.create_dt, n.update_dt
		FROM
		    note n
		INNER JOIN
			note_hierarchy h ON n.id = h.note_id
		WHERE
			n.delete_dt is NULL
			AND (n.is_private is FALSE OR (n.is_private is TRUE AND n.user_id = ?))`, userId,
	)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	notesMap := make(map[int64]*NoteDto)
	publicNotes := make([]*NoteDto, 0)
	privateNotes := make([]*NoteDto, 0)
	childNotes := make(map[int64][]*NoteDto)

	for rows.Next() {
		var note GetNoteResponse
		err := rows.StructScan(&note)
		if err != nil {
			return common.Response{
				Code:  http.StatusInternalServerError,
				Error: constant.ErrInternal,
			}
		}

		notesMap[note.Id] = &NoteDto{
			Id:         note.Id,
			UserId:     note.UserId,
			Title:      note.Title,
			Content:    note.Content,
			IsPrivate:  note.IsPrivate,
			CreateDt:   note.CreateDt,
			UpdateDt:   note.UpdateDt,
			ChildNotes: make([]*NoteDto, 0),
		}

		if note.ParentNoteId == note.Id {
			if note.IsPrivate {
				privateNotes = append(privateNotes, notesMap[note.Id])
			} else {
				publicNotes = append(publicNotes, notesMap[note.Id])
			}
		} else {
			childNotes[note.ParentNoteId] = append(childNotes[note.ParentNoteId], notesMap[note.Id])
		}
	}

	for _, note := range notesMap {
		if childNotes[note.Id] != nil {
			note.ChildNotes = childNotes[note.Id]
		}
	}

	notes := GetNotesResponse{
		Public:  publicNotes,
		Private: privateNotes,
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
	// FIXME: In schema, parent_note_id is nullable. But in code, it cannot be null
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
			h.order`, req.Id, userId).StructScan(&note); err != nil {
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
	// Public이거나 작성자가 본인일 경우에만 수정 가능
	query := `UPDATE note SET title = ?, content = ?, is_private = ?
            WHERE id = ? AND delete_dt is NULL AND (is_private = FALSE OR user_id = ?)`
	result, err := repo.Writer().Exec(query, req.Title, req.Content, req.IsPrivate, req.Id, userId)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	// 지워지지 않았고, Public이 아닐 경우에는 소유주 확인하고 아니라면 BadRequest 에러
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
	// Public이거나 작성자가 본인일 경우에만 삭제 가능
	query := `UPDATE note SET delete_dt = ?
				WHERE id IN (
					WITH RECURSIVE cte (note_id, note_parent_id) AS (
						SELECT	note_id, parent_note_id
						FROM	note_hierarchy
						WHERE	note_id = ?
						UNION ALL
						SELECT	n.note_id, n.parent_note_id
						FROM	note_hierarchy n
						INNER JOIN cte
								ON cte.note_id = n.parent_note_id AND cte.note_id != n.note_id
					)
					SELECT note_id FROM cte
				) AND delete_dt is NULL AND (is_private = FALSE OR user_id = ?)`
	result, err := repo.Writer().Exec(query, time.Now(), req.Id, userId)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	// 지워지지 않았고, Public이 아닐 경우에는 소유주 확인하고 아니라면 BadRequest 에러
	affectedRows, err := result.RowsAffected()
	if affectedRows < 1 {
		return common.Response{
			Code:  http.StatusBadRequest,
			Error: constant.ErrBadRequest,
		}
	}

	_, err = repo.Writer().Exec(
		`DELETE FROM note_hierarchy WHERE note_id = ?`, req.Id,
	)

	return common.Response{
		Data: nil,
		Code: http.StatusOK,
	}
}
