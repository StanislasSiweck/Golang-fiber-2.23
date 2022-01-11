package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Joins(statement *gorm.DB, joins []string) *gorm.DB {
	for _, join := range joins {
		if join == "All" {
			join = clause.Associations
		} else if join == "" {
			continue
		}
		statement.Preload(join, func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		})
	}
	return statement
}

func JoinsScoped(statement *gorm.DB, joins []string) *gorm.DB {
	for _, join := range joins {
		if join == "All" {
			join = clause.Associations
		} else if join == "" {
			continue
		}
		statement.Preload(join, func(db *gorm.DB) *gorm.DB {
			return db.Scopes()
		})
	}
	return statement
}
