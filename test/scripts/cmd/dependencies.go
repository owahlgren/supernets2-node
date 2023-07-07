package main

import (
	"github.com/0xPolygon/supernets2-node/test/scripts/cmd/dependencies"
	"github.com/urfave/cli/v2"
)

func updateDeps(ctx *cli.Context) error {
	cfg := &dependencies.Config{
		Images: &dependencies.ImagesConfig{
			Names:          []string{"hermeznetwork/geth-supernets2-contracts", "hermeznetwork/zkprover-local"},
			TargetFilePath: "../../../docker-compose.yml",
		},
		PB: &dependencies.PBConfig{
			TargetDirPath: "../../../proto/src",
			SourceRepo:    "https://github.com/0xPolygonHermez/supernets2-comms-protocol.git",
		},
		TV: &dependencies.TVConfig{
			TargetDirPath: "../../../test/vectors/src",
			SourceRepo:    "https://github.com/0xPolygonHermez/supernets2-testvectors.git",
		},
	}

	return dependencies.NewManager(cfg).Run()
}
