package consts

// ZeroAsynqDemo 示例类型
const ZeroAsynqDemo = "zeroasynq:demo"

// ZeroAsynqPayload 定义示例数据的负载数据结构
type ZeroAsynqPayload struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
