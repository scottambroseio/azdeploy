package logicapp

import (
	"testing"

	"github.com/scottrangerio/azdeploy/cmdtest"
)

func TestLogicAppCmd(t *testing.T) {
	cmd := newLogicAppCmd()

	cmdtest.TestCommand(t, cmd)
}
