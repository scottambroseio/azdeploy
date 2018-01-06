package cmdtest

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestCommand(t *testing.T, c *cobra.Command) {
	if c.Short == "" {
		t.Error("The Short field isn't populated")
	}

	if c.Long == "" {
		t.Error("The Long field isn't populated")
	}

	if c.Example == "" {
		t.Error("The Example field isn't populated")
	}
}
