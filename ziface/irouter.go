package ziface

// 路由抽象接口
// 路由里的数据都是IRequest
type IRouter interface {
	// 在处理conn业务之前的hook
	PreHandle(req IRequest)
	// 处理conn业务的主方法
	Handle(req IRequest)
	// 在处理conn业务之后的hook
	PostHandle(req IRequest)
}