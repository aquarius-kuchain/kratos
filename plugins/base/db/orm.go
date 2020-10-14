package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func regOrm(db *pg.DB) error {
	models := []interface{}{
		(*eventInDB)(nil),
		(*txInDB)(nil),
		(*messageInDB)(nil),
		(*kuTransferInDB)(nil),
	}

	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
