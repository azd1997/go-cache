package main

import (
	"go-learning/lec1-cache/cache"
	"go-learning/lec1-cache/myhttp"
)

func main() {
	c := cache.New("inmemory")
	myhttp.New(c).Listen()
}