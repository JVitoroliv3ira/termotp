package update

import (
	"errors"

	"github.com/JVitoroliv3ira/termotp/internal/version"
)

type Updater struct {
	Checker   Checker
	Installer Installer
}

func NewUpdater() *Updater {
	return &Updater{
		Checker:   NewGithubChecker("https://api.github.com/repos/JVitoroliv3ira/termotp/releases/latest"),
		Installer: NewGithubInstaller(),
	}
}

func (u *Updater) Update(targetVersion string) (bool, error) {
	if targetVersion == "" {
		latest, err := u.Checker.GetLatestVersion()
		if err != nil {
			return false, errors.New("não foi possível consultar a versão mais recente do TermOTP. Tente novamente mais tarde")
		}

		current := version.GetVersion()
		if current == latest {
			return false, nil
		}

		if err := u.Installer.InstallVersion("latest"); err != nil {
			return false, errors.New("não foi possível instalar a versão mais recente do TermOTP")
		}

		return true, nil
	}

	if targetVersion[0] != 'v' {
		targetVersion = "v" + targetVersion
	}

	if err := u.Installer.InstallVersion(targetVersion); err != nil {
		return false, errors.New("não foi possível instalar a versão solicitada do TermOTP. Verifique se o número da versão está correto ou tente novamente mais tarde")
	}

	return true, nil
}
