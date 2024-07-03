package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest-gin/controllers"
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest-gin/database"
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutesTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}
func CreateAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}
func DeleteAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}
func TestStatusCodeHello(t *testing.T) {
	r := SetupRoutesTest()
	r.GET("/:nome", controllers.Hello)
	req, _ := http.NewRequest("GET", "/bruno", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
func TestShowAllAlunos(t *testing.T) {
	database.ConectDB()
	CreateAlunoMock()
	defer DeleteAlunoMock()
	r := SetupRoutesTest()
	r.GET("/alunos", controllers.ShowAllAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
func TestSearchAluno(t *testing.T) {
	database.ConectDB()
	CreateAlunoMock()
	defer DeleteAlunoMock()
	r := SetupRoutesTest()
	r.GET("/alunos/cpf/:cpf", controllers.SearchAluno)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
