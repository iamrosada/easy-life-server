package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Candidate []string

type Call struct {
	gorm.Model
	CallID           string    `json:"callId"`
	OfferCandidates  Candidate `gorm:"type:VARCHAR(255)" json:"offer_candidates"`
	AnswerCandidates Candidate `gorm:"type:VARCHAR(255)" json:"answer_candidates"`
}

// type Candidate struct {
// 	gorm.Model
// 	CandidateData map[string]interface{} `json:"candidate"`
// }

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
	var err error

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

	r := gin.Default()
	// r.Use(gin.Recovery())
	r.Use(disableCors())

	err = gormDB.AutoMigrate(&Call{})
	if err != nil {
		panic(err)
	}
	r.POST("/api/createCall", createCall)
	r.POST("/api/calls/:callId/offerCandidates", addOfferCandidate)
	r.GET("/api/calls/:callId/answerCandidates", getAnswerCandidates)
	r.POST("/api/calls/:callId/answerCandidates", addAnswerCandidate)

	port := ":3001"
	fmt.Printf("Servidor rodando na porta %s\n", port)
	r.Run(port)

	err = http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println(err)
	}
}

func createCall(c *gin.Context) {
	var call Call
	if err := c.BindJSON(&call); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&call).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
		return
	}

	offerCandidatesEndpoint := fmt.Sprintf("/api/calls/%s/offerCandidates", call.CallID)
	answerCandidatesEndpoint := fmt.Sprintf("/api/calls/%s/answerCandidates", call.CallID)

	c.JSON(http.StatusOK, gin.H{"offerCandidatesEndpoint": offerCandidatesEndpoint, "answerCandidatesEndpoint": answerCandidatesEndpoint})
}

func addOfferCandidate(c *gin.Context) {
	callID := c.Param("callId")

	var candidate Candidate
	if err := c.BindJSON(&candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var call Call
	if err := db.Where("call_id = ?", callID).First(&call).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
		return
	}

	if err := db.Model(&call).Association("OfferCandidates").Append(&candidate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
		return
	}

	c.Status(http.StatusOK)
}

func getAnswerCandidates(c *gin.Context) {
	callID := c.Param("callId")

	var call Call
	if err := db.Where("call_id = ?", callID).Preload("AnswerCandidates").First(&call).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
		return
	}

	c.JSON(http.StatusOK, call.AnswerCandidates)
}

func addAnswerCandidate(c *gin.Context) {
	callID := c.Param("callId")

	var candidate Candidate
	if err := c.BindJSON(&candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var call Call
	if err := db.Where("call_id = ?", callID).First(&call).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
		return
	}

	if err := db.Model(&call).Association("AnswerCandidates").Append(&candidate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
		return
	}

	c.Status(http.StatusOK)
}
