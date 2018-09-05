package main

import (
	"fmt"
)

func work() {
	Log(10)

	// 调用panic函数相当于throw一个异常，逐层启动异常处理流程。在调用panic之前defer的操作会在调用panic后立即执行。
	// panic函数接受一个参数，任意类型，无返回值；
	panic("ERROR!") // 显示调用

	Log(100)
}

func Log(args ...interface{}) {
	fmt.Println(args...)
}

/*
func main(){
	println("HelloWorld")
	work()
}
*/

func main() {
	defer func() {
		// 调用recover函数相当于catch了异常，会中止异常处理流程，并可以返回这个异常。
		// recover函数没有参数，返回值就是异常本身；
		// 一般，recover函数放在defer后面的一个匿名函数中执行。个人认为还应该放在函数的首部；
		if r := recover(); r != nil {
			Log(r)
		}
	}()

	work()

	println("Exit main func!") // 不会打印, 程序在捕获异常后直接从panic.go文件中退出了!!!
}
