package initialize

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	r := initRouter()

	if g.Cfg.Mode == "release" {
		gin.SetMode(gin.DebugMode)
	} else if g.Cfg.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	//r.Run(g.Cfg.Server.Addr())

	// Graceful restart or stop
	// go must >= 1.18
	srv := &http.Server{
		Addr:    g.Cfg.Server.Addr(),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			g.Logger.Info("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	g.Logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		g.Logger.Info("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		g.Logger.Info("timeout of 5 seconds.")
	}
	g.Logger.Info("Server exiting")
}
