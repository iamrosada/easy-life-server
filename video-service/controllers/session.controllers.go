package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/iamrosada/easy-life-server/video-server/interfaces"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Session struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Host     string `json:"host"`
	Password string `json:"password"`
}

type Socket struct {
	SessionID string `json:"session_id"`
	HashedURL string `json:"hashed_url"`
	SocketURL string `json:"socket_url"`
}

func main() {
	// Initialize the database
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto-migrate the schema
	db.AutoMigrate(&Session{}, &Socket{})

	// Set up the Gin router
	router := gin.Default()
	router.Use(disableCors())

	// Define routes
	router.POST("/session", CreateSession)
	router.POST("/connect/:url", ConnectSession)
	router.GET("/connect", GetSession)

	router.GET("/ws/:socket", func(c *gin.Context) {
		socket := c.Param("socket")
		wshandler(c.Writer, c.Request, socket)
	})
	// Start the server
	router.Run(":9000")
}

// CreateSession handles the creation of a session
func CreateSession(c *gin.Context) {
	var session Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a UUID for the session ID
	session.ID = GenerateUUID()

	fmt.Println(session)

	// Check if the database is initialized
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Create the session in the database
	if err := db.Create(&session).Error; err != nil {
		fmt.Println("Error creating session:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Create the socket connection and get the hashed URL
	hashedURL := CreateSocket(session, c, session.ID)

	c.JSON(http.StatusOK, gin.H{"socket": hashedURL})

	// c.JSON(http.StatusOK, gin.H{"session": session, "hashed_url": hashedURL})
}

// GenerateUUID generates a new UUID
func GenerateUUID() string {
	return uuid.New().String()
}

// ConnectSession handles connecting to a session with a provided URL
func ConnectSession(ctx *gin.Context) {
	url := ctx.Param("url")

	// Assuming you have a 'sockets' table in your SQLite database
	var socket Socket
	if err := db.Table("sockets").Where("hashed_url = ?", url).First(&socket).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Socket connection not found."})
		return
	}

	var input Session
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assuming you have a 'sessions' table in your SQLite database
	var session Session
	if err := db.Table("sessions").Where("id = ?", socket.SessionID).First(&session).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Session not found."})
		return
	}

	// Check if the provided password matches the session password
	if session.Password != input.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":  session.Title,
		"socket": socket.SocketURL,
	})
}

// GetSession checks if a session exists
func GetSession(ctx *gin.Context) {
	// You may implement this according to your needs
	// This route checks if a socket connection exists based on a URL
	// Assuming you have a 'sockets' table in your SQLite database
	url := ctx.Query("url")
	if err := db.Table("sockets").Where("hashed_url = ?", url).First(&Socket{}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Socket connection not found."})
		return
	}

	ctx.Status(http.StatusOK)
}

// CreateSocket creates a socket connection with the given session
func CreateSocket(session Session, ctx *gin.Context, id string) string {
	// Assuming you have a 'sockets' table in your SQLite database
	var socket Socket
	hashURL := hashSession(session.Host + session.Title)
	socketURL := hashSession(session.Host + session.Password)
	socket.SessionID = id
	socket.HashedURL = hashURL
	socket.SocketURL = socketURL

	db.Create(&socket)

	return hashURL
}

func hashSession(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var sockets = make(map[string]map[string]*interfaces.Connection)

func wshandler(w http.ResponseWriter, r *http.Request, socket string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error handling websocket connection.")
		return
	}

	defer conn.Close()

	if sockets[socket] == nil {
		sockets[socket] = make(map[string]*interfaces.Connection)
	}

	clients := sockets[socket]

	var message interfaces.Message
	for {
		err = conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		if clients[message.UserID] == nil {
			connection := new(interfaces.Connection)
			connection.Socket = conn
			clients[message.UserID] = connection
		}

		switch message.Type {
		case "connect":
			message.Type = "session_joined"
			err := conn.WriteJSON(message)
			if err != nil {
				log.Printf("Websocket error: %s", err)
				delete(clients, message.UserID)
			}
			break
		case "disconnect":
			for user, client := range clients {
				err := client.Send(message)
				if err != nil {
					client.Socket.Close()
					delete(clients, user)
				}
			}
			delete(clients, message.UserID)
			break
		default:
			for user, client := range clients {
				err := client.Send(message)
				if err != nil {
					delete(clients, user)
				}
			}
		}
	}
}

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
