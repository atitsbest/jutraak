package repositories_test

import (
    . "github.com/atitsbest/jutraak/bugtracking/domain/entities"
    . "github.com/atitsbest/jutraak/ports/repositories"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestSqlProblemRepository(t *testing.T) {
    var (
        sut      *SqlProblemRepository
        err      error
        problems []*Problem
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
        execSql("fixtures/insert_3_test_problems.sql")

        Convey("When I query all problems", func() {
            problems, err = sut.All()

            Convey("I get an array with alle 3 problems", func() {
                So(len(problems), ShouldEqual, 3)
                So(err, ShouldBeNil)
            })
        })

        Reset(func() { removeAllProblems() })
    })
}

func execSql(file string) {
    db, err := sqlx.Connect("postgres", "user=jutraak dbname=jutraak_test sslmode=disable")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    db.LoadFile(file)
}

func removeAllProblems() {
    db, err := sqlx.Connect("postgres", "user=jutraak dbname=jutraak_test sslmode=disable")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    db.Execf("DELETE FROM problems")
}
