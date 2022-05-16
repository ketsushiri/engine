package engine

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type authData struct {
	login, pass string
}

func parseAuthForm(form url.Values) (*authData, error) {
	login, pass := form.Get("login"), form.Get("pass")
	if login == "" || pass == "" {
		return nil, fmt.Errorf("Auth invalid.")
	}
	data := authData{login, pass}
	return &data, nil
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("/auth", r.Form)
	if r.Form.Has("login") {
		data, err := parseAuthForm(r.Form)
		if err != nil {
			fmt.Fprintf(w, "%q", err)
			return
		}
		// set up cookies here
		fmt.Fprintf(w, "%s %q", authOk, data)
		return
	}
	if r.Header.Get("Set-Cookie") == "" {
		fmt.Fprintf(w, "%s", auth)
		return
	}
	fmt.Fprintf(w, "%s", authAlready)
}

func registerUser(data *authData) bool {
	if users.Has(data.login) {
		return false
	}
	users[data.login] = User{
		Login:      data.login,
		Password:   data.pass,
		AccessType: USER,
		UID:        0,
	}
	return true
}

func parseRegisterForm(form url.Values) bool {
	login, pass, key := form.Get("login"), form.Get("pass"), form.Get("key")
	pred := func(s string) bool {
		return s == ""
	}
	if pred(login) || pred(pass) || key != KEY {
		return false
	}
	return registerUser(&authData{login, pass})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("/register", r.Form)
	if r.Form.Has("key") {
		registered := parseRegisterForm(r.Form)
		if !registered {
			log.Println("/register: registration failed.")
			fmt.Fprintf(w, "%s", registerFailed)
		} else {
			log.Println("/register: registration done.")
			fmt.Fprintf(w, "%s", registerOk)
		}
		return
	}
	fmt.Fprintf(w, "%s", register)
}
