package db

import (
	"fmt"

	am "github.com/Kiarash-Parvizi/request-management-system/src/AccessManager"
)

//insert, err := db.Query("INSERT INTO test VALUES ( 2, 'Kiarash msq' )")
//if err != nil {
//	panic(err.Error())
//}
//defer insert.Close()

func Insert_user(FN, PN, PIN string, UP []byte, C int, pass string) error {
	str := `Insert into Users(
	FullName,
	PhoneNumber,
	PIN,
	UserPhoto,
	Credits, Pass, AccessLevel)
	VALUES(?,?,?,?,?,?,?);`
	q, err := db.Prepare(str)
	if err != nil {
		fmt.Println("1")
		return err
	}
	fmt.Println("2")
	//
	_, err = q.Exec(FN, PN, PIN, UP, C, pass, byte(0))
	return err
}

func Insert_tmp(name, pn string, age int) error {
	str := `Insert into Tmps(
	Name,
	PN,
	Age)
	VALUES(?,?,?);`
	q, err := db.Prepare(str)
	if err != nil {
		fmt.Println("1")
		return err
	}
	fmt.Println("2")
	//
	_, err = q.Exec(name, pn, age)
	return err
}

// returns User(Pid AccessLevel)
func Select_user_accesslevel_pin(fullName, pass string) (byte, string) {
	var AccessLevel byte
	var PIN string
	err := db.QueryRow("select AccessLevel, PIN from Users where FullName=? and Pass=?;",
		fullName, pass).Scan(&AccessLevel, &PIN)
	if err != nil {
		return 255, ""
	}
	return AccessLevel, PIN
}

// returns User(FullName,PhoneNumber,PIN,Credits)
func Select_userProfile(PIN string) (*am.User, error) {
	res := &am.User{}
	err := db.QueryRow("select FullName,PhoneNumber,PIN,Credits from Users where PIN=?;",
		PIN).Scan(&res.FullName, &res.PhoneNumber, &res.PIN, &res.Credits)
	return res, err
}
