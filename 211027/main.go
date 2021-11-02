package main

import "fmt"

type Sport struct{
	Name string
}

var s Sport
func DouInit(){
	s.Name = "排球"

}


func init(){
	s.Name = "游泳"
}



func main(){

	fmt.Println(s.Name)
}
