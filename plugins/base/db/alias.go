package db

import (
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/plugins/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Event       = types.Event
	KuTransfMsg = chainTypes.KuTransfMsg
	StdTx       = chainTypes.StdTx
	Msg         = sdk.Msg
	Logger      = log.Logger
)
