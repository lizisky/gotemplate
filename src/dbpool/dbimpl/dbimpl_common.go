package dbimpl

import (
	"github.com/golang/glog"
)

const (
	invalidParameter = "invalid parameter"
)

// add new record to db
func AddNewRecord(newRecord interface{}) error {
	db := localDB.currDB.Create(newRecord)
	if db.Error != nil {
		glog.Errorln("insert record to db error:", db.Error)
	}

	return db.Error
}

// update record from db
func UpdateRecord(record interface{}) error {
	db := localDB.currDB.Save(record)
	if db.Error != nil {
		glog.Errorln("update record to db error:", db.Error)
	}

	return db.Error
}
