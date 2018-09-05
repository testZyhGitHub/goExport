package main

import (
	"fmt"
)

// go type用法1 ---> 结构体定义
type person struct {
	name string // 注意后面不能有逗号
	age int
}

func main() {
	// 结构体初始化
	p := person {
		name: "taozs",  // 注意后面要加逗号
		age: 18,		// 或者下面的}提到这儿来可以省略逗号
	}
	
	fmt.Println(p.name)
	fmt.Println(p.age)
	
	/*
		// 初始化字段不一定要全部指定, 比如下面也是可以的, name默认取长度为0的空字符串
		p := person {
			age: 18,
		}		
	*/
}

