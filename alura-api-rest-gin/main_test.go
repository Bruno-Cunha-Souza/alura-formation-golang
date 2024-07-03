package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
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
func TestStatusCodeHelloHandler(t *testing.T) {
	r := SetupRoutesTest()
	r.GET("/:nome", controllers.Hello)
	req, _ := http.NewRequest("GET", "/bruno", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
func TestShowAllAlunosHandler(t *testing.T) {
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
func TestSearchAlunoHandler(t *testing.T) {
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
func TestAlunoHandler(t *testing.T) {
	database.ConectDB()
	CreateAlunoMock()
	defer DeleteAlunoMock()
	r := SetupRoutesTest()
	r.GET("/alunos/:id", controllers.AlunoId)
	pathID := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathID, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var nameMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &nameMock)
	assert.Equal(t, "Aluno Teste", nameMock.Nome)
	assert.Equal(t, "12345678901", nameMock.CPF)
	assert.Equal(t, "123456789", nameMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
func TestDeleteHandler(t *testing.T) {
	database.ConectDB()
	CreateAlunoMock()
	r := SetupRoutesTest()
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	pathID := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathID, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
func TestEditAlunoHandler(t *testing.T) {
	database.ConectDB()
	CreateAlunoMock()
	defer DeleteAlunoMock()
	r := SetupRoutesTest()
	r.PATCH("/alunos/:id", controllers.EditAluno)
	aluno := models.Aluno{Nome: "Aluno Testi", CPF: "44345678901", RG: "223456781"}
	valueJson, _ := json.Marshal(aluno)
	pathAluno := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathAluno, bytes.NewBuffer(valueJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockEdit models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockEdit)
	assert.Equal(t, "Aluno Testi", alunoMockEdit.Nome)
	assert.Equal(t, "44345678901", alunoMockEdit.CPF)
	assert.Equal(t, "223456781", alunoMockEdit.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
