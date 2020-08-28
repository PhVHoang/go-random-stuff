package account

import (
	"database/sql"
	"errors"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.logger
}
