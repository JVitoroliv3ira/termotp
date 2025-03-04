package update

import (
	"errors"
	"testing"
)

func TestGithubChecker_Success(t *testing.T) {

	checker := NewGithubChecker("https://api.github.com/repos/JVitoroliv3ira/termotp/releases/latest")
	got, err := checker.GetLatestVersion()

	if err != nil {
		t.Errorf("Expected nil but got error: '%v'", err)
	}
	if got == "" {
		t.Errorf("Expected a non-empty version string, but got: '%v'", got)
	}
}

func TestGithubChecker_InvalidURL(t *testing.T) {
	checker := NewGithubChecker("http://url-invalida-que-nao-existe-123.com")

	_, err := checker.GetLatestVersion()
	if err == nil {
		t.Error("Expected an error but got nil")
	}

	expected := errors.New("não foi possível se conectar ao servidor para buscar a versão. Verifique sua conexão com a internet e tente novamente")
	if err != nil && err.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, err)
	}
}
