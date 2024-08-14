package pages

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"real-time-forum-base/internal/chatsocket"
	"real-time-forum-base/internal/sqlite"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateSessionToken(w http.ResponseWriter) string {

	sessionToken := uuid.Must(uuid.NewV4()).String()
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Path:    "/",
		Expires: time.Now().Add(1000000 * time.Second),
	})

	return sessionToken
}

func DeleteCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "lole",
		Path:    "/",
		Expires: time.Now().Add(-1000 * time.Second),
	})
}

func CheckCookieExist(w http.ResponseWriter, r *http.Request) error {
	session, err := r.Cookie("session_token")
	if err != nil {
		fmt.Println(err, "error in checkcookieexsist")
		return err
	}
	if _, err = sqlite.CheckSessionExist(session.Value); err != nil {
		DeleteCookie(w)
		fmt.Println("fucked up cookie yo")
		return errors.New("invalid cookie")
	}
	return nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err1 := template.ParseFiles("index.html")
		if err1 != nil {
			fmt.Println("i dont know what i am doing at this point")
		}
		t.ExecuteTemplate(w, "index.html", nil)
	}

	if r.Method == "POST" {
		bu, _ := r.Cookie("session_token")
		if bu != nil {

			uid, _ := sqlite.CheckSessionExist(bu.Value)
			token := CreateSessionToken(w)
			sqlite.UpdateSessionToken(token, uid)
			fmt.Println("login successfull, yaaaaay")

			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Success"}`))
			return
		}

		r.ParseMultipartForm(100000)
		uid, password, dataExists := sqlite.CheckDataExistence(r.Form.Get("name"), "email")
		if !dataExists {
			uid, password, dataExists = sqlite.CheckDataExistence(r.Form.Get("name"), "username")
			fmt.Println("did first check", dataExists)

		}
		if !dataExists {
			fmt.Println("did second check", dataExists)
			return
		} else if dataExists {
			if password == (r.Form.Get("password")) || bcrypt.CompareHashAndPassword([]byte(password), []byte(r.Form.Get("password"))) == nil {
				token := CreateSessionToken(w)
				sqlite.UpdateSessionToken(token, uid)
				fmt.Println("login successfull, yaaaaay")

				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"status":"Success"}`))
				chatsocket.WsEndpoint(w, r)
			}
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Failed"}`))

		}
	}
}
