package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	student "github.com/iamrosada/easy-life-server/user-server/internal/usecase/student"
)

type StudentHandlers struct {
	CreateStudentUseCase          *student.CreateStudentUseCase
	ListStudentsUseCase           *student.GetAllStudentUseCase
	DeleteStudentUseCase          *student.DeleteStudentUseCase
	GetStudentByIDUseCase         *student.GetStudentByIDUseCase
	UpdateStudentUseCase          *student.UpdateStudentUseCase
	ApplyEventStudentUseCase      *student.ApplyEventStudentUseCase
	ListStudentsByTeacherIDUseCae *student.ListStudentsByTeacherIDUseCase
	GetStudentByEmailUseCase      *student.GetStudentByEmailUseCase
	GetStudentByEventIDUseCase    *student.GetStudentByEventIDUseCase
}

func NewStudentHandlers(
	createStudentUseCase *student.CreateStudentUseCase,
	listStudentsUseCase *student.GetAllStudentUseCase,
	deleteStudentUseCase *student.DeleteStudentUseCase,
	getStudentByIDUseCase *student.GetStudentByIDUseCase,
	updateStudentUseCase *student.UpdateStudentUseCase,
	applyEventStudentUseCase *student.ApplyEventStudentUseCase,
	listStudentsByTeacherIDUseCae *student.ListStudentsByTeacherIDUseCase,
	getStudentByEmailUseCase *student.GetStudentByEmailUseCase,
	getStudentByEventIDUseCase *student.GetStudentByEventIDUseCase,

) *StudentHandlers {
	return &StudentHandlers{
		CreateStudentUseCase:          createStudentUseCase,
		ListStudentsUseCase:           listStudentsUseCase,
		DeleteStudentUseCase:          deleteStudentUseCase,
		GetStudentByIDUseCase:         getStudentByIDUseCase,
		UpdateStudentUseCase:          updateStudentUseCase,
		ApplyEventStudentUseCase:      applyEventStudentUseCase,
		ListStudentsByTeacherIDUseCae: listStudentsByTeacherIDUseCae,
		GetStudentByEmailUseCase:      getStudentByEmailUseCase,
		GetStudentByEventIDUseCase:    getStudentByEventIDUseCase,
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

			Students.GET("/teacher-id/:id/students", p.GetStudentsUsingTeacherIDHandler)

			Students.GET("/email/:email", p.GetStudentByEmailHandler)

			Students.GET("/event/:id/students", p.GetStudentByEventIDHandler)

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
func (p *StudentHandlers) GetStudentByEmailHandler(c *gin.Context) {
	email := c.Param("email")

	input := student.GetStudentByEmailInputputDto{Email: email}
	output, err := p.GetStudentByEmailUseCase.Execute(input)
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

func (p *StudentHandlers) GetStudentByEventIDHandler(c *gin.Context) {
	id := c.Param("id")

	input := student.GetStudentByEventIDInputputDto{EventID: id}

	output, err := p.GetStudentByEventIDUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)

}
func (p *StudentHandlers) GetStudentsUsingTeacherIDHandler(c *gin.Context) {

	id := c.Param("id")

	output, err := p.ListStudentsByTeacherIDUseCae.ListStudentsByTeacherID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
