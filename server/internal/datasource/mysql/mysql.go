package mysql

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Config struct {
	User    string
	Passwd  string
	Host    string
	Port    string
	DBName  string
	MaxConn int
	Timeout time.Duration
}

func (c *Config) FormatDSN() string {
	config := &mysql.Config{
		User:                 c.User,
		Passwd:               c.Passwd,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", c.Host, c.Port),
		DBName:               c.DBName,
		Collation:            "utf8mb4_unicode_ci",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	return config.FormatDSN()
}

func NewMySQL(config Config) *sqlx.DB {
	db, err := sqlx.Connect("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db
}
