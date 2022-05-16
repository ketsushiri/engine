package engine

const auth = `<!DOCTYPE html>
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
</html>`

const authAlready = `<!DOCTYPE html>
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
</html>`

const authFailed = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
    <center><h2>403, ошибка авторизации.</h2></center>
</body>
</html>`

const authOk = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
    <center><h2>200, успешная авторизация.</h2></center>
</body>
</html>`

const register = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
    <center>
        <h2>Регистрация</h2>
        <form>
            <label for="key">Ключ приглашения</label><br>
            <input type="text" id="key" name="key"><br>
            <label for="login">Логин</label><br>
            <input type="text" id="login" name="login"><br>
            <label for="pass">Пароль</label><br>
            <input type="text" id="pass" name="pass"><br><br>
            <input type="submit" value="Зарегистрироваться">
        </form
    </center>
</body>
</html>`

const registerOk = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
    <center>
        <h2>200, успешная регистрация.</h2>
    </center>
</body>
</html>`

const registerFailed = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
    <center>
        <h2>
            403, регистрация не удалась.
        </h2>
    </center>
</body>
</html>`
