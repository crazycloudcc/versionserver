package main

import (
	"cherry/base"
	"cherry/dbproxy"
	"cherry/nethttp"
	"cherry/nettcp"
	"cherry/netwebsocket"
	"data"
	"handlers"
	"os"
	"os/signal"
	"time"
)

func main() {
	base.SetBigEndian()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	base.SetLogLevel(base.LOG_LEVEL_DEBUG)

	conf := data.LoadConfigFile()
	data.LoadVersionFile() // check file.

	base.LogInfo("Start Service Info: ", conf.AppConf.ID, conf.AppConf.Group)

	httpConf := conf.HTTPConf
	httpsConf := conf.HTTPSConf
	wsConf := conf.WSConf
	tcpConf := conf.TCPConf
	redisConfRemote := conf.RedisConfRemote
	redisConfLocal := conf.RedisConfLocal

	/************************************************************************/
	// database & dataproxy service module.
	/************************************************************************/
	if redisConfRemote.Flag == 1 {
		if !dbproxy.RedisConnectRemote(redisConfRemote) {
			os.Exit(1)
			return
		}
	}

	if redisConfLocal.Flag == 1 {
		if !dbproxy.RedisConnectLocal(redisConfLocal) {
			os.Exit(1)
			return
		}
	}

	/************************************************************************/
	// network service module.
	/************************************************************************/

	if wsConf.Flag == 1 {
		wsServ := netwebsocket.NewService(wsConf)
		wsServ.Start()
	}

	if tcpConf.Flag == 1 {
		tcpServ := nettcp.NewService(tcpConf)

		// TODO RegHandler.

		tcpServ.Start()
	}

	if httpConf.Flag == 1 {
		httpServ := nethttp.NewHTTP(httpConf)

		// httpServ.RegHandler("/", ATestsHandle)
		httpServ.RegHandler("/test", handlers.ATestsHandle)
		httpServ.RegHandler("/version", handlers.VersionFileHandle)
		httpServ.RegHandler("/notice", handlers.NoticeFileHandle)
		httpServ.RegHandler("/serverlist", handlers.ServerListHandle)

		httpServ.Start()
	}

	if httpsConf.Flag == 1 {
		httpsServ := nethttp.NewHTTPS(httpsConf)

		// httpsServ.RegHandler("/", ATestsHandle)
		// httpsServ.RegHandler("/test", handlers.ATestsHandle)
		// httpsServ.RegHandler("/version", handlers.VersionFileHandle)
		// httpsServ.RegHandler("/notice", handlers.NoticeFileHandle)
		// httpsServ.RegHandler("/serverlist", handlers.ServerListHandle)

		httpsServ.Start()
	}

	base.LogInfo("server start time:", time.Now(), time.Now().Unix())
	base.LogInfo("server start done.")

	/************************************************************************/
	// service stop.
	/************************************************************************/
	s := <-interrupt

	base.LogInfo("server stop with signal: ", s)
	base.LogInfo("server stop done.")
}
