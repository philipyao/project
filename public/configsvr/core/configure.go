package core

import (
    "github.com/philipyao/project/public/configsvr/def"
    "github.com/philipyao/project/public/configsvr/db"
    "fmt"
    //log "github.com/philipyao/kit/logging"
)

type FuncTraverse func(*def.Config) error
type Configure  map[uint]*def.Config

func (c Configure) Load(confs []*def.Config) error {
    for _, v := range confs {
        if ns.Exist(v.Namespace) == false {
            return fmt.Errorf("config %+v, namespace not found", v)
        }
        c[v.ID] = v
    }
    return nil
}

func (c Configure) Get(id uint) *def.Config {
    return c[id]
}

func (c Configure) GetBy(namespace, key string) *def.Config {
    for _, v := range c {
        if v.Namespace == namespace && v.Key == key {
            return v
        }
    }
    return nil
}

func (c Configure) Set(namespace, key, value string) (*def.Config, error) {
    var err error
    var addConf def.Config
    addConf.Namespace = namespace
    addConf.Key = key
    addConf.Value = value
    err = db.InsertConfig(&addConf)
    if err != nil {
        return nil, err
    }
    //ID由DB生成
    c[addConf.ID] = &addConf
    return &addConf, nil
}

func (c Configure) Foreach(f FuncTraverse) error {
    var e error
    for _, v := range c {
        e = f(v)
        if e != nil {
            return e
        }
    }
    return nil
}
