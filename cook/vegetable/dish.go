package vegetable

import "github.com/kyriefly/packagestudy/logprint"


func Makedish(name string){
	msg := "做的素菜" + name
	logprint.Info(msg)
}