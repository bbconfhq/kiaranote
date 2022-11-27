package dao

import "github.com/jmoiron/sqlx"

var repo Repository

type Repository interface {
	Writer() *sqlx.DB
	Reader() *sqlx.DB
}

type repositoryImpl struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func (r repositoryImpl) Writer() *sqlx.DB {
	return r.writer
}

func (r repositoryImpl) Reader() *sqlx.DB {
	return r.reader
}

func InitRepo(reader *sqlx.DB, writer *sqlx.DB) {
	repo = repositoryImpl{
		writer: writer,
		reader: reader,
	}
}

func GetRepo() Repository {
	return repo
}
