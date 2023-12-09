package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/iamrosada/easy-life-server/class-service/api"
	"github.com/iamrosada/easy-life-server/class-service/internal/entity"
	repository "github.com/iamrosada/easy-life-server/class-service/internal/infra"
	usecase "github.com/iamrosada/easy-life-server/class-service/internal/usecase/class"

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

	err = gormDB.AutoMigrate(&entity.Class{})
	if err != nil {
		panic(err)
	}
	// Callback functions for serialization and deserialization
	gormDB.Callback().Create().Before("gorm:before_create").Register("serializeStudentsIDs", serializeCodes)
	gormDB.Callback().Query().After("gorm:after_query").Register("deserializeStudentsIDs", deserializeCodes)

	// Create repositories and use cases
	ClassRepository := repository.NewClassRepositoryPostgres(gormDB)
	createClassUsecase := usecase.NewCreateClassUseCase(ClassRepository)
	listClasssUsecase := usecase.NewGetAllClassUseCase(ClassRepository)
	deleteClassUsecase := usecase.NewDeleteClassUseCase(ClassRepository)
	getClassByIDUsecase := usecase.NewGetClassByIDUseCase(ClassRepository)
	updateClassUsecase := usecase.NewUpdateClassUseCase(ClassRepository)

	// Create handlers
	ClassHandlers := api.NewClassHandlers(createClassUsecase, listClasssUsecase, deleteClassUsecase, getClassByIDUsecase, updateClassUsecase)

	// Set up Gin router
	router := gin.Default()

	// Set up Class routes
	ClassHandlers.SetupRoutes(router)

	// Start the server
	err = http.ListenAndServe(":8080", router)
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
