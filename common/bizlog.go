package common

import (
    "os"
    "path/filepath"
    "github.com/philipyao/app"
    "fmt"
    log "github.com/philipyao/kit/logging"
)

const (
    defaultLogSize  = 102400    //100M
)

//@1 业务log初始化
func InitBizlog() {
    config := `{"filename": "%v", "maxsize": %v, "maxbackup": 10}`
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    logPath := filepath.Join(wd, "log")
    if _, err := os.Stat(logPath); os.IsNotExist(err) {
        os.Mkdir(logPath, 0775)
    }
    logName := filepath.Join(logPath, app.ProcessName())
    config = fmt.Sprintf(config, logName, defaultLogSize)
    fmt.Printf("log config: %+v\n", config)
    err = log.AddAdapter(log.AdapterFile, config)
    if err != nil {
        panic(err)
    }
    log.SetLevel(log.LevelStringDebug)
    log.SetFlags(log.LogDate | log.LogTime | log.LogMicroTime | log.LogLongFile)
}

//@2 业务log清理
func FiniBizLog() {
    log.Flush()
}
