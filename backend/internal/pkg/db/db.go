package db
import (
	"github.com/jinzhu/gorm"
)

// this is the interface for our DB

type DB struct {
	GormDB *gorm.DB
}

type Database interface {
	Open() *DB
}

func (db *DB) Close() {
	db.GormDB.Close()
}

