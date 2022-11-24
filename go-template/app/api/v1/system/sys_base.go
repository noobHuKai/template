package system

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/common/response"
)

var (
	ExitMsg = []byte("exit")
)

type BaseApi struct{}

func (b *BaseApi) GetSysBasicInfo(c *gin.Context) {
	basicInfo := baseService.GetBasicInfo()

	response.OkWithData(c, basicInfo)
}

func (b *BaseApi) GetSysMonitorInfo(ctx *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := g.WSUpGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		g.Logger.Error(err.Error())
		return
	}
	defer ws.Close()
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			g.Logger.Error(err.Error())
			break
		}
		if bytes.Compare(ExitMsg, msg) == 0 {
			break
		}

		monitorInfo := baseService.GetMonitorInfo()
		byteData, err := monitorInfo.Bytes()
		if err != nil {
			g.Logger.Error(err.Error())
			continue
		}

		err = ws.WriteMessage(websocket.TextMessage, byteData)
		if err != nil {
			break
		}
	}
}
