package setting

import (
	"sync"
)

type SettingHandler struct {
	ServerType int // 0: 나, 1: 원효로1번서버, 2: 원효로2번서버
	Port       string
	LogPath    string
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
	st.ServerType = 0 // 0: 나, 1: 원효로1번서버, 2: 원효로2번서버, 3: 민석님

	switch st.ServerType {
	case 0:
		st.Port = ":3000"
		st.LogPath = "/data/DIPServerLog/WebServer/"
	case 1:
		st.Port = ":3000"
		st.LogPath = "/data/DIPServerLog/WebServer1/"
	case 2:
		st.Port = ":3000"
		st.LogPath = "/data/DIPServerLog/WebServer2/"
	case 3:
		st.Port = ":3000"
		st.LogPath = "/data/DIPServerLog/WebServer/"
	}
}
