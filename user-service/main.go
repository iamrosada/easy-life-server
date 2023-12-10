package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	// "github.com/iamrosada/easy-life-server/Student-server/internal/entity"
	// student "github.com/iamrosada/easy-life-server/class-service/internal/usecase/student"
	// teacher "github.com/iamrosada/easy-life-server/class-service/internal/usecase/teacher"

	"github.com/iamrosada/easy-life-server/user-server/api"
	"github.com/iamrosada/easy-life-server/user-server/internal/entity"
	repository "github.com/iamrosada/easy-life-server/user-server/internal/infra"
	"github.com/iamrosada/easy-life-server/user-server/internal/usecase/student"
	"github.com/iamrosada/easy-life-server/user-server/internal/usecase/teacher"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

	// Create handlers
	StudentHandlers := api.NewStudentHandlers(createStudentUsecase, listStudentsUsecase, deleteStudentUsecase, getStudentByIDUsecase, updateStudentUsecase, applyEventStudentUseCase)

	// Set up Gin router
	router := gin.Default()

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
