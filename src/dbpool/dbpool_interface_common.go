package dbpool

import (
	"lizisky.com/lizisky/src/dbpool/dbimpl"
)

func InitDatabase() {
	dbimpl.NewDBConnection(nil)
}

// general DB operations
func AddNewRecord(newRecord interface{}) error {
	return dbimpl.AddNewRecord(newRecord)
}

// for change some field of a record
func UpdateRecord(record interface{}) error {
	return dbimpl.UpdateRecord(record)
}
