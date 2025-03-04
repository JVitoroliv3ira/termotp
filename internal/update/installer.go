package update

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

type Installer interface {
	InstallVersion(version string) error
}

type GithubInstaller struct {
	LinuxURL   string
	WindowsURL string
}

func NewGithubInstaller() *GithubInstaller {
	return &GithubInstaller{
		LinuxURL:   "https://github.com/JVitoroliv3ira/termotp/releases/download/%s/totp-linux-amd64",
		WindowsURL: "https://github.com/JVitoroliv3ira/termotp/releases/download/%s/totp-windows-amd64.exe",
	}
}

func (g *GithubInstaller) InstallVersion(version string) error {
	osType := runtime.GOOS
	var err error

	switch osType {
	case "linux":
		err = g.installLinux(version)
	case "windows":
		err = g.installWindows(version)
	default:
		return errors.New("sistema operacional não suportado")
	}

	if err != nil {
		return errors.New("ocorreu um erro ao instalar o TermOTP")
	}

	return nil
}

func (g *GithubInstaller) installLinux(version string) error {
	url := fmt.Sprintf(g.LinuxURL, version)

	if err := runCmd(exec.Command("sudo", "rm", "-f", "/usr/local/bin/totp")); err != nil {
		return errors.New("não foi possível remover a versão antiga do TermOTP")
	}

	if err := runCmd(exec.Command("wget", url, "-O", "totp")); err != nil {
		return errors.New("não foi possível baixar o TermOTP. Verifique sua conexão e tente novamente")
	}

	if err := runCmd(exec.Command("chmod", "+x", "totp")); err != nil {
		return errors.New("não foi possível definir permissões de execução para o TermOTP")
	}

	if err := runCmd(exec.Command("sudo", "mv", "totp", "/usr/local/bin/")); err != nil {
		return errors.New("não foi possível mover o binário do TermOTP para /usr/local/bin")
	}

	return nil
}

func (g *GithubInstaller) installWindows(version string) error {
	url := fmt.Sprintf(g.WindowsURL, version)

	rmScript := `Remove-Item "C:\Program Files\TermOTP\totp.exe" -ErrorAction SilentlyContinue`
	if err := runPSScript(rmScript); err != nil {
		return errors.New("não foi possível remover a versão antiga do TermOTP")
	}

	mkScript := `mkdir "C:\Program Files\TermOTP" -ErrorAction SilentlyContinue`
	if err := runPSScript(mkScript); err != nil {
		return errors.New("não foi possível criar o diretório de instalação para o TermOTP")
	}

	dlScript := fmt.Sprintf(`Invoke-WebRequest -Uri "%s" -OutFile "C:\Program Files\TermOTP\totp.exe"`, url)
	if err := runPSScript(dlScript); err != nil {
		return errors.New("não foi possível baixar o TermOTP. Verifique sua conexão e tente novamente")
	}

	pathScript := `[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";C:\Program Files\TermOTP", [System.EnvironmentVariableTarget]::Machine)`
	if err := runPSScript(pathScript); err != nil {
		return errors.New("não foi possível adicionar o TermOTP ao PATH do sistema")
	}

	fmt.Println("Agora reinicie o terminal e rode totp de qualquer lugar!")
	return nil
}

func runCmd(cmd *exec.Cmd) error {
	return cmd.Run()
}

func runPSScript(script string) error {
	psCmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", script)
	return psCmd.Run()
}
