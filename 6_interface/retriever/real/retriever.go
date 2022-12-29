package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(interface{}(err))
	}

	result, err := httputil.DumpResponse(
		resp, true)

	err = resp.Body.Close()
	if err != nil {
		return ""
	}

	if err != nil {
		panic(interface{}(err))
	}

	return string(result)
}
