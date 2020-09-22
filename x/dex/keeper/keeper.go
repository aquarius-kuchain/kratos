package keeper

import (
	"fmt"

	assetKeeper "github.com/KuChainNetwork/kuchain/x/asset/keeper"
	"github.com/KuChainNetwork/kuchain/x/dex/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

// DexKeeper for asset state
type DexKeeper struct {
	// The (unexposed) key used to access the store from the Context.
	key sdk.StoreKey
	// The codec codec for binary encoding/decoding of accounts.
	cdc         *codec.Codec
	assetKeeper assetKeeper.AssetKeeper
}

// NewDexKeeper new keeper for a dex
func NewDexKeeper(cdc *codec.Codec,
	key sdk.StoreKey,
	keeper assetKeeper.AssetKeeper) DexKeeper {
	return DexKeeper{
		key:         key,
		cdc:         cdc,
		assetKeeper: keeper,
	}
}

// Logger returns a module-specific logger.
func (a DexKeeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetDex get dex data, if no found, return false
func (a DexKeeper) GetDex(ctx sdk.Context, creator types.Name) (*types.Dex, bool) {
	return a.getDex(ctx, creator)
}
