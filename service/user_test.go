package service

import "testing"

func TestUserExists(t *testing.T) {
	var name = "Tiffany"
	var email = "tiffany.gavrysh@gmail.com"
	if UserExists(email) {
		t.Errorf("User %s exists, before it was added to storage", name)
	}
}

func TestAddUser(t *testing.T) {
	ClearUsers()
	var name = "Tiffany"
	var email = "tiffany.gavrysh@gmail.com"

	u, e := AddUser(name, email)
	_, errorNamePresent := e["name"];
	_, errorEmailPresent := e["email"];
	if errorNamePresent || errorEmailPresent  {
		t.Errorf("Should add valid user without errors reporting")
	}

	if u.Name != name || u.Email != email {
		t.Errorf("User name or email does not match before and after adding user to storage")
	}
}

func TestAddUserWithEmptyName(t *testing.T) {
	ClearUsers()
	var name = ""
	var email = "tiffany.gavrysh@gmail.com"

	u, e := AddUser(name, email)

	if e == nil {
		t.Errorf("User added by error - name should not be empty for a new user")
	}

	if _, errorPresent := e["name"]; !errorPresent  {
		t.Errorf("name category is absent in errors map")
	}

	if u.Name != "" {
		t.Error("User name should be empty when name has incorrect value")
	}
}

func TestAddUserWithUnsupportedSymbolsName(t *testing.T) {
	ClearUsers()
	var name = "`1qa`()*"
	var email = "tiffany.gavrysh@gmail.com"

	u, e := AddUser(name, email)

	if e == nil {
		t.Errorf("User added by error - unsupported symbols in name")
	}

	if _, errorPresent := e["name"]; !errorPresent  {
		t.Errorf("name category is absent in errors map")
	}

	if u.Name != "" {
		t.Error("User name should be empty when name has incorrect value")
	}
}

func TestAddUserWithSupportedSymbolsNameName(t *testing.T) {
	ClearUsers()
	var name = "ABCa bc . "
	var email = "tiffany.gavrysh@gmail.com"

	u, e := AddUser(name, email)

	if e != nil {
		t.Errorf("User should be added, supported symbols in name")
	}

	if _, errorPresent := e["email"]; errorPresent  {
		t.Errorf("name category should be empty for supported symbols in name")
	}

	if u.Name != name {
		t.Error("User name should be equal to passed name")
	}
}

func TestAddUserWithSupportedSymbolsEmail(t *testing.T) {
	ClearUsers()
	var name = "tiffany"
	var email = "tiffany.gavrysh@gmail.com"

	u, e := AddUser(name, email)

	if e != nil {
		t.Errorf("User should be added, supported symbols in email")
	}

	if _, ok := e["email"]; ok {
		t.Errorf("email category should be empty for supported symbols in email")
	}

	if u.Email != email {
		t.Error("User email should be equal to passed email")
	}
}

func TestAddUserWithUnsupportedSymbolsEmail(t *testing.T) {
	ClearUsers()
	var name = "tiffany"
	var email = "tiffany.gavrysh"

	u, e := AddUser(name, email)

	if e == nil {
		t.Errorf("User should not be added, supported symbols in email")
	}

	if _, ok := e["email"]; !ok {
		t.Errorf("email category should be empty for supported symbols in email")
	}

	if u.Email != "" {
		t.Error("User email should be equal to passed email")
	}
}

func TestAddTwoUsers(t *testing.T) {
	ClearUsers()
	AddUser("Tiffany", "tiffany.gavrysh@gmail.com")
	AddUser("Samantha", "samantha.gavrysh@gmail.com")

	if UsersLen() != 2 {
		t.Errorf("Users length should be equal to two after 2 valid Users were added")
	}
}

func TestClearUsers(t *testing.T) {
	ClearUsers()
	if UsersLen() != 0 {
		t.Errorf("Users length should be 0 in the beginning of the text")
	}
	AddUser("Tiffany", "tiffany.gavrysh@gmail.com")
	AddUser("Samantha", "samantha.gavrysh@gmail.com")

	if UsersLen() != 2 {
		t.Errorf("Users length should be equal to 2 after 2 valid Users were added")
	}

	ClearUsers()
	if UsersLen() != 0 {
		t.Errorf("Users length should be 0 after all users are removed")
	}
}
