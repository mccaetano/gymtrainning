package main

import (
	"github.com/mccaetano/gymtranning/config"
	"github.com/mccaetano/gymtranning/models"
)

func main() {
	models.ConnectDatabase()

	config.Routes()
}
