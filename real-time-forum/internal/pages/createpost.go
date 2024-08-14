package pages

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"real-time-forum-base/internal/data"
	"real-time-forum-base/internal/sqlite"
	"strings"
	"time"
)

func gatherPostData(form url.Values) (data.Post, error) {

	post := data.Post{}
	post.Body = []byte(form.Get("content"))
	if len(post.Body) == 0 {
		return post, errors.New("No_Content")
	}
	post.Title = form.Get("title")
	if len(post.Title) == 0 {
		return post, errors.New("No_Title")
	}
	post.Categories = strings.Join(form["category"], "?")
	if len(post.Categories) == 0 {
		return post, errors.New("No_Category")
	}
	post.Created = time.Now()
	//post.ImagePath = form.Get("image")
	return post, nil
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err1 := template.ParseFiles("index.html")
		if err1 != nil {
			fmt.Println("i dont know what i am doing at this point")
		}
		t.ExecuteTemplate(w, "index.html", nil)
	}

	if r.Method == "POST" {
		err := CheckCookieExist(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Failed"}`))
			return
		}
		r.ParseMultipartForm(100000)
		fmt.Println(r.Form)
		postData, err := gatherPostData(r.Form)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":` + `"` + err.Error() + `"}`))
			return
		}
		session, _ := r.Cookie("session_token")
		uid, err := sqlite.CheckSessionExist(session.Value)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Failed"}`))
			return
		}
		postData.AuthorId = uid
		sqlite.AddPostToTable(postData)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"Success"}`))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"Failed"}`))
	}
}
