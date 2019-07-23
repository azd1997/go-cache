/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/6/21 14:50
* @Description: Server.process实现
***********************************************************************/

package mytcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

//
func readLen(r *bufio.Reader) (int, error) {
	// 空格作为停止读入的分隔。
	tmp, err := r.ReadString(' ')
	if err != nil {
		return 0, err
	}
	// tmp字符串， 去除掉前导和尾随的空格，再解析为十进制整数
	lenInt, err := strconv.Atoi(strings.TrimSpace(tmp))
	if err != nil {
		return 0, err
	}
	return lenInt, err
}

func (s *Server) readKey(r *bufio.Reader) (string, error) {
	keyLen, err := readLen(r)
	if err != nil {
		return "", err
	}

	// 将bufio.Reader读到内容传给key
	key := make([]byte, keyLen)
	_, err = io.ReadFull(r, key)
	if err != nil {
		return "", err
	}

	return string(key), nil
}

func (s *Server) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	// 获取请求的键值长度
	keyLen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}

	valueLen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}

	key := make([]byte, keyLen)
	_, err = io.ReadFull(r, key)
	if err != nil {
		return "", nil, err
	}

	value := make([]byte, valueLen)
	_, err = io.ReadFull(r, value)
	if err != nil {
		return "", nil, err
	}

	return string(key), value, nil
}

func sendResponse(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()	//Error()方法返回字符串
		tmp := fmt.Sprintf("-%d ", len(errString)) + errString
		_, e := conn.Write([]byte(tmp))
		return e
	}

	valueLen := fmt.Sprintf("%d", len(value))
	log.Println(valueLen)
	valueMsg := append([]byte(valueLen), value...)
	log.Println(string(valueMsg))
	_, e := conn.Write(valueMsg)
	log.Println("发送响应完毕")
	return e
}

func (s *Server) get(conn net.Conn, r *bufio.Reader) error {
	key, err := s.readKey(r)
	if err != nil {
		return err
	}
	value, err := s.Get(key)
	return sendResponse(value, err, conn)
}

func (s *Server) set(conn net.Conn, r *bufio.Reader) error {
	key, value, err := s.readKeyAndValue(r)
	if err != nil {
		return err
	}

	return sendResponse(nil, s.Set(key, value), conn)
}

func (s *Server) del(conn net.Conn, r *bufio.Reader) error {
	key, err := s.readKey(r)
	if err != nil {
		return err
	}

	return sendResponse(nil, s.Del(key), conn)
}

func (s *Server) process(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		log.Println("处理开始")
		op, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("close connection due to error: ", err)
			}
			return
		}

		//switch op {
		//case 'S':
		//	err = s.set(conn, r)
		//case 'G':
		//	err = s.get(conn, r)
		//case 'D':
		//	err = s.del(conn, r)
		//default:
		//	log.Println("close connection due to invalid operation: ", op)
		//	return
		//}

		if op == 'S' {
			log.Println("开始set")
			err = s.set(conn, r)
		} else if op == 'G' {
			log.Println("开始get")
			err = s.get(conn, r)
		} else if op == 'D' {
			log.Println("开始del")
			err = s.get(conn, r)
		} else {
			log.Println("close connection due to invalid operation: ", op)
			return
		}

		if err != nil {
			log.Println("close connection due to error: ", err)
			return
		}
	}
}

