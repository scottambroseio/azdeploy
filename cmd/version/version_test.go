package version

import (
	"bytes"
	"testing"
)

func TestVersionCmd_PrintsVersion(t *testing.T) {
	version = "1.2.3"
	expected := "1.2.3\n"

	b := bytes.NewBuffer([]byte{})

	cmd := newVersionCmd(b)

	cmd.Run(cmd, []string{})

	if s := b.String(); s != expected {
		t.Fatalf("expected %s, got %s", expected, s)
	}
}
