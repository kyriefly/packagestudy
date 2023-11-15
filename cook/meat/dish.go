package meat

import "github.com/kyriefly/packagestudy/logprint"

func  Makedish(name string){
	msg := "做的荤菜" + name
	logprint.Info(msg)
}