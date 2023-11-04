package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/bbconfhq/kiaranote/internal/constant"
	"github.com/bbconfhq/kiaranote/internal/dao"
	"github.com/bbconfhq/kiaranote/internal/handler"
	"github.com/bbconfhq/kiaranote/tests"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func loginGuest(e *echo.Echo) echo.Context {
	req := httptest.NewRequest(http.MethodGet,
		"/",
		strings.NewReader(""),
	)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("")))

	return c
}

func loginAdmin(e *echo.Echo, username string, password string) echo.Context {
	loginRequest := handler.LoginRequest{Username: username, Password: password}
	result, _ := json.Marshal(loginRequest)

	req := httptest.NewRequest(http.MethodPost,
		"/auth/login",
		bytes.NewReader(result),
	)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("")))

	handler.V1Login(&loginRequest, c)
	return c
}

func createAdmin(repo dao.Repository, username string, password string) {
	_, err := repo.Writer().Exec(
		`INSERT INTO user (username, password, role) VALUES (?, ?, ?)`,
		username, handler.EncodeHash(password), constant.RoleAdmin,
	)

	if err != nil {
		panic(errors.New("admin is not created"))
	}
}

func TestV1GetUsers(t *testing.T) {
	e, repo := tests.MockMain()
	tests.TruncateTable(repo.Reader(), []string{"audit_log", "user"})

	var adminUsername = "1"
	var adminPassword = "1"
	createAdmin(repo, adminUsername, adminPassword)
	c := loginAdmin(e, adminUsername, adminPassword)

	response := handler.V1GetUsers(nil, c)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, 1, len(response.Data.([]handler.GetUsersResponse)))

	user := response.Data.([]handler.GetUsersResponse)[0]
	assert.Equal(t, "1", user.Username)
}

func TestV1PostUser(t *testing.T) {
	e, repo := tests.MockMain()
	tests.TruncateTable(repo.Reader(), []string{"audit_log", "user"})

	var adminUsername = "1"
	var adminPassword = "1"
	createAdmin(repo, adminUsername, adminPassword)
	c := loginAdmin(e, adminUsername, adminPassword)

	response := handler.V1PostUser(&handler.PostUserRequest{Username: "2", Password: "2"}, c)
	assert.Equal(t, http.StatusCreated, response.Code)

	response = handler.V1GetUsers(nil, c)
	assert.Equal(t, http.StatusOK, response.Code)

	users := response.Data.([]handler.GetUsersResponse)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "1", users[0].Username)
	assert.Equal(t, "2", users[1].Username)

	response = handler.V1PostUser(&handler.PostUserRequest{Username: "2", Password: "2"}, c)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, constant.ErrBadRequest, response.Error)
}

func TestV1GetUser(t *testing.T) {
	e, repo := tests.MockMain()
	tests.TruncateTable(repo.Reader(), []string{"audit_log", "user"})

	c := loginGuest(e)
	c.SetPath("/user/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("0")

	response := handler.V1GetUser(&handler.GetUserRequest{Id: 0}, c)
	assert.Equal(t, http.StatusUnauthorized, response.Code)
	assert.Equal(t, constant.ErrUnauthorized, response.Error)

	var adminUsername = "1"
	var adminPassword = "1"
	createAdmin(repo, adminUsername, adminPassword)
	c = loginAdmin(e, adminUsername, adminPassword)

	response = handler.V1GetUsers(nil, c)
	assert.Equal(t, http.StatusOK, response.Code)

	users := response.Data.([]handler.GetUsersResponse)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "1", users[0].Username)

	c.SetPath("/user/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues(strconv.FormatInt(users[0].Id, 10))

	response = handler.V1GetUser(&handler.GetUserRequest{Id: users[0].Id}, c)
	assert.Equal(t, http.StatusOK, response.Code)

	user := response.Data.(handler.GetUserResponse)
	assert.Equal(t, users[0].Id, user.Id)
	assert.Equal(t, users[0].Username, user.Username)
}

func TestV1PutUser(t *testing.T) {
	e, repo := tests.MockMain()
	tests.TruncateTable(repo.Reader(), []string{"audit_log", "user"})

	c := loginGuest(e)
	c.SetPath("/user/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("1")

	response := handler.V1PutUser(&handler.PutUserRequest{
		Id:       1,
		Username: "1",
		Password: "1",
		Role:     "admin",
	}, c)
	assert.Equal(t, http.StatusUnauthorized, response.Code)
	assert.Equal(t, constant.ErrUnauthorized, response.Error)

	var adminUsername = "1"
	var adminPassword = "1"
	createAdmin(repo, adminUsername, adminPassword)
	c = loginAdmin(e, adminUsername, adminPassword)

	response = handler.V1GetUsers(nil, c)
	assert.Equal(t, http.StatusOK, response.Code)

	users := response.Data.([]handler.GetUsersResponse)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "1", users[0].Username)

	c.SetPath("/user/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues(strconv.FormatInt(users[0].Id, 10))

	response = handler.V1PutUser(&handler.PutUserRequest{
		Id:       users[0].Id,
		Username: "2",
	}, c)
	assert.Equal(t, http.StatusOK, response.Code)

	c.SetPath("/user/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues(strconv.FormatInt(users[0].Id, 10))

	response = handler.V1GetUser(&handler.GetUserRequest{Id: users[0].Id}, c)
	assert.Equal(t, http.StatusOK, response.Code)

	user := response.Data.(handler.GetUserResponse)
	assert.Equal(t, users[0].Id, user.Id)
	assert.Equal(t, "2", user.Username)
}

func TestV1DeleteUser(t *testing.T) {
	e, repo := tests.MockMain()
	tests.TruncateTable(repo.Reader(), []string{"audit_log", "user"})

	c := loginGuest(e)
	c.SetPath("/user/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("1")

	response := handler.V1DeleteUser(&handler.DeleteUserRequest{Id: 1}, c)
	assert.Equal(t, http.StatusUnauthorized, response.Code)
	assert.Equal(t, constant.ErrUnauthorized, response.Error)

	createAdmin(repo, "1", "1")
	createAdmin(repo, "2", "2")
	c = loginAdmin(e, "1", "1")

	response = handler.V1GetUsers(nil, c)
	assert.Equal(t, http.StatusOK, response.Code)

	users := response.Data.([]handler.GetUsersResponse)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "2", users[1].Username)

	c.SetPath("/user/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues(strconv.FormatInt(users[1].Id, 10))

	response = handler.V1DeleteUser(&handler.DeleteUserRequest{Id: users[1].Id}, c)
	assert.Equal(t, http.StatusOK, response.Code)

	response = handler.V1GetUsers(nil, c)
	assert.Equal(t, http.StatusOK, response.Code)

	users = response.Data.([]handler.GetUsersResponse)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "1", users[0].Username)
}
