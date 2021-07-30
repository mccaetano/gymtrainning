package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/gymtranning/models"
	"github.com/mccaetano/gymtranning/utils"
)

func UserProfileGet(c *gin.Context) {
	log.Println("(UserProfileGet - init)")

	name, _ := c.Params.Get("name")
	exercises, err := new(models.UserProfile).GetAll(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(UserProfileGet - finish)")
		return
	}
	c.JSON(http.StatusOK, exercises)
	log.Println("(UserProfileGet - finish)")
}

func UserProfileGetById(c *gin.Context) {
	log.Println("(UserProfileGetById - init)")

	paramId, res := c.Params.Get("id")
	if !res {
		paramId = "0"
	}
	id, _ := strconv.ParseInt(paramId, 10, 64)
	exercises := models.UserProfile{}
	err := exercises.GetById(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(UserProfileGetById - finish)")
		return
	}
	c.JSON(http.StatusOK, exercises)

	log.Println("(UserProfileGetById - finish)")
}

func UserProfilePost(c *gin.Context) {
	log.Println("(UserProfilePost - init)")
	exercises := models.UserProfile{}
	if err := c.BindJSON(&exercises); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErroAPIHandle(err))
		return
	}
	log.Println("(UserProfilePost - body):", exercises)
	err := exercises.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(UserProfilePost - finish)")
		return
	}
	c.JSON(http.StatusCreated, exercises)
	log.Println("(UserProfilePost - finish)")
}

func UserProfilePut(c *gin.Context) {
	log.Println("(UserProfilePut - init)")
	paramId, res := c.Params.Get("id")
	if !res {
		paramId = "0"
	}
	id, _ := strconv.ParseInt(paramId, 10, 64)

	exercises := models.UserProfile{}
	if err := c.BindJSON(&exercises); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErroAPIHandle(err))
		return
	}
	log.Println("(UserProfilePost - body):", exercises)
	err := exercises.Update(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(UserProfilePut - finish)")
		return
	}
	c.JSON(http.StatusCreated, exercises)

	log.Println("(UserProfilePut - finish)")
}

func UserProfileDelete(c *gin.Context) {
	log.Println("(UserProfileDelete - init)")
	c.JSON(http.StatusNoContent, gin.H{})
	log.Println("(UserProfileDelete - finish)")
}
