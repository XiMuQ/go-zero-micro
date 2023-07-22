package main

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"go-zero-micro/asynq/others_demo/async_task_demo/async_task_task"
	"go-zero-micro/common/utils"
	"log"
	"time"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	emailBody := fmt.Sprintf("定时任务邮件已发送，发送时间：%s", time.Now().Format(utils.DateTimeFormat))
	retryEmail := async_task_task.AsyncEmailPayload{
		To:      "user@example.com",
		Subject: "定时任务邮件",
		Body:    emailBody,
	}
	task, err := NewRetryEmailTask(retryEmail)
	if err != nil {
		log.Fatal(err)
	}
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	info, err := client.Enqueue(task, asynq.MaxRetry(5), asynq.Timeout(1*time.Minute))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

// NewRetryEmailTask 创建重试电子邮件任务的函数
func NewRetryEmailTask(asyncEmail async_task_task.AsyncEmailPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(asyncEmail)
	if err != nil {
		return nil, err
	}
	//任务级别：创建任务时设置重试次数、超时时间
	task := asynq.NewTask(async_task_task.AsyncEmailTask, payload, asynq.MaxRetry(5), asynq.Timeout(1*time.Minute))
	return task, err
}
