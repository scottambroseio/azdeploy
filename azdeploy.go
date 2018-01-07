package main

import (
	"github.com/scottrangerio/azdeploy/cmd"
	_ "github.com/scottrangerio/azdeploy/cmd/version"
)

func main() {
	cmd.Execute()
}
