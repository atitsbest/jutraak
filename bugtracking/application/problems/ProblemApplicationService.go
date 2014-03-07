package problems

import (
  "time"
  "dccs/jutraak/bugtracking/entities"
)

/**
  Interface zum Repository
 */
type ProblemRepository interface {
  Insert(*entities.Problem) error
}

/**
  Application für die Probleme.
 */
type ProblemApplicationService struct {
  problems ProblemRepository
}

/**
  CTR
 */
func NewProblemApplicationService(problems ProblemRepository) *ProblemApplicationService {
  return &ProblemApplicationService{
    problems: problems,
  }
}

/**
  Erstellt ein neues Problem.
  Liefert einen Error, wenn die Daten ungültig sind.
 */
func (self *ProblemApplicationService) CreateNewProblem(
  summary string, description string, createdBy string) (*entities.Problem, error) {

  cmd := CreateNewProblemCommand{
    Summary: summary,
    Description: description,
    CreatedBy: createdBy,
  }

  return self.createNewProblem(&cmd)
}


/**
 Erstellt ein neues Problem
 */
func (self *ProblemApplicationService) createNewProblem(cmd *CreateNewProblemCommand) (*entities.Problem, error) {
  result := &entities.Problem{
    Summary: cmd.Summary,
    Description: cmd.Description,
    CreatedBy: cmd.CreatedBy,
    CreatedAt: time.Now(),
  }

  err:= self.problems.Insert(result)

  return result, err
}
