package service

import "testing"

func TestUserExists(t *testing.T) {
	var name = "Ievgen"
	var email = "ievgen.gavrysh@gmail.com"
	if UserExists(email) {
		t.Errorf("User %s exists, before it was added to storage", name)
	}
}

func TestAddUser(t *testing.T) {
	ClearUsers()
	var name = "Ievgen"
	var email = "ievgen.gavrysh@gmail.com"

	u, e := AddUser(name, email)

	if e != nil {
		t.Errorf("Failed to add user")
	}

	if u.Name != name || u.Email != email {
		t.Errorf("User name or email does not match before and after adding user to storage")
	}
}


func TestAddUserWithEmptyName(t *testing.T) {
	ClearUsers()
	var name = ""
	var email = "ievgen.gavrysh@gmail.com"

	_, e := AddUser(name, email)

	if e == nil {
		t.Errorf("User added by error - name should not be empty for a new user")
	}

	if e["name"] == nil {
		t.Errorf("name error was not set")
	}
}
