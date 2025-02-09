package main

import (
	"net/http"

	socketio "github.com/doquangtan/socket.io/v4"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Handles Socket.IO events
func socketIoHandle(io *socketio.Io) {
	io.OnConnection(func(socket *socketio.Socket) {
		println("connect", socket.Nps, socket.Id)

		// When a client connects, they join the "demo" room
		socket.Join("demo")

		// Broadcast to all clients in "demo" that a new user has joined
		io.To("demo").Emit("test", socket.Id+" joined the room", "server message")

		// Listens for "test" event and responds with the same data (Echo)
		socket.On("test", func(event *socketio.EventPayload) {
			socket.Emit("test", event.Data...)
		})

		// Listens for "test-broadcast" event and responds with the same data (Echo) and send to all room clients
		socket.On("test-broadcast", func(event *socketio.EventPayload) {
			socket.To("demo").Emit("test", event.Data...)
		})

		// Logs when a client starts disconnecting
		socket.On("disconnecting", func(event *socketio.EventPayload) {
			println("disconnecting", socket.Nps, socket.Id)
		})

		// Logs when a client fully disconnects
		socket.On("disconnect", func(event *socketio.EventPayload) {
			println("disconnect", socket.Nps, socket.Id)
		})
	})
}

// Middleware to handle CORS (Cross-Origin Resource Sharing)
func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		// Handles preflight requests for CORS
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Remove "Origin" header to prevent conflicts
		c.Request.Header.Del("Origin")

		c.Next()
	}
}

// Initializes Gin web server and integrates Socket.IO
func usingWithGin() {
	io := socketio.New() // Create a new Socket.IO server
	socketIoHandle(io)   // Register event handlers

	router := gin.Default()

	// Apply CORS middleware to allow frontend (React) to connect
	router.Use(GinMiddleware("http://localhost:5173"))

	// Serve static files from the "./public" directory
	router.Use(static.Serve("/", static.LocalFile("./public", false)))

	// Route Socket.IO requests through Gin
	router.GET("/socket.io/", gin.WrapH(io.HttpHandler()))

	// Start the HTTP server on port 3300
	router.Run(":3300")
}

func main() {
	usingWithGin()
}
