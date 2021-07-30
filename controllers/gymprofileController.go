package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/gymtranning/models"
	"github.com/mccaetano/gymtranning/utils"
)

func GymProfileGet(c *gin.Context) {
	log.Println("(GymProfileGet - init)")
	name, _ := c.Params.Get("name")
	exercises, err := new(models.GymProfile).GetAll(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(GymProfileGet - finish)")
		return
	}
	c.JSON(http.StatusOK, exercises)
	log.Println("(GymProfileGet - finish)")
}

func GymProfileGetById(c *gin.Context) {
	log.Println("(GymProfileGetById - init)")
	paramId, res := c.Params.Get("id")
	if !res {
		paramId = "0"
	}
	id, _ := strconv.ParseInt(paramId, 10, 64)
	exercises := models.GymProfile{}
	err := exercises.GetById(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(GymProfileGetById - finish)")
		return
	}
	c.JSON(http.StatusOK, exercises)

	log.Println("(GymProfileGetById - finish)")
}

func GymProfilePost(c *gin.Context) {
	log.Println("(GymProfilePost - init)")
	exercises := models.GymProfile{}
	if err := c.BindJSON(&exercises); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErroAPIHandle(err))
		return
	}
	log.Println("(GymProfilePost - body):", exercises)
	err := exercises.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(GymProfilePost - finish)")
		return
	}
	c.JSON(http.StatusCreated, exercises)
	log.Println("(GymProfilePost - finish)")
}

func GymProfilePut(c *gin.Context) {
	log.Println("(GymProfilePut - init)")
	paramId, res := c.Params.Get("id")
	if !res {
		paramId = "0"
	}
	id, _ := strconv.ParseInt(paramId, 10, 64)

	exercises := models.GymProfile{}
	if err := c.BindJSON(&exercises); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErroAPIHandle(err))
		return
	}
	log.Println("(GymProfilePost - body):", exercises)
	err := exercises.Update(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(GymProfilePut - finish)")
		return
	}
	c.JSON(http.StatusCreated, exercises)

	log.Println("(GymProfilePut - finish)")
}

func GymProfileDelete(c *gin.Context) {
	log.Println("(GymProfileDelete - init)")
	c.JSON(http.StatusNoContent, gin.H{})
	log.Println("(GymProfileDelete - finish)")
}
