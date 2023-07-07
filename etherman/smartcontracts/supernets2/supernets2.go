// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package supernets2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Supernets2BatchData is an auto generated low-level Go binding around an user-defined struct.
type Supernets2BatchData struct {
	TransactionsHash   [32]byte
	GlobalExitRoot     [32]byte
	Timestamp          uint64
	MinForcedTimestamp uint64
}

// Supernets2ForcedBatchData is an auto generated low-level Go binding around an user-defined struct.
type Supernets2ForcedBatchData struct {
	Transactions       []byte
	GlobalExitRoot     [32]byte
	MinForcedTimestamp uint64
}

// Supernets2InitializePackedParameters is an auto generated low-level Go binding around an user-defined struct.
type Supernets2InitializePackedParameters struct {
	Admin                    common.Address
	TrustedSequencer         common.Address
	PendingStateTimeout      uint64
	TrustedAggregator        common.Address
	TrustedAggregatorTimeout uint64
}

// Supernets2MetaData contains all meta data concerning the Supernets2 contract.
var Supernets2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISupernets2GlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_matic\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"_rollupVerifier\",\"type\":\"address\"},{\"internalType\":\"contractISupernets2Bridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractISupernets2DataCommittee\",\"name\":\"_dataCommitteeAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_forkID\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ActivateForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"UpdateSupernets2Version\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequencedBatchNum\",\"type\":\"uint64\"}],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractISupernets2Bridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"checkStateRootInsidePrime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataCommitteeAddress\",\"outputs\":[{\"internalType\":\"contractISupernets2DataCommittee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maticAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forkID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractISupernets2GlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"trustedAggregatorTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structSupernets2.InitializePackedParameters\",\"name\":\"initializePackedParameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"genesisRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isForcedBatchDisallowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingState\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingStateConsolidated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matic\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingStateTransitions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifierRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structSupernets2.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signaturesAndAddrs\",\"type\":\"bytes\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structSupernets2.ForcedBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"sequencedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"setTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101606040523480156200001257600080fd5b50604051620061cb380380620061cb8339810160408190526200003591620000ac565b6001600160a01b0396871660c05294861660805292851660a05290841660e052909216610100526001600160401b039182166101205216610140526200014f565b6001600160a01b03811681146200008c57600080fd5b50565b80516001600160401b0381168114620000a757600080fd5b919050565b600080600080600080600060e0888a031215620000c857600080fd5b8751620000d58162000076565b6020890151909750620000e88162000076565b6040890151909650620000fb8162000076565b60608901519095506200010e8162000076565b6080890151909450620001218162000076565b92506200013160a089016200008f565b91506200014160c089016200008f565b905092959891949750929550565b60805160a05160c05160e051610100516101205161014051615f9c6200022f600039600081816106b501528181610ed101526130120152600081816108220152610ea70152600081816105c001526116750152600081816107d5015281816118c5015281816136ac015261430801526000818161098e0152818161145001528181611dc20152818161288c0152818161384901528181613aa40152614667015260008181610a4e015281816141f501526149d30152600081816108de01528181611893015281816123e301528181613a780152614ac20152615f9c6000f3fe608060405234801561001057600080fd5b50600436106103c55760003560e01c8063841b24d7116101ff578063c89e42df1161011a578063e7a7ed02116100ad578063f14916d61161007c578063f14916d614610ab0578063f2fde38b14610ac3578063f851a44014610ad6578063f8b823e414610af657600080fd5b8063e7a7ed0214610a19578063e8bf92ed14610a49578063eaeb077b14610a70578063ed6b010414610a8357600080fd5b8063d8d1091b116100e9578063d8d1091b146109c3578063d939b315146109d6578063dbc16976146109fe578063e6ad707e14610a0657600080fd5b8063c89e42df14610947578063cfa8ed471461095a578063d02103ca14610989578063d2e129f9146109b057600080fd5b8063ada8f91911610192578063b6b0b09711610161578063b6b0b097146108d9578063ba58ae3914610900578063c0ed84e014610913578063c754c7ed1461091b57600080fd5b8063ada8f9191461080a578063adc879e91461081d578063afd23cbe14610844578063b4d63f581461087257600080fd5b80639c9f3dfe116101ce5780639c9f3dfe146107aa578063a066215c146107bd578063a3c573eb146107d0578063a50a164b146107f757600080fd5b8063841b24d71461074c5780638c3d73011461077c5780638da5cb5b1461078457806399f5634e146107a257600080fd5b80634a1a89a7116102ef57806360469169116102825780637215541a116102515780637215541a146106895780637fcb36531461069c578063831c7ead146106b0578063837a4738146106d757600080fd5b806360469169146106465780636b8616ce1461064e5780636ff512cc1461066e578063715018a61461068157600080fd5b80634fd70464116102be5780634fd70464146105f55780635392c5e014610608578063542028d5146106365780635ec919581461063e57600080fd5b80634a1a89a7146105885780634a910e6a146105a85780634df61d24146105bb5780634e487706146105e257600080fd5b80632678224711610367578063423fa85611610336578063423fa85614610519578063438a539914610539578063456052671461054c578063458c04771461057457600080fd5b8063267822471461048257806329878983146104c7578063383b3be8146104f3578063394218e91461050657600080fd5b806315064c96116103a357806315064c961461042b5780631816b7e51461044857806319d8ac611461045b578063220d78991461046f57600080fd5b80630808270c146103ca5780630a0d9fbe146103df578063107bf28c14610416575b600080fd5b6103dd6103d836600461535a565b610aff565b005b606f546103f890610100900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b61041e610bb2565b60405161040d9190615461565b606f546104389060ff1681565b604051901515815260200161040d565b6103dd61045636600461547b565b610c40565b6073546103f89067ffffffffffffffff1681565b61041e61047d36600461549f565b610d58565b607b546104a29073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161040d565b6074546104a29068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b6104386105013660046154ec565b610f2f565b6103dd6105143660046154ec565b610f85565b6073546103f89068010000000000000000900467ffffffffffffffff1681565b6103dd61054736600461552b565b611109565b6073546103f890700100000000000000000000000000000000900467ffffffffffffffff1681565b6079546103f89067ffffffffffffffff1681565b6079546103f89068010000000000000000900467ffffffffffffffff1681565b6103dd6105b63660046154ec565b611986565b6104a27f000000000000000000000000000000000000000000000000000000000000000081565b6103dd6105f03660046154ec565b611a39565b6103dd6106033660046155dd565b611bbd565b6106286106163660046154ec565b60756020526000908152604090205481565b60405190815260200161040d565b61041e611f53565b6103dd611f60565b610628612060565b61062861065c3660046154ec565b60716020526000908152604090205481565b6103dd61067c366004615666565b612076565b6103dd61214b565b6103dd6106973660046154ec565b61215f565b6074546103f89067ffffffffffffffff1681565b6103f87f000000000000000000000000000000000000000000000000000000000000000081565b6107206106e5366004615681565b60786020526000908152604090208054600182015460029092015467ffffffffffffffff808316936801000000000000000090930416919084565b6040805167ffffffffffffffff958616815294909316602085015291830152606082015260800161040d565b6079546103f8907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b6103dd6122cf565b60335473ffffffffffffffffffffffffffffffffffffffff166104a2565b61062861239b565b6103dd6107b83660046154ec565b6124f4565b6103dd6107cb3660046154ec565b612670565b6104a27f000000000000000000000000000000000000000000000000000000000000000081565b6103dd6108053660046155dd565b612776565b6103dd610818366004615666565b61293b565b6103f87f000000000000000000000000000000000000000000000000000000000000000081565b606f5461085f906901000000000000000000900461ffff1681565b60405161ffff909116815260200161040d565b6108b36108803660046154ec565b6072602052600090815260409020805460019091015467ffffffffffffffff808216916801000000000000000090041683565b6040805193845267ffffffffffffffff928316602085015291169082015260600161040d565b6104a27f000000000000000000000000000000000000000000000000000000000000000081565b61043861090e366004615681565b6129ff565b6103f8612a89565b607b546103f89074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b6103dd61095536600461577d565b612ade565b606f546104a2906b010000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b6104a27f000000000000000000000000000000000000000000000000000000000000000081565b6103dd6109be3660046157b2565b612b6b565b6103dd6109d1366004615865565b6130b7565b6079546103f890700100000000000000000000000000000000900467ffffffffffffffff1681565b6103dd613659565b6103dd610a1436600461535a565b613732565b6073546103f8907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b6104a27f000000000000000000000000000000000000000000000000000000000000000081565b6103dd610a7e3660046158da565b613942565b607b54610438907c0100000000000000000000000000000000000000000000000000000000900460ff1681565b6103dd610abe366004615666565b613d38565b6103dd610ad1366004615666565b613e0a565b607a546104a29073ffffffffffffffffffffffffffffffffffffffff1681565b61062860705481565b606f5460ff1615610b3c576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b4c8888888888888888613ebe565b67ffffffffffffffff85166000908152607560209081526040918290205482519081529081018590527f1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010910160405180910390a1610ba8614306565b5050505050505050565b60778054610bbf90615926565b80601f0160208091040260200160405190810160405280929190818152602001828054610beb90615926565b8015610c385780601f10610c0d57610100808354040283529160200191610c38565b820191906000526020600020905b815481529060010190602001808311610c1b57829003601f168201915b505050505081565b607a5473ffffffffffffffffffffffffffffffffffffffff163314610c91576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88161ffff161080610caa57506103ff8161ffff16115b15610ce1576040517f4c2533c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffff0000ffffffffffffffffff16690100000000000000000061ffff8416908102919091179091556040519081527f7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5906020015b60405180910390a150565b67ffffffffffffffff8086166000818152607260205260408082205493881682529020546060929115801590610d8c575081155b15610dc3576040517f6818c29e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610dfa576040517f66385b5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e03846129ff565b610e39576040517f176b913c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080517fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003360601b166020820152603481019690965260548601929092527fffffffffffffffff00000000000000000000000000000000000000000000000060c098891b811660748701527f0000000000000000000000000000000000000000000000000000000000000000891b8116607c8701527f0000000000000000000000000000000000000000000000000000000000000000891b81166084870152608c86019490945260ac85015260cc840194909452509290931b90911660ec830152805180830360d401815260f4909201905290565b60795467ffffffffffffffff82811660009081526078602052604081205490924292610f7392700100000000000000000000000000000000909204811691166159a8565b67ffffffffffffffff16111592915050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314610fd6576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff8216111561101d576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff1661108c5760795467ffffffffffffffff780100000000000000000000000000000000000000000000000090910481169082161061108c576040517f401636df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6079805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527f1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a190602001610d4d565b606f5460ff1615611146576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f546b010000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1633146111a6576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8360008190036111e2576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e881111561121e576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff6801000000000000000082048116600081815260726020526040812054838516949293700100000000000000000000000000000000909304909216919082905b868110156116375760008c8c83818110611286576112866159d0565b90506080020180360381019061129c91906159ff565b606081015190915067ffffffffffffffff161561140d57846112bd81615a70565b955050600081600001518260200151836060015160405160200161131993929190928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89166000908152607190935291205490915081146113a2576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8087166000908152607160205260408082209190915560608401519084015190821691161015611407576040517f7f7ab87200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061150b565b6020810151158015906114d4575060208101516040517f257b363200000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063257b3632906024016020604051808303816000875af11580156114ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114d29190615a97565b155b1561150b576040517f73bd668d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8667ffffffffffffffff16816040015167ffffffffffffffff16108061153e575042816040015167ffffffffffffffff16115b15611575576040517fea82791600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b838160000151826020015183604001518e6040516020016116049594939291909485526020850193909352604084019190915260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060808401919091521b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166068820152607c0190565b6040516020818303038152906040528051906020012093508060400151965050808061162f90615ab0565b91505061126a565b506040517fc7a823e000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063c7a823e0906116ae9085908c908c90600401615b31565b60006040518083038186803b1580156116c657600080fd5b505afa1580156116da573d6000803e3d6000fd5b5050505085846116ea91906159a8565b60735490945067ffffffffffffffff780100000000000000000000000000000000000000000000000090910481169084161115611753576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061175f8285615b54565b6117739067ffffffffffffffff1688615b75565b604080516060810182528581524267ffffffffffffffff908116602080840191825260738054680100000000000000009081900485168688019081528d861660008181526072909552979093209551865592516001909501805492519585167fffffffffffffffffffffffffffffffff000000000000000000000000000000009384161795851684029590951790945583548c841691161793029290921790559091508281169085161461186957607380547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8716021790555b6118bb33308360705461187c9190615b8e565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692919061438e565b6118c3614470565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561192b57600080fd5b505af115801561193f573d6000803e3d6000fd5b505060405167ffffffffffffffff881692507f303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce9150600090a2505050505050505050505050565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314611a2d57606f5460ff16156119ee576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119f781610f2f565b611a2d576040517f0ce9e4a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611a368161451d565b50565b607a5473ffffffffffffffffffffffffffffffffffffffff163314611a8a576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115611ad1576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff16611b3c57607b5467ffffffffffffffff74010000000000000000000000000000000000000000909104811690821610611b3c576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b90602001610d4d565b606f5460ff1615611bfa576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff8681166000908152607260205260409020600101544292611c47927801000000000000000000000000000000000000000000000000909104811691166159a8565b67ffffffffffffffff161115611c89576040517f8a0704d300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8611c968787615b54565b67ffffffffffffffff161115611cd8576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ce787878787878787614730565b611cf085614ae9565b607954700100000000000000000000000000000000900467ffffffffffffffff16600003611e3857607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff87811691821790925560009081526075602052604090208490556079541615611d9357607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018590527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b158015611e1b57600080fd5b505af1158015611e2f573d6000803e3d6000fd5b50505050611f08565b611e40614470565b6079805467ffffffffffffffff16906000611e5a83615a70565b825467ffffffffffffffff9182166101009390930a92830292820219169190911790915560408051608081018252428316815288831660208083019182528284018a8152606084018a815260795487166000908152607890935294909120925183549251861668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169516949094171781559151600183015551600290910155505b604051838152339067ffffffffffffffff8716907f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966906020015b60405180910390a350505050505050565b60768054610bbf90615926565b607a5473ffffffffffffffffffffffffffffffffffffffff163314611fb1576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b547c0100000000000000000000000000000000000000000000000000000000900460ff1661200d576040517ff6ba91a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690556040517f854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f90600090a1565b600060705460646120719190615b8e565b905090565b607a5473ffffffffffffffffffffffffffffffffffffffff1633146120c7576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fff0000000000000000000000000000000000000000ffffffffffffffffffffff166b01000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610d4d565b612153614cc9565b61215d6000614d4a565b565b60335473ffffffffffffffffffffffffffffffffffffffff1633146122c7576000612188612a89565b90508067ffffffffffffffff168267ffffffffffffffff16116121d7576040517f812a372d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff680100000000000000009091048116908316118061221d575067ffffffffffffffff80831660009081526072602052604090206001015416155b15612254576040517f98c5c01400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80831660009081526072602052604090206001015442916122839162093a8091166159a8565b67ffffffffffffffff1611156122c5576040517fd257555a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b611a36614306565b607b5473ffffffffffffffffffffffffffffffffffffffff163314612320576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b54607a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600090819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561242a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061244e9190615a97565b9050600061245a612a89565b60735467ffffffffffffffff6801000000000000000082048116916124b29170010000000000000000000000000000000082048116917801000000000000000000000000000000000000000000000000900416615b54565b6124bc91906159a8565b6124c69190615b54565b67ffffffffffffffff169050806000036124e35760009250505090565b6124ed8183615bd4565b9250505090565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612545576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff8216111561258c576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff166125f35760795467ffffffffffffffff7001000000000000000000000000000000009091048116908216106125f3576040517f48a05a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607980547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c7590602001610d4d565b607a5473ffffffffffffffffffffffffffffffffffffffff1633146126c1576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620151808167ffffffffffffffff161115612708576040517fe067dfe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff1661010067ffffffffffffffff8416908102919091179091556040519081527f1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c2890602001610d4d565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1633146127d3576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6127e287878787878787614730565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8781169182179092556000908152607560205260409020849055607954161561285d57607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018590527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b1580156128e557600080fd5b505af11580156128f9573d6000803e3d6000fd5b505060405185815233925067ffffffffffffffff881691507fcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe90602001611f42565b607a5473ffffffffffffffffffffffffffffffffffffffff16331461298c576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610d4d565b600067ffffffff0000000167ffffffffffffffff8316108015612a37575067ffffffff00000001604083901c67ffffffffffffffff16105b8015612a58575067ffffffff00000001608083901c67ffffffffffffffff16105b8015612a6f575067ffffffff0000000160c083901c105b15612a7c57506001919050565b506000919050565b919050565b60795460009067ffffffffffffffff1615612acd575060795467ffffffffffffffff9081166000908152607860205260409020546801000000000000000090041690565b5060745467ffffffffffffffff1690565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612b2f576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6076612b3b8282615c36565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610d4d9190615461565b600054610100900460ff1615808015612b8b5750600054600160ff909116105b80612ba55750303b158015612ba5575060005460ff166001145b612c36576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015612c9457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b612ca16020880188615666565b607a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055612cf66040880160208901615666565b606f805473ffffffffffffffffffffffffffffffffffffffff929092166b010000000000000000000000027fff0000000000000000000000000000000000000000ffffffffffffffffffffff909216919091179055612d5b6080880160608901615666565b6074805473ffffffffffffffffffffffffffffffffffffffff9290921668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff9092169190911790556000805260756020527ff9e3fbf150b7a0077118526f473c53cb4734f166167e2c6213e3567dd390b4ad8690556076612de68682615c36565b506077612df38582615c36565b5062093a80612e086060890160408a016154ec565b67ffffffffffffffff161115612e4a576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612e5a60608801604089016154ec565b6079805467ffffffffffffffff92909216700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff90921691909117905562093a80612ebc60a0890160808a016154ec565b67ffffffffffffffff161115612efe576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f0e60a08801608089016154ec565b6079805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff939093169290920291909117905567016345785d8a0000607055606f80547fffffffffffffffffffffffffffffffffffffffffff00000000000000000000ff166a03ea000000000000070800179055607b80547fffffff000000000000000000ffffffffffffffffffffffffffffffffffffffff167c0100000000000697800000000000000000000000000000000000000000179055612fed614dc1565b7f13e60734641bd97ab6e444f419277d9983e2c3eab9bb631a578ffa07756b635160007f000000000000000000000000000000000000000000000000000000000000000085856040516130439493929190615d50565b60405180910390a180156130ae57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b607b547c0100000000000000000000000000000000000000000000000000000000900460ff1615613114576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff1615613151576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600081900361318d576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156131c9576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff78010000000000000000000000000000000000000000000000008204811691613214918491700100000000000000000000000000000000900416615d88565b111561324c576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff680100000000000000008204811660008181526072602052604081205491937001000000000000000000000000000000009004909216915b848110156134f65760008787838181106132ac576132ac6159d0565b90506020028101906132be9190615d9b565b6132c790615dd9565b9050836132d381615a70565b825180516020918201208185015160408087015190519499509194506000936133359386939101928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89166000908152607190935291205490915081146133be576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86166000908152607160205260408120556133e3600189615b75565b84036134525742607b60149054906101000a900467ffffffffffffffff16846040015161341091906159a8565b67ffffffffffffffff161115613452576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020838101516040805192830188905282018490526060808301919091524260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016608083015233901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c0160405160208183030381529060405280519060200120945050505080806134ee90615ab0565b915050613290565b5061350184846159a8565b6073805467ffffffffffffffff4281167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009092168217808455604080516060810182528781526020808201958652680100000000000000009384900485168284019081528589166000818152607290935284832093518455965160019390930180549151871686027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921693871693909317179091558554938916700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff938602939093167fffffffffffffffff00000000000000000000000000000000ffffffffffffffff90941693909317919091179093559151929550917f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a49190a2505050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff1633146136aa576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dbc169766040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561371257600080fd5b505af1158015613726573d6000803e3d6000fd5b5050505061215d614e61565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16331461378f576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61379f8888888888888888613ebe565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8781169182179092556000908152607560205260409020849055607954161561381a57607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018590527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b1580156138a257600080fd5b505af11580156138b6573d6000803e3d6000fd5b50506079805477ffffffffffffffffffffffffffffffffffffffffffffffff167a093a800000000000000000000000000000000000000000000000001790555050604051838152339067ffffffffffffffff8716907fcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf729060200160405180910390a35050505050505050565b607b547c0100000000000000000000000000000000000000000000000000000000900460ff161561399f576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff16156139dc576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006139e6612060565b905081811115613a22576040517f4732fdb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611388831115613a5e576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613aa073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633308461438e565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613b0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b319190615a97565b60738054919250780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16906018613b6b83615a70565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508484604051613ba2929190615e69565b60408051918290038220602083015281018290527fffffffffffffffff0000000000000000000000000000000000000000000000004260c01b166060820152606801604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815291815281516020928301206073547801000000000000000000000000000000000000000000000000900467ffffffffffffffff1660009081526071909352912055323303613cd257607354604080518381523360208201526060918101829052600091810191909152780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319060800160405180910390a2613d31565b607360189054906101000a900467ffffffffffffffff1667ffffffffffffffff167ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc93182338888604051613d289493929190615e79565b60405180910390a25b5050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314613d89576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607480547fffffffff0000000000000000000000000000000000000000ffffffffffffffff166801000000000000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca90602001610d4d565b613e12614cc9565b73ffffffffffffffffffffffffffffffffffffffff8116613eb5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401612c2d565b611a3681614d4a565b600067ffffffffffffffff891615613f8c5760795467ffffffffffffffff908116908a161115613f1a576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff8089166000908152607860205260409020600281015481549092898116680100000000000000009092041614613f86576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50614028565b5067ffffffffffffffff861660009081526075602052604090205480613fde576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60745467ffffffffffffffff9081169088161115614028576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff908116908916118061405a57508867ffffffffffffffff168867ffffffffffffffff1611155b80614081575060795467ffffffffffffffff68010000000000000000909104811690891611155b156140b8576040517fbfa7079f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff88811660009081526078602052604090205468010000000000000000900481169087161461411b576040517f32a2a77f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061412a8888888589610d58565b905060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160028360405161415f9190615eaf565b602060405180830381855afa15801561417c573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061419f9190615a97565b6141a99190615ec1565b6040805160208101825282815290517f8d8f8a5c00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691638d8f8a5c9161422c9189918991600401615ed5565b602060405180830381865afa158015614249573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061426d9190615f1d565b6142a3576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8a166000908152607860205260409020600201548690036142f9576040517fa47276bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632072f6c56040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561436e57600080fd5b505af1158015614382573d6000803e3d6000fd5b5050505061215d614ef0565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261446a9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152614f83565b50505050565b60795467ffffffffffffffff68010000000000000000820481169116111561215d576079546000906144b99068010000000000000000900467ffffffffffffffff1660016159a8565b90506144c481610f2f565b15611a36576079546000906002906144e790849067ffffffffffffffff16615b54565b6144f19190615f3f565b6144fb90836159a8565b905061450681610f2f565b15614518576145148161451d565b5050565b614514825b60795467ffffffffffffffff680100000000000000009091048116908216111580614557575060795467ffffffffffffffff908116908216115b1561458e576040517fd086b70b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff818116600081815260786020908152604080832080546074805468010000000000000000928390049098167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000090981688179055600282015487865260759094529382902092909255607980547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff169390940292909217909255600182015490517f33d6247d00000000000000000000000000000000000000000000000000000000815260048101919091529091907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b1580156146c057600080fd5b505af11580156146d4573d6000803e3d6000fd5b505050508267ffffffffffffffff168167ffffffffffffffff167f328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e846002015460405161472391815260200190565b60405180910390a3505050565b60008061473b612a89565b905067ffffffffffffffff89161561480b5760795467ffffffffffffffff908116908a161115614797576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff808a1660009081526078602052604090206002810154815490945090918a8116680100000000000000009092041614614805576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506148ac565b67ffffffffffffffff881660009081526075602052604090205491508161485e576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168867ffffffffffffffff1611156148ac576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168767ffffffffffffffff16116148f9576040517fb9b18f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000614908898989868a610d58565b905060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160028360405161493d9190615eaf565b602060405180830381855afa15801561495a573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061497d9190615a97565b6149879190615ec1565b6040805160208101825282815290517f8d8f8a5c00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691638d8f8a5c91614a0a918a918a91600401615ed5565b602060405180830381865afa158015614a27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a4b9190615f1d565b614a81576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6142f933614a8f858c615b54565b67ffffffffffffffff16614aa161239b565b614aab9190615b8e565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190615094565b6000614af3612a89565b905081600080614b038484615b54565b606f5467ffffffffffffffff9182169250600091614b279161010090041642615b75565b90505b8467ffffffffffffffff168467ffffffffffffffff1614614bb25767ffffffffffffffff80851660009081526072602052604090206001810154909116821015614b9057600181015468010000000000000000900467ffffffffffffffff169450614bac565b614b9a8686615b54565b67ffffffffffffffff16935050614bb2565b50614b2a565b6000614bbe8484615b75565b905083811015614c1557808403600c8111614bd95780614bdc565b600c5b9050806103e80a81606f60099054906101000a900461ffff1661ffff160a6070540281614c0b57614c0b615ba5565b0460705550614c85565b838103600c8111614c265780614c29565b600c5b90506000816103e80a82606f60099054906101000a900461ffff1661ffff160a670de0b6b3a76400000281614c6057614c60615ba5565b04905080607054670de0b6b3a76400000281614c7e57614c7e615ba5565b0460705550505b683635c9adc5dea000006070541115614caa57683635c9adc5dea000006070556130ae565b633b9aca0060705410156130ae57633b9aca0060705550505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461215d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401612c2d565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16614e58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401612c2d565b61215d33614d4a565b606f5460ff16614e9d576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b390600090a1565b606f5460ff1615614f2d576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a549790600090a1565b6000614fe5826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166150ea9092919063ffffffff16565b80519091501561508f57808060200190518101906150039190615f1d565b61508f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401612c2d565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff831660248201526044810182905261508f9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064016143e8565b60606150f98484600085615101565b949350505050565b606082471015615193576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401612c2d565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516151bc9190615eaf565b60006040518083038185875af1925050503d80600081146151f9576040519150601f19603f3d011682016040523d82523d6000602084013e6151fe565b606091505b509150915061520f8783838761521a565b979650505050505050565b606083156152b05782516000036152a95773ffffffffffffffffffffffffffffffffffffffff85163b6152a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401612c2d565b50816150f9565b6150f983838151156152c55781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c2d9190615461565b803567ffffffffffffffff81168114612a8457600080fd5b60008083601f84011261532357600080fd5b50813567ffffffffffffffff81111561533b57600080fd5b60208301915083602082850101111561535357600080fd5b9250929050565b60008060008060008060008060e0898b03121561537657600080fd5b61537f896152f9565b975061538d60208a016152f9565b965061539b60408a016152f9565b95506153a960608a016152f9565b94506080890135935060a0890135925060c089013567ffffffffffffffff8111156153d357600080fd5b6153df8b828c01615311565b999c989b5096995094979396929594505050565b60005b8381101561540e5781810151838201526020016153f6565b50506000910152565b6000815180845261542f8160208601602086016153f3565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006154746020830184615417565b9392505050565b60006020828403121561548d57600080fd5b813561ffff8116811461547457600080fd5b600080600080600060a086880312156154b757600080fd5b6154c0866152f9565b94506154ce602087016152f9565b94979496505050506040830135926060810135926080909101359150565b6000602082840312156154fe57600080fd5b615474826152f9565b803573ffffffffffffffffffffffffffffffffffffffff81168114612a8457600080fd5b60008060008060006060868803121561554357600080fd5b853567ffffffffffffffff8082111561555b57600080fd5b818801915088601f83011261556f57600080fd5b81358181111561557e57600080fd5b8960208260071b850101111561559357600080fd5b602083019750809650506155a960208901615507565b945060408801359150808211156155bf57600080fd5b506155cc88828901615311565b969995985093965092949392505050565b600080600080600080600060c0888a0312156155f857600080fd5b615601886152f9565b965061560f602089016152f9565b955061561d604089016152f9565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561564757600080fd5b6156538a828b01615311565b989b979a50959850939692959293505050565b60006020828403121561567857600080fd5b61547482615507565b60006020828403121561569357600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600067ffffffffffffffff808411156156e4576156e461569a565b604051601f85017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561572a5761572a61569a565b8160405280935085815286868601111561574357600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261576e57600080fd5b615474838335602085016156c9565b60006020828403121561578f57600080fd5b813567ffffffffffffffff8111156157a657600080fd5b6150f98482850161575d565b6000806000806000808688036101208112156157cd57600080fd5b60a08112156157db57600080fd5b5086955060a0870135945060c087013567ffffffffffffffff8082111561580157600080fd5b61580d8a838b0161575d565b955060e089013591508082111561582357600080fd5b61582f8a838b0161575d565b945061010089013591508082111561584657600080fd5b5061585389828a01615311565b979a9699509497509295939492505050565b6000806020838503121561587857600080fd5b823567ffffffffffffffff8082111561589057600080fd5b818501915085601f8301126158a457600080fd5b8135818111156158b357600080fd5b8660208260051b85010111156158c857600080fd5b60209290920196919550909350505050565b6000806000604084860312156158ef57600080fd5b833567ffffffffffffffff81111561590657600080fd5b61591286828701615311565b909790965060209590950135949350505050565b600181811c9082168061593a57607f821691505b602082108103615973577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff8181168382160190808211156159c9576159c9615979565b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060808284031215615a1157600080fd5b6040516080810181811067ffffffffffffffff82111715615a3457615a3461569a565b80604052508235815260208301356020820152615a53604084016152f9565b6040820152615a64606084016152f9565b60608201529392505050565b600067ffffffffffffffff808316818103615a8d57615a8d615979565b6001019392505050565b600060208284031215615aa957600080fd5b5051919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615ae157615ae1615979565b5060010190565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b838152604060208201526000615b4b604083018486615ae8565b95945050505050565b67ffffffffffffffff8281168282160390808211156159c9576159c9615979565b81810381811115615b8857615b88615979565b92915050565b8082028115828204841417615b8857615b88615979565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082615be357615be3615ba5565b500490565b601f82111561508f57600081815260208120601f850160051c81016020861015615c0f5750805b601f850160051c820191505b81811015615c2e57828155600101615c1b565b505050505050565b815167ffffffffffffffff811115615c5057615c5061569a565b615c6481615c5e8454615926565b84615be8565b602080601f831160018114615cb75760008415615c815750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555615c2e565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015615d0457888601518255948401946001909101908401615ce5565b5085821015615d4057878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600067ffffffffffffffff808716835280861660208401525060606040830152615d7e606083018486615ae8565b9695505050505050565b80820180821115615b8857615b88615979565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112615dcf57600080fd5b9190910192915050565b600060608236031215615deb57600080fd5b6040516060810167ffffffffffffffff8282108183111715615e0f57615e0f61569a565b816040528435915080821115615e2457600080fd5b50830136601f820112615e3657600080fd5b615e45368235602084016156c9565b82525060208301356020820152615e5e604084016152f9565b604082015292915050565b8183823760009101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201526000615d7e606083018486615ae8565b60008251615dcf8184602087016153f3565b600082615ed057615ed0615ba5565b500690565b604081526000615ee9604083018587615ae8565b905060208083018460005b6001811015615f1157815183529183019190830190600101615ef4565b50505050949350505050565b600060208284031215615f2f57600080fd5b8151801515811461547457600080fd5b600067ffffffffffffffff80841680615f5a57615f5a615ba5565b9216919091049291505056fea26469706673582212205f6db821da6f759d2be7ce71edb3709e45a5d7fdf189e8a24601801279bb362364736f6c63430008110033",
}

