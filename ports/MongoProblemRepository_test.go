package ports

import (
    "sort"
    "testing"
    "time"

    "github.com/atitsbest/jutraak/bugtracking/domain/entities"

    uuid "github.com/nu7hatch/gouuid"
    . "github.com/smartystreets/goconvey/convey"
    "labix.org/v2/mgo"
)

func TestMongoProblemRepository(t *testing.T) {
    var problem *entities.Problem
    var inserted *entities.Problem
    sut := NewMongoProblemRepository("localhost")

    // Only pass t into top-level Convey calls
    Convey("Given a Mongo-Problems-Repository", t, func() {
        removeAllProblems()

        Convey("When I insert a new Problem", func() {
            problem = &entities.Problem{
                Summary:     "Wir haben ein Problem",
                Description: "Nix geht mehr!",
                Tags:        []string{"Tag1", "Tag 2", "Bug"},
                CreatedBy:   "Tester",
                CreatedAt:   time.Now(),
            }
            err := sut.Insert(problem)
            So(err, ShouldBeNil)

            Convey("Then the problem should be in MongoDB", func() {
                inserted, _ = sut.GetById(problem.Id)
                So(inserted, ShouldNotBeNil)
                So(inserted.Summary, ShouldEqual, problem.Summary)
                So(inserted.Tags, ShouldResemble, problem.Tags)

                Convey("And the new Id should be a valid ProblemId", func() {

                    _, err := uuid.ParseHex(string(problem.Id))
                    So(err, ShouldBeNil)
                    So(problem.Id, ShouldEqual, inserted.Id)
                })
            })

            Convey("When I get the inserted Problem", func() {
                inserted, _ = sut.GetById(problem.Id)

                Convey("And update it with new values", func() {
                    inserted.Summary = "Hat sich geändert"
                    inserted.Tags = []string{"Bug", "CR"}
                    err := sut.Update(inserted)
                    So(err, ShouldBeNil)

                    Convey("Then the updated values should be persited", func() {
                        updated, _ := sut.GetById(problem.Id)
                        So(updated, ShouldNotBeNil)
                        So(updated.Summary, ShouldEqual, inserted.Summary)
                        So(updated.Tags, ShouldResemble, inserted.Tags)
                    })
                })
            })

            Reset(func() { problem = nil })
        })

        Convey("Given 3 problems in the db", func() {
            sut.Insert(&entities.Problem{Tags: []string{"Tag 2", "Tag1"}})
            sut.Insert(&entities.Problem{Tags: []string{"Bug"}})
            sut.Insert(&entities.Problem{Tags: []string{"Tag1"}})

            Convey("When I request all Problem tags", func() {
                tags, _ := sut.GetAllTags()
                sort.Strings(tags)

                Convey("Then I get a distinct list of all tags", func() {
                    So(tags, ShouldResemble, []string{"Bug", "Tag 2", "Tag1"})
                })

            })

            Convey("When I request them all", func() {
                problems, _ := sut.GetAllProblems()

                Convey("Then I get a list of all 3 problems", func() {
                    So(len(problems), ShouldEqual, 3)
                })
            })

            Convey("When I search for problems with a tag", func() {
                problems, _ := sut.GetProblemsByTag([]string{"Tag1"})

                Convey("Then I get a list of problems with this tag assigned", func() {
                    So(len(problems), ShouldEqual, 2)
                    So(problems[0].Tags, ShouldContain, "Tag1")
                    So(problems[1].Tags, ShouldContain, "Tag1")
                })
            })

            Convey("When I request a problem by Id", func() {
                problems, _ := sut.GetAllProblems()
                single, err := sut.GetById(problems[0].Id)

                Convey("Then I get that single problem", func() {
                    So(single, ShouldNotBeNil)
                    So(err, ShouldBeNil)
                    So(single.Id, ShouldEqual, problems[0].Id)
                })
            })
        })
    })
}

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
