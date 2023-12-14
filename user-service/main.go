package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	// Import the CORS middleware
	"github.com/gin-gonic/gin"

	"github.com/iamrosada/easy-life-server/user-server/api"
	"github.com/iamrosada/easy-life-server/user-server/internal/entity"
	repository "github.com/iamrosada/easy-life-server/user-server/internal/infra"
	"github.com/iamrosada/easy-life-server/user-server/internal/usecase/student"
	"github.com/iamrosada/easy-life-server/user-server/internal/usecase/teacher"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func disableCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		// I added this for another handler of mine,
		// but I do not think this is necessary for GraphQL's handler
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.Writer.WriteHeader(http.StatusOK)
			return
		}

		c.Next()
	}
}
func main() {
	dbPath := "./db/main.db"
	sqlDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	_, err = os.Stat(dbPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			panic(err)
		}

		file, err := os.Create(dbPath)
		if err != nil {
			panic(err)
		}
		file.Close()
	}

	// Create Gorm connection
	gormDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Set up Gin router with CORS middleware
	router := gin.Default()

	// Configure CORS middleware

	router.Use(disableCors())

	err = gormDB.AutoMigrate(&entity.Student{}, &entity.Teacher{})
	if err != nil {
		panic(err)
	}

	// Callback functions for serialization and deserialization
	gormDB.Callback().Create().Before("gorm:before_create").Register("serializeTeachersIDs", serializeCodes)
	gormDB.Callback().Query().After("gorm:after_query").Register("deserializeTeachersIDs", deserializeCodes)

	// Create repositories and use cases
	StudentRepository := repository.NewStudentRepositoryPostgres(gormDB)
	createStudentUsecase := student.NewCreateStudentUseCase(StudentRepository)
	listStudentsUsecase := student.NewGetAllStudentUseCase(StudentRepository)
	deleteStudentUsecase := student.NewDeleteStudentUseCase(StudentRepository)
	getStudentByIDUsecase := student.NewGetStudentByIDUseCase(StudentRepository)
	updateStudentUsecase := student.NewUpdateStudentUseCase(StudentRepository)
	applyEventStudentUseCase := student.NewCreateEventStudentUseCase(StudentRepository)
	listStudentsByTeacherIDUseCase := student.NewListStudentsByTeacherIDUseCase(StudentRepository)
	getStudentByEmailUseCase := student.NewGetStudentByEmailUseCase(StudentRepository)
	getStudentByEventIDUseCase := student.NewGetStudentByEventIDUseCase(StudentRepository)

	// Create handlers
	StudentHandlers := api.NewStudentHandlers(createStudentUsecase, listStudentsUsecase, deleteStudentUsecase, getStudentByIDUsecase, updateStudentUsecase, applyEventStudentUseCase, listStudentsByTeacherIDUseCase, getStudentByEmailUseCase, getStudentByEventIDUseCase)

	TeacherRepository := repository.NewTeacherRepositoryPostgres(gormDB)
	createTeacherUsecase := teacher.NewCreateTeacherUseCase(TeacherRepository)
	listTeachersUsecase := teacher.NewGetAllTeacherUseCase(TeacherRepository)
	deleteTeacherUsecase := teacher.NewDeleteTeacherUseCase(TeacherRepository)
	getTeacherByIDUsecase := teacher.NewGetTeacherByIDUseCase(TeacherRepository)
	updateTeacherUsecase := teacher.NewUpdateTeacherUseCase(TeacherRepository)

	TeacherHandlers := api.NewTeacherHandlers(createTeacherUsecase, listTeachersUsecase, deleteTeacherUsecase, getTeacherByIDUsecase, updateTeacherUsecase)

	// Set up Student routes
	StudentHandlers.SetupRoutes(router)
	TeacherHandlers.SetupRoutes(router)

	// Start the server

	err = http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err)
	}
}

func serializeCodes(db *gorm.DB) {
	if serializable, ok := db.Statement.Dest.(entity.Serializable); ok {
		serializable.BeforeSave()
	}
}

func deserializeCodes(db *gorm.DB) {
	if serializable, ok := db.Statement.Dest.(entity.Serializable); ok {
		serializable.AfterFind()
	}
}
