package mixins

import "testing"

func TestEncryptPassword(t *testing.T) {
	var password string = "password1234"
	passwordEncrypted, _ := EncryptPassword(password)

	if password == passwordEncrypted {
		t.Errorf("The password was not encrypted: expected: $2a got: %v", passwordEncrypted)
	}
}
