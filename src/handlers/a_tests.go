package handlers

import (
	"cherry/base"
	"net/http"
)

// ATestsHandle TODO
func ATestsHandle(w http.ResponseWriter, r *http.Request) {
	base.LogDebug("method:", r.Method)
	r.ParseForm() // 解析参数.
	base.LogDebug("ATestsHandle get request info:", r.Form)
	base.LogDebug("path:", r.URL.Path)
	base.LogDebug("scheme:", r.URL.Scheme)
	for k, v := range r.Form {
		base.LogDebug("info:", k, v)
	}

	b := []byte(r.Form.Encode())
	_, err := w.Write(b)
	if err != nil {
		base.LogError("ATestsHandle write error: ", err)
	}
}
