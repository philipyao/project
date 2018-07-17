package main

import (
    "github.com/philipyao/prpc/client"
    "github.com/philipyao/prpc/registry"
    "errors"
)

var (
    rpcClient *client.Client
)

func initCall() error {
    config := &registry.RegConfigZooKeeper{ZKAddr: argRegistry}
    client := client.New(config)
    if client == nil {
        return errors.New("new rpc client")
    }
    rpcClient = client

    return nil
}
