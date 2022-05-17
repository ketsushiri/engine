package engine

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type authData struct {
	login, pass string
}

func makeCookie(data *authData, duration time.Duration) *http.Cookie {
	return &http.Cookie{
		Name:    "userkey",
		Value:   fmt.Sprintf("%s %s", data.login, data.pass),
		Expires: time.Now().Add(duration),
	}
}

func validateCookie(cookie *http.Cookie) bool {
	values := strings.Split(cookie.Value, " ")
	if len(values) < 2 {
		return false
	}
	login, pass := values[0], strings.Join(values[1:], " ")
	if users.Has(login) && users[login].Password == pass {
		return true
	}
	return false
}

func parseAuthForm(form url.Values) (*authData, error) {
	login, pass := form.Get("login"), form.Get("pass")
	if login == "" || pass == "" || !users.Has(login) {
		return nil, fmt.Errorf("Auth invalid.")
	}
	if users[login].Password != pass {
		return nil, fmt.Errorf("Invalid password")
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
			log.Println("/auth", err)
			fmt.Fprintf(w, "%s", authFailed)
			return
		}
		http.SetCookie(w, makeCookie(data, time.Minute*30))
		fmt.Fprintf(w, "%s %q", authOk, data)
		return
	}
	if data, err := r.Cookie("userkey"); err != nil || !validateCookie(data) {
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
	if login == "" || pass == "" || key != KEY {
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
