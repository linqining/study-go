package _10904


func PanicRun() {
	println("=== begin ===")
	defer func() { // defer_0
		println("=== come in defer_0 ===")
	}()
	defer func() { // defer_1
		recover()
	}()
	defer func() { // defer_2
		panic("panic 2")
	}()
	panic("panic 1")
	println("=== end ===")
}

func PanicRun2() {
	println("=== begin ===")
	defer func() { // defer_0
		println("=== come in defer_0 ===")
	}()
	defer func() { // defer_1
		panic("panic 2")
	}()
	defer func() { // defer_2
		recover()
	}()
	panic("panic 1")
	println("=== end ===")
}