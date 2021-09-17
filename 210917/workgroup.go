package _10917

import (
	"fmt"
	"sync"
)

func NoDone(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		fmt.Println("1")
		wg.Done()
		wg.Add(1) // 没有done导致panic
	}()
	wg.Wait()
}
