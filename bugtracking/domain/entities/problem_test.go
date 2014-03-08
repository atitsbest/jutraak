package entities

import (
    "testing"

    . "github.com/smartystreets/goconvey/convey"
)

func TestCloseProblem(t *testing.T) {

    Convey("Given an open problem", t, func() {
        sut := Problem{
            Summary:     "Problem",
            Description: "Soll geschlossen werden!",
            CreatedBy:   "Tester",
        }

        Convey("When I resolve the problem", func() {
            err := sut.Resolve()

            Convey("Then the problem is resolved", func() {
                So(sut.IsResolved(), ShouldEqual, true)
            })
            Convey("And no error was returned", func() {
                So(err, ShouldBeNil)
            })
        })
    })

    Convey("Given a closed problem", t, func() {
        sut := Problem{
            Summary:     "Problem",
            Description: "Soll geschlossen werden!",
            CreatedBy:   "Tester",
        }
        sut.Resolve()

        Convey("When I resolve the problem", func() {
            err := sut.Resolve()

            Convey("Then an error is returned", func() {
                So(err, ShouldNotBeNil)
            })
        })
    })
}
