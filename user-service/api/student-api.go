package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	student "github.com/iamrosada/easy-life-server/user-server/internal/usecase/student"
)

type StudentHandlers struct {
	CreateStudentUseCase  *student.CreateStudentUseCase
	ListStudentsUseCase   *student.GetAllStudentUseCase
	DeleteStudentUseCase  *student.DeleteStudentUseCase
	GetStudentByIDUseCase *student.GetStudentByIDUseCase
	UpdateStudentUseCase  *student.UpdateStudentUseCase
}

func NewStudentHandlers(
	createStudentUseCase *student.CreateStudentUseCase,
	listStudentsUseCase *student.GetAllStudentUseCase,
	deleteStudentUseCase *student.DeleteStudentUseCase,
	getStudentByIDUseCase *student.GetStudentByIDUseCase,
	updateStudentUseCase *student.UpdateStudentUseCase,
) *StudentHandlers {
	return &StudentHandlers{
		CreateStudentUseCase:  createStudentUseCase,
		ListStudentsUseCase:   listStudentsUseCase,
		DeleteStudentUseCase:  deleteStudentUseCase,
		GetStudentByIDUseCase: getStudentByIDUseCase,
		UpdateStudentUseCase:  updateStudentUseCase,
	}
}

func (p *StudentHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		Students := api.Group("/Students")
		{
			Students.POST("/", p.CreateStudentHandler)
			Students.GET("/", p.ListStudentsHandler)
			Students.DELETE("/", p.DeleteStudentHandler)
			Students.GET("/:id", p.GetStudentByIDHandler)
			Students.PUT("/", p.UpdateStudentHandler)
		}

	}
}

func (p *StudentHandlers) CreateStudentHandler(c *gin.Context) {
	var input student.CreateStudentInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := p.CreateStudentUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

func (p *StudentHandlers) ListStudentsHandler(c *gin.Context) {
	output, err := p.ListStudentsUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (p *StudentHandlers) DeleteStudentHandler(c *gin.Context) {
	var input student.DeleteStudentInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := p.DeleteStudentUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (p *StudentHandlers) GetStudentByIDHandler(c *gin.Context) {
	id := c.Param("id")

	input := student.GetStudentByIDInputputDto{ID: id}
	output, err := p.GetStudentByIDUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (p *StudentHandlers) UpdateStudentHandler(c *gin.Context) {
	var input student.UpdateStudentInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := p.UpdateStudentUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
