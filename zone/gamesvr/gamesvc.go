package main

//implement app.Service
type serviceGame struct {}
func (sg *serviceGame) OnInit() error {
    return nil
}
func (sg *serviceGame) Serve() error {
    return nil
}
func (sg *serviceGame) Close() {}
func (sg *serviceGame) OnFini() {
}

