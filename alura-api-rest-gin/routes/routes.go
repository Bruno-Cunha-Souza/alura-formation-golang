package routes

import (
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ShowAllAlunos)
	r.GET("/:nome", controllers.Hello)
	r.Run()
}
