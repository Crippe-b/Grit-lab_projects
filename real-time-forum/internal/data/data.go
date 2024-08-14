package data

import (
	"database/sql"
	"time"
)

type Registration struct {
	Username         string
	Email            string
	Password         string
	Gender           string
	Age              int
	Firstname        string
	Lastname         string
	Cokkie           string
	RegistrationDate time.Time
}

type Post struct {
	AuthorId   int
	Title      string
	Body       []byte
	Created    time.Time
	Categories string
	//ImagePath  string
}

type Comment struct {
	PostId      int
	AuthorId    int
	DateCreated time.Time
	CommentText string
}

type Category struct {
	Category string
}

type Msg struct {
	Sender   string
	Reciever string
	Comment  string
}

var DB *sql.DB
