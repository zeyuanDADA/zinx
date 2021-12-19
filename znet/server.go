package znet

import (
	"github.com/zeyuanDADA/zinx/ziface"
)

// Server IServer的接口实现，定义一个Server的服务器模块
type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的ip
	IP string
	// 服务器监听的端口
	Port int
}

// Start 启动服务器
func (s *Server) Start() {

}

// Stop 停止服务器
func (s *Server) Stop() {

}

// Serve 运行服务器
func (s *Server) Serve() {

}

// 初始化Server模块
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
	}

	return s
}