package code

import (
	"fmt"

	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/totp"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

func CopyTOTP(name, password string) error {
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

	if err := utils.CopyToClipboard(code); err != nil {
		return err
	}

	fmt.Printf("Código TOTP para a conta [%s]: %s\n", account.Name, code)
	fmt.Printf("Este código foi copiado com sucesso e é válido por %d %s.\n", secondsRemaining, utils.Pluralize("segundo", "segundos", secondsRemaining))
	return nil
}

var copyName string

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Gere e copie um código TOTP",
	Long:  "Gere um código TOTP, visualize no terminal e tenha ele copiado automaticamente para a área de transferência.",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := utils.PromptPassword()
		utils.HandleError(err)

		utils.HandleError(CopyTOTP(copyName, password))
	},
}

func init() {
	CodeCmd.AddCommand(copyCmd)
	copyCmd.Flags().StringVarP(&copyName, "name", "n", "", "Nome da conta (ex: gitlab)")
}
