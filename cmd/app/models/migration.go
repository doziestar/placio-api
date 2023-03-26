package models

import "gorm.io/gorm"

func Migrate(Db *gorm.DB) error {
	// drop all tables
	// d.Migrator().DropTable(models...)
	db = Db.Debug()
	var modelList []interface{}
	modelList = append(modelList,
		&User{},
		&Event{},
		&Ticket{},
		&Booking{},
		&Payment{},
		Business{},
		Conversation{},
		Group{},
		Message{},
		Profile{},
		Token{},
		GeneralSettings{},
		Login{},
		Account{},
		NotificationsSettings{},
		AccountSettings{},
		ContentSettings{},
		ConnectedAccount{})
	//err := db.Migrator().DropTable(modelList...)
	//if err != nil {
	//	return err
	//}
	return db.AutoMigrate(modelList...)
}
