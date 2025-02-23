package cmd

import "github.com/spf13/cobra"

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Adiciona uma nova conta TOTP",
	Long:  "Cadastra um novo serviço TOTP, armazenando sua chave de forma segura.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().String("account", "", "Nome da conta (ex: gitlab)")
	setupCmd.Flags().String("secret", "", "Chave secreta TOTP")
}
