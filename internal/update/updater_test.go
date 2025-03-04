package update

import (
	"errors"
	"testing"
)

func TestUpdater_Update_InvalidSpecificVersion(t *testing.T) {
	u := NewUpdater()
	installed, err := u.Update("9999.9999.9999-nonexistent")
	if err == nil {
		t.Error("Expected an error but got nil")
	} else {
		expected := errors.New("não foi possível instalar a versão solicitada do TermOTP. Verifique se o número da versão está correto ou tente novamente mais tarde")
		if err.Error() != expected.Error() {
			t.Errorf("Expected: '%v' but got: '%v'", expected, err)
		}
	}

	if installed {
		t.Error("Should not have installed anything due to error")
	}
}
