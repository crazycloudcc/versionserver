package data

import (
	"cherry/base"
	"cherry/base/config"
	"cherry/dbproxy"
	"cherry/nethttp"
	"cherry/nettcp"
	"cherry/netwebsocket"
)

// ConfigFile TODO
type ConfigFile struct {
	AppConf struct {
		ID    int32
		Group string
	}
	HTTPConf        nethttp.Config
	HTTPSConf       nethttp.Config
	WSConf          netwebsocket.Config
	TCPConf         nettcp.Config // tcp
	RedisConfRemote dbproxy.RedisConfig
	RedisConfLocal  dbproxy.RedisConfig
}

// VersionFile TODO
type VersionFile struct {
	AppVersion        string
	NoticeFileURL     string
	ServerListFileURL string
	DownloadAppURL    string
}

// NoticeFile TODO
type NoticeFile struct {
	Title   string
	Context string
}

// ServerListFile TODO
type ServerListFile struct {
	ArrID    []string
	ArrGroup []string
	ArrHost  []string
	ArrPort  []string
}

var configData *ConfigFile
var versionData *VersionFile
var noticeData *NoticeFile
var serverlistData *ServerListFile

// LoadConfigFile TODO
func LoadConfigFile() *ConfigFile {
	if configData == nil {
		configData = new(ConfigFile)
		err := config.Read("./config/dev.json", "json", configData)
		if err != nil {
			base.LogFatal("LoadFile error [conf.json]: ", err)
		}
	}
	return configData
}

// LoadVersionFile TODO
func LoadVersionFile() *VersionFile {
	if versionData == nil {
		versionData = new(VersionFile)
		err := config.Read("./config/version.json", "json", versionData)
		if err != nil {
			base.LogFatal("LoadFile error [version.json]: ", err)
		}
	}
	return versionData
}

// LoadNoticeFile TODO
func LoadNoticeFile() *NoticeFile {
	if noticeData == nil {
		noticeData = new(NoticeFile)
		err := config.Read("./config/notice.json", "json", noticeData)
		if err != nil {
			base.LogFatal("LoadFile error [notice.json]: ", err)
		}
	}
	return noticeData
}

// LoadServerListFile TODO
func LoadServerListFile() *ServerListFile {
	if serverlistData == nil {
		serverlistData = new(ServerListFile)
		err := config.Read("./config/serverlist.json", "json", serverlistData)
		if err != nil {
			base.LogFatal("LoadFile error [serverlist.json]: ", err)
		}
	}
	return serverlistData
}
