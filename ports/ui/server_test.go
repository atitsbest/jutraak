package server

import (
    "github.com/codegangsta/martini"
    . "github.com/smartystreets/goconvey/convey"
    "net/http"
    "net/http/httptest"
    "testing"
)

// Basis Url.
var port = ":3002"
var baseUrl = "http://localhost" + port

func Test_Server(t *testing.T) {
    var sut *martini.ClassicMartini

    // Convey("Subject: Server", t, func() {
    //     Convey("A running server", func() {
    //         sut := initMartini()
    //         go http.ListenAndServe(":3001", sut)
    //
    //         Convey("should GET /", func() { shouldServe(t, baseUrl+"/") })
    //     })
    // })

    Convey("Subject: Problem API", t, func() {
        Convey("A running server", func() {
            sut = InitServer()
            go http.ListenAndServe(port, sut)

            Convey("should GET /api/problems", func() {
                shouldServe(t, sut, baseUrl+"/api/problems")
            })
            Convey("should GET /api/problems/12345", func() {
                shouldServe(t, sut, baseUrl+"/api/problems/12345")
            })

            // Convey("should POST /api/projects", func() {
            //     response := httptest.NewRecorder()
            //     json := `{
            //           "name":"Eriksson",
            //           "leader":"RatAn",
            //           "risk":"A",
            //           "accountingMode":"Fixpreis",
            //           "state":"beauftragt",
            //           "orderDate":"2014-02-19T07:48:33.833Z",
            //           "techs":["C#","F#","JavaScript","SharePoint","Progress"],
            //           "customer":"sldslkdh",
            //           "orderAmount":33.99,
            //           "orderAmountDays":44
            //         }`
            //     b := strings.NewReader(json)
            //     req, err := http.NewRequest("POST", baseUrl+"/api/projects", b)
            //     if err != nil {
            //         t.Error(err)
            //     }
            //
            //     // Act
            //     sut.ServeHTTP(response, req)
            //
            //     // Assert
            //     So(response.Code, ShouldEqual, http.StatusCreated)
            //     So(response.Body.Len(), ShouldBeGreaterThan, 0)
            // })
        })
    })
}

// ========================================================
func shouldServe(t *testing.T, server *martini.ClassicMartini, url string) {
    // Arrange
    response := httptest.NewRecorder()
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        t.Error(err)
    }

    // Act
    server.ServeHTTP(response, req)

    // Assert
    So(response.Code, ShouldEqual, http.StatusOK)
    So(response.Body.Len(), ShouldBeGreaterThan, 0)
}
