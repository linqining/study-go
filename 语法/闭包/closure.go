package 闭包

// https://books.studygolang.com/advanced-go-programming-book/ch1-basic/ch1-04-func-method-interface.html
// go语言高级编程
// 闭包引用for range的指针， ide就有提示
func Closure() {
	for i := 0; i < 3; i++ {
		defer func(){ println(i) } ()
	}
}