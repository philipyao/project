package main

import (
    "github.com/philipyao/project/public/configsvr/core"
)


type serviceConfiger struct {}
func (sc *serviceConfiger) OnInit() error {
    err := core.Init(argRegistry, argDBAddr, dbLogin)
    if err != nil {
        return err
    }
    return nil
}
func (sc *serviceConfiger) Serve() error {
    return nil
}
func (sc *serviceConfiger) Close() {}
func (sc *serviceConfiger) OnFini() {
    core.Fini()
}
