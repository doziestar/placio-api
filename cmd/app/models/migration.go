package models

import (
	"gorm.io/gorm"
)

func Migrate(Db *gorm.DB) error {
	// drop all tables
	// d.Migrator().DropTable(models...)
	// db := Db.Debug()
	var modelList []interface{}
	modelList = append(modelList,
		&Event{},
		&Ticket{},
		&Booking{},
		&Payment{},
		&Business{},
		&Conversation{},
		&Group{},
		&Message{},
		&Profile{},
		&Login{},
		&Post{},
		&Comment{},
		&Like{},
		&Media{},
		&Follow{},
		&Ticket{},
		&TicketOption{},
		&Attendee{},
		&Rating{})
	// err := db.Migrator().DropTable(modelList...)
	// if err != nil {
	// 	return err
	// }
	// // Migrate User model first
	// if err := db.AutoMigrate(modelList...); err != nil {
	// 	logger.Error(context.Background(), err.Error())
	// 	return err
	// }

	// Migrate other models
	//for _, model := range modelList[1:] {
	//	if err := db.AutoMigrate(model); err != nil {
	//		return err
	//	}
	//}

	return nil
}
