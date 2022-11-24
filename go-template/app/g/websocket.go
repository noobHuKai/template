package g

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var WSUpGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
