package problems

import (
    . "github.com/atitsbest/jutraak/bugtracking/domain/entities"
    . "github.com/atitsbest/jutraak/bugtracking/domain/valueobjects"
    "os"
    "time"
)

// Interface zum Repository
type (
  ProblemRepository interface {
      All() ([]*Problem, error)
      AllTags() ([]string, error)
      Insert(*Problem) error
      Update(*Problem) error
      GetById(ProblemId) (*Problem, error)
      Filtered(tags []string, q string) ([]*Problem, error)
  }

  ProblemApplicationServiceInterface interface {
      GetAllProblems() ([]*Problem, error)
      CreateNewProblem(summary string, description string, tags []string, createdBy string) (*Problem, error)
      ChangeProblemSummary(problemId ProblemId, summary string, description string, who string) error
      AttachFileToProblem(problemId ProblemId, fileName string, data []byte) error
      RemoveProblemAttachment(problemId ProblemId, filePath string) error
      CommentProblem(problemId ProblemId, text string, who string, attachments []*Attachment) error
  }

  // Application für die Probleme.
  ProblemApplicationService struct {
      problems ProblemRepository
  }
)

// CTR
func NewProblemApplicationService(problems ProblemRepository) *ProblemApplicationService {
    return &ProblemApplicationService{
        problems: problems,
    }
}

// Liefert eine Liste mit allen Problemen.
func (self *ProblemApplicationService) GetAllProblems() ([]*Problem, error) {
    return self.problems.All()
}

// Erstellt ein neues Problem.
// Liefert einen Error, wenn die Daten ungültig sind.
func (self *ProblemApplicationService) CreateNewProblem(
    summary string,
    description string,
    tags []string,
    createdBy string) (*Problem, error) {

    cmd := CreateNewProblemCommand{
        Summary:     summary,
        Description: description,
        Tags:        tags,
        CreatedBy:   createdBy,
    }

    return self.createNewProblem(&cmd)
}

func (self *ProblemApplicationService) ChangeProblemSummary(
    problemId ProblemId, summary string, description string, who string) error {

    problem, err := self.problems.GetById(problemId)
    if err != nil {
        return err
    }

    problem.Summary = summary
    problem.Description = description
    problem.LastChangeBy = who
    problem.LastChangeAt = time.Now()

    err = self.problems.Update(problem)
    if err != nil {
        return err
    }

    return nil
}

// Hängt eine Datei an ein Problem an.
func (self *ProblemApplicationService) AttachFileToProblem(problemId ProblemId, fileName string, data []byte) error {
    problem, err := self.problems.GetById(problemId)
    if err != nil {
        return err
    }
    attachment, err := NewAttachment(fileName, data)
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

// Ein Attachment wieder vom Problem entfernen.
func (self *ProblemApplicationService) RemoveProblemAttachment(problemId ProblemId, filePath string) error {
    problem, err := self.problems.GetById(problemId)
    if err != nil {
        return err
    }

    newAttachments := []*Attachment{}

    // Attachment mit dem angegebenen Pfad finden.
    for _, a := range problem.Attachments {
        if a.FilePath != filePath {
            // Kopieren.
            newAttachments = append(newAttachments, a)
        }
    }

    problem.Attachments = newAttachments
    err = self.problems.Update(problem)
    if err != nil {
        return err
    }

    // Datei löschen.
    os.Remove(filePath)

    return nil
}

// Ein Problem kommentieren.
func (self *ProblemApplicationService) CommentProblem(
    problemId ProblemId, text string, who string, attachments []*Attachment) error {

    problem, err := self.problems.GetById(problemId)
    if err != nil {
        return err
    }

    comment := NewComment(text, who, attachments)
    problem.Comments = append(problem.Comments, comment)

    err = self.problems.Update(problem)
    if err != nil {
        return err
    }
    return nil
}

// Erstellt ein neues Problem
func (self *ProblemApplicationService) createNewProblem(cmd *CreateNewProblemCommand) (*Problem, error) {
    result := &Problem{
        Summary:     cmd.Summary,
        Description: cmd.Description,
        Tags:        cmd.Tags,
        CreatedBy:   cmd.CreatedBy,
        CreatedAt:   time.Now(),
    }

    err := self.problems.Insert(result)

    return result, err
}
