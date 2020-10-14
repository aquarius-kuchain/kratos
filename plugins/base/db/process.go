package db

import (
	"reflect"

	"github.com/go-pg/pg/v10"
)

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
