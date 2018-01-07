package cmd

import (
	"testing"

	"github.com/scottrangerio/azdeploy/cmdtest"
)

func TestRootCmd(t *testing.T) {
	cmdtest.TestCommand(t, rootCmd)
}
