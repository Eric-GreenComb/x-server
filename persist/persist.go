package persist

import (
	"github.com/jinzhu/gorm"
)

// Persist struct
type Persist struct {
	db *gorm.DB
}

// GPersist global Persist
var GPersist *Persist

// GetPersist Get Persist
func GetPersist() *Persist {

	if GPersist != nil {
		return GPersist
	}

	GPersist = new(Persist)

	var err error
	GPersist.db, err = ConnectDb()

	if err != nil {
		GPersist = nil
	}

	return GPersist
}

// ReConnetDB ReConnetDB
func ReConnetDB() error {
	GPersist = new(Persist)

	var err error
	GPersist.db, err = ConnectDb()

	if err != nil {
		GPersist = nil
	}

	return err
}

// Close Close db
func (persist *Persist) Close() {
	if GPersist == nil {
		return
	}
	GPersist.db.Close()
}
