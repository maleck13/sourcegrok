package web

import (
	"net/http"

	"encoding/json"

	"github.com/maleck13/sourcegrok/pkg/sourcegrok"
)

type SysHandler struct {
	Logger sourcegrok.Logger
}

func (s *SysHandler) Ping(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("ok"))
}

func (s *SysHandler) Health(rw http.ResponseWriter, req *http.Request) {
	data := map[string]string{"http": "ok"}
	encoder := json.NewEncoder(rw)
	if err := encoder.Encode(data); err != nil {
		s.Logger.Log(sourcegrok.LogLevelError, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
