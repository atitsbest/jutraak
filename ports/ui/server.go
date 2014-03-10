package server

import (
    "github.com/codegangsta/martini"
    // "github.com/codegangsta/martini-contrib/binding"
    "github.com/atitsbest/jutraak/bugtracking/application/problems"
    . "github.com/atitsbest/jutraak/config"
    "github.com/atitsbest/jutraak/ports"
    "github.com/codegangsta/martini-contrib/render"
)

func InitServer() *martini.ClassicMartini {
    m := martini.Classic()

    problemRepository := ports.NewMongoProblemRepository(Config.ConnectionString)
    problemService := problems.NewProblemApplicationService(problemRepository)

    // DEPENDENCY-INJECTION
    m.MapTo(
        problemService,
        (*problems.ProblemApplicationServiceInterface)(nil))
    m.MapTo(
        problemRepository,
        (*problems.ProblemRepository)(nil))

    m.Use(render.Renderer(render.Options{
        Extensions: []string{".html"},
        Layout:     "_layout",
        Delims:     render.Delims{"{%", "%}"},
    }))

    m.Get("/", index)
    // m.Get("/projects", projectIndex)
    // m.Get("/projects/new", projectEdit)

    m.Get("/api/problems", ApiGetProblems)
    m.Get("/api/problems/tags", ApiGetProblemTags)
    m.Get("/api/problems/:id", ApiGetProblems)
    // m.Post("/api/projects", binding.Bind(handlers.ProjectPostModel{}), handlers.ApiPostProject)
    return m
}
