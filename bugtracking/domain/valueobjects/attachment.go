package valueobjects

import (
    "io/ioutil"
    "mime"
    "path"
    "path/filepath"

    . "github.com/atitsbest/jutraak/config"
    uuid "github.com/nu7hatch/gouuid"
)

type Attachment struct {
    FileName    string
    ContentType string
    FilePath    string // Pfad zur gespeicherten Datei.
}

// Neues Attachment erstellen.
// Dabei werden die Daten in eine Datei geschrieben.
func NewAttachment(fileName string, data []byte) (*Attachment, error) {
    // Dateiname erzeugen.
    filePath, err := generateAttachmentPath(fileName)
    if err != nil {
        return nil, err
    }

    // Daten in Datei speichern.
    err = ioutil.WriteFile(filePath, data, 0644)
    if err != nil {
        return nil, err
    }

    // Attachment erstellen und zurückgeben.
    return &Attachment{
        // Content-Type herausfinden.
        ContentType: contentType(fileName),
        FilePath:    filePath,
        FileName:    fileName,
    }, nil
}

// Liefert den Mime-Type anhand der Datei-Erweiterung.
func contentType(fileName string) string {
    ext := filepath.Ext(fileName)
    return mime.TypeByExtension(ext)
}

// Erzeugt einen Pfad aus /base-path/uuid.ext
func generateAttachmentPath(fileName string) (string, error) {
    ext := filepath.Ext(fileName)
    id, err := uuid.NewV4()
    if err != nil {
        return "", err
    }

    result := path.Join(Config.AttachmentsPath, id.String()+ext)
    return result, nil
}
