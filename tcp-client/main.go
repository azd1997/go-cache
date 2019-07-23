/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/6/21 14:48
* @Description: 实现简单的tcp客户端
***********************************************************************/

package main

import (
	"flag"
	"fmt"
	"golang-cache/cache-benchmark/cacheClient"
)

func main() {

	server := flag.String("h", "localhost", "cache server address")
	op := flag.String("c", "get", "command, could be get/set/del")
	key := flag.String("k", "", "key")
	value := flag.String("v", "", "value")
	flag.Parse()

	client := cacheClient.New("tcp", *server)
	cmd := &cacheClient.Cmd{*op, *key, *value, nil}
	client.Run(cmd)

	if cmd.Error != nil {
		fmt.Println("error: ", cmd.Error)
	} else {
		fmt.Println(cmd.Value)
	}
}