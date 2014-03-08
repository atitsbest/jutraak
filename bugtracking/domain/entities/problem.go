package entities

import (
    "time"

    "github.com/atitsbest/jutraak/bugtracking/domain"
)

type Problem struct {
    Id          string
    Summary     string
    Description string
    Tags        []string
    CreatedBy   string
    CreatedAt   time.Time

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
