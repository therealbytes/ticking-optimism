package derive

import (
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	TickFuncSignature    = "tick()"
	TickFuncBytes4       = crypto.Keccak256([]byte(TickFuncSignature))[:4]
	TickDepositerAddress = common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001")
	TickAddress          = predeploys.TickAddr
)

var TickGasLimit uint64 = 150_000_000

func MarshalBinary() ([]byte, error) {
	data := make([]byte, 4)
	offset := 0
	copy(data[offset:4], TickFuncBytes4)
	return data, nil
}

// TickDeposit creates a tick deposit transaction.
func TickDeposit(seqNumber uint64, block eth.BlockInfo) (*types.DepositTx, error) {
	data, err := MarshalBinary()
	if err != nil {
		return nil, err
	}
	source := L1InfoDepositSource{
		L1BlockHash: block.Hash(),
		SeqNumber:   seqNumber,
	}
	return &types.DepositTx{
		SourceHash:          source.SourceHash(),
		From:                TickDepositerAddress,
		To:                  &TickAddress,
		Mint:                nil,
		Value:               big.NewInt(0),
		Gas:                 TickGasLimit,
		IsSystemTransaction: true,
		Data:                data,
	}, nil
}

// TickDepositBytes returns a serialized tick transaction.
func TickDepositBytes(seqNumber uint64, Tick eth.BlockInfo) ([]byte, error) {
	dep, err := TickDeposit(seqNumber, Tick)
	if err != nil {
		return nil, fmt.Errorf("failed to create L1 info tx: %w", err)
	}
	l1Tx := types.NewTx(dep)
	opaqueL1Tx, err := l1Tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to encode L1 info tx: %w", err)
	}
	return opaqueL1Tx, nil
}
