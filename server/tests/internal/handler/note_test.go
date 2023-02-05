package handler

import (
	"github.com/bbconfhq/kiaranote/internal/dao"
	"github.com/bbconfhq/kiaranote/internal/handler"
	"github.com/bbconfhq/kiaranote/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createNote(repo dao.Repository, username string, title string, content string, isPrivate bool) {
	var userId int64
	if err := repo.Reader().QueryRow(`SELECT id FROM user WHERE username = ?`, username).Scan(&userId); err != nil {
		panic(err)
	}

	result, err := repo.Writer().Exec(
		`INSERT INTO note (user_id, title, content, is_private) VALUES (?, ?, ?, ?)`, userId, title, content, isPrivate,
	)
	if err != nil {
		panic(err)
	}

	noteId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	_, err = repo.Writer().Exec(
		"INSERT INTO note_hierarchy (`note_id`, `parent_note_id`, `order`) VALUES (?, ?, ?)", noteId, noteId, 1,
	)
	if err != nil {
		panic(err)
	}
}

func TestV1GetNotes(t *testing.T) {
	e, repo := tests.MockMain()
	tests.TruncateTable(repo.Reader(), []string{"audit_log", "note_hierarchy", "note", "user"})

	var adminUsername = "1"
	var adminPassword = "1"
	createAdmin(repo, adminUsername, adminPassword)
	c := loginAdmin(e, adminUsername, adminPassword)

	createNote(repo, adminUsername, "public", "public_note", false)
	createNote(repo, adminUsername, "private", "private_note", true)

	response := handler.V1GetNotes(nil, c)
	assert.NotEmpty(t, response.Data)

	data := response.Data.(handler.GetNotesResponse)
	assert.Len(t, data.Public, 1)
	assert.Len(t, data.Private, 1)
	assert.Equal(t, "public", data.Public[0].Title)
	assert.Equal(t, "private", data.Private[0].Title)
}

func TestV1PostNote(t *testing.T) {
	e, repo := tests.MockMain()
	tests.TruncateTable(repo.Reader(), []string{"audit_log", "note_hierarchy", "note", "user"})

	var adminUsername = "1"
	var adminPassword = "1"
	createAdmin(repo, adminUsername, adminPassword)
	c := loginAdmin(e, adminUsername, adminPassword)

	payload := handler.PostNoteRequest{
		ParentNoteId: 1,
		Title:        "",
		Content:      "",
		IsPrivate:    false,
		Order:        0,
	}
	response := handler.V1PostNote(&payload, c)

	assert.Equal(t, 1, response.Data.(int64))
}

//func TestV1GetNote(t *testing.T) {
//	e, repo := tests.MockMain()
//	tests.TruncateTable(repo.Reader(), []string{"audit_log", "user"})
//}

//func TestV1PutNote(t *testing.T) {
//	e, repo := tests.MockMain()
//	tests.TruncateTable(repo.Reader(), []string{"audit_log", "user"})
//}

//func TestV1DeleteNote(t *testing.T) {
//	e, repo := tests.MockMain()
//	tests.TruncateTable(repo.Reader(), []string{"audit_log", "user"})
//}
