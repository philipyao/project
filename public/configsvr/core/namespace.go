package core

import (
    "errors"
    log "github.com/philipyao/kit/logging"

    "github.com/philipyao/project/public/configsvr/def"
    "github.com/philipyao/project/public/configsvr/db"
)

var (
    ns *Namespace = new(Namespace)
)

type Namespace []string

func (n *Namespace) Load(ns []string) {
    *n = make([]string, len(ns))
    copy(*n, ns)
    log.Debug("ns load: %+v", ns)
}

func (n *Namespace) Exist(val string) bool {
    for _, entry := range *n {
        if entry == val {
            return true
        }
    }
    return false
}

//预先生成公共空间
func (n *Namespace) CreateCommon() (err error) {
    name := def.ConfNamespaceCommon
    if n.Exist(name) {
        return nil
    }
    defer func() {
        if err == nil {
            *n = append(*n, name)
        }
    }()
    log.Info("ceate namespace %v", name)
    return n.doCreate(def.AdminUsername, name, "公共配置区间，配置项可以被私有同名配置项覆盖")
}

//创建普通私有空间
func (n *Namespace) Create(creator, name, desc string) error {
    err := n.doCreate(creator, name, desc)
    if err != nil {
        return err
    }
    *n = append(*n, name)
    return nil
}

func (n *Namespace) doCreate(creator, name, desc string) error {
    if name == "" {
        return errors.New("error create namespace: empty name")
    }
    exist, err := db.ExistNamespace(name)
    if err != nil {
        return err
    }
    if exist {
        return errors.New("error create namespace: already exist")
    }
    namespace := &def.Namespace {
        Name: name,
        Desc: desc,
        Creator: creator,
    }
    err = db.InsertNamespace(namespace)
    if err != nil {
        return err
    }
    log.Info("create namespace [%v] ok", name)
    return nil
}
