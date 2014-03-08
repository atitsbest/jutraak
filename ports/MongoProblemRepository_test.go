package ports

import (
    "sort"
    "testing"
    "time"

    "github.com/atitsbest/jutraak/bugtracking/domain/entities"

    . "github.com/smartystreets/goconvey/convey"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

func TestMongoProblemRepository(t *testing.T) {

    // Only pass t into top-level Convey calls
    Convey("Given a Mongo-Problems-Repository", t, func() {
        sut := NewMongoProblemRepository("localhost")

        Convey("When I insert a new Problem", func() {
            problem := &entities.Problem{
                Summary:     "Wir haben ein Problem",
                Description: "Nix geht mehr!",
                Tags:        []string{"Tag1", "Tag 2", "Bug"},
                CreatedBy:   "Tester",
                CreatedAt:   time.Now(),
            }
            sut.Insert(problem)

            Convey("Then the problem should be in the MongoDB", func() {
                inserted := getMongoProblemById(problem.Id)
                So(inserted, ShouldNotBeNil)
                So(inserted.Summary, ShouldEqual, problem.Summary)
                So(inserted.Tags, ShouldResemble, problem.Tags)
            })
            Convey("Then the new Id should be set in the problem", func() {
                So(problem.Id, ShouldNotBeBlank)
            })

            Reset(func() { removeAllProblems() })
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

            Reset(func() { removeAllProblems() })
        })
    })
}

func getMongoProblemById(id string) *entities.Problem {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("jutraak_test").C("problems")

    result := &entities.Problem{}
    err = c.Find(bson.M{"id": id}).One(result)
    if err != nil {
        panic(err)
    }

    return result
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
