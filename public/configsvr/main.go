package main

import (
    "github.com/philipyao/app"
    log "github.com/philipyao/kit/logging"
    "github.com/philipyao/project/common"
    "fmt"
)

const (
    dbLogin     string  = "philip:tmpP001"
)

var (
    //http/rpc listening
    argPortBase int
    //rpc registry
    argRegistry string
    //configsvr data storage
    argDBAddr   string = "localhost:3306"
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
        app.WithCPUNum(1),
    )
    if err != nil {
        log.Fatal("init err %v, exit", err)
    }
    err = app.Run(
        //service: rpc
        app.ServeRpc(fmt.Sprintf(":%v", argPortBase), argRegistry, new(configerRpc), "Config"),
        //service: http
        app.ServeHttp(fmt.Sprintf(":%v", argPortBase + 1), serveHttp),
        //service: business
        new(serviceConfig),
    )
    if err != nil {
        log.Fatal("run err %v, exit", err)
    }

    common.FiniBizLog()
}

