package db

type User struct {
	FullName    string
	PhoneNumber string
	PIN         string
	UserPhoto   []byte
	//
	Credits     int
	Pass        string
	AccessLevel byte
}
