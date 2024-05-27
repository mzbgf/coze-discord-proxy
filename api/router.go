// @title COZE-DISCORD-PROXY
// @version 1.0.0
// @description COZE-DISCORD-PROXY 代理服务
// @BasePath
package handler

import (
	"context"
	"coze-discord-proxy/common"
	"coze-discord-proxy/common/config"
	"coze-discord-proxy/discord"
	"coze-discord-proxy/middleware"
	"coze-discord-proxy/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var server *gin.Engine

func init() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go discord.StartBot(ctx, discord.BotToken)

	common.SetupLogger()
	common.SysLog("COZE-DISCORD-PROXY " + common.Version + " started")
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	if config.DebugEnabled {
		common.SysLog("running in debug mode")
	}

	// Initialize HTTP server
	server = gin.New()

	server.Use(gin.Recovery())
	server.Use(middleware.RequestId())
	middleware.SetUpLogger(server)
	router.SetApiRouter(server)
}

func Listen(w http.ResponseWriter, r *http.Request) {
	server.ServeHTTP(w, r)
}
