package repositories

import (
    . "github.com/atitsbest/jutraak/bugtracking/domain/entities"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"

    "time"
    "strings"
)

type SqlProblemRepository struct {
    connectionString string
}

type problemDto struct {
  Id int
  Summary string
  Description string
  Tags string
  Created_at time.Time
  Created_by string
  Lastchange_at time.Time
  Lastchange_by string
}

func (self *problemDto) ToProblem() *Problem {
  result := &Problem{
    Id : ProblemId(string(self.Id)),
    Summary : self.Summary,
    Description : self.Description,
    Tags : self.TagArray(),
    CreatedAt : self.Created_at,
    CreatedBy : self.Created_by,
    LastChangeAt : self.Lastchange_at,
    LastChangeBy : self.Lastchange_by,
  }

  return result
}

func (self *problemDto) TagArray() []string {
  trimmed := strings.Trim(self.Tags, ",")
  return strings.Split(trimmed, ",")
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
    db, err := sqlx.Connect("postgres", self.connectionString)
    if err != nil {
        return nil, err
    }
    defer db.Close()
    ps := []problemDto{}
    db.Selectv(&ps, "SELECT * FROM problems ORDER BY id")

    result := []*Problem{}
    for _, p := range(ps) {
      result = append(result, p.ToProblem())
    }

    return result, nil
}
