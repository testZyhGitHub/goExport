package main

import (
	"fmt"
	"errors"
	"testing"
	"os"
)

const NotFound = "{\"Code\":990004,\"Desc\":\"未找到服务\"}"
var NotFoundErr = errors.New(NotFound)

type InvokeServiceRet struct {
	ret string
	err error
}

// Go关键字type
/*
	type有如下几种用法:
		1. 定义结构体
		2. 定义接口
		3. 类型定义
		4. 类型别名
		5. 类型查询
*/

// 定义结构体
/*
		结构体是用户自定义的一种抽象的数据结构, golang中struct类似于java语言中的class,
		在程序设计中, 有着举足轻重的地位!!!
		结构体的用法, 将会在struct关键字中详细的介绍。

		下边来看一下定义一个结构体的语法格式:

*/
type name struct {
	// Field1  dataType
	// Field2  dataType
	// Field3  dataType

	Field1  int
	Field2  string
	Field3  rune
}

// 定义接口
/*
		接口相关知识点, 将会在interface关键字中详细介绍, 下边来看一段定义接口的语法格式
*/
type nameInter interface {
	Read()
	Write()
}

// 类型定义
/*
	使用类型定义定义出来的类型与原类型不相同, 所以不能使用新类型变量赋值给原类型变量, 除非使用强制类型转换。
	下面来看一段示例代码, 根据string类型, 定义一种新的类型, 新类型名称是name
	为什么要使用类型定义呢?
	类型定义可以在原类型的基础上创造出新的类型, 有些场合下可以使代码更加简洁, 如下边示例代码:
	上边的示例是类型定义的一种简单应用场合, 如果不使用类型定义, 那么想要实现上边示例中的功能, 应该怎么书写这段代码呢?

	// exec函数, 接收handle类型的参数
	func exec(f func(str string)) {
		f("hello")
	}

	exec函数中的参数类型, 需要替换成func(str string)了, 咋一看去也不复杂, 但是假如exec接收一个需要5个参数的函数变量呢? 是不是感觉参数列表就会很长了!!!
	func exec(f func(str string, str2 string, num int, money float64, flag bool)) {
		f("hello")
	}

	从上边的代码可以发现, exec函数的参数列表可读性变差了, 下边再来看看使用类型定义是怎么实现这个功能!!!
*/
type nameNewType string

type handle func(str string)	// 定义一个接收一个字符串类型参数的函数类型

// 定义一个需要五个参数的函数类型
type handleFive func(str string, str2 string, num int, money float64, flag bool)

// exec函数, 接收handle类型的参数
func exec(f handle) {
	f("hello")
}

// execFive函数, 接收handle类型的参数
func execFive(f handleFive) {
	f("hello", "world", 10, 11.23, true)
}

func demo(str string, str2 string, num int, money float64, flag bool) {
	fmt.Println(str, str2, num, money, flag)
}

// 类型别名
/*
		类型别名这个特性在golang1.9中引入。
		使用类型别名定义出来的类型与原类型一样, 即可以与原类型变量互相赋值, 又拥有了原类型的所有方法集。
		给strng类型取一个别名, 别名名称是name:
				type name = string
		类型别名与类型定义不同之处在于, 使用类型别名需要在别名和原类型之间加上赋值符号(=);
		使用类型别名定义的类型与原类型等价, 而使用类型定义出来的类型是一种新的类型!!!
		如下边示例:

*/

type a = string			// 类型别名			------>			原类型
type b string			// 类型定义			------>			新类型

type S string			// 根据string类型, 定义类型S

func SayA(str a) {
	fmt.Println(str)
}

func SayB(str b) {
	fmt.Println(str)
}

func main1() {
	var str = "test"
	SayA(str)

	// 错误参数传递，str是字符串类型，不能赋值给b类型变量
	/*
			从错误信息可知, str为字符串类型, 不能当做b类型参数传入SayB函数中。
			而str却可以当做a类型参数传入到SayA函数中。
			由此可见, 使用类型别名定义的类型与原类型一致, 而类型定义定义出来的类型, 是一种新的类型!!!

			给类型别名新增方法, 会添加到原类型方法集中!!!

			给类型别名新增方法后, 原类型也能使用这个方法。

			下边请看一段示例代码:

	*/
	// SayB(str)
}

func (r *S) Hi() {
	fmt.Println("S hi")
}

// 定义S的类型别名为T
type T = S
func (r *T) Hello() {
	fmt.Println("T hello")
}

// 函数参数接收S类型的指针变量
func execS(obj *S) {
	obj.Hello()
	obj.Hi()
}

/*
		上边的示例中, S是原类型, T是S类型别名。
		在给T增加了Hello方法后, S类型的变量也可以使用Hello方法。
			说明给类型别名新增方法后, 原类型也能使用这个方法。
		从示例中可知, 变量t可以赋值给S类型变量s, 所以类型别名是给原类型取了一个小名, 本质上没有发生任何变化。

		类型别名, 只能对同一个包中的自定义类型产生作用。
		举个例子, golang sdk中有很多个包, 是不是我们可以使用类型别名, 给sdk包中的结构体类型新增方法呢?
						答案是：不行。
		请牢记一点:
						类型别名, 只能对包内的类型产生作用, 对包外的类型采用类型别名, 在编译时将会提示如下信息:
										cannot define new methods on non-local type string


*/
func main2() {
	t := new(T)
	s := new(S)
	execS(s)
	execS(t)			// 将T类型指针变量传递给S类型指针变量
}

// 类型查询
/*
		类型查询, 就是根据变量, 查询这个变量的类型。

		为什么会有这样的需求呢?

		goalng中有一个特殊的类型interface{}, 这个类型可以被任何类型的变量赋值, 如果想要知道到底是哪个类型的变量赋值给了interface{}类型变量, 就需要使用类型查询来解决这个需求,	示例代码如下:

		如果使用.(type)查询类型的变量不是interface{}类型, 则在编译时会报如下错误:
					cannot type switch on non-interface value a (type string)
		如果在switch以外地方使用.(type), 则在编译时会提示如下错误:
					use of .(type) outside type switch
		所以, 使用type进行类型查询时, 只能在switch中使用, 且使用类型查询的变量类型必须是interface{}!!!
*/
func main() {
	var a interface{} = "abc"			// 定义一个interface{}类型变量, 并使用string类型值"abc"初始化

	// 在switch中使用:				变量名.(type)查询变量是由哪个类型数据赋值。
	switch v := a.(type) {
	case string:
		fmt.Println("字符串")
	case int:
		fmt.Println("整型")
	default:
		fmt.Println("其他类型", v)
	}
}

func TestServer_billhelper(t *testing.T) {
	fmt.Printf("welcome to test")
	// os.Exit(code)
}

func TestServer_f1(t *testing.T) {
	var m *testing.M
	code := m.Run()
	os.Exit(code)
}

func Test_f1(t *testing.T) {
	fmt.Println(NotFoundErr)

	var a int

	a += 10

	fmt.Println(a)

	// var t1 handle
	// exec(t1)

	var p = func(str string) {					// 定义一个函数类型变量, 这个函数接收一个字符串类型的参数
		fmt.Println("first: ", str)
	}
	exec(p)

	// 匿名函数作为参数直接传递给exec函数
	exec(func(str string) {
		fmt.Println("second: ", str)
	})
}

func Test_f2(t *testing.T) {
	execFive(demo)
	main1()
}