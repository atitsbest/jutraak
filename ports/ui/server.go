package server

import (
    "github.com/codegangsta/martini"
    // "github.com/codegangsta/martini-contrib/binding"
    "github.com/codegangsta/martini-contrib/render"
)

func InitServer() *martini.ClassicMartini {
    m := martini.Classic()

    // DEPENDENCY-INJECTION

    m.Use(render.Renderer(render.Options{
        Extensions: []string{".html"},
        Layout:     "_layout",
        Delims:     render.Delims{"{%", "%}"},
    }))

    // m.Get("/", projectIndex)
    // m.Get("/projects", projectIndex)
    // m.Get("/projects/new", projectEdit)

    m.Get("/api/problems", ApiGetProblems)
    m.Get("/api/problems/:id", ApiGetProblems)
    // m.Post("/api/projects", binding.Bind(handlers.ProjectPostModel{}), handlers.ApiPostProject)
    return m
}
