package setting

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
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

	err := godotenv.Load("./setting/process.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	switch st.ServerType {
	case 0:
		st.Port = ":3000"
		st.LogPath = "../Server/ServerLog/WebServer/"
		st.DB = os.Getenv("CONTENT_DB0")
	case 1:
		st.Port = ":3000"
		st.LogPath = "/data/DIPServerLog/WebServer1/"
		st.DB = os.Getenv("CONTENT_DB01")
	case 2:
		st.Port = ":3000"
		st.LogPath = "/data/DIPServerLog/WebServer2/"
		st.DB = os.Getenv("CONTENT_DB1")
	}
}
