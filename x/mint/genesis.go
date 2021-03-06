package mint

import (
	"github.com/KuChainNetwork/kuchain/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis new mint genesis
func InitGenesis(ctx sdk.Context, keeper Keeper, supplyKeeper types.SupplyKeeper, data GenesisState) {
	keeper.SetMinter(ctx, data.Minter)
	keeper.SetParams(ctx, data.Params)

	supplyKeeper.GetModuleAccount(ctx, ModuleName)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	minter := keeper.GetMinter(ctx)
	params := keeper.GetParams(ctx)
	return NewGenesisState(minter, params)
}
