package domain

import "errors"

var (
	ErrWrongData = errors.New("некорректные данные")
	ErrDBConn    = errors.New("ошибка доступа к бд")
	ErrDBRead    = errors.New("ошибка чтения из бд")
)
