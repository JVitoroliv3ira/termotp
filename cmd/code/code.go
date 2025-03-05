package code

import "github.com/spf13/cobra"

var CodeCmd = &cobra.Command{
	Use:   "code",
	Short: "Gere, copie e liste códigos TOTP",
	Long:  "Gere, copie e liste códigos TOTP para autenticação em dois fatores (2FA).",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
