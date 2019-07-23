package myhttp

import (
	"encoding/json"
	"log"
	"net/http"
)

type statusHandler struct {
	*Server
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 状态只能查询
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	b, err := json.Marshal(h.GetStat())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}