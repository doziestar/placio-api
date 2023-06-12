package database

import (
	"context"
	"log"
	"placio-app/ent"
	"placio-pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct {
	*gorm.DB
}

func NewDatabase(dsn string) (*Database, error) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{DB}, nil
}

// Connect to database
func Connect(dsn string) (*Database, error) {
	var err error
	ctx := context.Background()
	logger.Info(ctx, "Connecting to database")
	db, err := NewDatabase(dsn)
	if err != nil {
		return nil, err
	}

	logger.Info(ctx, "======================================")
	logger.Info(ctx, "== Connected to database successfully ==")
	logger.Info(ctx, "======================================")

	if err != nil {
		return nil, err
	}

	return db, nil
}

// EntClient Migrate database
func EntClient(ctx context.Context) *ent.Client {
	//client, err := ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	client, err := ent.Open("postgres", "host=postgres-db port=5432 user=dozie dbname=placio password=918273645dozie")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	logger.Info(ctx, "======================================")
	logger.Info(ctx, "== Connected to database successfully ==")
	logger.Info(ctx, "======================================")
	//defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

// GetDB returns database connection

// LoadModels load models

// GetDB Make database connection available to all models and use it to perform database operations
func (d *Database) GetDB() *gorm.DB {
	return d.DB
}

//
//type Database struct {
//	*gorm.DB
//	pool *pgxpool.Pool
//}
//
//func NewDatabase(dsn string, poolSize int) (*Database, error) {
//	pool, err := pgxpool.Connect(context.Background(), dsn)
//	if err != nil {
//		return nil, err
//	}
//
//	db, err := gorm.Open(postgres.New(postgres.Config{
//		Conn: pool.Acquire,
//	}), &gorm.Config{})
//	if err != nil {
//		return nil, err
//	}
//
//	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
//	db.SetMaxIdleConns(poolSize)
//
//	// SetMaxOpenConns sets the maximum number of open connections to the database.
//	db.SetMaxOpenConns(poolSize)
//
//	return &Database{db, pool}, nil
//}
//
//func (d *Database) Migrate(models ...interface{}) error {
//	return d.AutoMigrate(models...)
//}
//
//// Connect to database
//func Connect(dsn string, poolSize int) (*Database, error) {
//	ctx := context.Background()
//	logger.Info(ctx, "Connecting to database")
//	db, err := NewDatabase(dsn, poolSize)
//	if err != nil {
//		return nil, err
//	}
//
//	logger.Info(ctx, "======================================")
//	logger.Info(ctx, "== Connected to database successfully ==")
//	logger.Info(ctx, "======================================")
//
//	return db, nil
//}
//
//// LoadModels load models
//func (d *Database) LoadModels() error {
//	var modelList []interface{}
//	modelList = append(modelList, &models.User{}, &models.Event{}, &models.Ticket{}, &models.Booking{}, &models.Payment{}, models.Business{}, models.Conversation{}, models.Group{}, models.Message{})
//	return d.Migrate(modelList...)
//}
//
//func (d *Database) Close() error {
//	d.pool.Close()
//	return nil
//}
