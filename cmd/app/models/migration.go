package models

import (
	"context"
	"gorm.io/gorm"
	"placio-pkg/logger"
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
		&ConnectedAccount{})
	//err := db.Migrator().DropTable(modelList...)
	//if err != nil {
	//	return err
	//}
	// Migrate User model first
	if err := db.AutoMigrate(&User{}); err != nil {
		logger.Error(context.Background(), err.Error())
		return err
	}

	// Migrate other models
	for _, model := range modelList[1:] {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}
