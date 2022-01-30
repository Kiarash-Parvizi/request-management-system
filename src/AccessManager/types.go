package accessmanager

type User struct {
	Uuid        string
	AccessLevel byte // 1 == admin
	//
	FullName    string
	PhoneNumber string
	PIN         string
	Credits     int
}