// Supernets2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Supernets2MetaData.ABI instead.
var Supernets2ABI = Supernets2MetaData.ABI

// Supernets2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Supernets2MetaData.Bin instead.
var Supernets2Bin = Supernets2MetaData.Bin

// DeploySupernets2 deploys a new Ethereum contract, binding an instance of Supernets2 to it.
func DeploySupernets2(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _matic common.Address, _rollupVerifier common.Address, _bridgeAddress common.Address, _dataCommitteeAddress common.Address, _chainID uint64, _forkID uint64) (common.Address, *types.Transaction, *Supernets2, error) {
	parsed, err := Supernets2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Supernets2Bin), backend, _globalExitRootManager, _matic, _rollupVerifier, _bridgeAddress, _dataCommitteeAddress, _chainID, _forkID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Supernets2{Supernets2Caller: Supernets2Caller{contract: contract}, Supernets2Transactor: Supernets2Transactor{contract: contract}, Supernets2Filterer: Supernets2Filterer{contract: contract}}, nil
}

// Supernets2 is an auto generated Go binding around an Ethereum contract.
type Supernets2 struct {
	Supernets2Caller     // Read-only binding to the contract
	Supernets2Transactor // Write-only binding to the contract
	Supernets2Filterer   // Log filterer for contract events
}

