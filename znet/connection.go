package znet

import (
	"net"
	"fmt"

	"github.com/zeyuanDADA/zinx/ziface"
)

// 链接模块
type Connection struct {
	// 当前链接的socket TCP套接字
	Conn *net.TCPConn
	
	// 链接的ID
	ConnID uint32

	// 当前的链接状态
	IsClosed bool

	// 告知当前链接已经退出/停止 channel
	ExitChan chan bool

	// 该链接处理的方法Router
	Router ziface.IRouter
}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		IsClosed: false,
		ExitChan: make(chan bool, 1),
		Router: router,
	}

	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer func ()  {
		fmt.Printf("connID = %v, Reader is exit, remote addr is %v\n", c.ConnID, c.RemoteAddr().String())
		c.Stop()
	}()

	for {
		// 读取客户端的数据到buf中，最大512字节
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			continue
		}

		// 得到当前conn数据的Request请求数据
		req := &Request{
			conn: c,
			data: buf,
		}
		// 执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(req)
	}
}

// 启动链接，让当前的链接准备开始工作
func (c *Connection) Start() {
	fmt.Println("conn start()... connID = ", c.ConnID)

	// 启动从当前链接读数据的业务
	go c.StartReader()
	// TODO 启动从当前链接写数据的业务

}

// 停止链接，结束当前链接的工作
func (c *Connection) Stop() {
	fmt.Println("conn stop()... connID = ", c.ConnID)

	if c.IsClosed {
		// 如果当前链接已经关闭直接返回
		return
	}
	c.IsClosed = true

	// 关闭socket链接
	c.Conn.Close()

	// 回收管道资源
	close(c.ExitChan)
}

// 获取当前链接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的TCP状态IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据，将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}