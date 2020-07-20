package models

import (
	"github.com/code-devel-cover/CodeCover/core"
	"github.com/jinzhu/gorm"
)

var (
	tables []interface{}
)

type databaseService struct {
	db *gorm.DB
}

// NewDatabaseService with GORM
func NewDatabaseService(db *gorm.DB) core.DatabaseService {
	return &databaseService{
		db: db,
	}
}

func (store *databaseService) Session() *gorm.DB {
	return store.db.New()
}

func (store *databaseService) Migrate() error {
	return migrate(store.db)
}

func init() {
	tables = append(tables,
		&Report{},
		&User{},
		&Repo{},
		&RepoSetting{},
	)
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(tables...).Error
}
