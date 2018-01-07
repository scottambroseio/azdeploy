package version

import (
	"fmt"
	"io"
	"os"

	"github.com/scottrangerio/azdeploy/cmd"
	"github.com/spf13/cobra"
)

var version = "version not set"

func newVersionCmd(w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "The CLI version",
		Long:    "The CLI version",
		Example: "azdeploy version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(w, version)
		},
	}
}

func init() {
	cmd.AddCommand(newVersionCmd(os.Stdout))
}
