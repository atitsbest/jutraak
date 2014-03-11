package repositories

import (
  _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

type SqlProblemRepository struct {
  db *sqlx.DB
}

// CTR
func NewSqlProblemRepsoitory(connectionString string) (*SqlProblemRepository, error) {
  db, err := sqlx.Connect("postgres", connectionString)
  if err != nil {return nil, err}
  return &SqlProblemRepository{ db: db }, nil
}
