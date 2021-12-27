package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
)

func MigrateDatabase() {
	m := gormigrate.New(DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "Initial",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(
					&User{},
					&Role{},
				)
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(
					&User{},
					&Role{},
				)
			},
		},
	})
	if err := m.Migrate(); err != nil {
		log.Printf("%s%s", "Erreur lors du rollback: ", err.Error())
		_ = m.RollbackLast()
	}
}
