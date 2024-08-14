package main

//works for now, might need some clean up stuff and add something

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"real-time-forum-base/internal/chatsocket"
	"real-time-forum-base/internal/data"
	"real-time-forum-base/internal/pages"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	data.DB, _ = sql.Open("sqlite3", "./database/forum.db")
	data.DB.SetMaxOpenConns(10)
	defer data.DB.Close()

	css := http.FileServer(http.Dir("static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	js := http.FileServer(http.Dir("static/js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	srv := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", pages.Pages)
	http.HandleFunc("/register", pages.Registration)
	http.HandleFunc("/loggedin", pages.Login)
	http.HandleFunc("/createpost", pages.CreatePost)
	http.HandleFunc("/post", pages.SinglePost)

	http.HandleFunc("/hello", pages.SendData)
	http.HandleFunc("/goodbye", pages.GetSinglePostJson)
	http.HandleFunc("/losers", pages.SendChatData)
	http.HandleFunc("/nodobylikesu", pages.SendMessages)

	http.HandleFunc("/ws", chatsocket.WsEndpoint)

	fmt.Println("Starting application on port " + srv.Addr)
	if srv.ListenAndServe() != nil {
		log.Fatalf("%v - Internal Server Error", http.StatusInternalServerError)
	}

}
