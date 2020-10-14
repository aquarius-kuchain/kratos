package db

import (
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/pkg/errors"
)

const (
	ChainIdx = 1
)

// DBCfg cfg for database connect
type DBCfg struct {
	Address  string `json:"address"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Cfg struct {
	DB DBCfg `json:"db"`
}

// SyncState sync state in pg database
type SyncState struct {
	tableName struct{} `pg:"sync_stat,alias:sync_stat"` // default values are the same

	ID       int // both "Id" and "ID" are detected as primary key
	BlockNum int64
	ChainID  string `pg:",unique"`
}

type dbWork struct {
	msg interface{}
}

func (w dbWork) IsStopped() bool {
	return w.msg == nil
}

type dbService struct {
	logger   Logger
	database *pg.DB

	dbChan chan dbWork
	wg     sync.WaitGroup
}

// NewDB create a connection commit event to db
func NewDB(cfg Cfg, logger Logger) *dbService {
	return &dbService{
		database: pg.Connect(&pg.Options{
			Addr:     cfg.DB.Address,
			User:     cfg.DB.User,
			Password: cfg.DB.Password,
			Database: cfg.DB.Database,
		}),
		logger: logger,
		dbChan: make(chan dbWork, 512),
	}
}

func (db *dbService) Start() error {
	db.logger.Info("Starting database service")

	db.wg.Add(1)
	go func() {
		defer db.wg.Done()

		if err := createSchema(db.database); err != nil {
			panic(err)
		}

		for {
			work, ok := <-db.dbChan
			if !ok {
				db.logger.Info("msg channel closed")
				return
			}

			if work.IsStopped() {
				db.logger.Info("db service stopped")
				return
			}

			if err := db.Process(&work); err != nil {
				db.logger.Error("db process error", "err", err)
			}
		}
	}()
	return nil
}

func (db *dbService) Process(work *dbWork) error {
	if err := Process(db.database, db.logger, work.msg); err != nil {
		return err
	}

	return nil
}

func (db *dbService) Emit(work dbWork) {
	db.dbChan <- work
}

func (db *dbService) Stop() error {
	db.logger.Info("Stopping database service")

	db.dbChan <- dbWork{}
	db.wg.Wait()

	db.logger.Info("Database service stopped")

	db.database.Close()

	db.logger.Info("Database connection closed")
	return nil
}

func (db *dbService) OnEvent(evt Event) error {
	if err := InsertEvent(db.database, db.logger, &evt); err != nil {
		db.logger.Error("insert event error", "err", err)
		return err
	}

	return nil
}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	if err := regOrm(db); err != nil {
		return err
	}

	models := []interface{}{
		(*SyncState)(nil),
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

func syncChainStat(db *pg.DB, logger Logger, num int64, chainID string) error {
	stat := &SyncState{
		ID: ChainIdx,
	}
	err := db.Select(stat)
	if err != nil {
		return errors.Wrapf(err, "get sync stat err")
	}

	logger.Info("get sync stat in %d", stat.BlockNum)

	return db.Update(SyncState{
		ID:       ChainIdx,
		BlockNum: num,
		ChainID:  chainID,
	})
}
