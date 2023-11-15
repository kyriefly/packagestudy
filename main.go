package main

import (
	// "fmt"
	"github.com/kyriefly/packagestudy/cook/meat"
	"github.com/kyriefly/packagestudy/cook/vegetable"
	// "packagestudy/logprint"
)

func main() {
  // fmt.Println("程序开始启动...")
  // fmt.Println("程序正在运行...")
  // fmt.Println("程序完成退出...")

  // logprint.Debug("这是一个debug日志, 用来描述程序的运行日志")
  // logprint.Info ("这是一个info日志, 用来描述程序的运行日志")
  // logprint.Error("这是一个err日志, 用来描述程序的运行日志")

  // 做的荤菜
  meat.Makedish("红烧肉")
  // 做的素菜
  vegetable.Makedish("土豆丝")
}