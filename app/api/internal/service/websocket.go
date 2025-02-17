package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gocument/app/api/global"
	"net/http"
	"sync"
)

var clients = make(map[*websocket.Conn]bool)
var lock = sync.Mutex{}
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		global.Logger.Error("websocket连接失败" + err.Error())
		return
	}
	defer func(conn *websocket.Conn) {
		if err := conn.Close(); err != nil {
			global.Logger.Error("关闭websocket连接失败" + err.Error())
		}
	}(conn)
	lock.Lock()
	clients[conn] = true
	lock.Unlock()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			global.Logger.Error("读取websocket消息失败" + err.Error())
			lock.Lock()
			delete(clients, conn)
			lock.Unlock()
			break
		}
		global.Logger.Info("收到消息：" + string(msg))
		broadcast(msg)
	}

}

func broadcast(msg []byte) {
	lock.Lock()
	defer lock.Unlock()
	for conn := range clients {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			global.Logger.Error("发送websocket消息失败" + err.Error())
			return
		}
	}
}
