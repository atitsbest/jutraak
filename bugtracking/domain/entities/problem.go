package entities

import (
    "time"

    "github.com/atitsbest/jutraak/bugtracking/domain"
    . "github.com/atitsbest/jutraak/bugtracking/domain/valueobjects"
    uuid "github.com/nu7hatch/gouuid"
)

type ProblemId string

// Neue ProblemId erzeugen.
func NewProblemId() ProblemId {
    id, err := uuid.NewV4()
    if err != nil {
        panic(err)
    }

    return ProblemId(id.String())
}

type Problem struct {
    Id          ProblemId
    Summary     string
    Description string
    Tags        []string
    CreatedBy   string
    CreatedAt   time.Time
    Attachments []*Attachment
    Comments    []string

    resolved bool
}

// Schließt ein Problem.
func (self *Problem) Resolve() error {
    if self.IsResolved() {
        return domain.NewDomainError("Das Problem ist bereits behoben.")
    }
    self.resolved = true
    return nil
}

// Ist das Problem behoben?
func (self *Problem) IsResolved() bool {
    return self.resolved == true
}

// Eine Datei an das Problem anhängen.
func (self *Problem) AddAttachment(file *Attachment) error {
    if len(self.Id) == 0 {
        return domain.NewDomainError("Datei kann nur bei einem gespeicherten Problem angehängt werden!")
    }
    self.Attachments = append(self.Attachments, file)
    return nil
}
