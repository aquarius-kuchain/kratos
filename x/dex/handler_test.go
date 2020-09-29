package dex_test

import (
	"testing"

	"github.com/tendermint/tendermint/libs/rand"

	"github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/test/simapp"
	dexTypes "github.com/KuChainNetwork/kuchain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/smartystreets/goconvey/convey"
)

func CreateDexForTest(t *testing.T, app *simapp.SimApp, isSuccess bool, account types.AccountID, stakings types.Coins, desc []byte) error {
	ctx := app.NewTestContext()
	wallet := app.GetWallet()

	var (
		acc     = account
		accName = account.MustName()
		auth    = wallet.GetAuth(acc)
	)

	msg := dexTypes.NewMsgCreateDex(
		auth,
		accName,
		stakings,
		desc)

	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	tx := simapp.NewTxForTest(
		acc,
		[]sdk.Msg{
			&msg,
		}, wallet.PrivKey(auth))

	if !isSuccess {
		tx = tx.WithCannotPass()
	}

	return simapp.CheckTxs(t, app, ctx, tx)
}

func TestHandleCreateDex(t *testing.T) {
	app, _ := createAppForTest()

	Convey("test create dex", t, func() {
		ctx := app.NewTestContext()

		var (
			acc     = account5
			accName = name5
			auth    = wallet.GetAuth(acc)
		)

		msg := dexTypes.NewMsgCreateDex(
			auth,
			accName,
			types.NewInt64CoreCoins(1000000),
			[]byte("dex for test"))

		So(msg.ValidateBasic(), ShouldBeNil)

		tx := simapp.NewTxForTest(
			acc,
			[]sdk.Msg{
				&msg,
			}, wallet.PrivKey(auth))

		err := simapp.CheckTxs(t, app, ctx, tx)
		So(err, ShouldBeNil)
	})

	Convey("test transfer no enough", t, func() {
		ctx := app.NewTestContext()

		var (
			acc     = account4
			accName = name4
			auth    = wallet.GetAuth(acc)
		)

		msg := dexTypes.NewMsgCreateDex(
			auth,
			accName,
			types.NewInt64CoreCoins(1000000),
			[]byte("dex for test"))

		msg.Transfers[0].Amount[0].Amount = types.NewInt(111)

		So(msg.ValidateBasic(), simapp.ShouldErrIs, types.ErrKuMsgFromNotEqual)

		tx := simapp.NewTxForTest(
			acc,
			[]sdk.Msg{
				&msg,
			}, wallet.PrivKey(auth)).WithCannotPass()

		err := simapp.CheckTxs(t, app, ctx, tx)
		So(err, simapp.ShouldErrIs, types.ErrKuMsgFromNotEqual)
	})

	Convey("test transfer no module", t, func() {
		ctx := app.NewTestContext()

		var (
			acc     = account4
			accName = name4
			auth    = wallet.GetAuth(acc)
		)

		msg := dexTypes.NewMsgCreateDex(
			auth,
			accName,
			types.NewInt64CoreCoins(1000000),
			[]byte("dex for test"))

		msg.Transfers[0].To = acc

		So(msg.ValidateBasic(), simapp.ShouldErrIs, types.ErrKuMsgFromNotEqual)

		tx := simapp.NewTxForTest(
			acc,
			[]sdk.Msg{
				&msg,
			}, wallet.PrivKey(auth)).WithCannotPass()

		err := simapp.CheckTxs(t, app, ctx, tx)
		So(err, simapp.ShouldErrIs, types.ErrKuMsgFromNotEqual)
	})

	Convey("test transfer from error", t, func() {
		ctx := app.NewTestContext()

		var (
			acc     = account4
			accName = name4
			auth    = wallet.GetAuth(acc)
		)

		msg := dexTypes.NewMsgCreateDex(
			auth,
			accName,
			types.NewInt64CoreCoins(1000000),
			[]byte("dex for test"))

		msg.Transfers[0].From = account5

		So(msg.ValidateBasic(), simapp.ShouldErrIs, types.ErrKuMsgFromNotEqual)

		tx := simapp.NewTxForTest(
			acc,
			[]sdk.Msg{
				&msg,
			}, wallet.PrivKey(auth)).WithCannotPass()

		err := simapp.CheckTxs(t, app, ctx, tx)
		So(err, simapp.ShouldErrIs, types.ErrKuMsgFromNotEqual)
	})

	Convey("test transfer desc too long", t, func() {
		So(CreateDexForTest(t, app,
			false,
			account4, types.NewInt64CoreCoins(111),
			make([]byte, 520)), simapp.ShouldErrIs, dexTypes.ErrDexDescTooLong)
	})

	Convey("test transfer create two times", t, func() {
		So(CreateDexForTest(t, app,
			false,
			account5, types.NewInt64CoreCoins(111),
			[]byte("hello")), simapp.ShouldErrIs, dexTypes.ErrDexHadCreated)
	})
}

