package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/gymtranning/models"
	"github.com/mccaetano/gymtranning/utils"
)

// UserProfileGet godoc
// @Summary Get User
// @Tags UserProfile
// @Description Get User
// @ID UserProfileGet
// @Accept  json
// @Produce  json
// @Param name query string false "Name of User"
// @Success 200 {array} models.UserProfile
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /gym/profile [get]
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

// UserProfileGetById godoc
// @Summary Get User by id
// @Tags UserProfile
// @Description Get User by id
// @ID UserProfileGetById
// @Accept  json
// @Produce  json
// @Param id path int true "Id of User"
// @Success 200 {object} models.UserProfile
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /user/profile/{id} [get]
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

// UserProfilePost godoc
// @Summary Create User
// @Tags UserProfile
// @Description Create User
// @ID UserProfilePost
// @Accept  json
// @Produce  json
// @Param body body models.UserProfile true "UserProfile Model"
// @Success 200 {object} models.UserProfile
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /user/profile [post]
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

// UserProfilePut godoc
// @Summary Updade User
// @Tags UserProfile
// @Description Updade User
// @ID UserProfilePut
// @Accept  json
// @Produce  json
// @Param body body models.UserProfile true "UserProfile Model"
// @Success 200 {object} models.UserProfile
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /user/profile [put]
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

// UserProfileDelete godoc
// @Summary Delete User by id
// @Tags UserProfile
// @Description Delete User by id
// @ID UserProfileDelete
// @Accept  json
// @Produce  json
// @Param id path int true "Id of User"
// @Success 200 {object} models.UserProfile
// @Failure 400,404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Failure default {object} utils.Error
// @Router /user/profile/{id} [delete]
func UserProfileDelete(c *gin.Context) {
	log.Println("(UserProfileDelete - init)")
	c.JSON(http.StatusNoContent, gin.H{})
	log.Println("(UserProfileDelete - finish)")
}
