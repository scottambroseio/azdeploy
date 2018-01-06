package version

import (
	"fmt"

	"github.com/scottrangerio/azuredeployment/cmd"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version of the CLI",
	Long:  "The version of the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.1.0")
	},
}

func init() {
	cmd.Register(versionCmd)
}
