package _10917

import (
	"fmt"
	"sync"
	"time"
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

func AssignMult(){
	var a []int =nil
	a, a[0] = []int{1,2},9
	fmt.Println(a)
}


type data struct{
	name string
}

func (p *data)print(){
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func Printer(){
	d1:= data{"one"}
	d1.print()

	var in printer = &data{"two"}
	in.print()
}


func AssignSlice (){
	a:=[3] int {0,1,2}
	fmt.Println(a[1:2])
	s:=a[1:2]
	s[0] = 11
	s=append(s,12)
	s=append(s, 13)
	s[0]=21
	fmt.Println(a)
	fmt.Println(s)
}

type foo struct {Val int}
type bar struct {Val int}

func TypeCompare(){
	a := &foo{Val:5}
	b := &foo{Val: 5}
	c := foo{Val: 5}
	d := bar{Val:5}
	e := bar{Val: 5}
	f := bar{Val:5}
	fmt.Println(a ==b, c ==foo(d), e==f)
}

func A() int{
	time.Sleep(1000*time.Millisecond)
	return 1
}

func B()int {
	time.Sleep(10000*time.Millisecond)
	return 2
}


// todo 为什么select 时间不一样为什么还是等了10秒
func RandomOutput(){
	ch := make(chan int, 1)
	go func(){
		select {
		case ch <-A():
		case ch <-B():
		default:
			ch<-3
		}
	}()
	fmt.Println(<-ch)
}

type Point struct {x,y int}

func SwapPoint(){
	s := []Point{
		{1,2},
		{3,4},
	}

	for  _, p:=range s{
		p.x,p.y  = p.y, p.x
	}
	fmt.Println(s)
}

func SwapPoint2(){
	s := []*Point{
		&Point{1,2},
		&Point{3,4},
	}

	for  _, p:=range s{
		p.x,p.y  = p.y, p.x
	}
	fmt.Println(*s[0])
	fmt.Println(*s[1])
}
