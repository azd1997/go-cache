package myhttp
// 尽管并不会和net/http冲突，还是改名了

import (
	"golang-cache/chapter2/server/cache"
	"net/http"
)

type Server struct {
	cache.Cache	//内嵌Cache接口即表示实现该接口
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status/", s.statusHandler())
	http.ListenAndServe(":12345", nil)
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

// 注册cacheHandler
func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}

// 注册statusHandler
func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}
