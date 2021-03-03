package scheduler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *Scheduler) ListenAndServe() error {
	app := gin.New()
	app.Use(gin.Recovery())
	//gin.SetMode(gin.ReleaseMode)
	if s.cfg.Debug {
		app.Use(gin.Logger())
	}

	s.registerRouter(app)

	fmt.Println("Run Plumber ...")
	return app.Run(s.cfg.ListenAddr)
}

func (s *Scheduler) registerRouter(app *gin.Engine) {
	// 建立新的监听
	app.POST("/new_monitor", s.newMonitor)

	// 获取所有进行的监听
	app.GET("/all_monitor", s.allMonitor)

	// 停止某个监听
	app.POST("/stop_monitor/:task_id", s.stopMonitor)
}
