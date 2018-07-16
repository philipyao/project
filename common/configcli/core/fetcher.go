package core

import (
	"fmt"
	"errors"
    log "github.com/philipyao/kit/logging"
    "github.com/philipyao/prpc/client"
    "github.com/philipyao/project/common/proto"
)

var (
	rpcSvc   *client.SvcClient
)

func InitFetcher(rpcClient *client.Client) error {
    rpcSvc = rpcClient.Service("Config", "public1")
	return nil
}

func FetchConfFromServer(namespace string, keys []string) ([]*proto.ConfigEntry, error){
	if rpcSvc == nil {
		panic("fetcher not init.")
	}
	if namespace == "" || len(keys) == 0 {
		panic("inv input for fetch.")
	}
	args := &proto.FetchConfigArg{
		Namespace: namespace,
		Keys: keys,
	}
	var response proto.FetchConfigRes
	err := rpcSvc.Call("FetchConfig", args, &response)
	if err != nil {
		return nil, fmt.Errorf("FetchConfig call error %v\n", err)
	}
    log.Debug("fetch response: %+v", response)
	if response.Errmsg != "" {
		return nil, errors.New(response.Errmsg)
	}
	return response.Confs, nil
}

