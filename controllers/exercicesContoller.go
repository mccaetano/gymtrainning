package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/gymtranning/models"
	"github.com/mccaetano/gymtranning/utils"
)

// ExercisesGet godoc
// @Summary Get exercises
// @Tags Exercises
// @Description Get exercises
// @ID ExercisesGet
// @Accept  json
// @Produce  json
// @Param name query string false "Name of exercises"
// @Success 200 {array} models.Exercises
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /exercises [get]
func ExercisesGet(c *gin.Context) {
	log.Println("(ExercisesGet - init)")
	name, _ := c.Params.Get("name")
	exercises, err := new(models.Exercises).GetAll(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(ExercisesGet - finish)")
		return
	}
	c.JSON(http.StatusOK, exercises)
	log.Println("(ExercisesGet - finish)")
}

// ExercisesGetById godoc
// @Summary Get exercises by id
// @Tags Exercises
// @Description Get exercises by id
// @ID ExercisesGetById
// @Accept  json
// @Produce  json
// @Param id path int true "Id of exercises"
// @Success 200 {object} models.Exercises
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /exercises/{id} [get]
func ExercisesGetById(c *gin.Context) {
	log.Println("(ExercisesGetById - init)")
	paramId, res := c.Params.Get("id")
	if !res {
		paramId = "0"
	}
	id, _ := strconv.ParseInt(paramId, 10, 64)
	exercises := models.Exercises{}
	err := exercises.GetById(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(ExercisesGetById - finish)")
		return
	}
	c.JSON(http.StatusOK, exercises)

	log.Println("(ExercisesGetById - finish)")
}

// ExercisesPost godoc
// @Summary Create exercises
// @Tags Exercises
// @Description Create exercises
// @ID ExercisesPost
// @Accept  json
// @Produce  json
// @Param body body models.Exercises true "Exercises Model"
// @Success 200 {object} models.Exercises
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /exercises [post]
func ExercisesPost(c *gin.Context) {
	log.Println("(ExercisesPost - init)")
	exercises := models.Exercises{}
	if err := c.BindJSON(&exercises); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErroAPIHandle(err))
		return
	}
	log.Println("(ExercisesPost - body):", exercises)
	err := exercises.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(ExercisesPost - finish)")
		return
	}
	c.JSON(http.StatusCreated, exercises)
	log.Println("(ExercisesPost - finish)")
}

// ExercisesPut godoc
// @Summary Updade exercises
// @Tags Exercises
// @Description Updade exercises
// @ID ExercisesPut
// @Accept  json
// @Produce  json
// @Param body body models.Exercises true "Exercises Model"
// @Success 200 {object} models.Exercises
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /exercises [put]
func ExercisesPut(c *gin.Context) {
	log.Println("(ExercisesPut - init)")
	paramId, res := c.Params.Get("id")
	if !res {
		paramId = "0"
	}
	id, _ := strconv.ParseInt(paramId, 10, 64)

	exercises := models.Exercises{}
	if err := c.BindJSON(&exercises); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErroAPIHandle(err))
		return
	}
	log.Println("(ExercisesPost - body):", exercises)
	err := exercises.Update(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(ExercisesPut - finish)")
		return
	}
	c.JSON(http.StatusCreated, exercises)

	log.Println("(ExercisesPut - finish)")
}

// ExercisesDelete godoc
// @Summary Delete exercises by id
// @Tags Exercises
// @Description Delete exercises by id
// @ID ExercisesDelete
// @Accept  json
// @Produce  json
// @Param id path int true "Id of exercises"
// @Success 200 {object} models.Exercises
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /exercises/{id} [delete]
func ExercisesDelete(c *gin.Context) {
	log.Println("(ExercisesDelete - init)")
	c.JSON(http.StatusNoContent, gin.H{})
	log.Println("(ExercisesDelete - finish)")
}
