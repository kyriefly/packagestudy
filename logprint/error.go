package logprint

import (
	"fmt"
	"time"
)

func Error (err interface{}){
	t := time.Now()
	fmt.Printf("[Info], %s: %s\n", t.Format("2006-01-02 15:04:05"), err)
}