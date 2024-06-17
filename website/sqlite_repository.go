package website

import (
	"database/sql"
	"errors"
)

var (
    ErrDuplicate    = errors.New("record already exists")
    ErrNotExists    = errors.New("row not exists")
    ErrUpdateFailed = errors.New("update failed")
    ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
    db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
    return &SQLiteRepository{
        db: db,
    }
}

func (r *SQLiteRepository) Migrate() error {
    query := `
    CREATE TABLE IF NOT EXISTS websites(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        url TEXT NOT NULL,
        rank INTEGER NOT NULL
    );
    `

    _, err := r.db.Exec(query)
    return err
}
