package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/iamrosada/easy-life-server/class-service/internal/usecase/class"
)

type ClassHandlers struct {
	CreateClassUseCase  *usecase.CreateClassUseCase
	ListClasssUseCase   *usecase.GetAllClassUseCase
	DeleteClassUseCase  *usecase.DeleteClassUseCase
	GetClassByIDUseCase *usecase.GetClassByIDUseCase
	UpdateClassUseCase  *usecase.UpdateClassUseCase
}

func NewClassHandlers(
	createClassUseCase *usecase.CreateClassUseCase,
	listClasssUseCase *usecase.GetAllClassUseCase,
	deleteClassUseCase *usecase.DeleteClassUseCase,
	getClassByIDUseCase *usecase.GetClassByIDUseCase,
	updateClassUseCase *usecase.UpdateClassUseCase,
) *ClassHandlers {
	return &ClassHandlers{
		CreateClassUseCase:  createClassUseCase,
		ListClasssUseCase:   listClasssUseCase,
		DeleteClassUseCase:  deleteClassUseCase,
		GetClassByIDUseCase: getClassByIDUseCase,
		UpdateClassUseCase:  updateClassUseCase,
	}
}

func (p *ClassHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		Classs := api.Group("/class")
		{
			Classs.POST("/", p.CreateClassHandler)
			Classs.GET("/", p.ListClasssHandler)
			Classs.DELETE("/", p.DeleteClassHandler)
			Classs.GET("/:id", p.GetClassByIDHandler)
			Classs.PUT("/", p.UpdateClassHandler)
		}

	}
}

func (p *ClassHandlers) CreateClassHandler(c *gin.Context) {
	var input usecase.CreateClassInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := p.CreateClassUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

func (p *ClassHandlers) ListClasssHandler(c *gin.Context) {
	output, err := p.ListClasssUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (p *ClassHandlers) DeleteClassHandler(c *gin.Context) {
	var input usecase.DeleteClassInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := p.DeleteClassUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (p *ClassHandlers) GetClassByIDHandler(c *gin.Context) {
	id := c.Param("id")

	input := usecase.GetClassByIDInputputDto{ID: id}
	output, err := p.GetClassByIDUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (p *ClassHandlers) UpdateClassHandler(c *gin.Context) {
	var input usecase.UpdateClassInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := p.UpdateClassUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

// func (p *ClassHandlers) GetClassEventsIDHandler(c *gin.Context) {
// 	class, err := p.ListClasssUseCase.Execute()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	for classId,_ := range class{
// 		if classId.students_ids

// 	}
// 	c.JSON(http.StatusOK, output)
// }
