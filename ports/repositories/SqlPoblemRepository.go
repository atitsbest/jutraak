package repositories

import (
    . "github.com/atitsbest/jutraak/bugtracking/domain/entities"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

type SqlProblemRepository struct {
    connectionString string
}

// CTR
func NewSqlProblemRepsoitory(connectionString string) (*SqlProblemRepository, error) {
    db, err := sqlx.Connect("postgres", connectionString)
    if err != nil {
        return nil, err
    }
    db.Close()
    return &SqlProblemRepository{
        connectionString: connectionString,
    }, nil
}

// Alle Probleme
func (self *SqlProblemRepository) All() ([]*Problem, error) {
    return nil, nil
}
