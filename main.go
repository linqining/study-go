package main

import (
	"fmt"
	"runtime"
)

func main(){
	// 没有抢占调度前的版本是不会输出，后面的版本可以抢占输出。

	runtime.GOMAXPROCS(1)
	go func(){
		for i:=0;i<10;i++{
			fmt.Println(i)
		}
	}()
}
