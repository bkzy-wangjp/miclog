// test project main.go
package main

import (
	"fmt"
	"time"

	"github.com/bkzy-wangjp/miclog"
)

func main() {
	miclog.Config("./log", "logtest", 1024, 3)
	for i := 0; i < 200; i++ {
		fmt.Printf("Hello World! i=%d\n", i)
		miclog.Info("Hello World! i=%d", i)
	}
	time.Sleep(10 * time.Second)
}
