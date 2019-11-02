package coinswap

import (
	"github.com/PirosB3/coinswap/incubator/coinswap/internal/keeper"
	"github.com/PirosB3/coinswap/incubator/coinswap/internal/types"
)

type (
	Keeper             = keeper.Keeper
	MsgSwapOrder       = types.MsgSwapOrder
	MsgAddLiquidity    = types.MsgAddLiquidity
	MsgRemoveLiquidity = types.MsgRemoveLiquidity
)

var (
	ErrInvalidDeadline  = types.ErrInvalidDeadline
	ErrNotPositive      = types.ErrNotPositive
	ErrConstraintNotMet = types.ErrConstraintNotMet
)

const (
	DefaultCodespace = types.DefaultCodespace
	ModuleName       = types.ModuleName
)
