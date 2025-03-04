package update

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Checker interface {
	GetLatestVersion() (string, error)
}

type GithubChecker struct {
	Client *http.Client
	APIURL string
}

func NewGithubChecker(apiURL string) *GithubChecker {
	return &GithubChecker{
		Client: http.DefaultClient,
		APIURL: apiURL,
	}
}

func (g *GithubChecker) GetLatestVersion() (string, error) {
	resp, err := g.Client.Get(g.APIURL)
	if err != nil {
		return "", errors.New("não foi possível se conectar ao servidor para buscar a versão. Verifique sua conexão com a internet e tente novamente")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("o servidor retornou um status de erro ao tentar obter a versão mais recente. Tente novamente mais tarde")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("não foi possível ler a resposta do servidor ao tentar obter a versão. Tente novamente mais tarde")
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", errors.New("não foi possível interpretar a resposta do servidor ao tentar obter a versão. Tente novamente mais tarde")
	}

	tag, ok := data["tag_name"].(string)
	if !ok {
		return "", errors.New("não foi possível encontrar a versão mais recente. Tente novamente mais tarde")
	}

	return tag, nil
}
