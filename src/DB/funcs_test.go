package db

import (
	"fmt"
	"testing"
)

func Test_Insert_user(t *testing.T) {
	fmt.Println("Test_Insert_user:")
	err := Insert_user("Fn", "PN", "PIN123", []byte("bytesMine:)"), 10, "maPass")
	if err != nil {
		fmt.Println("-<=>-> ERR: ", err)
		t.Errorf("failed: Insert")
	}
}

func Test_Insert_tmps(t *testing.T) {
	fmt.Println("Test_Insert_user:")
	err := Insert_tmp("Fn", "PN", 102)
	if err != nil {
		fmt.Println("-<=>-> ERR: ", err)
		t.Errorf("failed: Insert")
	}
}

func Test_Select_user_accesslevel(t *testing.T) {
	aLevel, pin := Select_user_accesslevel_pin("Fn", "maPass")
	fmt.Println(aLevel, pin)
}