func TestHandleCreateDexNumber(t *testing.T) {
	app, _ := createAppForTest()

	Convey("test create dex number", t, func() {
		So(CreateDexForTest(t, app,
			true,
			account4, types.NewInt64CoreCoins(111),
			[]byte("account4")), ShouldBeNil)

		So(CreateDexForTest(t, app,
			true,
			account5, types.NewInt64CoreCoins(111),
			[]byte("account5")), ShouldBeNil)

		simapp.AfterBlockCommitted(app, 1)

		ctx := app.NewTestContext()
		dex1, ok := app.DexKeeper().GetDex(ctx, name4)

		So(dex1, ShouldNotBeNil)
		So(ok, ShouldBeTrue)
		So(dex1.Creator, simapp.ShouldEq, name4)
		So(dex1.Number, ShouldEqual, 0)

		dex2, ok := app.DexKeeper().GetDex(ctx, name5)

		So(dex2, ShouldNotBeNil)
		So(ok, ShouldBeTrue)
		So(dex2.Creator, simapp.ShouldEq, name5)
		So(dex2.Number, ShouldEqual, 1)
	})
}

func TestHandleUpdateDexDescription(t *testing.T) {
	app, _ := createAppForTest()

	Convey("test update dex description", t, func() {
		var (
			acc     = account4
			accName = name4
			auth    = wallet.GetAuth(acc)
		)

		So(CreateDexForTest(t, app,
			true,
			acc, types.NewInt64CoreCoins(111),
			[]byte("account4")), ShouldBeNil)

		simapp.AfterBlockCommitted(app, 1)

		ctx := app.NewTestContext()
		dex, ok := app.DexKeeper().GetDex(ctx, accName)

		So(dex, ShouldNotBeNil)
		So(ok, ShouldBeTrue)
		So(dex.Creator, simapp.ShouldEq, accName)

		msgUpdateDexDescription := dexTypes.NewMsgUpdateDexDescription(auth,
			accName,
			[]byte("xxx.yyy.zzz"))
		So(msgUpdateDexDescription.ValidateBasic(), ShouldBeNil)

		tx := simapp.NewTxForTest(
			acc,
			[]sdk.Msg{
				&msgUpdateDexDescription,
			}, wallet.PrivKey(auth))
		ctx = app.NewTestContext()
		err := simapp.CheckTxs(t, app, ctx, tx)
		So(err, ShouldBeNil)

		simapp.AfterBlockCommitted(app, 1)

		ctx = app.NewTestContext()
		dex, ok = app.DexKeeper().GetDex(ctx, accName)
		So(ok, ShouldBeTrue)
		So(dex, ShouldNotBeNil)
		So(dex.Description, ShouldEqual, "xxx.yyy.zzz")

		msgUpdateDexDescription = dexTypes.NewMsgUpdateDexDescription(auth,
			accName,
			[]byte(rand.Str(dexTypes.MaxDexDescriptorLen)))
		So(msgUpdateDexDescription.ValidateBasic(), ShouldNotBeNil)
	})
}

func TestHandleDestroyDex(t *testing.T) {
	app, _ := createAppForTest()

	Convey("test destroy dex", t, func() {
		var (
			acc     = account4
			accName = name4
			auth    = wallet.GetAuth(acc)
		)

		So(CreateDexForTest(t, app,
			true,
			acc, types.NewInt64CoreCoins(1000),
			[]byte("account4")), ShouldBeNil)

		simapp.AfterBlockCommitted(app, 1)

		ctx := app.NewTestContext()
		dex, ok := app.DexKeeper().GetDex(ctx, accName)
		So(dex, ShouldNotBeNil)
		So(ok, ShouldBeTrue)
		So(dex.Creator, simapp.ShouldEq, accName)
		So(dex.Number, ShouldEqual, 0)

		coins := app.AssetKeeper().GetCoinPowers(ctx, acc)
		So(coins.IsZero(), ShouldBeTrue)

		msgDestroyDex := dexTypes.NewMsgDestroyDex(
			auth,
			accName)
		So(msgDestroyDex.ValidateBasic(), ShouldBeNil)

		tx := simapp.NewTxForTest(
			acc,
			[]sdk.Msg{
				&msgDestroyDex,
			}, wallet.PrivKey(auth))
		ctx = app.NewTestContext()
		err := simapp.CheckTxs(t, app, ctx, tx)
		So(err, ShouldBeNil)

		simapp.AfterBlockCommitted(app, 1)

		ctx = app.NewTestContext()
		dex, ok = app.DexKeeper().GetDex(ctx, accName)
		So(ok, ShouldBeFalse)
		So(dex, ShouldBeNil)
		coins = app.AssetKeeper().GetCoinPowers(ctx, acc)
		So(coins.IsEqual(types.NewInt64CoreCoins(1000)), ShouldBeTrue)
	})
}
