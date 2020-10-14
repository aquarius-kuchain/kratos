package eventDB

import (
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/plugins/db_event_history/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

// plugin for test
type plugin struct {
	logger log.Logger
}

func (t *plugin) Init(ctx types.Context) error {
	t.logger.Info("plugin init", "name", types.PluginName)

	return nil
}

func (t *plugin) Start(ctx types.Context) error {
	t.logger.Info("plugin start", "name", types.PluginName)

	return nil
}

func (t *plugin) Stop(ctx types.Context) error {
	t.logger.Info("plugin stop", "name", types.PluginName)

	return nil
}

func (t *plugin) OnEvent(ctx types.Context, evt types.Event) {
	t.logger.Info("on event", "type", evt.Type)
}

func (t *plugin) OnTx(ctx types.Context, tx chainTypes.StdTx) {
	t.logger.Info("on tx", "tx", tx)
}

func (t *plugin) OnMsg(ctx types.Context, msg sdk.Msg) {
	t.logger.Info("on msg", "msg", msg)
}

func (t *plugin) MsgHandler() types.PluginMsgHandler {
	return func(ctx types.Context, msg sdk.Msg) {
		t.OnMsg(ctx, msg)
	}
}

func (t *plugin) TxHandler() types.PluginTxHandler {
	return func(ctx types.Context, tx chainTypes.StdTx) {
		t.OnTx(ctx, tx)
	}
}

func (t *plugin) EvtHandler() types.PluginEvtHandler {
	return func(ctx types.Context, evt types.Event) {
		t.OnEvent(ctx, evt)
	}
}

func (t *plugin) Logger() log.Logger {
	return t.logger
}

func (t *plugin) Name() string {
	return types.PluginName
}

// New new test plugin
func New(ctx types.Context, cfg types.BaseCfg) *plugin {
	return &plugin{
		logger: ctx.Logger().With("module", types.PluginName),
	}
}
