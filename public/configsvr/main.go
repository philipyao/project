package main

import (
    "github.com/philipyao/app"
    syslog "log"
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
    err = app.Init(
        app.WithArgInt(&argPortBase, "portbase", 0, "listening base port"),
        app.WithArgString(&argRegistry, "registry", "", "rpc registry"),
        //todo
        //app.WithArgString(&argDBAddr, "db", "", "db addr"),
    )
    if err != nil {
        syslog.Fatalf("init err %v, exit", err)
    }

    common.InitBizlog()

    err = app.Run(
        //rpc
        app.ServeRpc(fmt.Sprintf(":%v", argPortBase), argRegistry, new(configerRpc), "Config"),
        //http
        app.ServeHttp(fmt.Sprintf(":%v", argPortBase+1), serveHttp),
        //business
        new(serviceConfiger),
    )
    if err != nil {
        log.Fatal("run err %v, exit", err)
    }
    log.Flush()
}

