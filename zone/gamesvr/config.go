package main

import (
    "github.com/philipyao/project/common/configcli"
    log "github.com/philipyao/kit/logging"
)

var (
    configer     *ConfDefGame = new(ConfDefGame)
)

// gamesvr config define
type ConfDefGame struct {
    loglv         string  `config:"loglv"`
    openTime      string  `config:"open_time"`
}
func (cdf *ConfDefGame) Loglv() string {
    return cdf.loglv
}
func (cdf *ConfDefGame) SetLoglv(v string) error {
    log.Debug("SetLoglv: %v", v)
    cdf.loglv = v
    return nil
}
func (cdf *ConfDefGame) OnUpdateLoglv(val, oldVal string) {
    log.Debug("OnUpdateLoglv: %v %v", val, oldVal)
    err := log.SetLevel(cdf.loglv)
    if err != nil {
        log.Error("log.SetLevel(%v): %v", cdf.loglv, err)
    }
    return
}
func (cdf *ConfDefGame) OpenTime() string {
    return cdf.openTime
}
func (cdf *ConfDefGame) SetOpenTime(v string) error {
    log.Debug("SetOpenTime: %v", v)
    cdf.openTime = v
    return nil
}
func (cdf *ConfDefGame) OnUpdateOpenTime(val, oldVal string) {
    log.Debug("OnUpdateOpenTime: %v %v", val, oldVal)
    return
}

func initConfig() error {
    err := configcli.RegisterConfig("game", configer, rpcClient)
    if err != nil {
        return err
    }
    done := make(chan struct{})
    err = configcli.Load(done, argRegistry)
    if err != nil {
        return err
    }
    return nil
}
