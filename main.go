package main

import (
	"os"
	"github.com/RaphaelTaibi/Go-BlockChain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()

}
