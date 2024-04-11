package dbimpl

import (
	"fmt"
	"path"

	"lizisky.com/lizisky/src/basictypes/accounts"
	orgtype "lizisky.com/lizisky/src/basictypes/orgtype_common"
	"lizisky.com/lizisky/src/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
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

func CloseDB() {
	if localDB != nil {
		localDB.Close()
		localDB = nil
	}
}

// ----------------------------------------------------------------------------
func NewDBConnection(cfg *config.Configuration) *DBPool {
	if cfg == nil {
		cfg = config.GetConfig()
	}
	switch cfg.DBType {
	case 1: // SQLite3
		localDB = initDb_SQLite(&cfg.SQLite)
	case 2: // MySQL
		localDB = initDb_mysql(&cfg.MySQL)
	}
	initDB_autoMigration()

	return localDB
}

// build mysql DB dsn
// func build_DB_dsn(dbcfg *config.DBConfigMySQL) string {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbcfg.DBUser, dbcfg.DBUserPwd, dbcfg.DBHost, dbcfg.DBDatabase)
// 	return dsn
// }

// init mysql DB connection
func initDb_mysql(dbcfg *config.DBConfigMySQL) *DBPool {
	// dsn := build_DB_dsn(dbcfg)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbcfg.DBUser, dbcfg.DBUserPwd, dbcfg.DBHost, dbcfg.DBDatabase)
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

// init SQLite DB connection
func initDb_SQLite(dbcfg *config.DBConfigSQLite) *DBPool {

	dbpath := path.Join(config.GetDataDir(), dbcfg.DBPath)
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})

	if err != nil {
		panic("open SQLite error:" + err.Error())
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
