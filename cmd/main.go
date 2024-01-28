package main

import (
	"CodeWithAzri/internal/app"
)

//	@title			CodeWithAzri API
//	@version		1.0
//	@description	API documentation for CodeWithAzri, an educational platform that offers a free and collaborative environment for learning coding. Provides resources, exercises, and a community for both mobile app and web app users.
//	@contact.email	support@codewithazri.com
//	@license.name	MIT

//	@host		localhost:8080
//	@BasePath	/api/v1
func main() {
	// Initialize the CodeWithAzri application
	a := app.NewApp()

	// Run the server
	a.Run()
}
