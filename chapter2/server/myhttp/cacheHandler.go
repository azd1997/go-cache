package myhttp

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//多重内嵌。 cacheHandler内嵌*Server， Server内嵌cache.Cache接口。 所以cacheHandler可以直接访问Cache的方法
type cacheHandler struct {
	*Server
}

// 实现net/http中的Handler接口的ServeHTTP(ResponseWriter, *Request)方法
func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.EscapedPath(), "/")[2]	//从请求URL中截取出查询的Key
	if len(key) == 0 {		//没有key则表示请求有误
		w.WriteHeader(http.StatusBadRequest)
	}

	m := r.Method		//检查请求是哪一种并做对应处理

	if m == http.MethodPut {
		b, _ := ioutil.ReadAll(r.Body)
		if len(b) != 0 {
			err := h.Set(key, b)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write([]byte("PUT成功！"))
		}
		return
	}

	if m == http.MethodGet {
		b, err := h.Get(key)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(b) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write([]byte("GET成功！"))
		w.Write(b)
		return
	}

	if m == http.MethodDelete {
		err := h.Del(key)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
		w.Write([]byte("DEL成功！"))
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}
