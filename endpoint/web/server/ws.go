package server

import (
	"feature-distributor/common/subscribe"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var wsHandle gin.HandlerFunc = func(c *gin.Context) {
	upgrade := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "upgrade to websocket failed",
		})
		c.Abort()
		return
	}
	clientName := "myName"
	subscribe.Sub(clientName, func(event subscribe.ChannelEvent) {
		err := ws.WriteJSON(event)
		if err != nil {
			logrus.Warnf("send event to ws failed: %v", err)
		}
	})
	for {
		_, _, err = ws.ReadMessage()
		if err != nil {
			subscribe.Unsub(clientName)
			err = ws.Close()
			if err != nil {
				logrus.Warnf("close ws failed: %v", err)
			}
			break
		}
	}
}
