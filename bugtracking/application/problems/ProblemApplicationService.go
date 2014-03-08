package problems

import (
    "github.com/atitsbest/jutraak/bugtracking/domain/entities"
    "github.com/atitsbest/jutraak/bugtracking/domain/valueobjects"
    "time"
)

// Interface zum Repository
type ProblemRepository interface {
    Insert(*entities.Problem) error
    Update(*entities.Problem) error
    GetById(entities.ProblemId) (*entities.Problem, error)
}

// Application für die Probleme.
type ProblemApplicationService struct {
    problems ProblemRepository
}

// CTR
func NewProblemApplicationService(problems ProblemRepository) *ProblemApplicationService {
    return &ProblemApplicationService{
        problems: problems,
    }
}

// Erstellt ein neues Problem.
// Liefert einen Error, wenn die Daten ungültig sind.
func (self *ProblemApplicationService) CreateNewProblem(
    summary string,
    description string,
    tags []string,
    createdBy string) (*entities.Problem, error) {

    cmd := CreateNewProblemCommand{
        Summary:     summary,
        Description: description,
        Tags:        tags,
        CreatedBy:   createdBy,
    }

    return self.createNewProblem(&cmd)
}

// Hängt eine Datei an ein Problem an.
func (self *ProblemApplicationService) AttachFileToProblem(problemId entities.ProblemId, fileName string, data []byte) error {
    problem, err := self.problems.GetById(problemId)
    if err != nil {
        return err
    }
    attachment, err := valueobjects.NewAttachment(fileName, data)
    if err != nil {
        return err
    }
    problem.AddAttachment(attachment)
    err = self.problems.Update(problem)
    if err != nil {
        return err
    }

    return nil
}

// Erstellt ein neues Problem
func (self *ProblemApplicationService) createNewProblem(cmd *CreateNewProblemCommand) (*entities.Problem, error) {
    result := &entities.Problem{
        Summary:     cmd.Summary,
        Description: cmd.Description,
        Tags:        cmd.Tags,
        CreatedBy:   cmd.CreatedBy,
        CreatedAt:   time.Now(),
    }

    err := self.problems.Insert(result)

    return result, err
}
