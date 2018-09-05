package main

import (
	"time"
	"fmt"
)

func main() {

}

// Go语言使用error类型来返回函数执行过程中遇到的错误, 如果返回的error值为nil, 则表示未遇到错误, 否则error会返回一个字符串, 用于说明遇到了什么错误!!!
// error只是一个接口, 定义如下:
/*
		type error interface {
			Error() string
		}
*/

// 你可以用任何类型去实现它(只要添加一个Error()方法即可), 也就是说, error可以是任何类型, 这意味着, 函数返回的error值实际可以包含任意信息, 不一定是字符串(当然字符串是必须的)。
// error不一定表示一个错误, 它可以表示任何信息, 比如io包中就用error类型的io.EOF表示数据读取结束, 而不是遇到了什么错误。
//			再比如path/filepath包中用error类型的filepath.SkipDir表示跳过当前目录, 继续遍历下一个目录, 而不是遇到了什么错误。
// errors包实现了一个最简单的error类型, 只包含一个字符串, 它可以记录大多数情况下遇到的错误信息!!!
// errors包的用法也很简单, 只有一个New 函数, 用于生成一个最简单的error对象：
//
func main1() {

}

// 将字符串 text 包装成一个 error 对象返回
// func New(text string) error

// 示例
/*
func SomeFunc() error {
	if 遇到错误 {
		return errors.New("遇到了某某错误")
	}
	return nil
}
*/

// 如果你的程序需要记录更多的错误信息, 比如时间、数值等信息, 可以声明一个自定义的 error 类型。
type myError struct {
	err   string
	time  time.Time
	count int
}

func (m *myError) Error() string {
	return fmt.Sprintf("%s %d 次。时间：%v", m.err, m.count, m.time)
}

func newErr(s string, i int) *myError {
	return &myError{
		err:   s,
		time:  time.Now(),
		count: i,
	}
}

var count int

func SomeFunc() error {
	if true {
		count++
		return newErr("遇到某某情况", count)
	}
	return nil
}

func main2() {
	for i := 0; i < 5; i++ {
		if err := SomeFunc(); err != nil {
			fmt.Println(err)
		}
	}
}