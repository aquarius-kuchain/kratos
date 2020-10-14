package db

import (
	"sync"

	"github.com/go-pg/pg/v10"
)

const (
	ChainIdx = 1
)

type dbWork struct {
	msg interface{}
}

func NewWork(msg interface{}) dbWork {
	return dbWork{msg: msg}
}

func (w dbWork) IsStopped() bool {
	return w.msg == nil
}

type DBService struct {
	logger   Logger
	database *pg.DB

	dbChan chan dbWork
	wg     sync.WaitGroup
}

// NewDB create a connection commit event to db
func NewDB(cfg Cfg, logger Logger) *DBService {
	return &DBService{
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

func (db *DBService) Start() error {
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

func (db *DBService) Process(work *dbWork) error {
	if err := Process(db.database, db.logger, work.msg); err != nil {
		return err
	}

	return nil
}

func (db *DBService) Emit(work dbWork) {
	db.dbChan <- work
}

func (db *DBService) Stop() error {
	db.logger.Info("Stopping database service")

	db.dbChan <- dbWork{}
	db.wg.Wait()

	db.logger.Info("Database service stopped")

	db.database.Close()

	db.logger.Info("Database connection closed")
	return nil
}

func (db *DBService) OnEvent(evt Event) error {
	if err := InsertEvent(db.database, db.logger, &evt); err != nil {
		db.logger.Error("insert event error", "err", err)
		return err
	}

	return nil
}
