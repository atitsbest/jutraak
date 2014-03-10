package config

type Settings struct {
    ConnectionString string
    AttachmentsPath  string
}

var Config = &Settings{
    AttachmentsPath:  "/tmp/jutraak/",
    ConnectionString: "localhost",
}
