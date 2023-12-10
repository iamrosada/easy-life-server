package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	student "github.com/iamrosada/easy-life-server/user-server/internal/usecase/student"
)

type StudentHandlers struct {
	CreateStudentUseCase     *student.CreateStudentUseCase
	ListStudentsUseCase      *student.GetAllStudentUseCase
	DeleteStudentUseCase     *student.DeleteStudentUseCase
	GetStudentByIDUseCase    *student.GetStudentByIDUseCase
	UpdateStudentUseCase     *student.UpdateStudentUseCase
	ApplyEventStudentUseCase *student.ApplyEventStudentUseCase
}

func NewStudentHandlers(
	createStudentUseCase *student.CreateStudentUseCase,
	listStudentsUseCase *student.GetAllStudentUseCase,
	deleteStudentUseCase *student.DeleteStudentUseCase,
	getStudentByIDUseCase *student.GetStudentByIDUseCase,
	updateStudentUseCase *student.UpdateStudentUseCase,
	applyEventStudentUseCase *student.ApplyEventStudentUseCase,

) *StudentHandlers {
	return &StudentHandlers{
		CreateStudentUseCase:     createStudentUseCase,
		ListStudentsUseCase:      listStudentsUseCase,
		DeleteStudentUseCase:     deleteStudentUseCase,
		GetStudentByIDUseCase:    getStudentByIDUseCase,
		UpdateStudentUseCase:     updateStudentUseCase,
		ApplyEventStudentUseCase: applyEventStudentUseCase,
	}
}

func (p *StudentHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		Students := api.Group("/students")
		{
			Students.POST("/", p.CreateStudentHandler)
			Students.GET("/", p.ListStudentsHandler)
			Students.DELETE("/", p.DeleteStudentHandler)
			Students.GET("/:id", p.GetStudentByIDHandler)
			Students.PUT("/", p.UpdateStudentHandler)

			Students.POST("/event/:id", p.ApplyEventToStudentHandler)

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

func (p *StudentHandlers) ApplyEventToStudentHandler(c *gin.Context) {
	eventID := c.Param("id")

	var request struct {
		StudentsIDs []string `json:"students_ids"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := p.ApplyEventStudentUseCase.ApplyEvent(eventID, request.StudentsIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event applied successfully"})
}
