package sqlite

import (
	"database/sql"
	"fmt"
	"net/url"
)

func New(path string) (*sql.DB, error) {
	params := url.Values{}

	params.Add("cache", "shared")
	params.Add("mode", "rwc")
	params.Add("_journal", "WAL")

	conn, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?%s", path, params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	// https://github.com/mattn/go-sqlite3/issues/209
	conn.SetMaxOpenConns(1)

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to check connection: %w", err)
	}

	return conn, nil
}
