package service

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

//ServiceError ...
type ServiceError struct {
	HTTPStatus int
	Message    string
}

func (e *ServiceError) Error() string {
	return e.Message
}

//Log ...
var Log = func(msg string) {}

//NewServiceError ...
func NewServiceError(err error) *ServiceError {

	de := ServiceError{}

	switch err.(type) {

	case *mysql.MySQLError:

		Log(err.(error).Error())
		de.HTTPStatus = 500
		de.Message = "Unknown error"

	case error:

		Log(err.(error).Error())

		if err == sql.ErrNoRows {

			de.HTTPStatus = 404
			de.Message = "Not Found"
		} else {

			de.HTTPStatus = 500
			de.Message = "Unknown error"
		}

	default:

		return nil
	}

	return &de
}
