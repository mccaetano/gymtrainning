package config

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/gymtranning/controllers"
	"github.com/mccaetano/gymtranning/utils"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type headers struct {
	Authorization string `header:"Authorization"`
}

// Routes inicializa as rotas da api
func Routes() {
	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "5000"
	}
	routes := gin.New()
	routes.Use(gin.Logger())
	routes.Use(gin.CustomRecovery(recovery))

	routes.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	basePathNoAuth := routes.Group("/gymtranning/v1/api")
	basePathNoAuth.POST("/login", controllers.LoginPost)
	basePathNoAuth.POST("/user/profile", controllers.UserProfilePost)
	basePathNoAuth.PUT("/user/profile/:id", controllers.UserProfilePut)

	basePathAuth := routes.Group("/gymtranning/v1/api")
	basePathAuth.Use(authorize())
	basePathAuth.GET("/exercises", controllers.ExercisesGet)
	basePathAuth.GET("/exercises/:id", controllers.ExercisesGetById)
	basePathAuth.POST("/exercises", controllers.ExercisesPost)
	basePathAuth.PUT("/exercises/:id", controllers.ExercisesPut)
	basePathAuth.DELETE("/exercises/:id", controllers.ExercisesDelete)

	basePathAuth.GET("/gym/profile", controllers.GymProfileGet)
	basePathAuth.GET("/gym/profile/:id", controllers.GymProfileGetById)
	basePathAuth.POST("/gym/profile", controllers.GymProfilePost)
	basePathAuth.PUT("/gym/profile/:id", controllers.GymProfilePut)
	basePathAuth.DELETE("/gym/profile/:id", controllers.GymProfileDelete)

	basePathAuth.GET("/gym/teachers", controllers.GymProfileGet)
	basePathAuth.GET("/gym/teachers/:id", controllers.GymProfileGetById)
	basePathAuth.POST("/gym/teachers", controllers.GymProfilePost)
	basePathAuth.PUT("/gym/teachers/:id", controllers.GymProfilePut)
	basePathAuth.DELETE("/gym/teachers/:id", controllers.GymProfileDelete)

	basePathAuth.GET("/user/profile", controllers.UserProfileGet)
	basePathAuth.GET("/user/profile/:id", controllers.UserProfileGetById)
	basePathAuth.DELETE("/user/profile/:id", controllers.UserProfileDelete)

	routes.Run(":" + port)

}

func recovery(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}

func authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := headers{}
		if err := c.ShouldBindHeader(&h); err != nil {
			log.Println("error:", "Missing Authorization")
			c.JSON(http.StatusForbidden, gin.H{"error": "Missing Authorization"})
			c.AbortWithStatus(http.StatusForbidden)
		}
		if h.Authorization == "" {
			log.Println("error:", "Missing Authorization")
			c.JSON(http.StatusForbidden, gin.H{"error": "Missing Authorization"})
			c.AbortWithStatus(http.StatusForbidden)
		}
		if err := utils.ValidateToken(strings.Replace(h.Authorization, "Bearer ", "", 1)); err != nil {
			log.Println("error:", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}
