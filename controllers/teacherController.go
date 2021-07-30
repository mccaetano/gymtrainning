package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/gymtranning/models"
	"github.com/mccaetano/gymtranning/utils"
)

// TeacherGet godoc
// @Summary Get Teacher
// @Tags Teacher
// @Description Get Teacher
// @ID TeacherGet
// @Accept  json
// @Produce  json
// @Param name query string false "Name of Teacher"
// @Success 200 {array} models.Teacher
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /gym [get]
func TeacherGet(c *gin.Context) {
	log.Println("(TeacherGet - init)")

	name, _ := c.Params.Get("name")
	teachers, err := new(models.Teacher).GetAll(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(TeacherGet - finish)")
		return
	}
	c.JSON(http.StatusOK, teachers)
	log.Println("(TeacherGet - finish)")
}

// TeacherGetById godoc
// @Summary Get Teacher by id
// @Tags Teacher
// @Description Get Teacher by id
// @ID TeacherGetById
// @Accept  json
// @Produce  json
// @Param id path int true "Id of Teacher"
// @Success 200 {object} models.Teacher
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /teacher/{id} [get]
func TeacherGetById(c *gin.Context) {
	log.Println("(TeacherGetById - init)")

	paramId, res := c.Params.Get("id")
	if !res {
		paramId = "0"
	}
	id, _ := strconv.ParseInt(paramId, 10, 64)
	teacher := models.Teacher{}
	err := teacher.GetById(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(TeacherGetById - finish)")
		return
	}
	c.JSON(http.StatusOK, teacher)

	log.Println("(TeacherGetById - finish)")
}

// TeacherPost godoc
// @Summary Create Teacher
// @Tags Teacher
// @Description Create Teacher
// @ID TeacherPost
// @Accept  json
// @Produce  json
// @Param body body models.Teacher true "Teacher Model"
// @Success 200 {object} models.Teacher
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /teacher [post]
func TeacherPost(c *gin.Context) {
	log.Println("(TeacherPost - init)")
	teacher := models.Teacher{}
	if err := c.BindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErroAPIHandle(err))
		return
	}
	log.Println("(TeacherPost - body):", teacher)
	err := teacher.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(TeacherPost - finish)")
		return
	}
	c.JSON(http.StatusCreated, teacher)
	log.Println("(TeacherPost - finish)")
}

// TeacherPut godoc
// @Summary Updade Teacher
// @Tags Teacher
// @Description Updade Teacher
// @ID TeacherPut
// @Accept  json
// @Produce  json
// @Param body body models.Teacher true "Teacher Model"
// @Success 200 {object} models.Teacher
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /teacher [put]
func TeacherPut(c *gin.Context) {
	log.Println("(TeacherPut - init)")
	paramId, res := c.Params.Get("id")
	if !res {
		paramId = "0"
	}
	id, _ := strconv.ParseInt(paramId, 10, 64)

	teacher := models.Teacher{}
	if err := c.BindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErroAPIHandle(err))
		return
	}
	log.Println("(TeacherPost - body):", teacher)
	err := teacher.Update(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErroAPIHandle(err))
		log.Println("(TeacherPut - finish)")
		return
	}
	c.JSON(http.StatusCreated, teacher)

	log.Println("(TeacherPut - finish)")
}

// TeacherDelete godoc
// @Summary Delete Teacher by id
// @Tags Teacher
// @Description Delete Teacher by id
// @ID TeacherDelete
// @Accept  json
// @Produce  json
// @Param id path int true "Id of Teacher"
// @Success 200 {object} models.Teacher
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /teacher/{id} [delete]
func TeacherDelete(c *gin.Context) {
	log.Println("(TeacherDelete - init)")
	c.JSON(http.StatusNoContent, gin.H{})
	log.Println("(TeacherDelete - finish)")
}
