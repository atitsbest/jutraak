package repositories_test

import (
    . "github.com/atitsbest/jutraak/ports/repositories"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestSqlProblemRepository(t *testing.T) {
    var (
        sut *SqlProblemRepository
        err error //
    )

    Convey("When I connect to database", t, func() {
        Convey("with a valid connection string", func() {
            sut, err = NewSqlProblemRepsoitory("user=jutraak dbname=jutraak_test sslmode=disable")

            Convey("it should create repository", func() {
                So(err, ShouldBeNil)
                So(sut, ShouldNotBeNil)
            })
        })

        Convey("with an invalid connection string", func() {
            sut, err = NewSqlProblemRepsoitory("wrong=verywrong")

            Convey("it should return an error an no repository", func() {
                So(err, ShouldNotBeNil)
                So(sut, ShouldBeNil)
            })
        })
    })

    Convey("Given 3 problems", t, func() {
        sut, err = NewSqlProblemRepsoitory("user=jutraak dbname=jutraak_test sslmode=disable")
        So(err, ShouldBeNil)

        Convey("When I query all problems", func() {
            Convey("I get an arry with alle 3 problems", nil)
        })
    })
}
