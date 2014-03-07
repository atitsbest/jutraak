package ports

import (
  "time"
  "testing"

  "dccs/jutraak/bugtracking/entities"

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
              Summary: "Wir haben ein Problem",
              Description: "Nix geht mehr!",
              CreatedBy: "Tester",
              CreatedAt: time.Now(),
            }
            sut.Insert(problem)

            Convey("Then the problem should be in the MongoDB", func() {
              inserted := getMongoProblemById(problem.Id)
              So(inserted, ShouldNotBeNil)
              So(inserted.Summary, ShouldEqual, problem.Summary)
            })
            Convey("Then the new Id should be set in the problem", func() {
              So(problem.Id, ShouldNotBeBlank)
            })
        })
    })
}

func getMongoProblemById(id string) *entities.Problem {
  session, err := mgo.Dial("localhost")
  if err != nil { panic(err) }
  defer session.Close()

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)

  c := session.DB("jutraak_test").C("problems")

  result := &entities.Problem{}
  err = c.Find(bson.M{"id": id}).One(result)
  if err != nil { panic(err) } 

  return result
}


