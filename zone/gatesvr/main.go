package main

import (
    "github.com/philipyao/app"
    log "github.com/philipyao/kit/logging"
    "github.com/philipyao/project/common"
    "runtime"
)

var (
    argPortBase int
    //rpc registry
    argRegistry string
)

func main() {
    var err error

    app.ReadArgs(
        app.WithArgInt(&argPortBase, "portbase", 0, "listening base port"),
        app.WithArgString(&argRegistry, "registry", "", "rpc registry"),
    )
    common.InitBizlog()

    err = app.Init(
        app.WithLogger(log.Info),
        app.WithCPUNum(runtime.NumCPU()),
    )
    if err != nil {
        log.Fatal("init err %v, exit", err)
    }
    err = app.Run(
        //service: business
        new(serviceGate),
    )
    if err != nil {
        log.Fatal("run err %v, exit", err)
    }

    common.FiniBizLog()
}

