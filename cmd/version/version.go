package version

import (
	"fmt"
	"io"

	"github.com/scottrangerio/azuredeployment/cmd"
	"github.com/spf13/cobra"
)

var version = "version not set"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version of the CLI",
	Long:  "The version of the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func newVersionCmd(w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "The CLI version",
		Long:  "The CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(w, version)
		},
	}
}

func init() {
	cmd.AddCommand(newVersionCmd)
}
