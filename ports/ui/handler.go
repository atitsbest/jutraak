package server

import (
    "github.com/codegangsta/martini-contrib/render"
)

func index(r render.Render) {
    r.HTML(200, "index", nil)
}
