package problems

import (
    "testing"
    "time"

    "github.com/atitsbest/jutraak/bugtracking/domain/entities"

    . "github.com/smartystreets/goconvey/convey"
)

func TestCreateNewProblem(t *testing.T) {
    repository := new(ProblemsMock)
    sut := NewProblemApplicationService(repository)

    // Only pass t into top-level Convey calls
    Convey("Given all the properties for a problem", t, func() {
        summary := "Wir haben ein Problem"
        description := "Nix geht mehr"
        createdBy := "Tester"
        tags := []string{"Tag1", "T A G 2", "Tags 3"}

        Convey("When the problem is posted", func() {
            problem, _ := sut.CreateNewProblem(
                summary, description, tags, createdBy)

            Convey("Then the created Problem should be returned", func() {
                So(problem, ShouldNotBeNil)
            })
            Convey("Then the problem should be created with the current time", func() {
                So(time.Since(problem.CreatedAt), ShouldBeLessThan, oneSecond)
            })
            Convey("Then the problem should contain the passed parameters", func() {
                So(problem.Summary, ShouldEqual, summary)
                So(problem.Description, ShouldEqual, description)
                So(problem.CreatedBy, ShouldEqual, createdBy)
                So(problem.Tags, ShouldResemble, tags)
            })
            Convey("Then the problem should be saved by the repository", func() {
                So(repository.InsertCount, ShouldEqual, 4) // 4 Conveys == 4x speichern.
            })
        })
    })
}

type ProblemsMock struct {
    InsertCount int
}

func (self *ProblemsMock) Insert(problem *entities.Problem) error {
    self.InsertCount += 1
    return nil
}

var oneSecond, _ = time.ParseDuration("1s")
