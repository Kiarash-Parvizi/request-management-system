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
func Insert_loanRequest(UPIN, LoanType, Amount, AdditionalNotes string, RefObjectId int) error {
	str := `Insert into LoanRequests(
	UPIN,
	LoanType,
	Amount,
	AdditionalNotes,
	RefObjectId)
	VALUES(?,?,?,?,?);`
	q, err := db.Prepare(str)
	if err != nil {
		fmt.Println("1")
		return err
	}
	fmt.Println("2")
	//
	_, err = q.Exec(UPIN, LoanType, Amount, AdditionalNotes, RefObjectId)
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

// returns User joined LoanRequests
func Select_Next_UserLoanRequests() interface{} {
	res := struct {
		FullName, PIN string
		Credits       int
		AccessLevel   byte
		//
		RID                               int
		LoanType, Amount, AdditionalNotes string
		RefObjectId                       int
	}{}
	err := db.QueryRow(`
	select u.FullName, u.PIN, u.Credits, u.AccessLevel, l.RID, l.LoanType, l.Amount, l.AdditionalNotes, l.RefObjectId
	from Users u inner join LoanRequests l on u.PIN=l.UPIN
	where l.Reviewed = False;`).
		Scan(&res.FullName, &res.PIN, &res.Credits, &res.AccessLevel,
			&res.RID, &res.LoanType, &res.Amount, &res.AdditionalNotes, &res.RefObjectId)
	if err != nil {
		fmt.Println("nxtErr: ", err)
		return nil
	}
	return res
}

// set result of the reviewed request
func SetResultOfReviewedRequest(RID, RequestAccepted string) error {
	addTableName := "DeclinedRequests"
	if RequestAccepted == "True" {
		addTableName = "AcceptedRequests"
	}
	// funny stuff can happen here :) TODO
	// I couldn't find a way to exec these 2 queries in one OP
	_, err := db.Exec("Update LoanRequests set Reviewed=True where RID=?;", RID)
	if err != nil {
		return err
	}
	_, err = db.Exec("Insert INTO "+addTableName+"(RID) VALUES(?);", RID)
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
