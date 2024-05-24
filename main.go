package main

import (
	"os"
	"github.com/RaphaelTaibi/Go-BlockChain/tree/main/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()

}
