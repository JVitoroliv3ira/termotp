package cmd

import (
	"fmt"
	"os"

	"github.com/JVitoroliv3ira/termotp/cmd/account"
	"github.com/JVitoroliv3ira/termotp/cmd/code"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "totp",
	Short: "Gerenciador de códigos TOTP para autenticação 2FA",
	Long:  "Ferramenta CLI para gerar e gerenciar códigos TOTP de forma segura.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(account.AccountCmd)
	rootCmd.AddCommand(code.CodeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}
}
