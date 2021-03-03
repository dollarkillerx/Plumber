package scheduler

import (
	"fmt"
	"log"

	"github.com/dollarkillerx/plumber/internal/mq_manager"
	"github.com/dollarkillerx/plumber/internal/task"
	"github.com/dollarkillerx/plumber/pkg/newsletter"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Scheduler) newMonitor(ctx *gin.Context) {
	var taskConfig newsletter.TaskConfig
	if err := ctx.BindJSON(&taskConfig); err != nil {
		ctx.JSON(400, gin.H{"error": fmt.Sprintf("%+v", err)})
		return
	}

	mq, err := mq_manager.MQManager.InitMQManager(taskConfig)
	if err != nil {
		log.Printf("%+v \n", err)
		ctx.JSON(400, gin.H{"error": fmt.Sprintf("%+v", err)})
		return
	}

	t := task.New(mq, taskConfig)

	taskID := uuid.New().String()
	s.tasks[taskID] = t
	go func() {
		if err := t.Synchronize(); err != nil {
			log.Printf("%+v \n", err)
			delete(s.tasks, taskID)
		}
	}()

	ctx.JSON(200, newsletter.TaskResponse{TaskID: taskID, Success: true})
}

func (s *Scheduler) allMonitor(ctx *gin.Context) {
	r := make([]gin.H, 0)
	for k, v := range s.tasks {
		r = append(r, gin.H{
			"TaskID": k,
			"DB":     v.Cfg.DBConfig.DBName,
			"Table":  v.Cfg.DBConfig.TableName,
		})
	}

	ctx.JSON(200, gin.H{"body": r, "success": true})
}

func (s *Scheduler) stopMonitor(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	t, ex := s.tasks[taskID]
	if !ex {
		ctx.JSON(400, gin.H{"error": "not found task"})
		return
	}

	if err := t.Close(); err != nil {
		log.Printf("%+v \n", err)
		ctx.JSON(500, gin.H{"error": fmt.Sprintf("%+v", err)})
		return
	}

	delete(s.tasks, taskID)
	ctx.JSON(200, gin.H{"success": true})
}
