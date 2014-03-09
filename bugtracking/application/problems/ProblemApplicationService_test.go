package problems

import (
    "io/ioutil"
    "os"
    "testing"
    "time"

    "github.com/atitsbest/jutraak/bugtracking/domain/entities"
    "github.com/atitsbest/jutraak/ports"

    . "github.com/smartystreets/goconvey/convey"
    "labix.org/v2/mgo"
)

func TestProblemApplicationService(t *testing.T) {

    // Only pass t into top-level Convey calls
    Convey("Given all the properties for a problem", t, func() {
        repository := ports.NewMongoProblemRepository("localhost")
        sut := NewProblemApplicationService(repository)

        summary := "Wir haben ein Problem"
        description := "Nix geht mehr"
        createdBy := "Tester"
        tags := []string{"Tag1", "T A G 2", "Tags 3"}
        var problem *entities.Problem

        Convey("When the problem is posted", func() {
            removeAllProblems()
            problem, _ = sut.CreateNewProblem(
                summary, description, tags, createdBy)

            Convey("Then the problem should be saved by the repository", func() {
                p, err := repository.GetById(problem.Id)
                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)

                Convey("And the created Problem should be returned", func() {
                    So(problem, ShouldNotBeNil)
                })
                Convey("And the problem should be created with the current time", func() {
                    So(time.Since(problem.CreatedAt), ShouldBeLessThan, oneSecond)
                })
                Convey("And the problem should contain the passed parameters", func() {
                    So(p.Summary, ShouldEqual, summary)
                    So(p.Description, ShouldEqual, description)
                    So(p.CreatedBy, ShouldEqual, createdBy)
                    So(p.Tags, ShouldResemble, tags)
                })
            })

            Convey("Given a file", func() {
                data, _ := ioutil.ReadFile("/Users/stephan/dev/go/src/github.com/atitsbest/jutraak/fixtures/image.png")

                Convey("When I attach that file to the problem", func() {
                    err := sut.AttachFileToProblem(problem.Id, "image.png", data)
                    So(err, ShouldBeNil)

                    Convey("Then the problem contains one more attachment", func() {
                        updated, _ := repository.GetById(problem.Id)
                        So(len(updated.Attachments), ShouldEqual, 1)
                        attachData, _ := ioutil.ReadFile(updated.Attachments[0].FilePath)
                        So(attachData, ShouldResemble, data)

                        Reset(func() {
                            os.Remove(updated.Attachments[0].FilePath)

                        })
                    })

                    Convey("When I remove the attachment", func() {
                        Convey("Then the file on disk is gone", nil)

                        Convey("And the poblem has one attachement less", nil)
                    })

                })
            })

            Reset(func() { problem = nil })
        })
    })

}

var oneSecond, _ = time.ParseDuration("1s")

func removeAllProblems() {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("jutraak_test").C("problems")

    _, err = c.RemoveAll(nil)
    if err != nil {
        panic(err)
    }
}
