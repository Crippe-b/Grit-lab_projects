package pages

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"real-time-forum-base/internal/sqlite"
)

func SendData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err1 := template.ParseFiles("index.html")
		if err1 != nil {
			fmt.Println("i dont know what i am doing at this point")
		}
		t.ExecuteTemplate(w, "index.html", nil)
	}

	if r.URL.Path != "/hello" {
		return
	}
	if r.Method == "POST" {

		err := CheckCookieExist(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Failed"}`))
			return
		}

		session, err := r.Cookie("session_token")
		if err != nil {
			return
		}
		_, err = sqlite.CheckSessionExist(session.Value)
		if err != nil {
			return
		}
		var pageDataJson []byte
		pageDataJson, err1 := json.Marshal(sqlite.MainPageDataGathering())
		if err1 != nil {
			fmt.Println("fuckfuckfuckfuckfuckfuckfuckfuck", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(pageDataJson)
	}
}

func Pages(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Println("should get 404")
	}
	DeleteCookie(w)
	fmt.Println("deleted cookie")
	t, err1 := template.ParseFiles("index.html")
	if err1 != nil {
		fmt.Println("i dont know what i am doing at this point")
	}
	t.ExecuteTemplate(w, "index.html", nil)
}
