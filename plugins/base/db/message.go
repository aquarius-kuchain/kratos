package db

import (
	"reflect"

	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

type eventInDB struct {
	tableName struct{} `pg:"events,alias:events"` // default values are the same

	ID   int // both "Id" and "ID" are detected as primary key
	Type string

	Attributes map[string]string
}

type messageInDB struct {
	tableName struct{} `pg:"messages,alias:messages"` // default values are the same

	ID int64 // both "Id" and "ID" are detected as primary key

	Msg // FIXME: use petty show msg
}

type kuTransferInDB struct {
	tableName struct{} `pg:"transfer,alias:transfer"` // default values are the same

	ID     int64 // both "Id" and "ID" are detected as primary key
	Route  string
	Type   string
	From   string
	To     string
	Amount int64
	Symbol string
}

type txInDB struct {
	tableName struct{} `pg:"tx,alias:tx"` // default values are the same

	ID int64 // both "Id" and "ID" are detected as primary key

	StdTx
}

// SyncState sync state in pg database
type syncState struct {
	tableName struct{} `pg:"sync_stat,alias:sync_stat"` // default values are the same

	ID       int // both "Id" and "ID" are detected as primary key
	BlockNum int64
	ChainID  string `pg:",unique"`
}

func InsertEvent(db *pg.DB, logger Logger, evt *Event) error {
	return db.Insert(&eventInDB{
		Type:       evt.Type,
		Attributes: evt.Attributes,
	})
}

func newTxInDB(tx StdTx) *txInDB {
	return &txInDB{
		StdTx: tx,
	}
}

func newMsgToDB(msg Msg) *messageInDB {
	return &messageInDB{
		Msg: msg,
	}
}

func processMsg(db *pg.DB, msg Msg) error {
	if err := db.Insert(newMsgToDB(msg)); err != nil {
		return errors.Wrapf(err, "insert msg")
	}

	if msg, ok := msg.(KuTransfMsg); ok {
		transfers := msg.GetTransfers()

		for _, t := range transfers {
			amounts := t.Amount

			in := &kuTransferInDB{
				Route: msg.Route(),
				Type:  msg.Type(),
				From:  t.From.String(),
				To:    t.To.String(),
			}

			for _, amount := range amounts {
				in.Amount = amount.Amount.BigInt().Int64()
				in.Symbol = amount.Denom
				if err := db.Insert(in); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func syncChainStat(db *pg.DB, logger Logger, num int64, chainID string) error {
	stat := &syncState{
		ID: ChainIdx,
	}
	err := db.Select(stat)
	if err != nil {
		return errors.Wrapf(err, "get sync stat err")
	}

	logger.Info("get sync stat in %d", stat.BlockNum)

	return db.Update(syncState{
		ID:       ChainIdx,
		BlockNum: num,
		ChainID:  chainID,
	})
}

func Process(db *pg.DB, logger Logger, msg interface{}) error {
	logger.Debug("process msg", "typ", reflect.TypeOf(msg), "msg", msg)
	switch msg := msg.(type) {
	case Event:
		return InsertEvent(db, logger, &msg)
	case StdTx:
		return db.Insert(newTxInDB(msg))
	}

	if msg, ok := msg.(Msg); ok {
		return processMsg(db, msg)
	}

	return nil
}
