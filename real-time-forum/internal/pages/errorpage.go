package pages

import (
	"fmt"
	"html/template"
	"net/http"
)

var errorcount = 0

func ErrorHandler(w http.ResponseWriter, status int) {

	t, err1 := template.ParseFiles("index.html")
	if err1 != nil {
		fmt.Println("i dont know what i am doing at this point")
	}
	t.ExecuteTemplate(w, "index.html", nil)
	errorcount++
	fmt.Println(errorcount)
}
