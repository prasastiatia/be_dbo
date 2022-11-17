package main

import (
	"be_dbo/database"
	Routes "be_dbo/routes"
)

func main() {
	// Initialize Database
	// database.Connect("tia@mail.com:secret@tcp(be_dbo_db)/be_dbo?parseTime=true")
	database.Connect("root:@tcp(localhost:3306)/be_dbo?parseTime=true")
	database.Migrate()

	// Initialize Router
	router := Routes.SetupRouter()
	router.Run(":8080")
}
