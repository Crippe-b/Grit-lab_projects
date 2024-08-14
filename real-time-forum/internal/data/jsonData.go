package data

import (
	"time"
)

type MainPage struct {
	Posts      []SPostMain `json:"posts"`
	Categories []Topic     `json:"categories"`
}
type PostPage struct {
	Post     Post          `json:"post"`
	Comments []PostComment `json:"comments"`
}
type PostJ struct {
	Author     string    `json:"author"`
	Title      string    `json:"title"`
	Body       []byte    `json:"body"`
	Created    time.Time `json:"created"`
	Categories string    `json:"categories"`
}

type SPostMain struct {
	PostId     int       `json:"postid"`
	Author     string    `json:"author"`
	Title      string    `json:"title"`
	Created    time.Time `json:"created"`
	Categories []string  `json:"categories"`
}

type SinglePost struct {
	PostId     int           `json:"postid"`
	Author     string        `json:"author"`
	Title      string        `json:"title"`
	Body       string        `json:"body"`
	Created    time.Time     `json:"created"`
	Categories []string      `json:"categories"`
	Comments   []PostComment `json:"comments"`
}
type PostComment struct {
	Id          int    `json:"id"`
	Author      string `json:"author"`
	Created     string `json:"created"`
	CommentText string `json:"commenttext"`
}

type Topic struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
}

type Users struct {
	Username   string    `json:"username"`
	Loggedin   bool      `json:"loggedin"`
	HasSentmsg bool      `json:"hassentmsg"`
	LastMsg    time.Time `json:"lastmsg"`
	You        bool      `json:"you"`
}

type MsgJ struct {
	Sender string    `json:"sender"`
	Msg    string    `json:"message"`
	Date   time.Time `json:"timestamp"`
}
