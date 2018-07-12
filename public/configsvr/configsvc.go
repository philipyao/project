package main

import (
    "github.com/philipyao/project/public/configsvr/core"
)

//implement app.Service
type serviceConfig struct {}
func (sc *serviceConfig) OnInit() error {
    err := core.Init(argRegistry, argDBAddr, dbLogin)
    if err != nil {
        return err
    }
    return nil
}
func (sc *serviceConfig) Serve() error {
    return nil
}
func (sc *serviceConfig) Close() {}
func (sc *serviceConfig) OnFini() {
    core.Fini()
}
