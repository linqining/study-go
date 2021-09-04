package 闭包

// https://books.studygolang.com/advanced-go-programming-book/ch1-basic/ch1-04-func-method-interface.html
// go语言高级编程
// 闭包引用for range的指针， ide就有提示
func Closure() {
	for i := 0; i < 3; i++ {
		defer func(){ println(i) } ()
	}
}



func ClosureFix1() {
	for i := 0; i < 3; i++ {
		i := i // 定义一个循环体内局部变量i
		defer func(){ println(i) } ()
	}
}

func ClosureFix2() {
	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int){ println(i) } (i)
	}
}