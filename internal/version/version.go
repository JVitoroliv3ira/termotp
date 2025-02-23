package version

import "fmt"

var Version = "dev"

func GetVersion() string {
	return fmt.Sprintf("TermOTP - Versão %s", Version)
}
