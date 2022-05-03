# Статьи
Отличная статья про работу с [MongoDB](https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver-ru),
[habr mongodb](https://habr.com/ru/post/433776/)
<br><br>

# Создание проекта
<div>
go mod init lenarbatdalov/app <br>
Надо было именовать github.com/lenarbatdalov/app, для этого выполнил: <br>
go mod edit -module github.com/lenarbatdalov/app <br>
</div><br><br>

# Зависимости
<div>
go get -u github.com/gin-gonic/gin <br>
go get -u github.com/gin-contrib/sessions <br>
go get github.com/gin-contrib/static <br>
go get go.mongodb.org/mongo-driver <br>
go get go.mongodb.org/mongo-driver/bson/primitive <br>
go get github.com/joho/godotenv <br>
go mod tidy
</div><br><br>

# other
Пакет [cli](https://github.com/urfave/cli) для оказания помощи в разработке программы диспетчера задач: <br>
go get github.com/urfave/cli/v2 <br>

Пакет [color](https://github.com/gookit/color) для придания цвета выводимым результатам: <br>
go get gopkg.in/gookit/color.v1 <br><br>

# Выполнить cli задачи
go run main.go <br>
go run main.go add "Learn Go" <br>
go run main.go add "Read a book" <br>
go run main.go all <br>
go run main.go done "Learn Go" <br>
go run main.go finished <br>
go run main.go rm "Read a book" <br><br>


# Go-gin simple usage
```
	r := gin.New()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(static.Serve("/", static.LocalFile("./assets/dist", false)))

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", gin.H{"title": "Home Page"})
	})

	// auth := r.Group("/")
	// {
	// 	auth.GET("/", controller.DefaultController)
	// 	auth.GET("/login", s.loginController.LoginPage)
	// 	auth.POST("/login", s.loginController.Login)
	// 	auth.GET("/logout", s.loginController.Logout)
	// }

	// err := r.Run(":8080")
	// if err != nil {
	// 	panic(err)
	// }
```
