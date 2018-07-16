package main

import (
    log "github.com/philipyao/kit/logging"

    "github.com/philipyao/project/common/proto"
    "github.com/philipyao/project/public/configsvr/core"
)

type configerRpc struct {}

//根据特定namespace获取配置键值对
func (configerRpc) FetchConfig(args *proto.FetchConfigArg, response *proto.FetchConfigRes) error {
    log.Debug("[rpc]FetchConfig: args %+v", args)
    confMap, err := core.ConfigWithNamespaceKey(args.Namespace, args.Keys)
    if err != nil {
        response.Errmsg = err.Error()
        return nil
    }
    for k, v := range confMap {
        response.Confs = append(response.Confs, &proto.ConfigEntry{
            Namespace: v[0],
            Key:    k,
            Value:  v[1],
        })
    }
    log.Debug("rsp confs: %+v", response)
    return nil
}