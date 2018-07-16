package configcli

import (
	"testing"
	"strconv"
    "github.com/philipyao/prpc/registry"
    "github.com/philipyao/prpc/client"
)

type SampleConfDef struct {
	foo         string  `config:"foo"`
	bar         int     `config:"bar"`
}
func (scd *SampleConfDef) Foo() string {
	return scd.foo
}
func (scd *SampleConfDef) SetFoo(v string) error {
	scd.foo = v
	return nil
}
func (scd *SampleConfDef) OnUpdateFoo(val, oldVal string) {
	return
}
func (scd *SampleConfDef) Bar() int {
	return scd.bar
}
func (scd *SampleConfDef) SetBar(v string) error {
	ival, err := strconv.Atoi(v)
	if err != nil {
		return err
	}
	scd.bar = ival
	return nil
}
func (scd *SampleConfDef) OnUpdateBar(val, oldVal string) {
	return
}

func TestRegisterConfDef(t *testing.T) {
    zkAddr := "localhost:2181"
    config := &registry.RegConfigZooKeeper{ZKAddr: zkAddr}
    client := client.New(config)
    if client == nil {
        t.Fatal("error new client")
    }

	err := RegisterConfig("game", new(SampleConfDef), client)
	if err != nil {
		t.Error(err)
	}
    done := make(chan struct{})
    err = Load(done, zkAddr)
    if err != nil {
        t.Fatal(err)
    }
    <- done
	t.Log("test normal ok.")
}
