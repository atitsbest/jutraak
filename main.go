package main

import (
    server "github.com/atitsbest/jutraak/ports/ui"
)

func main() {
    m := server.InitServer()
    m.Run()
}
