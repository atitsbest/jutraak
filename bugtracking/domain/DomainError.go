package domain

type DomainError struct {
    Message string
}

// CTR
func NewDomainError(msg string) *DomainError {
    return &DomainError{Message: msg}
}

// error Interface
func (self *DomainError) Error() string {
    return self.Message
}
