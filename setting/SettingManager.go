package setting

import (
	"sync"
)

type SettingHandler struct {
	ServerType int // 0: 나, 1: 원효로1번서버, 2: 원효로2번서버
	Port       string
	LogPath    string
	DB         string
}

var St_Ins *SettingHandler
var St_once sync.Once

func GetStManager() *SettingHandler {
	St_once.Do(func() {
		St_Ins = &SettingHandler{}
	})
	return St_Ins
}

func (st *SettingHandler) Init() {
	st.ServerType = 0 // 0: 나, 1: 원효로1번서버, 2: 원효로2번서버

	switch st.ServerType {
	case 0:
		st.Port = ":3000"
		st.LogPath = "../Server/ServerLog/WebServer/"
		st.DB = "root:Tomatosoup22!@tcp(127.0.0.1:3306)/dip?charset=utf8mb4&parseTime=True&loc=Local"
	case 1:
		st.Port = ":3000"
		st.LogPath = "/data/DIPServerLog/WebServer1/"
		st.DB = "DIPADM:P!ssw0rd@tcp(10.5.147.148:3306)/dipdb?charset=utf8mb4&parseTime=True&loc=Local"
	case 2:
		st.Port = ":3000"
		st.LogPath = "/data/DIPServerLog/WebServer2/"
		st.DB = "DIPADM:P!ssw0rd@tcp(10.5.147.148:3306)/dipdb?charset=utf8mb4&parseTime=True&loc=Local"
	}
}
