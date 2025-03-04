package code

import (
	"fmt"

	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/totp"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

func GenerateTOTP(name, password string) error {
	if err := utils.ValidateServiceName(name); err != nil {
		return err
	}
	if err := utils.ValidatePassword(password); err != nil {
		return err
	}

	accounts, err := storage.LoadEncrypted(password)
	if err != nil {
		return err
	}

	account, err := accounts.GetAccount(name)
	if err != nil {
		return err
	}

	code, secondsRemaining, err := totp.GenerateTOTP(account.Secret)
	if err != nil {
		return err
	}

	fmt.Printf("Código TOTP para a conta [%s]: %s\n", account.Name, code)
	fmt.Printf("Este código é válido por %d %s.\n", secondsRemaining, utils.Pluralize("segundo", "segundos", secondsRemaining))
	return nil
}

var generateName string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Gere um código TOTP",
	Long:  "Gere um código TOTP válido para autenticação em dois fatores (2FA).",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := utils.PromptPassword()
		utils.HandleError(err)

		utils.HandleError(GenerateTOTP(generateName, password))
	},
}

func init() {
	CodeCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&generateName, "name", "n", "", "Nome da conta (ex: gitlab)")
}
