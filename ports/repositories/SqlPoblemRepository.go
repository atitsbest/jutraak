package repositories

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

type SqlProblemRepository struct {
    db *sqlx.DB
}

// CTR
func NewSqlProblemRepsoitory(connectionString string) (*SqlProblemRepository, error) {
    db, err := sqlx.Connect("postgres", connectionString)
    if err != nil {
        return nil, err
    }
    return &SqlProblemRepository{db: db}, nil
}
