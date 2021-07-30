package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/gymtranning/models"
	"github.com/mccaetano/gymtranning/utils"
)

type Token struct {
	Token string `json:"token"`
}

// LoginPost godoc
// @Summary Login user is Applocation
// @Tags Login
// @Description Login by username and password
// @ID LoginPost
// @Accept  json
// @Produce  json
// @Param body body models.UserLogin true "User Login"
// @Success 200 {object} controllers.Token
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /login [post]
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
	response := Token{Token: tokem}
	c.JSON(http.StatusOK, response)
	log.Println("(loginPost - finish)")
}
