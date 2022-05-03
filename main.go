package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lenarbatdalov/app/configs"
	"github.com/lenarbatdalov/app/routes"
)

func main() {
	r := gin.New()

	configs.ConnectDB()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("application_session", store))
	r.Use(static.Serve("/", static.LocalFile("./assets/dist", false)))

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", gin.H{"title": "Home Page"})
	})

	routes.UserRoute(r)
	routes.ProductRoute(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
