package main

import (
	"fmt"
)

// 不过, 要注意的是, type绝不只是用于定义一系列的别名!!! 还可以针对新类型定义方法!!!
// name类型可以像下面这样定义方法:
type name string

func (n name) len() int {
	return len(n)
}

func main() {
	var myname name = "taozs"	// 其实就是字符串类型	
	l := []byte(myname)		// 字符串转字节数组

	fmt.Println(len(l))		// 字节长度
	
	var myStr name = "123456789"
	fmt.Println(myStr.len())	// 调用对象的方法
}