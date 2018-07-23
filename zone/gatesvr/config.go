package main

import (
    "github.com/philipyao/project/common/configcli"
    log "github.com/philipyao/kit/logging"
    "strconv"
    "fmt"
)

var (
    configer     *ConfDefGate = new(ConfDefGate)
)

// gamesvr config define
type ConfDefGate struct {
    loglv         string  `config:"loglv"`
    maxOnline     int     `config:"max_online"`
}
func (cdg *ConfDefGate) Loglv() string {
    return cdg.loglv
}
func (cdg *ConfDefGate) SetLoglv(v string) error {
    log.Info("SetLoglv: %v", v)
    cdg.loglv = v
    return nil
}
func (cdg *ConfDefGate) OnUpdateLoglv(val, oldVal string) {
    log.Debug("OnUpdateLoglv: %v %v", val, oldVal)
    err := log.SetLevel(cdg.loglv)
    if err != nil {
        log.Error("log.SetLevel(%v): %v", cdg.loglv, err)
    }
    return
}
func (cdg *ConfDefGate) MaxOnline() int {
    return cdg.maxOnline
}
func (cdg *ConfDefGate) SetMaxOnline(v string) error {
    log.Info("SetMaxOnline: %v", v)
    iVal, err := strconv.Atoi(v)
    if err != nil {
        log.Error(err.Error())
        return err
    }
    if iVal <= 0 {
        err = fmt.Errorf("invalid maxOnline setting: %v", iVal)
        log.Error(err.Error())
        return err
    }
    cdg.maxOnline = iVal
    return nil
}
func (cdg *ConfDefGate) OnUpdateMaxOnline(val, oldVal string) {
    log.Debug("OnUpdateMaxOnline: %v %v", val, oldVal)
    return
}

func initConfig() error {
    err := configcli.RegisterConfig("gate", configer, rpcClient)
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
