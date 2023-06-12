package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"placio-app/api"
	"placio-app/database"
	"placio-app/models"
	"placio-app/start"
)

// @title Placio Application Api
// @version 0.01
// @description This is the documentation for the Placio Application Api
// @termsOfService https://placio.io/terms
// @privacyPolicy https://placio.io/privacy-policy
// @contact.name Darc Technologies
// @contact.url https://placio.io
// @contact.email support@placio.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host https://api.palnight.com
// @BasePath /qpi/v1
// @schemes http https
func main() {
	// get port from env
	port := os.Getenv("PORT")
	// if port is not set, set it to 3000

	// initialize fiber app
	app := gin.Default()

	start.Middleware(app)

	//initialize database
	//env, _ := config.LoadConfig("./config")
	databaseInstance, err := database.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		return
	}
	db := databaseInstance.GetDB()
	err = models.Migrate(db)
	if err != nil {
		return
	}

	//client := database.EntClient(context.Background())
	//log.Printf("client: %v", client)

	// initialize routes
	api.InitializeRoutes(app, db)
	// set port
	start.Initialize(port, app)

}
