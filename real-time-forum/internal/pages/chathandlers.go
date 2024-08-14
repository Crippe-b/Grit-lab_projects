package pages

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"real-time-forum-base/internal/chatsocket"
	"real-time-forum-base/internal/sqlite"
)

func SendChatData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err1 := template.ParseFiles("index.html")
		if err1 != nil {
			fmt.Println("i dont know what i am doing at this point")
		}
		t.ExecuteTemplate(w, "index.html", nil)
	}

	if r.Method == "POST" && r.URL.Path == "/losers" {

		err := CheckCookieExist(w, r)
		if err != nil {
			fmt.Println("failed at /losers")
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Failed"}`))
			return
		}

		session, _ := r.Cookie("session_token")
		uid, err := sqlite.CheckSessionExist(session.Value)
		if err != nil {
			return
		}
		var pageDataJson []byte
		username := sqlite.GetUsername(uid)
		pageDataJson, err = json.Marshal(chatsocket.GetAllUsers(username))
		if err != nil {
			fmt.Println("fuckfuckfuckfuckfuckfuckfuckfuck", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(pageDataJson)
	}
}

func SendMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err1 := template.ParseFiles("index.html")
		if err1 != nil {
			fmt.Println("i dont know what i am doing at this point")
		}
		t.ExecuteTemplate(w, "index.html", nil)
	}

	if r.Method == "POST" && r.URL.Path == "/nodobylikesu" {

		err := CheckCookieExist(w, r)
		if err != nil {
			fmt.Println("failed at /nobodylikesu")
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Failed"}`))
			return
		}

		session, err := r.Cookie("session_token")
		if err != nil {
			return
		}
		uid, err := sqlite.CheckSessionExist(session.Value)
		if err != nil {
			return
		}
		username := sqlite.GetUsername(uid)

		var pageDataJson []byte
		pageDataJson, err = json.Marshal(sqlite.GetallMessages(r.Header.Get("Body"), username))
		if err != nil {
			fmt.Println("fuckfuckfuckfuckfuckfuckfuckfuck", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(pageDataJson)
	}
}
