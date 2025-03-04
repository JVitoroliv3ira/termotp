package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualize o TermOTP",
	Long:  "Verifique e instale a versão mais recente do TermOTP disponível.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runUpdate(); err != nil {
			fmt.Printf("Erro ao atualizar: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func runUpdate() error {
	switch runtime.GOOS {
	case "windows":
		return runPowerShellScript("scripts/update_windows.ps1")
	default:
		return runShellScript("scripts/update_linux.sh")
	}
}

func runShellScript(scriptPath string) error {
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runPowerShellScript(scriptPath string) error {
	cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-File", scriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
