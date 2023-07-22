package async_task_task

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"golang.org/x/net/context"
)

// 创建一个新任务类型
const AsyncEmailTask = "async_email_task"

// AsyncEmailPayload 定义发送邮件的负载数据结构
type AsyncEmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// NewAsyncEmailTask 创建异步电子邮件任务的函数
func NewAsyncEmailTask(asyncEmail AsyncEmailPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(asyncEmail)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(AsyncEmailTask, payload)
	return task, err
}

// HandleAsyncEmailTask 处理异步电子邮件任务的函数
func HandleAsyncEmailTask(ctx context.Context, task *asynq.Task) error {
	payload := AsyncEmailPayload{}
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}
	// TODO: 模拟发送邮件
	fmt.Printf("\nAsync Server：Start handle AsyncTask!\n")
	fmt.Printf("Sending email to %s, subject: %s, body: %s\n", payload.To, payload.Subject, payload.Body)
	fmt.Println("Async Server：End handle AsyncTask!")
	return nil
}

// AsyncEmailProcessor implements asynq.Handler interface.
type AsyncEmailProcessor struct {
	// ... fields for struct
}

func NewAsyncEmailProcessor() *AsyncEmailProcessor {
	return &AsyncEmailProcessor{}
}

func (processor *AsyncEmailProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload AsyncEmailPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		//return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
		return err
	}
	// TODO: 模拟发送邮件
	fmt.Printf("\nAsync Server：ProcessTask：Start handle AsyncTask!\n")
	fmt.Printf("Sending email to %s, subject: %s, body: %s\n", payload.To, payload.Subject, payload.Body)
	fmt.Println("Async Server：ProcessTask：End handle AsyncTask!")
	return nil
}
