package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/gymtranning/models"
	"github.com/mccaetano/gymtranning/utils"
)

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

func ExercisesDelete(c *gin.Context) {
	log.Println("(ExercisesDelete - init)")
	c.JSON(http.StatusNoContent, gin.H{})
	log.Println("(ExercisesDelete - finish)")
}
