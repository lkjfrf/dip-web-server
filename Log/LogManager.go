package setting

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type LogHandler struct {
}

var Log_Ins *LogHandler
var Log_once sync.Once

func GetLogManager() *LogHandler {
	Log_once.Do(func() {
		Log_Ins = &LogHandler{}
	})
	return Log_Ins
}

func (l *LogHandler) SetLogFile() {
	// 현재시간
	startDate := time.Now().Format("2006-01-02")
	// log 폴더 위치
	logFolderPath := "/dipnas/DIPServer/ServerLog/DIPWebServer_log"
	// log 파일 경로
	logFilePath := fmt.Sprintf("%s/logFile-%s.log", logFolderPath, startDate)
	// log 폴더가 없을 경우 log 폴더 생성
	if _, err := os.Stat(logFolderPath); os.IsNotExist(err) {
		os.MkdirAll(logFolderPath, 0777)
	}

	// log 파일이 없을 경우 log 파일 생성
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		os.Create(logFilePath)
	}
	// log 파일 열기
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// log2 파일 열기
	logFile2, err := os.OpenFile("DIP_Log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// log 패키지를 활요하여 작성할 경우 log 파일에 작성되도록 설정
	mw := io.MultiWriter(os.Stdout, logFile, logFile2)
	log.SetOutput(mw)
}
