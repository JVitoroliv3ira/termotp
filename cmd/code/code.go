package code

import "github.com/spf13/cobra"

var CodeCmd = &cobra.Command{
	Use:   "code",
	Short: "Gerencie códigos TOTP de autenticação",
	Long:  "Permite gerar, listar e copiar códigos TOTP, facilitando o uso de autenticação em duas etapas (2FA).",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
