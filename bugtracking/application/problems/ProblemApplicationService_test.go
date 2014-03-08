package problems

import (
    "io/ioutil"
    "testing"
    "time"

    "github.com/atitsbest/jutraak/bugtracking/domain/entities"

    . "github.com/smartystreets/goconvey/convey"
)

func TestProblemApplicationService(t *testing.T) {
    repository := &ProblemsMock{
        P: &entities.Problem{Id: entities.NewProblemId()},
    }
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

            Convey("Then the problem should be saved by the repository", func() {
                So(repository.InsertCount, ShouldEqual, 1)
            })
            Convey("And the created Problem should be returned", func() {
                So(problem, ShouldNotBeNil)
            })
            Convey("And the problem should be created with the current time", func() {
                So(time.Since(problem.CreatedAt), ShouldBeLessThan, oneSecond)
            })
            Convey("And the problem should contain the passed parameters", func() {
                So(problem.Summary, ShouldEqual, summary)
                So(problem.Description, ShouldEqual, description)
                So(problem.CreatedBy, ShouldEqual, createdBy)
                So(problem.Tags, ShouldResemble, tags)
            })
        })
    })

    Convey("Given a problem and a file", t, func() {
        data, _ := ioutil.ReadFile("fixtures/image.png")
        problemId := entities.NewProblemId()

        Convey("When I attach that file to the problem", func() {
            err := sut.AttachFileToProblem(problemId, "image.png", data)
            So(err, ShouldBeNil)
            So(repository.GetByIdCount, ShouldEqual, 1)

            Convey("And the problem contains one more attachment", func() {
                So(repository.UpdateCount, ShouldEqual, 1)
                So(len(repository.P.Attachments), ShouldEqual, 1)
            })
        })
    })
}

// ProblemRepository Mock

type ProblemsMock struct {
    InsertCount  int
    UpdateCount  int
    GetByIdCount int

    P   *entities.Problem
}

func (self *ProblemsMock) Insert(problem *entities.Problem) error {
    self.InsertCount += 1
    return nil
}

func (self *ProblemsMock) Update(problem *entities.Problem) error {
    self.UpdateCount += 1
    return nil
}

func (self *ProblemsMock) GetById(id entities.ProblemId) (*entities.Problem, error) {
    self.GetByIdCount += 1
    return self.P, nil
}

var oneSecond, _ = time.ParseDuration("1s")
