package models

import (
	"context"
	"placio-pkg/logger"

	"gorm.io/gorm"
)

func Migrate(Db *gorm.DB) error {
	// drop all tables
	// d.Migrator().DropTable(models...)
	db = Db.Debug()
	var modelList []interface{}
	modelList = append(modelList,
		&User{},
		&Account{},
		&Event{},
		&Ticket{},
		&Booking{},
		&Payment{},
		&Business{},
		&Conversation{},
		&Group{},
		&Message{},
		&Profile{},
		&Token{},
		&GeneralSettings{},
		&Login{},
		&NotificationsSettings{},
		&AccountSettings{},
		&ContentSettings{},
		&ConnectedAccount{},
		&Post{},
		&Comment{},
		&Like{},
		&Media{},
		&Follow{},
		&Ticket{},
		&TicketOption{},
		&Attendee{},
		&Rating{})
	err := db.Migrator().DropTable(modelList...)
	if err != nil {
		return err
	}
	// Migrate User model first
	if err := db.AutoMigrate(modelList...); err != nil {
		logger.Error(context.Background(), err.Error())
		return err
	}

	// Migrate other models
	//for _, model := range modelList[1:] {
	//	if err := db.AutoMigrate(model); err != nil {
	//		return err
	//	}
	//}

	return nil
}
