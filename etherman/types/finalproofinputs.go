package types

import "github.com/0xPolygon/supernets2-node/aggregator/pb"

// FinalProofInputs struct
type FinalProofInputs struct {
	FinalProof       *pb.FinalProof
	NewLocalExitRoot []byte
	NewStateRoot     []byte
}
