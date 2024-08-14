package pages

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"real-time-forum-base/internal/data"
	"real-time-forum-base/internal/sqlite"
	"strconv"
	"time"
)

func GetSinglePostJson(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err1 := template.ParseFiles("index.html")
		if err1 != nil {
			fmt.Println("i dont know what i am doing at this point")
		}
		t.ExecuteTemplate(w, "index.html", nil)
	}

	if r.Method == "POST" {
		fmt.Println("singlePost data?")

		err := CheckCookieExist(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Failed"}`))
			return
		}

		var pageDataJson []byte
		postId, err2 := strconv.Atoi(r.URL.Query().Get("id"))
		if err2 != nil {
			return
		}
		pageDataJson, err = json.Marshal(sqlite.SinglePostDataGathering(postId))
		if err != nil {
			fmt.Println("fuckfuckfuckfuckfuckfuckfuckfuck", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(pageDataJson)
	}
}

func gatherCommentData(form url.Values) (data.Comment, error) {
	res := data.Comment{}
	res.DateCreated = time.Now()
	res.CommentText = form.Get("comment")
	var err error
	res.PostId, err = strconv.Atoi(form.Get("id"))
	return res, err
}

func SinglePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err1 := template.ParseFiles("index.html")
		if err1 != nil {
			fmt.Println("i dont know what i am doing at this point")
		}
		t.ExecuteTemplate(w, "index.html", nil)
	}

	fmt.Println("get's to post")
	if r.Method == http.MethodPost {
		r.ParseMultipartForm(100000)
		postData, err2 := gatherCommentData(r.Form)
		if err2 != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Failed"}`))
			return
		}
		session, err := r.Cookie("session_token")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Failed"}`))
			return
		}
		uid, err := sqlite.CheckSessionExist(session.Value)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Failed"}`))
			return
		}
		fmt.Println("first single post?")
		postData.AuthorId = uid
		sqlite.AddCommentToTable(postData)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"Success"}`))
	}
}
