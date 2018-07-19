package main

import log "github.com/philipyao/kit/logging"

//implement app.Service
type serviceGame struct {}
func (sg *serviceGame) OnInit() error {
    var err error

    //初始化调用系统
    err = initCall()
    if err != nil {
        return err
    }
    //拉取本 server 配置
    err = initConfig()
    if err != nil {
        return err
    }
    //设置日志级别
    loglv := configer.Loglv()
    err = log.SetLevel(loglv)
    if err != nil {
        return err
    }
    log.Info("set loglv: %v", loglv)

    return nil
}
func (sg *serviceGame) Serve() error {
    return nil
}
func (sg *serviceGame) Close() {}
func (sg *serviceGame) OnFini() {
}
