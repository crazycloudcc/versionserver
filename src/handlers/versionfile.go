package handlers

import (
	"cherry/base"
	"data"
	"fmt"
	"net/http"
)

// VersionFileHandle TODO
func VersionFileHandle(w http.ResponseWriter, r *http.Request) {
	base.LogDebug("VersionFileHandle method:", r.Method)
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
	// base.LogDebug("VersionFileHandle get request info:", r.Form)
	// for k, v := range r.Form {
	// 	base.LogDebug("info:", k, v)
	// }

	// clientVersion := r.Form.Get("AppVersion")
	// if cache.AppVersion != clientVersion {
	// 	base.LogDebug("client version error: ", clientVersion, cache.AppVersion)
	// }

	cache := data.LoadVersionFile()
	resData := fmt.Sprintf("0|%v|%v|%v|%v", cache.AppVersion, cache.NoticeFileURL, cache.ServerListFileURL, cache.DownloadAppURL)
	_, err := w.Write([]byte(resData))
	if err != nil {
		base.LogError("VersionFileHandle write error: ", err)
	}
}
