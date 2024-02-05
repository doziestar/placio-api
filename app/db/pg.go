package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"math"
	"math/rand"
	"os"
	"placio-app/ent"
	"placio-pkg/logger"
	"time"

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

// Connect to db
func Connect(dsn string) (*Database, error) {
	var err error
	ctx := context.Background()
	logger.Info(ctx, "Connecting to db")
	db, err := NewDatabase(dsn)
	if err != nil {
		return nil, err
	}

	logger.Info(ctx, "======================================")
	logger.Info(ctx, "== Connected to db successfully ==")
	logger.Info(ctx, "======================================")

	if err != nil {
		return nil, err
	}

	return db, nil
}

// EntClient migrates the db and returns the ent client
func EntClient(ctx context.Context) *ent.Client {
	var (
		client *ent.Client
		err    error
	)

	// Load environment variables from .env file if exists
	_ = godotenv.Load()

	host := getEnv("DB_HOST", "monorail.proxy.rlwy.net")
	//host := getEnv("DB_HOST", "postgres-db")
	port := getEnv("DB_PORT", "20871")
	//user := getEnv("DB_USER", "dozie")
	user := getEnv("DB_USER", "postgres")
	//user := getEnv("DB_USER", "dozie")
	dbName := getEnv("DB_NAME", "postgres")
	//password := getEnv("DB_PASSWORD", "918273645dozie")
	password := getEnv("DB_PASSWORD", "c4Eg3g5BggGa4bAED13CbEaCb13GDFd1")

	log.Println("Connecting to db", host, port, user, dbName, password)

	maxRetries := 5
	for i := 1; i <= maxRetries; i++ {
		client, err = ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password))
		if err != nil {
			log.Println("Error connecting to db: ", err)
			if i < maxRetries {
				waitTime := time.Duration(math.Pow(2, float64(i))) * time.Second
				waitTime += time.Duration(rand.Intn(1000)) * time.Millisecond
				log.Printf("Retrying db connection in %v", waitTime)
				time.Sleep(waitTime)
				continue
			}
			log.Fatalf("Failed to connect to db after %v retries", maxRetries)
		}
		break
	}

	log.Println("======================================")
	log.Println("== Connected to db successfully ==")
	log.Println("======================================")

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

// getEnv gets an environment variable or returns a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

//// EntClient Migrate db
//func EntClient(ctx context.Context) *ent.Client {
//	//client, err := ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<db> password=<pass>")
//	client, err := ent.Open("postgres", "host=postgres-db port=5432 user=dozie dbname=placio password=918273645dozie")
//
//	if err != nil {
//		log.Fatalf("failed opening connection to postgres: %v", err)
//	}
//	logger.Info(ctx, "======================================")
//	logger.Info(ctx, "== Connected to db successfully ==")
//	logger.Info(ctx, "======================================")
//	//defer client.Close()
//	// Run the auto migration tool.
//	if err := client.Schema.Create(context.Background()); err != nil {
//		log.Fatalf("failed creating schema resources: %v", err)
//	}
//	return client
//}

// GetDB returns db connection

// LoadModels load models

// GetDB Make db connection available to all models and use it to perform db operations
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
//	// SetMaxOpenConns sets the maximum number of open connections to the db.
//	db.SetMaxOpenConns(poolSize)
//
//	return &Database{db, pool}, nil
//}
//
//func (d *Database) Migrate(models ...interface{}) error {
//	return d.AutoMigrate(models...)
//}
//
//// Connect to db
//func Connect(dsn string, poolSize int) (*Database, error) {
//	ctx := context.Background()
//	logger.Info(ctx, "Connecting to db")
//	db, err := NewDatabase(dsn, poolSize)
//	if err != nil {
//		return nil, err
//	}
//
//	logger.Info(ctx, "======================================")
//	logger.Info(ctx, "== Connected to db successfully ==")
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
