package controllers

import (
	"net/http"

	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest-gin/database"
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest-gin/models"
	"github.com/gin-gonic/gin"
)

func ShowAllAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Hello(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API:": "Hello " + nome,
	})
}
func CreateAluno(c *gin.Context) {
	// cria a variavel com base na struct auno
	var aluno models.Aluno

	// faz a checagem de erro
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: ": err.Error()})
		return
	}
	if err := models.ValidaAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: ": err.Error()})
		return
	}

	// se não retornar erro, salva o novo aluno no DB
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
func AlunoId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
func DeleteAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}
func EditAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	// Atualizar os campos do aluno com os dados da solicitação JSON
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validação do aluno
	if err := models.ValidaAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func SearchAluno(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
