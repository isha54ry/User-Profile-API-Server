package models

type User struct {
    ID       string
    Name     string
    ImageURL string
    PDFPath  string
}

var Users = make(map[string]*User)