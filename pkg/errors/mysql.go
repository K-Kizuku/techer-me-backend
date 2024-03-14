package errors

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func HandleError(err error) *Error {
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return New(http.StatusNotFound, err)
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1062:
			return New(http.StatusBadRequest, err)
		default:
			return New(http.StatusInternalServerError, err)
		}
	}

	return New(http.StatusInternalServerError, err)
}
