package main

import (
	"projectmanagement/backend-api/config"
	"projectmanagement/backend-api/database"
	"projectmanagement/backend-api/routes"
)

func main() {

	config.LoadEnv()
	database.InitDB()
	database.SeedAdminUser(database.DB)
	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
