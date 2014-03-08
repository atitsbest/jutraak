package entities

import (
    "testing"

    "github.com/atitsbest/jutraak/bugtracking/domain/valueobjects"
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

    Convey("Given an not yet posted problem", t, func() {
        sut := Problem{
            Summary:     "Problem",
            Description: "Hier kommen Dateien dran",
            CreatedBy:   "Tester",
        }

        Convey("When I attach a file to the problem", func() {
            err := sut.AddAttachment(&valueobjects.Attachment{})

            Convey("Then I get an error", func() {
                So(err, ShouldNotBeNil)
            })
        })
    })

    Convey("Given a problem", t, func() {
        sut := Problem{
            Id:          "CR1",
            Summary:     "Problem",
            Description: "Hier kommen Dateien dran",
            CreatedBy:   "Tester",
        }

        Convey("When I attach a file to the problem", func() {
            file := &valueobjects.Attachment{
                FileName:    "image.jpg",
                ContentType: "image/jpeg",
            }
            err := sut.AddAttachment(file)

            Convey("Then the file is attached to the problem", func() {
                So(len(sut.Attachments), ShouldEqual, 1)
            })

            Convey("And I dont get an error", func() {
                So(err, ShouldBeNil)
            })

            Convey("And the Content-Type is stored", func() {
                So(sut.Attachments[0].ContentType, ShouldEqual, "image/jpeg")
            })

            Convey("And the original file name is stored", func() {
                So(sut.Attachments[0].FileName, ShouldEqual, "image.jpg")
            })
        })
    })
}
