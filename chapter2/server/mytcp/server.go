/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/6/21 14:21
* @Description: The file is for
***********************************************************************/

package mytcp

import (
	"golang-cache/chapter2/server/cache"
	"net"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	listener, err := net.Listen("tcp", ":12346")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// 开辟协程，处理连接事务, 有多少个连接就开多少个协程
		go s.process(conn)
	}
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

/**
TCP的ABNF表达
SP表示空格  DIGIT取值范围0~9 1*表示1个或更多个
OCTET取值范围0x00~0xFF, *表示0个或更多个
command = op key | key-value
op = 'S' | 'G' | 'D'
key = bytes-array
bytes-array = length SP content
length = 1*DIGIT
content = *OCTET
key-value = length SP length SP content content
response = error | bytes-array
error = '-' bytes-array
*/