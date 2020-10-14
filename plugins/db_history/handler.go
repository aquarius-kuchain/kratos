package dbHistory

import (
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/plugins/base/db"
	"github.com/KuChainNetwork/kuchain/plugins/db_history/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (t *plugin) OnEvent(ctx types.Context, evt types.Event) {
	t.logger.Info("on event", "type", evt.Type)

	t.db.Emit(db.NewWork(evt))
}

func (t *plugin) OnTx(ctx types.Context, tx chainTypes.StdTx) {
	t.logger.Info("on tx", "tx", tx)

	t.db.Emit(db.NewWork(tx))
}

func (t *plugin) OnMsg(ctx types.Context, msg sdk.Msg) {
	t.logger.Info("on msg", "msg", msg)

	t.db.Emit(db.NewWork(msg))
}
