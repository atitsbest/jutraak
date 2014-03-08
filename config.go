package main

type Configuration struct {
    AttachmentsPath string
}

var Config = Configuration{
    AttachmentsPath: "/tmp/jutraak/",
}
