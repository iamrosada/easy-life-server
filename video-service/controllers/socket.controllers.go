// package controllers

// import (
// 	"crypto/sha1"
// 	"encoding/hex"
// 	"net/http"

// 	"github.com/iamrosada/easy-life-server/video-server/interfaces"
// 	"github.com/iamrosada/easy-life-server/video-server/utils"

// 	"github.com/gin-gonic/gin"
// 	// "gorm.io/driver/sqlite"
// )

// // var db *gorm.DB

// ConnectSession - Given a host and a password returns the session object.
func ConnectSession(ctx *gin.Context) {
	// db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
	// 	return
	// }

	url := ctx.Param("url")

	// Assuming you have a 'sockets' table in your SQLite database
	var socket interfaces.Socket
	if err := db.Table("sockets").Where("hashedurl = ?", url).First(&socket).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Socket connection not found."})
		return
	}

	var input interfaces.Session
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assuming you have a 'sessions' table in your SQLite database
	var session interfaces.Session
	if err := db.Table("sessions").Where("id = ?", socket.SessionID).First(&session).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Session not found."})
		return
	}

	if !utils.ComparePasswords(session.Password, []byte(input.Password)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":  session.Title,
		"socket": socket.SocketURL,
	})
}

// GetSession - Checks if session exists.
func GetSession(ctx *gin.Context) {
	// db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
	// 	return
	// }

	id := ctx.Request.URL.Query()["url"][0]

	// Assuming you have a 'sockets' table in your SQLite database
	if err := db.Table("sockets").Where("hashedurl = ?", id).First(&struct{}{}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Socket connection not found."})
		return
	}

	ctx.Status(http.StatusOK)
}

// CreateSocket - Creates socket connection with given session
func CreateSocket(session interfaces.Session, ctx *gin.Context, id string) string {
	// db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
	// 	return ""
	// }

	// Assuming you have a 'sockets' table in your SQLite database
	var socket interfaces.Socket
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
