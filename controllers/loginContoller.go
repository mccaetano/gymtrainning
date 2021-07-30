package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/gymtranning/models"
	"github.com/mccaetano/gymtranning/utils"
)

func LoginPost(c *gin.Context) {
	log.Println("(loginPost - init)")

	userLogin := models.UserLogin{}
	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErroAPIHandle(err))
		return
	}

	tokem, err := userLogin.Login()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(loginPost - finish)")
		return
	}
	c.JSON(http.StatusOK, gin.H{"tokem": tokem})
	log.Println("(loginPost - finish)")
}
