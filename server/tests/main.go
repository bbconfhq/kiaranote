package tests

import (
	"fmt"
	"github.com/bbconfhq/kiaranote/internal/dao"
	"github.com/bbconfhq/kiaranote/internal/datasource/mysql"
	"github.com/bbconfhq/kiaranote/internal/handler"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"path"
	"runtime"
)

func MockMain() (*echo.Echo, dao.Repository) {
	_, filename, _, _ := runtime.Caller(0)
	env := path.Join(path.Dir(filename), "..", ".env.test")

	if err := godotenv.Load(env); err != nil {
		panic(err)
	}

	readerConfig := mysql.Config{
		User:   os.Getenv("READER_DB_USER"),
		Passwd: os.Getenv("READER_DB_PASS"),
		Host:   os.Getenv("READER_DB_HOST"),
		Port:   os.Getenv("READER_DB_PORT"),
		DBName: os.Getenv("READER_DB_NAME"),
	}
	reader := mysql.NewMySQL(readerConfig)
	writerConfig := mysql.Config{
		User:   os.Getenv("WRITER_DB_USER"),
		Passwd: os.Getenv("WRITER_DB_PASS"),
		Host:   os.Getenv("WRITER_DB_HOST"),
		Port:   os.Getenv("WRITER_DB_PORT"),
		DBName: os.Getenv("WRITER_DB_NAME"),
	}
	writer := mysql.NewMySQL(writerConfig)
	dao.InitRepo(reader, writer)

	e := echo.New()
	secret := ""
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))
	e.Validator = &handler.RequestValidator{
		Validator: validator.New(),
	}
	handler.InitV1Handler(e)

	return e, dao.GetRepo()
}

func TruncateTable(writer *sqlx.DB, tables []string) {
	for _, table := range tables {
		_, err := writer.Exec(`DELETE FROM ` + table)
		if err != nil {
			panic(fmt.Sprintf("failed to delete all rows from table %s", table))
		}
	}
}
