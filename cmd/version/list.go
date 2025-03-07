package version

import (
	"fmt"

	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/JVitoroliv3ira/termotp/internal/version"
	"github.com/spf13/cobra"
)

func ListVersions() error {
	client := version.NewVersionClient()
	lines, err := client.ListFormattedReleases()
	if err != nil {
		return err
	}

	for _, l := range lines {
		fmt.Println(l)
	}

	return nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste as versões disponíveis do TermOTP",
	Long:  "Exiba todas as versões disponíveis do TermOTP publicadas no GitHub Releases, incluindo pré-releases, com suas respectivas datas de publicação.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.HandleError(ListVersions())
	},
}

func init() {
	VersionCmd.AddCommand(listCmd)
}
