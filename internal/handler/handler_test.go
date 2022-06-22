package handler

import (
	"testing"
)

func TestConfig(t *testing.T) {
	// responseJson()
	GetUser := GetUser()
	if GetUser == nil {
		t.Errorf("error getUser func")
	}

	AddUser := AddUser()
	if AddUser == nil {
		t.Errorf("error AddUser func")
	}

	ShowUser := ShowUser()
	if ShowUser == nil {
		t.Errorf("error ShowUser func")
	}

	UpdateUser := UpdateUser()
	if UpdateUser == nil {
		t.Errorf("error UpdateUser func")
	}

	DeleteUser := DeleteUser()
	if DeleteUser == nil {
		t.Errorf("error DeleteUser func")
	}
}
