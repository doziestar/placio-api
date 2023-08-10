package models

import (
	"context"
	"placio-pkg/logger"

	_ "github.com/lib/pq"
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
		&Login{},
		&Post{},
		&Comment{},
		&Like{},
		&Media{},
		&Follow{},
		&Ticket{},
		&TicketOption{},
		&Attendee{},
		&Rating{},
		&User{},
		&BusinessAccount{},
		&UserBusinessRelationship{},
		&AccountSettings{},
		&GeneralSettings{},
		&NotificationsSettings{},
		&ContentSettings{},
	)
	err := Db.Migrator().DropTable(modelList...)
	if err != nil {
		return err
	}
	// // Migrate User model first
	if err := Db.AutoMigrate(modelList...); err != nil {
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
