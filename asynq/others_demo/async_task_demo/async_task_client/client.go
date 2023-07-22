package main

import (
	"fmt"
	"github.com/hibiken/asynq"
	"go-zero-micro/asynq/others_demo/async_task_demo/async_task_task"
	"go-zero-micro/common/utils"
	"log"
	"time"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	//emailBody := fmt.Sprintf("异步任务邮件已发送，发送时间：%s", time.Now().Format(utils.DateTimeFormat))
	//asyncEmail := async_task_task.AsyncEmailPayload{
	//	To:      "user@example.com",
	//	Subject: "异步任务邮件",
	//	Body:    emailBody,
	//}
	//task, err := async_task_task.NewAsyncEmailTask(asyncEmail)
	//if err != nil {
	//	log.Fatalf("could not create task: %v", err)
	//}
	//info, err := client.Enqueue(task)
	//if err != nil {
	//	log.Fatalf("could not enqueue task: %v", err)
	//}
	//log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	for i := 0; i < 5; i++ {
		emailBody := fmt.Sprintf("异步任务邮件已发送，发送时间：%s", time.Now().Format(utils.DateTimeFormat))
		asyncEmail := async_task_task.AsyncEmailPayload{
			To:      "user@example.com",
			Subject: "异步任务邮件",
			Body:    emailBody,
		}
		task, err := async_task_task.NewAsyncEmailTask(asyncEmail)
		time.Sleep(time.Second)
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}
		info, err := client.Enqueue(task)
		if err != nil {
			log.Fatalf("could not enqueue task: %v", err)
		}
		log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	}
}
