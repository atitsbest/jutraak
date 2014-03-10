package problems

import (
    "io/ioutil"
    "os"
    "testing"
    "time"

    "github.com/atitsbest/jutraak/bugtracking/domain/entities"
    . "github.com/atitsbest/jutraak/bugtracking/domain/valueobjects"
    . "github.com/atitsbest/jutraak/config"
    "github.com/atitsbest/jutraak/ports/repositories"

    . "github.com/smartystreets/goconvey/convey"
    "labix.org/v2/mgo"
)

func TestProblemApplicationService(t *testing.T) {
    var (
        repository *repositories.MongoProblemRepository
        sut        *ProblemApplicationService
        problem    *entities.Problem
        err        error
    )
    data, _ := ioutil.ReadFile("/Users/stephan/dev/go/src/github.com/atitsbest/jutraak/fixtures/image.png")

    // Only pass t into top-level Convey calls
    Convey("Given all the properties for a problem", t, func() {
        repository = repositories.NewMongoProblemRepository(Config.ConnectionString)
        sut = NewProblemApplicationService(repository)

        summary := "Wir haben ein Problem"
        description := "Nix geht mehr"
        createdBy := "Tester"
        tags := []string{"Tag1", "T A G 2", "Tags 3"}

        Convey("When the problem is posted", func() {
            removeAllProblems()
            problem, err = sut.CreateNewProblem(summary, description, tags, createdBy)
            So(err, ShouldBeNil)

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

            Convey("When I change the problem Description/Summary", func() {
                sut.ChangeProblemSummary(problem.Id, "Neue Summary", "Neue Description", "Changer")

                Convey("Then the changes are persisted", func() {
                    problem, err = repository.GetById(problem.Id)
                    So(err, ShouldBeNil)
                    So(problem.Summary, ShouldEqual, "Neue Summary")
                    So(problem.Description, ShouldEqual, "Neue Description")

                    Convey("And the change date/user have been updated", func() {
                        So(problem.LastChangeBy, ShouldEqual, "Changer")
                        So(time.Since(problem.LastChangeAt), ShouldBeLessThan, oneSecond)
                    })
                })
            })

            Convey("When I request all problems", func() {
                problems, _ := sut.GetAllProblems()

                Convey("Then I get at least the posted problem", func() {
                    So(len(problems), ShouldBeGreaterThan, 0)
                })
            })

            Convey("Given a file", func() {

                Convey("When I attach that file to the problem", func() {
                    err := sut.AttachFileToProblem(problem.Id, "image.png", data)
                    So(err, ShouldBeNil)

                    Convey("Then the problem contains one more attachment", func() {
                        problem, _ = repository.GetById(problem.Id)
                        So(len(problem.Attachments), ShouldEqual, 1)
                        attachData, _ := ioutil.ReadFile(problem.Attachments[0].FilePath)
                        So(attachData, ShouldResemble, data)

                        Reset(func() {
                            os.Remove(problem.Attachments[0].FilePath)
                        })
                    })

                    Convey("When I remove the attachment", func() {
                        problem, _ = repository.GetById(problem.Id)
                        filePath := problem.Attachments[0].FilePath
                        sut.RemoveProblemAttachment(problem.Id, filePath)

                        Convey("Then the file on disk is gone", func() {
                            _, err := os.Stat(filePath)
                            So(err, ShouldNotBeNil)
                        })

                        Convey("And the poblem has one attachement less", func() {
                            problem, _ := repository.GetById(problem.Id)
                            So(len(problem.Attachments), ShouldEqual, 0)
                        })
                    })

                })
            })

            Convey("When I add a comment", func() {
                sut.CommentProblem(problem.Id, "Comment", "Tester", nil)

                Convey("Then the problem contains a comment", func() {
                    problem, _ = repository.GetById(problem.Id)
                    So(len(problem.Comments), ShouldEqual, 1)

                    var comment *entities.Comment
                    comment = problem.Comments[0]

                    Convey("And the comment contains who and when", func() {
                        So(comment.CreatedBy, ShouldEqual, "Tester")
                        So(time.Since(comment.CreatedAt), ShouldBeLessThan, oneSecond)
                    })

                })
            })

            Convey("When I add a comment with attachment", func() {
                ca, err := NewAttachment("tester.jpg", data)
                So(err, ShouldBeNil)
                sut.CommentProblem(problem.Id, "Comment", "Tester", []*Attachment{ca})

                Convey("Then the problem contains a comment with attachment", func() {
                    problem, _ = repository.GetById(problem.Id)
                    So(len(problem.Comments[0].Attachments), ShouldEqual, 1)
                    cdata, _ := ioutil.ReadFile(problem.Comments[0].Attachments[0].FilePath)
                    So(cdata, ShouldResemble, data)

                    Reset(func() {
                        os.Remove(problem.Comments[0].Attachments[0].FilePath)
                    })
                })
            })

        })

    })

}

var oneSecond, _ = time.ParseDuration("1s")

func removeAllProblems() {
    session, err := mgo.Dial(Config.ConnectionString)
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
