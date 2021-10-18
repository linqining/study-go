package _10917

import "fmt"

func f(){
	defer func(){
		if r:=recover();r!=nil{
			fmt.Printf("revocer:%#v", r)
		}
	}()
	panic(1)
	panic(2)
}


func main(){
	f()
}
