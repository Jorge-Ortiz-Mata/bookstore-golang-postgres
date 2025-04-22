package mixins

import "testing"

func TestEncryptPassword(t *testing.T) {
	var password string = "password1234"
	passwordEncrypted, _ := EncryptPassword(password)

	if password == passwordEncrypted {
		t.Errorf("The password was not encrypted: expected: $2a got: %v", passwordEncrypted)
	}
}

func TestValidPassordOKResponse(t *testing.T) {
	var password string = "pass1234"
	passwordHash, _ := EncryptPassword(password)

	result := ValidPassword(password, passwordHash)

	if result != true {
		t.Errorf("The password must be valid")
	}
}

func TestValidPassordNOKResponse(t *testing.T) {
	var password string = "pass1234"
	passwordHash, _ := EncryptPassword(password)

	result := ValidPassword("invalidPass", passwordHash)

	if result != false {
		t.Errorf("The password must be invalid")
	}
}
