package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iamrosada/easy-life-server/video-server/controllers"
	"github.com/iamrosada/easy-life-server/video-server/interfaces"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gorilla/websocket"
)

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

	gormDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	router := gin.Default()

	// Configure CORS middleware

	router.Use(disableCors())

	err = gormDB.AutoMigrate(&interfaces.Session{}, &interfaces.Socket{})
	if err != nil {
		panic(err)
	}
	// middleware - intercept requests to use our db controller
	// router.Use(func(context *gin.Context) {
	// 	context.Set("db", client)
	// 	context.Next()
	// })

	// REST API
	router.POST("/session", controllers.CreateSession)
	router.GET("/connect", controllers.GetSession)
	router.POST("/connect/:url", controllers.ConnectSession)

	// Websocket connection
	router.GET("/ws/:socket", func(c *gin.Context) {
		socket := c.Param("socket")
		wshandler(c.Writer, c.Request, socket)
	})

	// router.Run("0.0.0.0:" + getenv("PORT", "9000"))

	port := ":9000"
	fmt.Printf("Servidor rodando na porta %s\n", port)
	router.Run(port)

	err = http.ListenAndServe(port, router)
	if err != nil {
		fmt.Println(err)
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
