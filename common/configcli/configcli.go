package configcli

import (
    "fmt"
    "reflect"
    "errors"

    "github.com/philipyao/prpc/client"
    "github.com/philipyao/project/common/commdef"
    "github.com/philipyao/project/common/configcli/core"
)

var (
    currNamespace   string
)

// 配置客户端向配置中心注册用到的配置项
func RegisterConfig(namespace string, confDef interface{}, rpcClient *client.Client) error {
    var err error

    err = core.InitFetcher(rpcClient)
    if err != nil {
        return err
    }

    if namespace == "" {
        return errors.New("empty conf namespace.")
    }
    currNamespace = namespace

    t := reflect.TypeOf(confDef)
    v := reflect.ValueOf(confDef)
    if t.Kind() != reflect.Ptr {
        return errors.New("confdef should be pointer.")
    }
    t = t.Elem()
    if t.Kind() != reflect.Struct {
        return fmt.Errorf("confdef should be struct pointer. %v", reflect.TypeOf(t.Elem()).Kind())
    }
    if t.NumField() == 0 {
        return errors.New("confdef with no fields.")
    }
    //找出'config' tag的字段
    tagFound := false
    for i := 0; i < t.NumField(); i++ {
        sf := t.Field(i)
        tag, ok := sf.Tag.Lookup(commdef.CongigTagName)
        if !ok { continue }
        if tag == "" {
            return fmt.Errorf("empty value of '%v' tag for field <%v> is not allowed.",
                commdef.CongigTagName, sf.Name)
        }
        tagFound = true
        goName := tag2GoName(tag)
        err = core.RegisterEntry(tag, goName, v)
        if err != nil {
            return err
        }
    }
    if tagFound == false {
        return fmt.Errorf("no '%v' tag found in provided confDef", commdef.CongigTagName)
    }
    return nil
}

func SetLogger(l func(int, string, ...interface{})) {
    core.SetLogger(l)
}

func Load(done chan struct{}, zkAddr string) error {
    //开始从远程服务器加载需要的配置
    keys := core.EntryKeys()
    core.Log("start loading confs: count %v", len(keys))
    confs, err := core.FetchConfFromServer(currNamespace, keys)
    if err != nil {
        return err
    }
    core.Log("fetch confs from confsvr ok. namespace %v", currNamespace)

    err = core.InitWatcher(zkAddr)
    if err != nil {
        return err
    }

    notify := make(chan string)
    for _, c := range confs {
        err = core.InitEntry(c.Key, c.Value)
        if err != nil {
            return err
        }
        err = core.WatchEntryUpdate(c.Namespace, c.Key, notify, done)
        if err != nil {
            return err
        }
        core.Log("watch entry <%v %v>.", c.Namespace, c.Key)
    }

    // listen updates
    go handleWatch(notify, done)

    core.Log("finished loading confs.")
    return nil
}

func handleWatch(notify chan string, done chan struct{}) {
    for {
        select {
        case <- done:
            return
        case updateKey := <- notify:
            handleUpdate(updateKey)
        }
    }
}

func handleUpdate(key string) {
    confs, err := core.FetchConfFromServer(currNamespace, []string{key})
    if err != nil {
        core.Log("FetchConfFromServer: %v", err)
        return
    }
    if len(confs) != 1 {
        core.Log("inv conf counts")
        return
    }
    if confs[0].Key != key {
        core.Log("mismatch key %v %+v", key, confs[0])
        return
    }
    core.Log("handleUpdate: key %v", key)
    //TODO
    core.UpdateEntry(key, "", confs[0].Value)
}