// Supernets2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Supernets2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Supernets2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Supernets2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Supernets2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Supernets2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Supernets2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Supernets2Session struct {
	Contract     *Supernets2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Supernets2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Supernets2CallerSession struct {
	Contract *Supernets2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Supernets2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Supernets2TransactorSession struct {
	Contract     *Supernets2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Supernets2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Supernets2Raw struct {
	Contract *Supernets2 // Generic contract binding to access the raw methods on
}

// Supernets2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Supernets2CallerRaw struct {
	Contract *Supernets2Caller // Generic read-only contract binding to access the raw methods on
}

// Supernets2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Supernets2TransactorRaw struct {
	Contract *Supernets2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSupernets2 creates a new instance of Supernets2, bound to a specific deployed contract.
func NewSupernets2(address common.Address, backend bind.ContractBackend) (*Supernets2, error) {
	contract, err := bindSupernets2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Supernets2{Supernets2Caller: Supernets2Caller{contract: contract}, Supernets2Transactor: Supernets2Transactor{contract: contract}, Supernets2Filterer: Supernets2Filterer{contract: contract}}, nil
}

// NewSupernets2Caller creates a new read-only instance of Supernets2, bound to a specific deployed contract.
func NewSupernets2Caller(address common.Address, caller bind.ContractCaller) (*Supernets2Caller, error) {
	contract, err := bindSupernets2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Supernets2Caller{contract: contract}, nil
}

// NewSupernets2Transactor creates a new write-only instance of Supernets2, bound to a specific deployed contract.
func NewSupernets2Transactor(address common.Address, transactor bind.ContractTransactor) (*Supernets2Transactor, error) {
	contract, err := bindSupernets2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Supernets2Transactor{contract: contract}, nil
}

// NewSupernets2Filterer creates a new log filterer instance of Supernets2, bound to a specific deployed contract.
func NewSupernets2Filterer(address common.Address, filterer bind.ContractFilterer) (*Supernets2Filterer, error) {
	contract, err := bindSupernets2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Supernets2Filterer{contract: contract}, nil
}

// bindSupernets2 binds a generic wrapper to an already deployed contract.
func bindSupernets2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Supernets2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Supernets2 *Supernets2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Supernets2.Contract.Supernets2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Supernets2 *Supernets2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.Contract.Supernets2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Supernets2 *Supernets2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Supernets2.Contract.Supernets2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Supernets2 *Supernets2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Supernets2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Supernets2 *Supernets2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Supernets2 *Supernets2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Supernets2.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Supernets2 *Supernets2Caller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Supernets2 *Supernets2Session) Admin() (common.Address, error) {
	return _Supernets2.Contract.Admin(&_Supernets2.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Supernets2 *Supernets2CallerSession) Admin() (common.Address, error) {
	return _Supernets2.Contract.Admin(&_Supernets2.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Supernets2 *Supernets2Caller) BatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "batchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Supernets2 *Supernets2Session) BatchFee() (*big.Int, error) {
	return _Supernets2.Contract.BatchFee(&_Supernets2.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Supernets2 *Supernets2CallerSession) BatchFee() (*big.Int, error) {
	return _Supernets2.Contract.BatchFee(&_Supernets2.CallOpts)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2Caller) BatchNumToStateRoot(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "batchNumToStateRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2Session) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Supernets2.Contract.BatchNumToStateRoot(&_Supernets2.CallOpts, arg0)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2CallerSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Supernets2.Contract.BatchNumToStateRoot(&_Supernets2.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Supernets2 *Supernets2Caller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Supernets2 *Supernets2Session) BridgeAddress() (common.Address, error) {
	return _Supernets2.Contract.BridgeAddress(&_Supernets2.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Supernets2 *Supernets2CallerSession) BridgeAddress() (common.Address, error) {
	return _Supernets2.Contract.BridgeAddress(&_Supernets2.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Supernets2 *Supernets2Caller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Supernets2 *Supernets2Session) CalculateRewardPerBatch() (*big.Int, error) {
	return _Supernets2.Contract.CalculateRewardPerBatch(&_Supernets2.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Supernets2 *Supernets2CallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Supernets2.Contract.CalculateRewardPerBatch(&_Supernets2.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Supernets2 *Supernets2Caller) ChainID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Supernets2 *Supernets2Session) ChainID() (uint64, error) {
	return _Supernets2.Contract.ChainID(&_Supernets2.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) ChainID() (uint64, error) {
	return _Supernets2.Contract.ChainID(&_Supernets2.CallOpts)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Supernets2 *Supernets2Caller) CheckStateRootInsidePrime(opts *bind.CallOpts, newStateRoot *big.Int) (bool, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "checkStateRootInsidePrime", newStateRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Supernets2 *Supernets2Session) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Supernets2.Contract.CheckStateRootInsidePrime(&_Supernets2.CallOpts, newStateRoot)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Supernets2 *Supernets2CallerSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Supernets2.Contract.CheckStateRootInsidePrime(&_Supernets2.CallOpts, newStateRoot)
}

// DataCommitteeAddress is a free data retrieval call binding the contract method 0x4df61d24.
//
// Solidity: function dataCommitteeAddress() view returns(address)
func (_Supernets2 *Supernets2Caller) DataCommitteeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "dataCommitteeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataCommitteeAddress is a free data retrieval call binding the contract method 0x4df61d24.
//
// Solidity: function dataCommitteeAddress() view returns(address)
func (_Supernets2 *Supernets2Session) DataCommitteeAddress() (common.Address, error) {
	return _Supernets2.Contract.DataCommitteeAddress(&_Supernets2.CallOpts)
}

// DataCommitteeAddress is a free data retrieval call binding the contract method 0x4df61d24.
//
// Solidity: function dataCommitteeAddress() view returns(address)
func (_Supernets2 *Supernets2CallerSession) DataCommitteeAddress() (common.Address, error) {
	return _Supernets2.Contract.DataCommitteeAddress(&_Supernets2.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Caller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Session) ForceBatchTimeout() (uint64, error) {
	return _Supernets2.Contract.ForceBatchTimeout(&_Supernets2.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) ForceBatchTimeout() (uint64, error) {
	return _Supernets2.Contract.ForceBatchTimeout(&_Supernets2.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2Caller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2Session) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Supernets2.Contract.ForcedBatches(&_Supernets2.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2CallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Supernets2.Contract.ForcedBatches(&_Supernets2.CallOpts, arg0)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Supernets2 *Supernets2Caller) ForkID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "forkID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Supernets2 *Supernets2Session) ForkID() (uint64, error) {
	return _Supernets2.Contract.ForkID(&_Supernets2.CallOpts)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) ForkID() (uint64, error) {
	return _Supernets2.Contract.ForkID(&_Supernets2.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Supernets2 *Supernets2Caller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Supernets2 *Supernets2Session) GetForcedBatchFee() (*big.Int, error) {
	return _Supernets2.Contract.GetForcedBatchFee(&_Supernets2.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Supernets2 *Supernets2CallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Supernets2.Contract.GetForcedBatchFee(&_Supernets2.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Supernets2 *Supernets2Caller) GetInputSnarkBytes(opts *bind.CallOpts, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "getInputSnarkBytes", initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Supernets2 *Supernets2Session) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Supernets2.Contract.GetInputSnarkBytes(&_Supernets2.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Supernets2 *Supernets2CallerSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Supernets2.Contract.GetInputSnarkBytes(&_Supernets2.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2Caller) GetLastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "getLastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2Session) GetLastVerifiedBatch() (uint64, error) {
	return _Supernets2.Contract.GetLastVerifiedBatch(&_Supernets2.CallOpts)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) GetLastVerifiedBatch() (uint64, error) {
	return _Supernets2.Contract.GetLastVerifiedBatch(&_Supernets2.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Supernets2 *Supernets2Caller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Supernets2 *Supernets2Session) GlobalExitRootManager() (common.Address, error) {
	return _Supernets2.Contract.GlobalExitRootManager(&_Supernets2.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Supernets2 *Supernets2CallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Supernets2.Contract.GlobalExitRootManager(&_Supernets2.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Supernets2 *Supernets2Caller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Supernets2 *Supernets2Session) IsEmergencyState() (bool, error) {
	return _Supernets2.Contract.IsEmergencyState(&_Supernets2.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Supernets2 *Supernets2CallerSession) IsEmergencyState() (bool, error) {
	return _Supernets2.Contract.IsEmergencyState(&_Supernets2.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Supernets2 *Supernets2Caller) IsForcedBatchDisallowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "isForcedBatchDisallowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Supernets2 *Supernets2Session) IsForcedBatchDisallowed() (bool, error) {
	return _Supernets2.Contract.IsForcedBatchDisallowed(&_Supernets2.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Supernets2 *Supernets2CallerSession) IsForcedBatchDisallowed() (bool, error) {
	return _Supernets2.Contract.IsForcedBatchDisallowed(&_Supernets2.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Supernets2 *Supernets2Caller) IsPendingStateConsolidable(opts *bind.CallOpts, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "isPendingStateConsolidable", pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Supernets2 *Supernets2Session) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Supernets2.Contract.IsPendingStateConsolidable(&_Supernets2.CallOpts, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Supernets2 *Supernets2CallerSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Supernets2.Contract.IsPendingStateConsolidable(&_Supernets2.CallOpts, pendingStateNum)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastBatchSequenced() (uint64, error) {
	return _Supernets2.Contract.LastBatchSequenced(&_Supernets2.CallOpts)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastBatchSequenced() (uint64, error) {
	return _Supernets2.Contract.LastBatchSequenced(&_Supernets2.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastForceBatch() (uint64, error) {
	return _Supernets2.Contract.LastForceBatch(&_Supernets2.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastForceBatch() (uint64, error) {
	return _Supernets2.Contract.LastForceBatch(&_Supernets2.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastForceBatchSequenced() (uint64, error) {
	return _Supernets2.Contract.LastForceBatchSequenced(&_Supernets2.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Supernets2.Contract.LastForceBatchSequenced(&_Supernets2.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastPendingState(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastPendingState")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastPendingState() (uint64, error) {
	return _Supernets2.Contract.LastPendingState(&_Supernets2.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastPendingState() (uint64, error) {
	return _Supernets2.Contract.LastPendingState(&_Supernets2.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastPendingStateConsolidated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastPendingStateConsolidated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastPendingStateConsolidated() (uint64, error) {
	return _Supernets2.Contract.LastPendingStateConsolidated(&_Supernets2.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastPendingStateConsolidated() (uint64, error) {
	return _Supernets2.Contract.LastPendingStateConsolidated(&_Supernets2.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastTimestamp() (uint64, error) {
	return _Supernets2.Contract.LastTimestamp(&_Supernets2.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastTimestamp() (uint64, error) {
	return _Supernets2.Contract.LastTimestamp(&_Supernets2.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastVerifiedBatch() (uint64, error) {
	return _Supernets2.Contract.LastVerifiedBatch(&_Supernets2.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastVerifiedBatch() (uint64, error) {
	return _Supernets2.Contract.LastVerifiedBatch(&_Supernets2.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Supernets2 *Supernets2Caller) Matic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "matic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Supernets2 *Supernets2Session) Matic() (common.Address, error) {
	return _Supernets2.Contract.Matic(&_Supernets2.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Supernets2 *Supernets2CallerSession) Matic() (common.Address, error) {
	return _Supernets2.Contract.Matic(&_Supernets2.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Supernets2 *Supernets2Caller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Supernets2 *Supernets2Session) MultiplierBatchFee() (uint16, error) {
	return _Supernets2.Contract.MultiplierBatchFee(&_Supernets2.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Supernets2 *Supernets2CallerSession) MultiplierBatchFee() (uint16, error) {
	return _Supernets2.Contract.MultiplierBatchFee(&_Supernets2.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Supernets2 *Supernets2Caller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Supernets2 *Supernets2Session) NetworkName() (string, error) {
	return _Supernets2.Contract.NetworkName(&_Supernets2.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Supernets2 *Supernets2CallerSession) NetworkName() (string, error) {
	return _Supernets2.Contract.NetworkName(&_Supernets2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Supernets2 *Supernets2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Supernets2 *Supernets2Session) Owner() (common.Address, error) {
	return _Supernets2.Contract.Owner(&_Supernets2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Supernets2 *Supernets2CallerSession) Owner() (common.Address, error) {
	return _Supernets2.Contract.Owner(&_Supernets2.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Supernets2 *Supernets2Caller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Supernets2 *Supernets2Session) PendingAdmin() (common.Address, error) {
	return _Supernets2.Contract.PendingAdmin(&_Supernets2.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Supernets2 *Supernets2CallerSession) PendingAdmin() (common.Address, error) {
	return _Supernets2.Contract.PendingAdmin(&_Supernets2.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Caller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Session) PendingStateTimeout() (uint64, error) {
	return _Supernets2.Contract.PendingStateTimeout(&_Supernets2.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) PendingStateTimeout() (uint64, error) {
	return _Supernets2.Contract.PendingStateTimeout(&_Supernets2.CallOpts)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Supernets2 *Supernets2Caller) PendingStateTransitions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "pendingStateTransitions", arg0)

	outstruct := new(struct {
		Timestamp         uint64
		LastVerifiedBatch uint64
		ExitRoot          [32]byte
		StateRoot         [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.LastVerifiedBatch = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ExitRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StateRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Supernets2 *Supernets2Session) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Supernets2.Contract.PendingStateTransitions(&_Supernets2.CallOpts, arg0)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Supernets2 *Supernets2CallerSession) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Supernets2.Contract.PendingStateTransitions(&_Supernets2.CallOpts, arg0)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Supernets2 *Supernets2Caller) RollupVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "rollupVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Supernets2 *Supernets2Session) RollupVerifier() (common.Address, error) {
	return _Supernets2.Contract.RollupVerifier(&_Supernets2.CallOpts)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Supernets2 *Supernets2CallerSession) RollupVerifier() (common.Address, error) {
	return _Supernets2.Contract.RollupVerifier(&_Supernets2.CallOpts)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Supernets2 *Supernets2Caller) SequencedBatches(opts *bind.CallOpts, arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "sequencedBatches", arg0)

	outstruct := new(struct {
		AccInputHash               [32]byte
		SequencedTimestamp         uint64
		PreviousLastBatchSequenced uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccInputHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.SequencedTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.PreviousLastBatchSequenced = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Supernets2 *Supernets2Session) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Supernets2.Contract.SequencedBatches(&_Supernets2.CallOpts, arg0)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Supernets2 *Supernets2CallerSession) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Supernets2.Contract.SequencedBatches(&_Supernets2.CallOpts, arg0)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Supernets2 *Supernets2Caller) TrustedAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "trustedAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Supernets2 *Supernets2Session) TrustedAggregator() (common.Address, error) {
	return _Supernets2.Contract.TrustedAggregator(&_Supernets2.CallOpts)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Supernets2 *Supernets2CallerSession) TrustedAggregator() (common.Address, error) {
	return _Supernets2.Contract.TrustedAggregator(&_Supernets2.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Caller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Session) TrustedAggregatorTimeout() (uint64, error) {
	return _Supernets2.Contract.TrustedAggregatorTimeout(&_Supernets2.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Supernets2.Contract.TrustedAggregatorTimeout(&_Supernets2.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Supernets2 *Supernets2Caller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Supernets2 *Supernets2Session) TrustedSequencer() (common.Address, error) {
	return _Supernets2.Contract.TrustedSequencer(&_Supernets2.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Supernets2 *Supernets2CallerSession) TrustedSequencer() (common.Address, error) {
	return _Supernets2.Contract.TrustedSequencer(&_Supernets2.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Supernets2 *Supernets2Caller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Supernets2 *Supernets2Session) TrustedSequencerURL() (string, error) {
	return _Supernets2.Contract.TrustedSequencerURL(&_Supernets2.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Supernets2 *Supernets2CallerSession) TrustedSequencerURL() (string, error) {
	return _Supernets2.Contract.TrustedSequencerURL(&_Supernets2.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Supernets2 *Supernets2Caller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Supernets2 *Supernets2Session) VerifyBatchTimeTarget() (uint64, error) {
	return _Supernets2.Contract.VerifyBatchTimeTarget(&_Supernets2.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Supernets2.Contract.VerifyBatchTimeTarget(&_Supernets2.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Supernets2 *Supernets2Transactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Supernets2 *Supernets2Session) AcceptAdminRole() (*types.Transaction, error) {
	return _Supernets2.Contract.AcceptAdminRole(&_Supernets2.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Supernets2 *Supernets2TransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Supernets2.Contract.AcceptAdminRole(&_Supernets2.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Supernets2 *Supernets2Transactor) ActivateEmergencyState(opts *bind.TransactOpts, sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "activateEmergencyState", sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Supernets2 *Supernets2Session) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.ActivateEmergencyState(&_Supernets2.TransactOpts, sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Supernets2 *Supernets2TransactorSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.ActivateEmergencyState(&_Supernets2.TransactOpts, sequencedBatchNum)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Supernets2 *Supernets2Transactor) ActivateForceBatches(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "activateForceBatches")
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Supernets2 *Supernets2Session) ActivateForceBatches() (*types.Transaction, error) {
	return _Supernets2.Contract.ActivateForceBatches(&_Supernets2.TransactOpts)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Supernets2 *Supernets2TransactorSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Supernets2.Contract.ActivateForceBatches(&_Supernets2.TransactOpts)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Supernets2 *Supernets2Transactor) ConsolidatePendingState(opts *bind.TransactOpts, pendingStateNum uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "consolidatePendingState", pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Supernets2 *Supernets2Session) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.ConsolidatePendingState(&_Supernets2.TransactOpts, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Supernets2 *Supernets2TransactorSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.ConsolidatePendingState(&_Supernets2.TransactOpts, pendingStateNum)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Supernets2 *Supernets2Transactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Supernets2 *Supernets2Session) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Supernets2.Contract.DeactivateEmergencyState(&_Supernets2.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Supernets2 *Supernets2TransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Supernets2.Contract.DeactivateEmergencyState(&_Supernets2.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Supernets2 *Supernets2Transactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "forceBatch", transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Supernets2 *Supernets2Session) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Supernets2.Contract.ForceBatch(&_Supernets2.TransactOpts, transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Supernets2 *Supernets2TransactorSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Supernets2.Contract.ForceBatch(&_Supernets2.TransactOpts, transactions, maticAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Supernets2 *Supernets2Transactor) Initialize(opts *bind.TransactOpts, initializePackedParameters Supernets2InitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "initialize", initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Supernets2 *Supernets2Session) Initialize(initializePackedParameters Supernets2InitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Supernets2.Contract.Initialize(&_Supernets2.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Supernets2 *Supernets2TransactorSession) Initialize(initializePackedParameters Supernets2InitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Supernets2.Contract.Initialize(&_Supernets2.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0xe6ad707e.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2Transactor) OverridePendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "overridePendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0xe6ad707e.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2Session) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.OverridePendingState(&_Supernets2.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0xe6ad707e.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2TransactorSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.OverridePendingState(&_Supernets2.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x0808270c.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2Transactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "proveNonDeterministicPendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x0808270c.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2Session) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.ProveNonDeterministicPendingState(&_Supernets2.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x0808270c.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2TransactorSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.ProveNonDeterministicPendingState(&_Supernets2.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Supernets2 *Supernets2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Supernets2 *Supernets2Session) RenounceOwnership() (*types.Transaction, error) {
	return _Supernets2.Contract.RenounceOwnership(&_Supernets2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Supernets2 *Supernets2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Supernets2.Contract.RenounceOwnership(&_Supernets2.TransactOpts)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes signaturesAndAddrs) returns()
func (_Supernets2 *Supernets2Transactor) SequenceBatches(opts *bind.TransactOpts, batches []Supernets2BatchData, l2Coinbase common.Address, signaturesAndAddrs []byte) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "sequenceBatches", batches, l2Coinbase, signaturesAndAddrs)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes signaturesAndAddrs) returns()
func (_Supernets2 *Supernets2Session) SequenceBatches(batches []Supernets2BatchData, l2Coinbase common.Address, signaturesAndAddrs []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.SequenceBatches(&_Supernets2.TransactOpts, batches, l2Coinbase, signaturesAndAddrs)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes signaturesAndAddrs) returns()
func (_Supernets2 *Supernets2TransactorSession) SequenceBatches(batches []Supernets2BatchData, l2Coinbase common.Address, signaturesAndAddrs []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.SequenceBatches(&_Supernets2.TransactOpts, batches, l2Coinbase, signaturesAndAddrs)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Supernets2 *Supernets2Transactor) SequenceForceBatches(opts *bind.TransactOpts, batches []Supernets2ForcedBatchData) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Supernets2 *Supernets2Session) SequenceForceBatches(batches []Supernets2ForcedBatchData) (*types.Transaction, error) {
	return _Supernets2.Contract.SequenceForceBatches(&_Supernets2.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Supernets2 *Supernets2TransactorSession) SequenceForceBatches(batches []Supernets2ForcedBatchData) (*types.Transaction, error) {
	return _Supernets2.Contract.SequenceForceBatches(&_Supernets2.TransactOpts, batches)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Supernets2 *Supernets2Transactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Supernets2 *Supernets2Session) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetForceBatchTimeout(&_Supernets2.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Supernets2 *Supernets2TransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetForceBatchTimeout(&_Supernets2.TransactOpts, newforceBatchTimeout)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Supernets2 *Supernets2Transactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Supernets2 *Supernets2Session) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Supernets2.Contract.SetMultiplierBatchFee(&_Supernets2.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Supernets2 *Supernets2TransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Supernets2.Contract.SetMultiplierBatchFee(&_Supernets2.TransactOpts, newMultiplierBatchFee)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Supernets2 *Supernets2Transactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Supernets2 *Supernets2Session) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetPendingStateTimeout(&_Supernets2.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Supernets2 *Supernets2TransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetPendingStateTimeout(&_Supernets2.TransactOpts, newPendingStateTimeout)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Supernets2 *Supernets2Transactor) SetTrustedAggregator(opts *bind.TransactOpts, newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setTrustedAggregator", newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Supernets2 *Supernets2Session) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedAggregator(&_Supernets2.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Supernets2 *Supernets2TransactorSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedAggregator(&_Supernets2.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Supernets2 *Supernets2Transactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Supernets2 *Supernets2Session) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedAggregatorTimeout(&_Supernets2.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Supernets2 *Supernets2TransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedAggregatorTimeout(&_Supernets2.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Supernets2 *Supernets2Transactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Supernets2 *Supernets2Session) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedSequencer(&_Supernets2.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Supernets2 *Supernets2TransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedSequencer(&_Supernets2.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Supernets2 *Supernets2Transactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Supernets2 *Supernets2Session) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedSequencerURL(&_Supernets2.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Supernets2 *Supernets2TransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedSequencerURL(&_Supernets2.TransactOpts, newTrustedSequencerURL)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Supernets2 *Supernets2Transactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Supernets2 *Supernets2Session) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetVerifyBatchTimeTarget(&_Supernets2.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Supernets2 *Supernets2TransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetVerifyBatchTimeTarget(&_Supernets2.TransactOpts, newVerifyBatchTimeTarget)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Supernets2 *Supernets2Transactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Supernets2 *Supernets2Session) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.TransferAdminRole(&_Supernets2.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Supernets2 *Supernets2TransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.TransferAdminRole(&_Supernets2.TransactOpts, newPendingAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Supernets2 *Supernets2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Supernets2 *Supernets2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.TransferOwnership(&_Supernets2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Supernets2 *Supernets2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.TransferOwnership(&_Supernets2.TransactOpts, newOwner)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x4fd70464.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2Transactor) VerifyBatches(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "verifyBatches", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x4fd70464.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2Session) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.VerifyBatches(&_Supernets2.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x4fd70464.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2TransactorSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.VerifyBatches(&_Supernets2.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0xa50a164b.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2Transactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "verifyBatchesTrustedAggregator", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0xa50a164b.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2Session) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.VerifyBatchesTrustedAggregator(&_Supernets2.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0xa50a164b.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes proof) returns()
func (_Supernets2 *Supernets2TransactorSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.VerifyBatchesTrustedAggregator(&_Supernets2.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// Supernets2AcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Supernets2 contract.
type Supernets2AcceptAdminRoleIterator struct {
	Event *Supernets2AcceptAdminRole // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2AcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2AcceptAdminRole)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2AcceptAdminRole)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2AcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2AcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2AcceptAdminRole represents a AcceptAdminRole event raised by the Supernets2 contract.
type Supernets2AcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Supernets2 *Supernets2Filterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*Supernets2AcceptAdminRoleIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &Supernets2AcceptAdminRoleIterator{contract: _Supernets2.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Supernets2 *Supernets2Filterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *Supernets2AcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2AcceptAdminRole)
				if err := _Supernets2.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Supernets2 *Supernets2Filterer) ParseAcceptAdminRole(log types.Log) (*Supernets2AcceptAdminRole, error) {
	event := new(Supernets2AcceptAdminRole)
	if err := _Supernets2.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2ActivateForceBatchesIterator is returned from FilterActivateForceBatches and is used to iterate over the raw logs and unpacked data for ActivateForceBatches events raised by the Supernets2 contract.
type Supernets2ActivateForceBatchesIterator struct {
	Event *Supernets2ActivateForceBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2ActivateForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2ActivateForceBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2ActivateForceBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2ActivateForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2ActivateForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2ActivateForceBatches represents a ActivateForceBatches event raised by the Supernets2 contract.
type Supernets2ActivateForceBatches struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivateForceBatches is a free log retrieval operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Supernets2 *Supernets2Filterer) FilterActivateForceBatches(opts *bind.FilterOpts) (*Supernets2ActivateForceBatchesIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return &Supernets2ActivateForceBatchesIterator{contract: _Supernets2.contract, event: "ActivateForceBatches", logs: logs, sub: sub}, nil
}

// WatchActivateForceBatches is a free log subscription operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Supernets2 *Supernets2Filterer) WatchActivateForceBatches(opts *bind.WatchOpts, sink chan<- *Supernets2ActivateForceBatches) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2ActivateForceBatches)
				if err := _Supernets2.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActivateForceBatches is a log parse operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Supernets2 *Supernets2Filterer) ParseActivateForceBatches(log types.Log) (*Supernets2ActivateForceBatches, error) {
	event := new(Supernets2ActivateForceBatches)
	if err := _Supernets2.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2ConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Supernets2 contract.
type Supernets2ConsolidatePendingStateIterator struct {
	Event *Supernets2ConsolidatePendingState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2ConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2ConsolidatePendingState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2ConsolidatePendingState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2ConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2ConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2ConsolidatePendingState represents a ConsolidatePendingState event raised by the Supernets2 contract.
type Supernets2ConsolidatePendingState struct {
	NumBatch        uint64
	StateRoot       [32]byte
	PendingStateNum uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterConsolidatePendingState is a free log retrieval operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Supernets2 *Supernets2Filterer) FilterConsolidatePendingState(opts *bind.FilterOpts, numBatch []uint64, pendingStateNum []uint64) (*Supernets2ConsolidatePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2ConsolidatePendingStateIterator{contract: _Supernets2.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Supernets2 *Supernets2Filterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *Supernets2ConsolidatePendingState, numBatch []uint64, pendingStateNum []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2ConsolidatePendingState)
				if err := _Supernets2.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsolidatePendingState is a log parse operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Supernets2 *Supernets2Filterer) ParseConsolidatePendingState(log types.Log) (*Supernets2ConsolidatePendingState, error) {
	event := new(Supernets2ConsolidatePendingState)
	if err := _Supernets2.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2EmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Supernets2 contract.
type Supernets2EmergencyStateActivatedIterator struct {
	Event *Supernets2EmergencyStateActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2EmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2EmergencyStateActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2EmergencyStateActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2EmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2EmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2EmergencyStateActivated represents a EmergencyStateActivated event raised by the Supernets2 contract.
type Supernets2EmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Supernets2 *Supernets2Filterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*Supernets2EmergencyStateActivatedIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &Supernets2EmergencyStateActivatedIterator{contract: _Supernets2.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Supernets2 *Supernets2Filterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *Supernets2EmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2EmergencyStateActivated)
				if err := _Supernets2.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyStateActivated is a log parse operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Supernets2 *Supernets2Filterer) ParseEmergencyStateActivated(log types.Log) (*Supernets2EmergencyStateActivated, error) {
	event := new(Supernets2EmergencyStateActivated)
	if err := _Supernets2.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2EmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Supernets2 contract.
type Supernets2EmergencyStateDeactivatedIterator struct {
	Event *Supernets2EmergencyStateDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2EmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2EmergencyStateDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2EmergencyStateDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2EmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2EmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2EmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Supernets2 contract.
type Supernets2EmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Supernets2 *Supernets2Filterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*Supernets2EmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &Supernets2EmergencyStateDeactivatedIterator{contract: _Supernets2.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Supernets2 *Supernets2Filterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *Supernets2EmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2EmergencyStateDeactivated)
				if err := _Supernets2.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyStateDeactivated is a log parse operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Supernets2 *Supernets2Filterer) ParseEmergencyStateDeactivated(log types.Log) (*Supernets2EmergencyStateDeactivated, error) {
	event := new(Supernets2EmergencyStateDeactivated)
	if err := _Supernets2.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2ForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Supernets2 contract.
type Supernets2ForceBatchIterator struct {
	Event *Supernets2ForceBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2ForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2ForceBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2ForceBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2ForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2ForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2ForceBatch represents a ForceBatch event raised by the Supernets2 contract.
type Supernets2ForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Supernets2 *Supernets2Filterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*Supernets2ForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2ForceBatchIterator{contract: _Supernets2.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Supernets2 *Supernets2Filterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *Supernets2ForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2ForceBatch)
				if err := _Supernets2.contract.UnpackLog(event, "ForceBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForceBatch is a log parse operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Supernets2 *Supernets2Filterer) ParseForceBatch(log types.Log) (*Supernets2ForceBatch, error) {
	event := new(Supernets2ForceBatch)
	if err := _Supernets2.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Supernets2 contract.
type Supernets2InitializedIterator struct {
	Event *Supernets2Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2Initialized represents a Initialized event raised by the Supernets2 contract.
type Supernets2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Supernets2 *Supernets2Filterer) FilterInitialized(opts *bind.FilterOpts) (*Supernets2InitializedIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Supernets2InitializedIterator{contract: _Supernets2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Supernets2 *Supernets2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Supernets2Initialized) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2Initialized)
				if err := _Supernets2.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Supernets2 *Supernets2Filterer) ParseInitialized(log types.Log) (*Supernets2Initialized, error) {
	event := new(Supernets2Initialized)
	if err := _Supernets2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2OverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Supernets2 contract.
type Supernets2OverridePendingStateIterator struct {
	Event *Supernets2OverridePendingState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2OverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2OverridePendingState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2OverridePendingState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2OverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2OverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2OverridePendingState represents a OverridePendingState event raised by the Supernets2 contract.
type Supernets2OverridePendingState struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOverridePendingState is a free log retrieval operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) FilterOverridePendingState(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*Supernets2OverridePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2OverridePendingStateIterator{contract: _Supernets2.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *Supernets2OverridePendingState, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2OverridePendingState)
				if err := _Supernets2.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOverridePendingState is a log parse operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) ParseOverridePendingState(log types.Log) (*Supernets2OverridePendingState, error) {
	event := new(Supernets2OverridePendingState)
	if err := _Supernets2.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Supernets2 contract.
type Supernets2OwnershipTransferredIterator struct {
	Event *Supernets2OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2OwnershipTransferred represents a OwnershipTransferred event raised by the Supernets2 contract.
type Supernets2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Supernets2 *Supernets2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Supernets2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2OwnershipTransferredIterator{contract: _Supernets2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Supernets2 *Supernets2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Supernets2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2OwnershipTransferred)
				if err := _Supernets2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Supernets2 *Supernets2Filterer) ParseOwnershipTransferred(log types.Log) (*Supernets2OwnershipTransferred, error) {
	event := new(Supernets2OwnershipTransferred)
	if err := _Supernets2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2ProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Supernets2 contract.
type Supernets2ProveNonDeterministicPendingStateIterator struct {
	Event *Supernets2ProveNonDeterministicPendingState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2ProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2ProveNonDeterministicPendingState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2ProveNonDeterministicPendingState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2ProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2ProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2ProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Supernets2 contract.
type Supernets2ProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Supernets2 *Supernets2Filterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*Supernets2ProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &Supernets2ProveNonDeterministicPendingStateIterator{contract: _Supernets2.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Supernets2 *Supernets2Filterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *Supernets2ProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2ProveNonDeterministicPendingState)
				if err := _Supernets2.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProveNonDeterministicPendingState is a log parse operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Supernets2 *Supernets2Filterer) ParseProveNonDeterministicPendingState(log types.Log) (*Supernets2ProveNonDeterministicPendingState, error) {
	event := new(Supernets2ProveNonDeterministicPendingState)
	if err := _Supernets2.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Supernets2 contract.
type Supernets2SequenceBatchesIterator struct {
	Event *Supernets2SequenceBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2SequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SequenceBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2SequenceBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2SequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SequenceBatches represents a SequenceBatches event raised by the Supernets2 contract.
type Supernets2SequenceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Supernets2 *Supernets2Filterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*Supernets2SequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2SequenceBatchesIterator{contract: _Supernets2.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Supernets2 *Supernets2Filterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *Supernets2SequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SequenceBatches)
				if err := _Supernets2.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequenceBatches is a log parse operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Supernets2 *Supernets2Filterer) ParseSequenceBatches(log types.Log) (*Supernets2SequenceBatches, error) {
	event := new(Supernets2SequenceBatches)
	if err := _Supernets2.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Supernets2 contract.
type Supernets2SequenceForceBatchesIterator struct {
	Event *Supernets2SequenceForceBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2SequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SequenceForceBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2SequenceForceBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2SequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SequenceForceBatches represents a SequenceForceBatches event raised by the Supernets2 contract.
type Supernets2SequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Supernets2 *Supernets2Filterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*Supernets2SequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2SequenceForceBatchesIterator{contract: _Supernets2.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Supernets2 *Supernets2Filterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *Supernets2SequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SequenceForceBatches)
				if err := _Supernets2.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequenceForceBatches is a log parse operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Supernets2 *Supernets2Filterer) ParseSequenceForceBatches(log types.Log) (*Supernets2SequenceForceBatches, error) {
	event := new(Supernets2SequenceForceBatches)
	if err := _Supernets2.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Supernets2 contract.
type Supernets2SetForceBatchTimeoutIterator struct {
	Event *Supernets2SetForceBatchTimeout // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2SetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetForceBatchTimeout)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2SetForceBatchTimeout)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2SetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Supernets2 contract.
type Supernets2SetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Supernets2 *Supernets2Filterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*Supernets2SetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetForceBatchTimeoutIterator{contract: _Supernets2.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Supernets2 *Supernets2Filterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *Supernets2SetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetForceBatchTimeout)
				if err := _Supernets2.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetForceBatchTimeout is a log parse operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Supernets2 *Supernets2Filterer) ParseSetForceBatchTimeout(log types.Log) (*Supernets2SetForceBatchTimeout, error) {
	event := new(Supernets2SetForceBatchTimeout)
	if err := _Supernets2.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Supernets2 contract.
type Supernets2SetMultiplierBatchFeeIterator struct {
	Event *Supernets2SetMultiplierBatchFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2SetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetMultiplierBatchFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2SetMultiplierBatchFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2SetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Supernets2 contract.
type Supernets2SetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Supernets2 *Supernets2Filterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*Supernets2SetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetMultiplierBatchFeeIterator{contract: _Supernets2.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Supernets2 *Supernets2Filterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *Supernets2SetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetMultiplierBatchFee)
				if err := _Supernets2.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMultiplierBatchFee is a log parse operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Supernets2 *Supernets2Filterer) ParseSetMultiplierBatchFee(log types.Log) (*Supernets2SetMultiplierBatchFee, error) {
	event := new(Supernets2SetMultiplierBatchFee)
	if err := _Supernets2.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Supernets2 contract.
type Supernets2SetPendingStateTimeoutIterator struct {
	Event *Supernets2SetPendingStateTimeout // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2SetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetPendingStateTimeout)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2SetPendingStateTimeout)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2SetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Supernets2 contract.
type Supernets2SetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Supernets2 *Supernets2Filterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*Supernets2SetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetPendingStateTimeoutIterator{contract: _Supernets2.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Supernets2 *Supernets2Filterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *Supernets2SetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetPendingStateTimeout)
				if err := _Supernets2.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPendingStateTimeout is a log parse operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Supernets2 *Supernets2Filterer) ParseSetPendingStateTimeout(log types.Log) (*Supernets2SetPendingStateTimeout, error) {
	event := new(Supernets2SetPendingStateTimeout)
	if err := _Supernets2.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Supernets2 contract.
type Supernets2SetTrustedAggregatorIterator struct {
	Event *Supernets2SetTrustedAggregator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2SetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetTrustedAggregator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2SetTrustedAggregator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2SetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetTrustedAggregator represents a SetTrustedAggregator event raised by the Supernets2 contract.
type Supernets2SetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Supernets2 *Supernets2Filterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*Supernets2SetTrustedAggregatorIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetTrustedAggregatorIterator{contract: _Supernets2.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Supernets2 *Supernets2Filterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *Supernets2SetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetTrustedAggregator)
				if err := _Supernets2.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedAggregator is a log parse operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Supernets2 *Supernets2Filterer) ParseSetTrustedAggregator(log types.Log) (*Supernets2SetTrustedAggregator, error) {
	event := new(Supernets2SetTrustedAggregator)
	if err := _Supernets2.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Supernets2 contract.
type Supernets2SetTrustedAggregatorTimeoutIterator struct {
	Event *Supernets2SetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2SetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetTrustedAggregatorTimeout)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2SetTrustedAggregatorTimeout)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2SetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Supernets2 contract.
type Supernets2SetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Supernets2 *Supernets2Filterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*Supernets2SetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetTrustedAggregatorTimeoutIterator{contract: _Supernets2.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Supernets2 *Supernets2Filterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *Supernets2SetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetTrustedAggregatorTimeout)
				if err := _Supernets2.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedAggregatorTimeout is a log parse operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Supernets2 *Supernets2Filterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*Supernets2SetTrustedAggregatorTimeout, error) {
	event := new(Supernets2SetTrustedAggregatorTimeout)
	if err := _Supernets2.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Supernets2 contract.
type Supernets2SetTrustedSequencerIterator struct {
	Event *Supernets2SetTrustedSequencer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2SetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetTrustedSequencer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2SetTrustedSequencer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2SetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetTrustedSequencer represents a SetTrustedSequencer event raised by the Supernets2 contract.
type Supernets2SetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Supernets2 *Supernets2Filterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*Supernets2SetTrustedSequencerIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetTrustedSequencerIterator{contract: _Supernets2.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Supernets2 *Supernets2Filterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *Supernets2SetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetTrustedSequencer)
				if err := _Supernets2.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedSequencer is a log parse operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Supernets2 *Supernets2Filterer) ParseSetTrustedSequencer(log types.Log) (*Supernets2SetTrustedSequencer, error) {
	event := new(Supernets2SetTrustedSequencer)
	if err := _Supernets2.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Supernets2 contract.
type Supernets2SetTrustedSequencerURLIterator struct {
	Event *Supernets2SetTrustedSequencerURL // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2SetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetTrustedSequencerURL)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2SetTrustedSequencerURL)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2SetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Supernets2 contract.
type Supernets2SetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Supernets2 *Supernets2Filterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*Supernets2SetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetTrustedSequencerURLIterator{contract: _Supernets2.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Supernets2 *Supernets2Filterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *Supernets2SetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetTrustedSequencerURL)
				if err := _Supernets2.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedSequencerURL is a log parse operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Supernets2 *Supernets2Filterer) ParseSetTrustedSequencerURL(log types.Log) (*Supernets2SetTrustedSequencerURL, error) {
	event := new(Supernets2SetTrustedSequencerURL)
	if err := _Supernets2.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Supernets2 contract.
type Supernets2SetVerifyBatchTimeTargetIterator struct {
	Event *Supernets2SetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2SetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetVerifyBatchTimeTarget)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2SetVerifyBatchTimeTarget)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2SetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Supernets2 contract.
type Supernets2SetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Supernets2 *Supernets2Filterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*Supernets2SetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetVerifyBatchTimeTargetIterator{contract: _Supernets2.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Supernets2 *Supernets2Filterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *Supernets2SetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetVerifyBatchTimeTarget)
				if err := _Supernets2.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetVerifyBatchTimeTarget is a log parse operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Supernets2 *Supernets2Filterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*Supernets2SetVerifyBatchTimeTarget, error) {
	event := new(Supernets2SetVerifyBatchTimeTarget)
	if err := _Supernets2.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2TransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Supernets2 contract.
type Supernets2TransferAdminRoleIterator struct {
	Event *Supernets2TransferAdminRole // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2TransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2TransferAdminRole)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2TransferAdminRole)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2TransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2TransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2TransferAdminRole represents a TransferAdminRole event raised by the Supernets2 contract.
type Supernets2TransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Supernets2 *Supernets2Filterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*Supernets2TransferAdminRoleIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &Supernets2TransferAdminRoleIterator{contract: _Supernets2.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Supernets2 *Supernets2Filterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *Supernets2TransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2TransferAdminRole)
				if err := _Supernets2.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Supernets2 *Supernets2Filterer) ParseTransferAdminRole(log types.Log) (*Supernets2TransferAdminRole, error) {
	event := new(Supernets2TransferAdminRole)
	if err := _Supernets2.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2UpdateSupernets2VersionIterator is returned from FilterUpdateSupernets2Version and is used to iterate over the raw logs and unpacked data for UpdateSupernets2Version events raised by the Supernets2 contract.
type Supernets2UpdateSupernets2VersionIterator struct {
	Event *Supernets2UpdateSupernets2Version // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2UpdateSupernets2VersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2UpdateSupernets2Version)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2UpdateSupernets2Version)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2UpdateSupernets2VersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2UpdateSupernets2VersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2UpdateSupernets2Version represents a UpdateSupernets2Version event raised by the Supernets2 contract.
type Supernets2UpdateSupernets2Version struct {
	NumBatch uint64
	ForkID   uint64
	Version  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateSupernets2Version is a free log retrieval operation binding the contract event 0x13e60734641bd97ab6e444f419277d9983e2c3eab9bb631a578ffa07756b6351.
//
// Solidity: event UpdateSupernets2Version(uint64 numBatch, uint64 forkID, string version)
func (_Supernets2 *Supernets2Filterer) FilterUpdateSupernets2Version(opts *bind.FilterOpts) (*Supernets2UpdateSupernets2VersionIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "UpdateSupernets2Version")
	if err != nil {
		return nil, err
	}
	return &Supernets2UpdateSupernets2VersionIterator{contract: _Supernets2.contract, event: "UpdateSupernets2Version", logs: logs, sub: sub}, nil
}

// WatchUpdateSupernets2Version is a free log subscription operation binding the contract event 0x13e60734641bd97ab6e444f419277d9983e2c3eab9bb631a578ffa07756b6351.
//
// Solidity: event UpdateSupernets2Version(uint64 numBatch, uint64 forkID, string version)
func (_Supernets2 *Supernets2Filterer) WatchUpdateSupernets2Version(opts *bind.WatchOpts, sink chan<- *Supernets2UpdateSupernets2Version) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "UpdateSupernets2Version")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2UpdateSupernets2Version)
				if err := _Supernets2.contract.UnpackLog(event, "UpdateSupernets2Version", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateSupernets2Version is a log parse operation binding the contract event 0x13e60734641bd97ab6e444f419277d9983e2c3eab9bb631a578ffa07756b6351.
//
// Solidity: event UpdateSupernets2Version(uint64 numBatch, uint64 forkID, string version)
func (_Supernets2 *Supernets2Filterer) ParseUpdateSupernets2Version(log types.Log) (*Supernets2UpdateSupernets2Version, error) {
	event := new(Supernets2UpdateSupernets2Version)
	if err := _Supernets2.contract.UnpackLog(event, "UpdateSupernets2Version", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2VerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Supernets2 contract.
type Supernets2VerifyBatchesIterator struct {
	Event *Supernets2VerifyBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2VerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2VerifyBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2VerifyBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2VerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2VerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2VerifyBatches represents a VerifyBatches event raised by the Supernets2 contract.
type Supernets2VerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*Supernets2VerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2VerifyBatchesIterator{contract: _Supernets2.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *Supernets2VerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2VerifyBatches)
				if err := _Supernets2.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifyBatches is a log parse operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) ParseVerifyBatches(log types.Log) (*Supernets2VerifyBatches, error) {
	event := new(Supernets2VerifyBatches)
	if err := _Supernets2.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2VerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Supernets2 contract.
type Supernets2VerifyBatchesTrustedAggregatorIterator struct {
	Event *Supernets2VerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Supernets2VerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2VerifyBatchesTrustedAggregator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Supernets2VerifyBatchesTrustedAggregator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Supernets2VerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2VerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2VerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Supernets2 contract.
type Supernets2VerifyBatchesTrustedAggregator struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatchesTrustedAggregator is a free log retrieval operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*Supernets2VerifyBatchesTrustedAggregatorIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2VerifyBatchesTrustedAggregatorIterator{contract: _Supernets2.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *Supernets2VerifyBatchesTrustedAggregator, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2VerifyBatchesTrustedAggregator)
				if err := _Supernets2.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifyBatchesTrustedAggregator is a log parse operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*Supernets2VerifyBatchesTrustedAggregator, error) {
	event := new(Supernets2VerifyBatchesTrustedAggregator)
	if err := _Supernets2.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
