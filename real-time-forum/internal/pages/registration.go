package pages

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"real-time-forum-base/internal/data"
	"real-time-forum-base/internal/sqlite"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	byt, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		fmt.Println("Could not generate password", err.Error())
	}
	return string(byt)
}

func gatherRegData(form url.Values) data.Registration {

	data := data.Registration{}
	data.Username = form.Get("username")
	data.Password = HashPassword(form.Get("password"))
	data.Age, _ = strconv.Atoi(form.Get("age"))
	data.Email = form.Get("email")
	data.Firstname = form.Get("firstname")
	data.Lastname = form.Get("lastname")
	data.RegistrationDate = time.Now()
	data.Gender = form.Get("gender")
	return data

}

func Registration(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err1 := template.ParseFiles("index.html")
		if err1 != nil {
			fmt.Println("i dont know what i am doing at this point")
		}
		t.ExecuteTemplate(w, "index.html", nil)
	}

	err := CheckCookieExist(w, r)
	fmt.Println("cookie?")
	if err == nil {
		fmt.Println("get here")
		return
	}

	if r.Method == "POST" {
		r.ParseMultipartForm(100000)
		fmt.Println(r.Form)
		regData := gatherRegData(r.Form)
		_, _, userExists := sqlite.CheckDataExistence(regData.Username, "username")
		_, _, userEmail := sqlite.CheckDataExistence(regData.Email, "email")
		if !userExists && !userEmail {
			fmt.Println("woho")
			mu := CreateSessionToken(w)
			fmt.Println(mu, "mu")
			regData.Cokkie = mu
			sqlite.AddUserToTable(regData)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Success"}`))
		} else {
			if err != nil {
				fmt.Println("fuck", err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"You suck"}`))

		}
	}

}
