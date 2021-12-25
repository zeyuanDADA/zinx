package znet

import (
	"fmt"
	"net"

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
	// 当前的server添加一个router,server注册的链接对应的处理业务
	Router ziface.IRouter
}

// Start 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP: %s, Port: %d\n", s.IP, s.Port)
	go func ()  {
		// 1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Printf("resolve tcp addr error: %v\n", err)
			return
		}

		// 2 监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("listen %s err: %v\n", s.IPVersion, err)
			return
		}
		fmt.Printf("start Zinx server %s succ, listenning...\n", s.Name)

		var cid uint32 = 0

		// 3 阻塞的等待客户端连接，处理客户端连接业务（读写）
		for {
			// 如果有客户端连接过来，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error", err)
				continue
			}

			// 将处理新链接的业务方法和conn进行绑定 得到我们的链接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			// 启动当前的链接业务处理
			go dealConn.Start()
		}
	}()
}

// Stop 停止服务器
func (s *Server) Stop() {
	// TODO 将一些服务器的资源、状态或者一些已经开辟的链接信息进行停止或回收
}

// Serve 运行服务器
func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	//TODO 做一些启动服务器之后的额外任务

	// 阻塞状态
	select{}
}

// 注册路由
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succ!!")
}

// 初始化Server模块
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
		Router: nil,
	}

	return s
}