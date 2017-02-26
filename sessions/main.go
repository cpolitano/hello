package main

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID - user
var dbSessions = map[string]string{} // session ID - user ID

func getUser(res http.ResponseWriter, req *http.Request) user {
	cookie, err := req.Cookie("session")
	if err != nil {
		sessionId := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sessionId.String(),
		}

	}
	http.SetCookie(res, cookie)

	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/profile", profile)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func profile(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if u.Role != "admin" {
		http.Error(w, "You must be admin to view this page", http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(res, "profile.gohtml", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var u user

	// process form submission
	if req.Method == http.MethodPost {

		// get form values
		username := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		// username taken?
		if _, ok := dbUsers[username]; ok {
			http.Error(res, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sessionId := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sessionId.String(),
		}
		http.SetCookie(res, cookie)
		dbSessions[cookie.Value] = username

		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = user{username, bs, f, l, r}
		dbUsers[username] = u

		// redirect
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "signup.gohtml", u)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		username := req.FormValue("username")
		p := req.FormValue("password")

		// is there a username?
		user, ok := dbUsers[username]
		if !ok {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(p))
		if err != nil {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sessionId := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sessionId.String(),
		}
		http.SetCookie(res, cookie)
		dbSessions[cookie.Value] = username
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	cookie, _ := req.Cookie("session")
	delete(dbSessions, cookie.Value)
	cookie := &http.Cookie{
		Name:  "session",
		Value: "",
		MaxAge: -1
	}
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", http.StatusSeeOther)
	return
}
