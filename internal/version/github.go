package version

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Release struct {
	TagName     string    `json:"tag_name"`
	Prerelease  bool      `json:"prerelease"`
	PublishedAt time.Time `json:"published_at"`
}

type VersionClient struct {
	Url    string
	Client *http.Client
}

func NewVersionClient() *VersionClient {
	return &VersionClient{
		Url:    "https://api.github.com/repos/JVitoroliv3ira/termotp/releases",
		Client: http.DefaultClient,
	}
}

func (vc *VersionClient) FetchReleases() ([]Release, error) {
	res, err := vc.Client.Get(vc.Url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("ocorreu um erro ao buscar as versões disponíveis")
	}

	var releases []Release
	if err := json.NewDecoder(res.Body).Decode(&releases); err != nil {
		return nil, err
	}

	return releases, nil
}

func (vc *VersionClient) ListFormattedReleases() ([]string, error) {
	releases, err := vc.FetchReleases()
	if err != nil {
		return nil, err
	}

	var formatted []string
	for _, release := range releases {
		date := release.PublishedAt.Format("2006-01-02")
		label := ""
		if release.Prerelease {
			label = "(pré-release)"
		}
		formatted = append(formatted, fmt.Sprintf("- %s %s (%s)", release.TagName, label, date))
	}

	return formatted, nil
}
