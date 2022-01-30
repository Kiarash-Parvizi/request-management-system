package main

import (
	"fmt"
	"net/http"

	accessmanager "github.com/Kiarash-Parvizi/request-management-system/src/AccessManager"
	am "github.com/Kiarash-Parvizi/request-management-system/src/AccessManager"
	db "github.com/Kiarash-Parvizi/request-management-system/src/DB"
	uuid "github.com/nu7hatch/gouuid"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	sesKey = []byte("skd*&^$7shfkh9787xcv12&^%")
	store  = sessions.NewCookieStore(sesKey)
)

type Server struct {
	router *mux.Router
}

func (s *Server) home(rw http.ResponseWriter, r *http.Request) {
	user := &am.User{AccessLevel: 255}
	//
	fmt.Println("HOME")
	session, err := store.Get(r, "uuid")
	fmt.Println("1", err)
	if err == nil {
		fmt.Println("2")
		if uuid, ok := session.Values["uuid"].(string); ok {
			user.AccessLevel = accessmanager.GetUser_accessLevel(uuid)
			fmt.Println("user.AccessLevel:", uuid, user.AccessLevel)
		}
	}
	user.Render("home", rw)
}

func (s *Server) login(rw http.ResponseWriter, r *http.Request) {
	(*am.User)(nil).Render("login", rw)
}
func (s *Server) loginData(rw http.ResponseWriter, r *http.Request) {
	accessLevel, pin := db.Select_user_accesslevel_pin(r.FormValue("name"),
		r.FormValue("pass"))
	uuid, err := uuid.NewV4()
	if err != nil {
		rw.Write([]byte("Error"))
	}
	user := &am.User{
		Uuid:        uuid.String(),
		AccessLevel: accessLevel,
	}
	// store session
	session, err := store.Get(r, "uuid")
	if err == nil {
		session.Values["uuid"] = user.Uuid
		session.Save(r, rw)
		fmt.Println("session was set")
	}
	accessmanager.SetUuid_accessLevel(user.Uuid, accessLevel)
	accessmanager.SetUuid_pin(user.Uuid, pin)
	//
	fmt.Println(user)
	http.Redirect(rw, r, "/", http.StatusSeeOther)
}

func (s *Server) signup(rw http.ResponseWriter, r *http.Request) {
	(*am.User)(nil).Render("signup", rw)
}
func (s *Server) signupData(rw http.ResponseWriter, r *http.Request) {
	err := db.Insert_user(r.FormValue("FN"), r.FormValue("PN"),
		r.FormValue("PIN"), []byte(r.FormValue("UP")), 0, r.FormValue("Pass"))
	if err != nil {
		rw.Write([]byte("Error"))
	}
	http.Redirect(rw, r, "/login", http.StatusSeeOther)
}

//-----------

func (s *Server) requestLoan(rw http.ResponseWriter, r *http.Request) {
	user := &am.User{AccessLevel: 255}
	//
	session, err := store.Get(r, "uuid")
	fmt.Println("1", err)
	if err == nil {
		fmt.Println("2")
		if uuid, ok := session.Values["uuid"].(string); ok {
			user.AccessLevel = accessmanager.GetUser_accessLevel(uuid)
			fmt.Println("user.AccessLevel:", uuid, user.AccessLevel)
		}
	}
	if user.AccessLevel != 0 {
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
		return
	}
	user.Render("requestLoan", rw)
}

func (s *Server) verifyLoan(rw http.ResponseWriter, r *http.Request) {
	user := &am.User{AccessLevel: 255}
	//
	session, err := store.Get(r, "uuid")
	fmt.Println("1", err)
	if err == nil {
		fmt.Println("2")
		if uuid, ok := session.Values["uuid"].(string); ok {
			user.AccessLevel = accessmanager.GetUser_accessLevel(uuid)
			fmt.Println("user.AccessLevel:", uuid, user.AccessLevel)
		}
	}
	if user.AccessLevel != 1 {
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
		return
	}
	user.Render("verifyLoan", rw)
}

func (s *Server) profile(rw http.ResponseWriter, r *http.Request) {
	user := &am.User{AccessLevel: 255}
	//
	session, err := store.Get(r, "uuid")
	fmt.Println("1", err)
	if err == nil {
		fmt.Println("2")
		if uuid, ok := session.Values["uuid"].(string); ok {
			user.AccessLevel = accessmanager.GetUser_accessLevel(uuid)
			user.Uuid = uuid
			fmt.Println("user.AccessLevel:", uuid, user.AccessLevel)
		}
	}
	if user.AccessLevel != 0 {
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
		return
	}
	// get profile from db
	user, err = db.Select_userProfile(am.GetUuid_pin(user.Uuid))
	if err != nil {
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
		return
	}
	//
	user.Render("profile", rw)
}
