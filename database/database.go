package database

import "Golang_Fiber/model"

func DropTables() {
	migrator := DB.Migrator()
	_ = migrator.DropTable(
		"migrations",
	)

	_ = migrator.DropTable(
		&model.User{},
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
