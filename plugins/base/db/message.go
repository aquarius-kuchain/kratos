package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

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
