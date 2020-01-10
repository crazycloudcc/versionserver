package handlers

import (
	"cherry/base"
	"data"
	"fmt"
	"net/http"
)

// NoticeFileHandle TODO
func NoticeFileHandle(w http.ResponseWriter, r *http.Request) {
	base.LogDebug("NoticeFileHandle method:", r.Method)
	// if r.Method != "POST" {
	// 	return
	// }

	base.LogDebug("client info: ---------------------- S")
	base.LogDebug("Host: ", r.Host)
	base.LogDebug("RemoteAddr: ", r.RemoteAddr)
	base.LogDebug("Proto: ", r.Proto)
	base.LogDebug("RequestURI: ", r.RequestURI)
	base.LogDebug("ContentLength: ", r.ContentLength)
	base.LogDebug("client info: ---------------------- E")

	r.ParseForm()
	base.LogDebug("NoticeFileHandle get request info:", r.Form)
	// for k, v := range r.Form {
	// 	base.LogDebug("info:", k, v)
	// }
	cache := data.LoadNoticeFile()
	base.LogDebug(*cache)

	resData := fmt.Sprintf("0|%v|%v", cache.Title, cache.Context)

	base.LogDebug("resData: ", resData)
	_, err := w.Write([]byte(resData))
	if err != nil {
		base.LogError("NoticeFileHandle write error: ", err)
	}
}
