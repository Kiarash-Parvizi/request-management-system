package accessmanager

import (
	"fmt"
	"html/template"
	"net/http"
)

func init() {
}

var (
	tmplMp = map[string]string{
		"home":        "frontend/home.html",
		"login":       "frontend/login.html",
		"signup":      "frontend/signup.html",
		"requestLoan": "frontend/request-loan.html",
		"verifyLoan":  "frontend/verify-loan.html",
		"profile":     "frontend/profile.html",
	}
)

func (u *User) Render(page string, rw http.ResponseWriter) {
	fmt.Println("render", page)
	parsedTemplate, err := template.ParseFiles(tmplMp[page])
	if err != nil {
		rw.Write([]byte("Error"))
		return
	}
	err = parsedTemplate.Execute(rw, u)
	if err != nil {
		rw.Write([]byte("Error"))
	}
}
