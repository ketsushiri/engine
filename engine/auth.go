package engine

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const auth = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Auth</title>
</head>
<body>
    <center><h2>Авторизация</h2></center>
    <center>
        <form>
            <label for="login">Логин</label><br>
            <input type="text" id="login" name="login"><br>
            <label for="pass">Пароль</label><br>
            <input type="text" id="pass" name="pass"><br><br>
            <input type="submit" value="Войти">
        </form>
    </center>
</body>
</html>
`

const authAlready = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Авторизация</title>
</head>
<body>
    <center><h2>Вы уже авторизованы</h2>
    <input type="submit" value="Выйти">
	</center>
</body>
</html>
`

const authFailed = `failed`

const authOk = `ok`

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
	log.Println(r.Form)
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
