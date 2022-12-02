package middleware

import (
	"github.com/bbconfhq/kiaranote/internal/common"
	"github.com/bbconfhq/kiaranote/internal/constant"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func validateUserRole(session *sessions.Session, role constant.Role) bool {
	if role == constant.RoleGuest {
		return true
	}

	if session.Values["user_id"] == nil {
		return false
	}

	userRole := constant.Role(session.Values["user_role"].(string))
	if userRole == constant.RoleAdmin {
		return true
	}

	return userRole == role
}

func AuthMiddleware(minRole constant.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("session", c)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.Response{
					Code:  http.StatusInternalServerError,
					Error: constant.ErrSession,
				})
			}

			if !validateUserRole(sess, minRole) {
				return c.JSON(http.StatusUnauthorized, common.Response{
					Code:  http.StatusUnauthorized,
					Error: constant.ErrUnauthorized,
				})
			}
			return next(c)
		}
	}
}
