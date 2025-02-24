package cmd

import (
	"fmt"

	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/totp"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

var generateName string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Gera um código TOTP",
	Long:  "Gera um código TOTP válido para autenticação de dois fatores.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.HandleError(utils.ValidateServiceName(generateName))

		password, err := utils.PromptPassword()
		utils.HandleError(err)
		utils.HandleError(utils.ValidatePassword(password))

		account, err := storage.GetAccount(generateName, password)
		utils.HandleError(err)

		code, err := totp.GenerateTOTP(account.Secret)
		utils.HandleError(err)

		fmt.Printf("Código TOTP para %s: %s\n", account.Name, code)
		fmt.Printf("Este código é válido por 30 segundos.\n")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&generateName, "name", "n", "", "Nome da conta (ex: gitlab)")
}
