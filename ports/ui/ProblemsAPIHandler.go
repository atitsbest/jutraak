package server

import (
    // "log"
    "net/http"
    // "time"

    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/render"

    "github.com/atitsbest/jutraak/bugtracking/application/problems"
)

type (
    ApiGetProblemsParams struct {
        Query string   `form:"q"`
        Tags  []string `form:"tags"`
    }
)

func ApiGetProblems(params ApiGetProblemsParams, r render.Render, app problems.ProblemRepository) {
    p, err := app.Filtered(params.Tags, params.Query)
    if err != nil {
        r.JSON(http.StatusInternalServerError, err)
    } else {
        r.JSON(200, p)
    }
}

func ApiGetProblem(params martini.Params, r render.Render) {
    ps := []string{"test1", "test2"}
    r.JSON(200, ps[0])
}

func ApiGetProblemTags(params martini.Params, r render.Render, app problems.ProblemRepository) {
    ts, err := app.AllTags()
    if err != nil {
        r.JSON(http.StatusInternalServerError, err)
    } else {
        r.JSON(200, ts)
    }
}

// func ApiPostProject(prj ProjectPostModel, rep application.ProjectRepository, l *log.Logger) (int, string) {
//     p := prj.toProject()
//     l.Printf("%#v", p)
//     rep.Add(p)
//     return http.StatusCreated, prj.Name
// }
