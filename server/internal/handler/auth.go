package handler

import (
	"encoding/hex"
	"github.com/bbconfhq/kiaranote/internal/dao"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/scrypt"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const salt = ""

func encodeHash(value string) string {
	key, _ := scrypt.Key([]byte(value), []byte(salt), 32768, 8, 1, 32)
	return hex.EncodeToString(key)
}

// V1Login   godoc
// @Summary      Login
// @Description  Request login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200	{object}	response
// @Failure      401	{object}	response
// @Failure      500	{object}	response
// @Router       /auth/login [get]
func V1Login(req *LoginRequest, c echo.Context) Response {
	repo := dao.GetRepo()
	rows, err := repo.Reader().Query("SELECT id, username, role FROM user WHERE username = ? AND password = ?", req.Username, encodeHash(req.Password))
	if err != nil {
		panic(err)
	}

	var (
		id       int64
		username string
		role     string
	)
	if rows.Next() {
		err := rows.Scan(&id, &username, role)
		if err != nil {
			panic(err)
		}

		sess, err := session.Get("session", c)
		if err != nil {
			return Response{
				Code:  http.StatusInternalServerError,
				Error: ErrSession,
			}
		}
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["user_id"] = id
		sess.Values["username"] = username
		sess.Values["user_role"] = role
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return Response{
				Code:  http.StatusInternalServerError,
				Error: ErrInternal,
			}
		}

		return Response{
			Code: http.StatusOK,
		}
	} else {
		return Response{
			Code:  http.StatusUnauthorized,
			Error: ErrUnauthorized,
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
// @Failure      500	{object}	response
// @Router       /auth/logout [get]
func V1Logout(_ *LogoutRequest, c echo.Context) Response {
	sess, err := session.Get("session", c)
	if err != nil {
		return Response{
			Code:  http.StatusInternalServerError,
			Error: ErrSession,
		}
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return Response{
			Code:  http.StatusInternalServerError,
			Error: ErrInternal,
		}
	}
	return Response{
		Code: http.StatusOK,
	}
}
