package main

import (
	"fmt"
	"net/http"
	"strconv"

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

func getSessionUuid(r *http.Request) (string, error) {
	session, err := store.Get(r, "uuid")
	if err == nil {
		if uuid, ok := session.Values["uuid"].(string); ok {
			return uuid, nil
		}
	}
	return "", err
}

func (s *Server) home(rw http.ResponseWriter, r *http.Request) {
	user := &am.User{AccessLevel: 255}
	//
	var err error
	user.Uuid, err = getSessionUuid(r)
	if err == nil {
		user.AccessLevel = accessmanager.GetUser_accessLevel(user.Uuid)
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
		return
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
		return
	}
	http.Redirect(rw, r, "/login", http.StatusSeeOther)
}

//-----------

func (s *Server) requestLoan(rw http.ResponseWriter, r *http.Request) {
	user := &am.User{AccessLevel: 255}
	//
	var err error
	user.Uuid, err = getSessionUuid(r)
	if err == nil {
		user.AccessLevel = accessmanager.GetUser_accessLevel(user.Uuid)
	}
	if user.AccessLevel != 0 {
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
		return
	}
	user.Render("requestLoan", rw)
}
func (s *Server) requestLoanData(rw http.ResponseWriter, r *http.Request) {
	user := &am.User{AccessLevel: 255}
	//
	var err error
	user.Uuid, err = getSessionUuid(r)
	if err == nil {
		user.AccessLevel = accessmanager.GetUser_accessLevel(user.Uuid)
	}
	if user.AccessLevel != 0 {
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
		return
	}
	fmt.Println()
	objId, err := strconv.Atoi(r.FormValue("ObjectId"))
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	fmt.Println(": ", am.GetUuid_pin(user.Uuid))
	err = db.Insert_loanRequest(am.GetUuid_pin(user.Uuid),
		r.FormValue("loanType"),
		r.FormValue("loanAmount"),
		r.FormValue("AdditionalNotes"),
		objId)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	http.Redirect(rw, r, "/success", http.StatusSeeOther)
}

func (s *Server) reviewLoan(rw http.ResponseWriter, r *http.Request) {
	user := &am.User{AccessLevel: 255}
	//
	var err error
	user.Uuid, err = getSessionUuid(r)
	if err == nil {
		user.AccessLevel = accessmanager.GetUser_accessLevel(user.Uuid)
	}
	if user.AccessLevel != 1 {
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
		return
	}
	// special render
	nxt := db.Select_Next_UserLoanRequests()
	if nxt == nil {
		http.Redirect(rw, r, "/empty-list", http.StatusSeeOther)
		return
	}
	am.RenderSpecial("reviewLoan", rw, nxt)
	//---------------
}

func (s *Server) reviewLoanData(rw http.ResponseWriter, r *http.Request) {
	user := &am.User{AccessLevel: 255}
	//
	var err error
	user.Uuid, err = getSessionUuid(r)
	if err == nil {
		user.AccessLevel = accessmanager.GetUser_accessLevel(user.Uuid)
	}
	if user.AccessLevel != 1 {
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
		return
	}
	// set result in db
	fmt.Println("-=>", r.FormValue("RID"), r.FormValue("RequestAccepted"))
	err = db.SetResultOfReviewedRequest(r.FormValue("RID"), r.FormValue("RequestAccepted"))
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}
	http.Redirect(rw, r, "/review-loan", http.StatusSeeOther)
}

func (s *Server) profile(rw http.ResponseWriter, r *http.Request) {
	user := &am.User{AccessLevel: 255}
	//
	var err error
	user.Uuid, err = getSessionUuid(r)
	if err == nil {
		user.AccessLevel = accessmanager.GetUser_accessLevel(user.Uuid)
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

func (s *Server) success(rw http.ResponseWriter, r *http.Request) {
	//
	(*am.User)(nil).Render("success", rw)
}
func (s *Server) emptyList(rw http.ResponseWriter, r *http.Request) {
	//
	(*am.User)(nil).Render("emptyList", rw)
}
