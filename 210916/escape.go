package main
import "fmt"

func main() {
	f := foo("Ding")
	fmt.Println(f)
}

type bar struct {
	s string
}

func foo(s string) bar {
	f := new(bar) // 这里的new(bar)会不会发生逃逸？？？
	defer func() {
		f = nil
	}()
	f.s = s
	return *f
}