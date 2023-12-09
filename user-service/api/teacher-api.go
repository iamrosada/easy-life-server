package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	teacher "github.com/iamrosada/easy-life-server/user-server/internal/usecase/teacher"
)

type TeacherHandlers struct {
	CreateTeacherUseCase  *teacher.CreateTeacherUseCase
	ListTeachersUseCase   *teacher.GetAllTeacherUseCase
	DeleteTeacherUseCase  *teacher.DeleteTeacherUseCase
	GetTeacherByIDUseCase *teacher.GetTeacherByIDUseCase
	UpdateTeacherUseCase  *teacher.UpdateTeacherUseCase
}

func NewTeacherHandlers(
	createTeacherUseCase *teacher.CreateTeacherUseCase,
	listTeachersUseCase *teacher.GetAllTeacherUseCase,
	deleteTeacherUseCase *teacher.DeleteTeacherUseCase,
	getTeacherByIDUseCase *teacher.GetTeacherByIDUseCase,
	updateTeacherUseCase *teacher.UpdateTeacherUseCase,
) *TeacherHandlers {
	return &TeacherHandlers{
		CreateTeacherUseCase:  createTeacherUseCase,
		ListTeachersUseCase:   listTeachersUseCase,
		DeleteTeacherUseCase:  deleteTeacherUseCase,
		GetTeacherByIDUseCase: getTeacherByIDUseCase,
		UpdateTeacherUseCase:  updateTeacherUseCase,
	}
}

func (p *TeacherHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		Teachers := api.Group("/Teachers")
		{
			Teachers.POST("/", p.CreateTeacherHandler)
			Teachers.GET("/", p.ListTeachersHandler)
			Teachers.DELETE("/", p.DeleteTeacherHandler)
			Teachers.GET("/:id", p.GetTeacherByIDHandler)
			Teachers.PUT("/", p.UpdateTeacherHandler)
		}

	}
}

func (p *TeacherHandlers) CreateTeacherHandler(c *gin.Context) {
	var input teacher.CreateTeacherInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := p.CreateTeacherUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

func (p *TeacherHandlers) ListTeachersHandler(c *gin.Context) {
	output, err := p.ListTeachersUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (p *TeacherHandlers) DeleteTeacherHandler(c *gin.Context) {
	var input teacher.DeleteTeacherInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := p.DeleteTeacherUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (p *TeacherHandlers) GetTeacherByIDHandler(c *gin.Context) {
	id := c.Param("id")

	input := teacher.GetTeacherByIDInputputDto{ID: id}
	output, err := p.GetTeacherByIDUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (p *TeacherHandlers) UpdateTeacherHandler(c *gin.Context) {
	var input teacher.UpdateTeacherInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := p.UpdateTeacherUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
