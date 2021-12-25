package znet

import "github.com/zeyuanDADA/zinx/ziface"

// 实现router时，先嵌入这个BaseRouter基类，然后根据需要对这个基类的方法进行重写就好了
type BaseRouter struct {}

// 在处理conn业务之前的hook
func (b *BaseRouter) PreHandle(req ziface.IRequest) {}
// 处理conn业务的主方法
func (b *BaseRouter) Handle(req ziface.IRequest) {}
// 在处理conn业务之后的hook
func (b *BaseRouter) PostHandle(req ziface.IRequest) {}
