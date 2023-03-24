package models

import "gorm.io/gorm"

func Migrate(Db *gorm.DB) error {
	// drop all tables
	// d.Migrator().DropTable(models...)
	var modelList []interface{}
	modelList = append(modelList, &User{}, &Event{}, &Ticket{}, &Booking{}, &Payment{}, Business{}, Conversation{}, Group{}, Message{}, Profile{}, Token{}, GeneralSettings{}, Login{}, Account{})
	return Db.AutoMigrate(modelList...)
}
