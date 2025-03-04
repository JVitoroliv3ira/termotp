package account

import (
	"fmt"
	"time"

	"github.com/JVitoroliv3ira/termotp/internal/models"
	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

func AddAccount(name, secret, password string) error {
	if err := utils.ValidateServiceName(name); err != nil {
		return err
	}
	if err := utils.ValidateServiceSecret(secret); err != nil {
		return err
	}
	if err := utils.ValidatePassword(password); err != nil {
		return err
	}

	account := models.Account{
		Name:      name,
		Secret:    secret,
		CreatedAt: time.Now(),
	}

	accounts, err := storage.LoadEncrypted(password)
	if err != nil {
		return err
	}

	if err := accounts.AddAccount(account); err != nil {
		return err
	}

	if err := storage.SaveEncrypted(accounts, password); err != nil {
		return err
	}

	fmt.Printf("\nConta '%s' cadastrada e armazenada com segurança!\n", account.Name)
	return nil
}

var addName string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adicione uma nova conta TOTP",
	Long:  "Cadastre uma nova conta TOTP e armazene sua chave com segurança.",
	Run: func(cmd *cobra.Command, args []string) {
		secret, err := utils.PromptSecret()
		utils.HandleError(err)

		password, err := utils.PromptPassword()
		utils.HandleError(err)

		utils.HandleError(AddAccount(addName, secret, password))
	},
}

func init() {
	AccountCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addName, "name", "n", "", "Nome da conta a ser cadastrada (ex: gitlab)")
}
