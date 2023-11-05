package handler

import (
	"encoding/hex"
	"github.com/bbconfhq/kiaranote/internal/common"
	"github.com/bbconfhq/kiaranote/internal/constant"
	"github.com/bbconfhq/kiaranote/internal/dao"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/scrypt"
	"net/http"
	"time"
)

const salt = ""

func EncodeHash(value string) string {
	key, _ := scrypt.Key([]byte(value), []byte(salt), 32768, 8, 1, 32)
	return hex.EncodeToString(key)
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,alphanumunicode,lte=20"`
	Password string `json:"password" validate:"required"`
}

// V1Register    godoc
// @Summary      Register
// @Description  Request register
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        req	body		RegisterRequest	true	"Username and password"
// @Success      200	{object}	response
// @Failure      400	{object}	response
// @Failure      500	{object}	response
// @Router       /auth/register [post]
func V1Register(req *RegisterRequest, c echo.Context) common.Response {
	repo := dao.GetRepo()

	rows, err := repo.Reader().Query(`SELECT COUNT(*) FROM user`)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrInternal,
		}
	}

	// First user is admin
	role := constant.RoleGuest
	if rows.Next() {
		var count int64
		err := rows.Scan(&count)
		if err != nil {
			panic(err)
		}

		if count == 0 {
			role = constant.RoleAdmin
		}
	}

	_, err = repo.Writer().Exec(
		`INSERT INTO user (username, password, role) VALUES (?, ?, ?)`,
		req.Username, EncodeHash(req.Password), role,
	)
	if err != nil {
		return common.Response{
			Code:  http.StatusBadRequest,
			Error: constant.ErrBadRequest,
		}
	}

	return common.Response{
		Code: http.StatusOK,
	}
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanumunicode,lte=20"`
	Password string `json:"password" validate:"required"`
}
type LoginResponse struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// V1Login   godoc
// @Summary      Login
// @Description  Request login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        req	body		LoginRequest	true	"Username and password"
// @Success      200	{object}	response
// @Failure      400	{object}	response
// @Failure      401	{object}	response
// @Failure      500	{object}	response
// @Router       /auth/login [post]
func V1Login(req *LoginRequest, c echo.Context) common.Response {
	repo := dao.GetRepo()
	rows, err := repo.Reader().Queryx("SELECT id, username, role FROM user WHERE username = ? AND password = ?", req.Username, EncodeHash(req.Password))
	if err != nil {
		panic(err)
	}

	var user LoginResponse
	if rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			panic(err)
		}

		sess, err := session.Get("session", c)
		if err != nil {
			return common.Response{
				Code:  http.StatusInternalServerError,
				Error: constant.ErrSession,
			}
		}
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["user_id"] = user.Id
		sess.Values["username"] = user.Username
		sess.Values["user_role"] = user.Role
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return common.Response{
				Code:  http.StatusInternalServerError,
				Error: constant.ErrInternal,
			}
		}

		_, err = repo.Writer().Exec(
			`UPDATE user SET last_login_dt = ? WHERE id = ?`, time.Now(), user.Id,
		)
		if err != nil {
			panic(err)
		}

		return common.Response{
			Data: user,
			Code: http.StatusOK,
		}
	} else {
		return common.Response{
			Code:  http.StatusUnauthorized,
			Error: constant.ErrUnauthorized,
		}
	}
}

type LogoutRequest struct{}

// V1Logout	     godoc
// @Summary      Logout
// @Description  Request logout
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200	{object}	response
// @Failure      400	{object}	response
// @Failure      500	{object}	response
// @Router       /auth/logout [get]
func V1Logout(_ *LogoutRequest, c echo.Context) common.Response {
	sess, err := session.Get("session", c)
	if err != nil {
		return common.Response{
			Code:  http.StatusInternalServerError,
			Error: constant.ErrSession,
		}
	}
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
	return common.Response{
		Code: http.StatusOK,
	}
}
