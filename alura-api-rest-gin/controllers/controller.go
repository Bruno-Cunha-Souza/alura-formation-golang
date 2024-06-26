package controllers

import (
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest-gin/models"
	"github.com/gin-gonic/gin"
)

func ShowAllAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Hello(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API:": "Hello " + nome,
	})
}
