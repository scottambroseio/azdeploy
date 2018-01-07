package logicapp

import (
	"github.com/spf13/cobra"

	"github.com/scottrangerio/azdeploy/cmd"
)

func newLogicAppCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "logicapp",
		Short:   "Provides commands for deploying Logic Apps",
		Long:    "Provides commands for deploying Logic Apps",
		Example: "logicapp",
	}
}

func init() {
	cmd.AddCommand(newLogicAppCmd())
}
