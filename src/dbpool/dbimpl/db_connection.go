package dbimpl

import (
	"fmt"

	"gorm.io/driver/mysql"
	"lizisky.com/lizisky/src/basictypes/accounts"
	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/config"

	"gorm.io/gorm"
)

// ----------------------------------------------------------------------------
var (
	localDB *DBPool
)

// initialize local DB pool
func init() {
	localDB = &DBPool{}
}

// define DBPool struct
type DBPool struct {
	currDB *gorm.DB
	isConn bool
	err    error
}

// check connection status
func (pool *DBPool) IsConn() bool {
	return pool.isConn
}

// get current DB pool error status
func (pool *DBPool) Err() error {
	return pool.err
}

// close db connection
func (pool *DBPool) Close() {
	sqldb, err := pool.currDB.DB()
	if err == nil {
		defer sqldb.Close()
	}

	pool.isConn = false
	pool.currDB = nil
}

func GetDB() *gorm.DB {
	return localDB.currDB
}

// ----------------------------------------------------------------------------
func NewDBConnection(cfg *config.Configuration) *DBPool {
	if cfg == nil {
		cfg = config.GetConfig()
	}
	localDB = initDb(&cfg.MySQL)
	initDB_autoMigration()

	return localDB
}

// build mysql DB dsn
func build_DB_dsn(dbcfg *config.DBConfig) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbcfg.DBUser, dbcfg.DBUserPwd, dbcfg.DBHost, dbcfg.DBDatabase)
	return dsn
}

// init mysql DB connection
func initDb(dbcfg *config.DBConfig) *DBPool {
	dsn := build_DB_dsn(dbcfg)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("open mysql error:" + err.Error())
	}

	sqldb, err := db.DB()
	if (err == nil) && (sqldb != nil) {
		sqldb.SetMaxOpenConns(dbcfg.MaxConn)
		sqldb.SetMaxIdleConns(dbcfg.MaxConn)
	}

	return &DBPool{
		currDB: db,
		isConn: true,
	}
}

// auto migration
func initDB_autoMigration() {
	localDB.currDB.AutoMigrate(&accounts.Account{})

	localDB.currDB.AutoMigrate(&orgtype.Organization{})
	localDB.currDB.AutoMigrate(&orgtype.OrgStaff{})
}
