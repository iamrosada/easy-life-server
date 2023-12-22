// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"

// 	"context"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// type Call struct {
// 	gorm.Model
// 	CallID           string `json:"callId"`
// 	OfferCandidates  []Candidate
// 	AnswerCandidates []Candidate
// }

// type Candidate struct {
// 	gorm.Model
// 	CandidateData map[string]interface{} `json:"candidate"`
// }

// var sampleCollection *mongo.Collection

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(
// 		ctx,
// 		options.Client().ApplyURI("mongodb://michael:secret@localhost:27017/"),
// 	)
// 	if err != nil {
// 		log.Fatalf("connection error: %v", err)
// 	}

// 	// Check the connection
// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatalf("ping MongoDB error: %v", err)
// 	}
// 	fmt.Println("ping success")

// 	// Set up the sampleCollection
// 	database := client.Database("demo")
// 	sampleCollection = database.Collection("sampleCollection")
// 	sampleCollection.Drop(ctx)

// 	// Create a new Gin router
// 	router := gin.Default()

// 	// Define routes
// 	router.POST("/api/createCall", createCall)
// 	router.POST("/api/calls/:callId/offerCandidates", addOfferCandidate)
// 	router.GET("/api/calls/:callId/answerCandidates", getAnswerCandidates)
// 	router.POST("/api/calls/:callId/answerCandidates", addAnswerCandidate)

// 	router.POST("/api/insertOne", insertOne)
// 	router.GET("/api/queryAll", queryAll)
// 	router.POST("/api/insertMany", insertMany)
// 	router.GET("/api/querySpecific", querySpecific)
// 	router.PUT("/api/update", update)
// 	router.DELETE("/api/deleteMany", deleteMany)

// 	// Run the server
// 	port := ":3000"
// 	fmt.Printf("Server running on port %s\n", port)
// 	router.Run(port)
// }

// // Handler to insert one document
// func insertOne(c *gin.Context) {
// 	var document bson.M
// 	if err := c.BindJSON(&document); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	insertedResult, err := sampleCollection.InsertOne(context.Background(), document)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"insertedID": insertedResult.InsertedID})
// }

// // Handler to query all documents
// func queryAll(c *gin.Context) {
// 	cursor, err := sampleCollection.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}
// 	defer cursor.Close(context.Background())

// 	var queryResult []bson.M
// 	if err := cursor.All(context.Background(), &queryResult); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, queryResult)
// }

// // Handler to insert many documents
// func insertMany(c *gin.Context) {
// 	var documents []interface{}
// 	if err := c.BindJSON(&documents); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	insertedResult, err := sampleCollection.InsertMany(context.Background(), documents)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"insertedIDs": insertedResult.InsertedIDs})
// }

// // Handler to query specific documents
// func querySpecific(c *gin.Context) {
// 	var filter bson.M
// 	if err := c.BindJSON(&filter); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	cursor, err := sampleCollection.Find(context.Background(), filter)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}
// 	defer cursor.Close(context.Background())

// 	var queryResult []bson.M
// 	if err := cursor.All(context.Background(), &queryResult); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, queryResult)
// }

// // Handler to update documents
// func update(c *gin.Context) {
// 	var filter, update bson.M
// 	if err := c.BindJSON(&filter); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := c.BindJSON(&update); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	updateResult, err := sampleCollection.UpdateMany(context.Background(), filter, bson.M{"$set": update})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"modifiedCount": updateResult.ModifiedCount})
// }

// // Handler to delete documents
// func deleteMany(c *gin.Context) {
// 	var filter bson.M
// 	if err := c.BindJSON(&filter); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	deleteResult, err := sampleCollection.DeleteMany(context.Background(), filter)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"deletedCount": deleteResult.DeletedCount})
// }

// // ----
// func createCall(c *gin.Context) {
// 	var call Call
// 	if err := c.BindJSON(&call); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := db.Create(&call).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
// 		return
// 	}

// 	offerCandidatesEndpoint := fmt.Sprintf("/api/calls/%s/offerCandidates", call.CallID)
// 	answerCandidatesEndpoint := fmt.Sprintf("/api/calls/%s/answerCandidates", call.CallID)

// 	c.JSON(http.StatusOK, gin.H{"offerCandidatesEndpoint": offerCandidatesEndpoint, "answerCandidatesEndpoint": answerCandidatesEndpoint})
// }

// func addOfferCandidate(c *gin.Context) {
// 	callID := c.Param("callId")

// 	var candidate Candidate
// 	if err := c.BindJSON(&candidate); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var call Call
// 	if err := db.Where("call_id = ?", callID).First(&call).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
// 		return
// 	}

// 	if err := db.Model(&call).Association("OfferCandidates").Append(&candidate).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }

// func getAnswerCandidates(c *gin.Context) {
// 	callID := c.Param("callId")

// 	var call Call
// 	if err := db.Where("call_id = ?", callID).Preload("AnswerCandidates").First(&call).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, call.AnswerCandidates)
// }

// func addAnswerCandidate(c *gin.Context) {
// 	callID := c.Param("callId")

// 	var candidate Candidate
// 	if err := c.BindJSON(&candidate); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var call Call
// 	if err := db.Where("call_id = ?", callID).First(&call).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
// 		return
// 	}

// 	if err := db.Model(&call).Association("AnswerCandidates").Append(&candidate).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }
