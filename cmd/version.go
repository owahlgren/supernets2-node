package main

import (
	"os"

	"github.com/0xPolygon/supernets2-node"
	"github.com/urfave/cli/v2"
)

func versionCmd(*cli.Context) error {
	supernets2.PrintVersion(os.Stdout)
	return nil
}
