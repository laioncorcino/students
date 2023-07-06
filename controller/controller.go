package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/laioncorcino/students/database"
	"github.com/laioncorcino/students/entity"
	"net/http"
)

func GetAll(c *gin.Context) {
	var students []entity.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetById(c *gin.Context) {
	var s entity.Student
	studentID := c.Params.ByName("studentID")
	database.DB.First(&s, studentID)

	if s.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, s)
}

func Create(c *gin.Context) {
	var s entity.Student
	err := c.ShouldBindJSON(&s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}
	database.DB.Create(&s)
	c.JSON(http.StatusCreated, s)
}

func Delete(c *gin.Context) {
	var s entity.Student
	studentID := c.Params.ByName("studentID")
	database.DB.Delete(&s, studentID)
	c.JSON(http.StatusNoContent, nil)
}
