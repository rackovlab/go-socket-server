# **Go Socket.IO v4 Server with Gin**

This project is a **Go server** that integrates **Socket.IO v4** for real-time WebSocket communication and **Gin** for HTTP routing and middleware. It allows a React.js frontend to connect and exchange messages in real time. Find the React JS code here `https://github.com/rackovlab/reactjs-socket-io-v4-client`

## **Features**

- WebSocket communication using **Socket.IO** utilizing `github.com/doquangtan/socket.io/v4` Socket.IO v4
- HTTP server using **Gin**.
- CORS middleware for cross-origin access.
- Static file serving.
- Real-time event handling (clients can send and receive messages).

---

## **Setup Instructions**

### **1. Clone the Repository**

```sh
git clone https://github.com/rackovlab/go-socket-server.git
cd go-socketio-server
```

### **2. Install Dependencies**

Ensure you have Go installed (Go 1.18+ recommended). Then, install required dependencies:

```sh
go mod tidy
```

### **3. Run the Server**

Start the server with:

```sh
go run main.go
```

The server will run on http://localhost:3300.
Ensure the React JS client is running at http://localhost:5173 (or update the client Url in main.go to match your client address).
