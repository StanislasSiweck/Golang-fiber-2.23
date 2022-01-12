package database

import "Golang_Fiber/model"

func DropTables() {
	migrator := DB.Migrator()
	_ = migrator.DropTable(
		"migrations",
		"casbin_rule",
		"user_roles",
	)

	_ = migrator.DropTable(
		&model.User{},
		&model.Role{},
	)
}

func CloseConnections() {
	//Close mysql
	db, _ := DB.DB()
	_ = db.Close()
}

//Filter Database filter
type Filter struct {
	Column  string
	Value   string
	Like    bool
	OrderBy bool
}
