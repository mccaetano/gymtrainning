package main

import (
	"github.com/mccaetano/gymtranning/config"
	"github.com/mccaetano/gymtranning/docs"
	"github.com/mccaetano/gymtranning/models"
)

// @title Gym Trainning
// @version 1.0
// @description Show your trainning and manage your trainning
// @termsOfService http://swagger.io/terms/

// @contact.name MCCAETANO
// @contact.url http://www.swagger.io/support
// @contact.email marcelo.cheruti@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host gymtrainning-api.herokuapp.com/
// @BasePath /gymtranning/v1/api
// @query.collection.format multi

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

func main() {
	models.ConnectDatabase()

	docs.SwaggerInfo.Description = "teste"

	config.Routes()
}
