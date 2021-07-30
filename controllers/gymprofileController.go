package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/gymtranning/models"
	"github.com/mccaetano/gymtranning/utils"
)

// GymProfileGet godoc
// @Summary Get Gym
// @Tags GymProfile
// @Description Get Gym
// @ID GymProfileGet
// @Accept  json
// @Produce  json
// @Param name query string false "Name of Gym"
// @Success 200 {array} models.GymProfile
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /gym/profile [get]
func GymProfileGet(c *gin.Context) {
	log.Println("(GymProfileGet - init)")
	name, _ := c.Params.Get("name")
	gymProfile, err := new(models.GymProfile).GetAll(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(GymProfileGet - finish)")
		return
	}
	c.JSON(http.StatusOK, gymProfile)
	log.Println("(GymProfileGet - finish)")
}

// GymProfileGetById godoc
// @Summary Get Gym by id
// @Tags GymProfile
// @Description Get Gym by id
// @ID GymProfileGetById
// @Accept  json
// @Produce  json
// @Param id path int true "Id of Gym"
// @Success 200 {object} models.GymProfile
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /gym/profile/{id} [get]
func GymProfileGetById(c *gin.Context) {
	log.Println("(GymProfileGetById - init)")
	paramId, res := c.Params.Get("id")
	if !res {
		paramId = "0"
	}
	id, _ := strconv.ParseInt(paramId, 10, 64)
	gymProfile := models.GymProfile{}
	err := gymProfile.GetById(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(GymProfileGetById - finish)")
		return
	}
	c.JSON(http.StatusOK, gymProfile)

	log.Println("(GymProfileGetById - finish)")
}

// GymProfilePost godoc
// @Summary Create Gym
// @Tags GymProfile
// @Description Create Gym
// @ID GymProfilePost
// @Accept  json
// @Produce  json
// @Param body body models.GymProfile true "GymProfile Model"
// @Success 200 {object} models.GymProfile
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /gym/profile [post]
func GymProfilePost(c *gin.Context) {
	log.Println("(GymProfilePost - init)")
	gymProfile := models.GymProfile{}
	if err := c.BindJSON(&gymProfile); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErroAPIHandle(err))
		return
	}
	log.Println("(GymProfilePost - body):", gymProfile)
	err := gymProfile.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(GymProfilePost - finish)")
		return
	}
	c.JSON(http.StatusCreated, gymProfile)
	log.Println("(GymProfilePost - finish)")
}

// GymProfilePut godoc
// @Summary Updade Gym
// @Tags GymProfile
// @Description Updade Gym
// @ID GymProfilePut
// @Accept  json
// @Produce  json
// @Param body body models.GymProfile true "GymProfile Model"
// @Success 200 {object} models.GymProfile
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /gym/profile [put]
func GymProfilePut(c *gin.Context) {
	log.Println("(GymProfilePut - init)")
	paramId, res := c.Params.Get("id")
	if !res {
		paramId = "0"
	}
	id, _ := strconv.ParseInt(paramId, 10, 64)

	gymProfile := models.GymProfile{}
	if err := c.BindJSON(&gymProfile); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErroAPIHandle(err))
		return
	}
	log.Println("(GymProfilePost - body):", gymProfile)
	err := gymProfile.Update(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(GymProfilePut - finish)")
		return
	}
	c.JSON(http.StatusCreated, gymProfile)

	log.Println("(GymProfilePut - finish)")
}

// GymProfileDelete godoc
// @Summary Delete Gym by id
// @Tags GymProfile
// @Description Delete Gym by id
// @ID GymProfileDelete
// @Accept  json
// @Produce  json
// @Param id path int true "Id of Gym"
// @Success 200 {object} models.GymProfile
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /gym/profile/{id} [delete]
func GymProfileDelete(c *gin.Context) {
	log.Println("(GymProfileDelete - init)")
	c.JSON(http.StatusNoContent, gin.H{})
	log.Println("(GymProfileDelete - finish)")
}
