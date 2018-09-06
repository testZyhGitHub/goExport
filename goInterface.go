package main

import "fmt"

func main() {
	main_tst_interface()
}

// GO interface显示类型转换方法
func main_tst_interface() {
	// var i interface{} = "TT"
	var i interface{} = 77

	value, ok := i.(int)
	if ok {
		fmt.Printf("类型匹配int:%d\n", value)
	} else {
		fmt.Println("类型不匹配int\n")
	}

	if value, ok := i.(int); ok {
		fmt.Println("类型匹配整型：%d\n", value)
	} else if value, ok := i.(string); ok {
		fmt.Printf("类型匹配字符串:%s\n", value)
	}
}