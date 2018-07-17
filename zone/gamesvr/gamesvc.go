package main

//implement app.Service
type serviceGame struct {}
func (sg *serviceGame) OnInit() error {
    var err error
    err = initCall()
    if err != nil {
        return err
    }
    err = initConfig()
    if err != nil {
        return err
    }
    return nil
}
func (sg *serviceGame) Serve() error {
    return nil
}
func (sg *serviceGame) Close() {}
func (sg *serviceGame) OnFini() {
}
