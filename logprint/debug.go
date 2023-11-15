// 包名定义，最好是小写字母，不要带下划线，见名知意
package logprint

import (
	"fmt"
	"time"
)

// interface 代表任意类型
func Debug (msg interface{}){
	t := time.Now()
	fmt.Printf("[Debug], %s: %s\n", t.Format("2006-01-02 15:04:05"), msg)
}