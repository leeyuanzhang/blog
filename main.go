package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo"
	_ "github.com/joho/godotenv/autoload"
	"blog/db"
	"blog/routes"
	"os"
)

func main() {
	router := gin.New()
	loadMiddlewares(router)
	loadRouters(router)
	router.Run(":" + os.Getenv("PORT"))
}

func loadMiddlewares(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/**/**/*")

	// session
	store := mongo.NewStore(db.Sessions, 3600, true, []byte("secret"))
	r.Use(sessions.Sessions("goulang", store))
}

func loadRouters(r *gin.Engine) {
	r.GET("", routes.Index)

}
