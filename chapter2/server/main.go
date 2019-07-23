/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/6/21 14:01
* @Description: main.go file
***********************************************************************/

package main

import (
	"golang-cache/chapter2/server/cache"
	"golang-cache/chapter2/server/myhttp"
	"golang-cache/chapter2/server/mytcp"
)

func main() {

	// Start your code here
	ca := cache.New("inmemory")

	// 在创建ca和调用myhttp.Server.Listen之间开辟协程，调用mytcp.Server.Listen
	go mytcp.New(ca).Listen()

	myhttp.New(ca).Listen()

}

// Set流程： client -> mytcp.Server.set() -> cache.Cache$inMemoryCache$.Set(key, value) -> memory
// Get流程： client -> mytcp.Server.get() -> cache.Cache$inMemoryCache$.Get(key) -> memory
// Del流程： client -> mytcp.Server.del() -> cache.Cache$inMemoryCache$.Del(key) -> memory