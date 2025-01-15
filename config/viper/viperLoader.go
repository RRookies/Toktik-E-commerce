package viper

type ViperLoader struct {

}

func NewViperLoader() *ViperLoader {
	return &ViperLoader{}
}

func (vp *ViperLoader) initConfig(){
	var err error
	if err = vp.AddRemoteProvider("etcd3","http://127.0.0.1:2379","/config/viper-test/con")
}
