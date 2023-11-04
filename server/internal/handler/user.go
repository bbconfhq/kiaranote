package handler

import (
	"github.com/bbconfhq/kiaranote/internal/common"
	"github.com/bbconfhq/kiaranote/internal/constant"
	"github.com/bbconfhq/kiaranote/internal/dao"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type GetUsersRequest struct{}
type GetUsersResponse struct {
	Id          int64     `json:"id"`
	Username    string    `json:"username"`
	Role        string    `json:"role"`
	LastLoginDt time.Time `db:"last_login_dt" json:"last_login_dt"`
	CreateDt    time.Time `db:"create_dt" json:"create_dt"`
	UpdateDt    time.Time `db:"update_dt" json:"update_dt"`
}

// V1GetUsers   godoc
// @Summary      Get users
// @Description  Get list of users, role >= ADMIN
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200	{object}	[]GetUsersResponse
// @Router       /user [get]
func V1GetUsers(_ *GetUsersRequest, _ echo.Context) common.Response {
	repo := dao.GetRepo()
	rows, err := repo.Reader().Queryx(
		`SELECT
			id, username, role, last_login_dt, create_dt, update_dt
		FROM user
		WHERE delete_dt is NULL`,
	)
	if err != nil {
		panic(err)
	}

	users := make([]GetUsersResponse, 0)

	for rows.Next() {
		var user GetUsersResponse
		err := rows.StructScan(&user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return common.Response{
		Data: users,
		Code: http.StatusOK,
	}
}

type PostUserRequest struct {
	Username string `json:"username" validate:"required,alphanumunicode,lte=20"`
	Password string `json:"password" validate:"required"`
}

// V1PostUser   godoc
// @Summary      Post user
// @Description  Register new user, role >= ADMIN
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        req	body		PostUserRequest	true	"Username and password"
// @Success      201	{object}	int
// @Failure      400	{object}	int
// @Router       /user [post]
func V1PostUser(req *PostUserRequest, _ echo.Context) common.Response {
	repo := dao.GetRepo()
	result, err := repo.Writer().Exec(
		`INSERT INTO user (username, password) VALUES (?, ?)`, req.Username, EncodeHash(req.Password),
	)

	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == constant.MysqlDuplicateEntry {
			return common.Response{
				Code:  http.StatusBadRequest,
				Error: constant.ErrBadRequest,
			}
		}
	}

	if err != nil {
		panic(err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return common.Response{
		Data: userId,
		Code: http.StatusCreated,
	}
}

type GetUserRequest struct {
	Id int64 `param:"user_id" validate:"required,gte=1"`
}
type GetUserResponse struct {
	Id          int64     `json:"id"`
	Username    string    `json:"username"`
	Role        string    `json:"role"`
	LastLoginDt time.Time `db:"last_login_dt" json:"last_login_dt"`
	CreateDt    time.Time `db:"create_dt" json:"create_dt"`
	UpdateDt    time.Time `db:"update_dt" json:"update_dt"`
}

// V1GetUser   godoc
// @Summary      Get user
// @Description  Get user by user_id, user itself or role >= ADMIN
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user_id	path		uint			true	"User Id"
// @Success      200		{object}	GetUserResponse
// @Failure      400		{object}	int
// @Failure      401		{object}	int
// @Failure      500		{object}	int
// @Router       /user/{user_id} [get]
func V1GetUser(req *GetUserRequest, c echo.Context) common.Response {
	sess, err := session.Get("session", c)

	if !validateItselfOrAdmin(sess, req.Id) {
		return common.Response{
			Code:  http.StatusUnauthorized,
			Error: constant.ErrUnauthorized,
		}
	}

	repo := dao.GetRepo()
	rows, err := repo.Reader().Queryx(
		`SELECT
			id, username, role, last_login_dt, create_dt, update_dt
		FROM user
		WHERE
		    delete_dt is NULL AND id = ?`, req.Id,
	)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var user GetUserResponse
		err := rows.StructScan(&user)
		if err != nil {
			panic(err)
		}

		return common.Response{
			Data: user,
			Code: http.StatusOK,
		}
	}

	return common.Response{
		Code:  http.StatusInternalServerError,
		Error: constant.ErrInternal,
	}
}

type PutUserRequest struct {
	Id       int64  `param:"user_id" validate:"required,gte=1"`
	Username string `json:"username" validate:"alphanumunicode,lte=20"`
	Password string `json:"password"`
	Role     string `json:"role" validate:"eq=admin|eq=user|eq=guest"`
}

// V1PutUser   godoc
// @Summary      Put user
// @Description  Put user by user_id, user itself or role >= ADMIN
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user_id	path		uint			true	"User Id"
// @Param        req		body		PutUserRequest	true	"Username and password"
// @Success      200		{object}	uint
// @Failure      400		{object}	int
// @Failure      401		{object}	int
// @Failure      500		{object}	int
// @Router       /user/{user_id} [put]
func V1PutUser(req *PutUserRequest, c echo.Context) common.Response {
	sess, err := session.Get("session", c)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrSession,
		}
	}

	if !validateItselfOrAdmin(sess, req.Id) {
		return common.Response{
			Code:  http.StatusUnauthorized,
			Error: constant.ErrUnauthorized,
		}
	}

	fields := make([]string, 0)
	values := make([]interface{}, 0)
	if req.Username != "" {
		fields = append(fields, "username = ?")
		values = append(values, req.Username)
	}
	if req.Role != "" && constant.Role(sess.Values["user_role"].(string)) == constant.RoleAdmin {
		fields = append(fields, "role = ?")
		values = append(values, req.Role)
	}
	if req.Password != "" {
		fields = append(fields, "password = ?")
		values = append(values, EncodeHash(req.Password))
	}
	values = append(values, strconv.FormatInt(req.Id, 10))

	repo := dao.GetRepo()
	query := "UPDATE user SET " + strings.Join(fields, ", ") + " WHERE id = ?"
	_, err = repo.Writer().Exec(query, values...)

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

type DeleteUserRequest struct {
	Id int64 `param:"user_id" validate:"required,gte=1"`
}

// V1DeleteUser   godoc
// @Summary      Delete user
// @Description  Delete user by user_id, user itself or role >= ADMIN
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user_id	path		uint			true	"User Id"
// @Success      200		{object}	nil
// @Failure      400		{object}	int
// @Failure      401		{object}	int
// @Failure      500		{object}	int
// @Router       /user/{user_id} [delete]
func V1DeleteUser(req *DeleteUserRequest, c echo.Context) common.Response {
	sess, err := session.Get("session", c)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrSession,
		}
	}

	if !validateItselfOrAdmin(sess, req.Id) {
		return common.Response{
			Code:  http.StatusUnauthorized,
			Error: constant.ErrUnauthorized,
		}
	}

	repo := dao.GetRepo()

	// 자기 자신을 지울 경우 세션 값 제거
	if sess.Values["user_id"] == req.Id {
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		}
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return common.Response{
				Code:  http.StatusInternalServerError,
				Error: constant.ErrInternal,
			}
		}
	}

	_, err = repo.Writer().Exec(
		`DELETE FROM user WHERE id = ?`, req.Id,
	)

	return common.Response{
		Code: http.StatusOK,
	}
}

func validateItselfOrAdmin(sess *sessions.Session, reqUserId int64) bool {
	userId := sess.Values["user_id"]

	if userId == nil {
		return false
	}

	userRole := constant.Role(sess.Values["user_role"].(string))

	if userId == reqUserId || userRole == constant.RoleAdmin {
		return true
	}

	return false
}
