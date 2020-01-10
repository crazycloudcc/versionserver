package handlers

import (
	"cherry/base"
	"data"
	"fmt"
	"net/http"
)

// ServerListHandle TODO
func ServerListHandle(w http.ResponseWriter, r *http.Request) {
	base.LogDebug("ServerListHandle method:", r.Method)
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

	// r.ParseForm()
	// base.LogDebug("ServerListHandle get request info:", r.Form)
	// for k, v := range r.Form {
	// 	base.LogDebug("info:", k, v)
	// }

	cache := data.LoadServerListFile()

	resData := "0|"
	_count := len(cache.ArrID)
	for i := 0; i < _count; i++ {
		resData += fmt.Sprintf("%v|%v|%v|%v", cache.ArrID[i], cache.ArrGroup[i], cache.ArrHost[i], cache.ArrPort[i])
		if i != _count-1 {
			resData += "|"
		}
	}

	base.LogDebug("resData: ", resData)
	_, err := w.Write([]byte(resData))
	if err != nil {
		base.LogError("ServerListHandle write error: ", err)
	}
}
