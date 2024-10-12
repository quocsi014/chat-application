package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool{
    return true
  },
}

func HandleMessageWebSocket() func(*gin.Context){
  return func(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
      c.JSON(http.StatusInternalServerError, "could not upgrade connection")
      return
    }
    defer conn.Close()
    for{
       messageType, msg, err := conn.ReadMessage()
            if err != nil {
                log.Println("Error reading message:", err)
                break // Thoát vòng lặp nếu có lỗi
            }

            log.Printf("Received: %s", msg)

            // Xử lý tin nhắn ở đây (ví dụ: bạn có thể phân tích cú pháp và thực hiện các hành động dựa trên nội dung)

            // Gửi phản hồi lại cho client
            err = conn.WriteMessage(messageType, msg) // Echo lại tin nhắn
            if err != nil {
                log.Println("Error writing message:", err)
                break // Thoát vòng lặp nếu có lỗi
            }
    }
  }
}
