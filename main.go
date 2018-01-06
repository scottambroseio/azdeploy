package main

import (
	"github.com/scottrangerio/azuredeployment/cmd"
	_ "github.com/scottrangerio/azuredeployment/cmd/version"
)

func main() {
	cmd.Execute()
}
