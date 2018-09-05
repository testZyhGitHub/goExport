package main

import (
	"fmt"
	"math"
)

// go语言练习:   幂、函授接收和返回参数、转义字符、变量和常量
func main() {
	// main1()
	// main2()
	// main3()
	// main4()
	// main5()
	// main6()
	main7()
}

// 实现a^b次方
func main1() {
	// r2 := power1(2, 4)
	// r2 := power1(2, 5)
	// r2 := power1(2, 6)
	// r2 := power1(2, 7)
	// r2 := power1(2, 8)
	// r2 := power1(2, 9)
	// r2 := power1(2, 10)

	//r2 := power1(10, 4)
	//r2 := power1(10, 8)
	// r2 := power1(10, 10)
	r2 := power1(10, 16)
	println(r2)
}

// a是底数, b是幂, 返回r
func power1(a uint64, b uint64) (r uint64) {
	var i		uint64
	var temp	uint64

	if a != 0 {
		temp = 1

		for i = 1; i <= b; i++ {
			temp = temp * a
		}

		return temp
	}

	return		// 这里return只是占一个位置, 不然会报错, 实际不返回任何值, 并且这个函数在调用时候, 只能接收一个值, 否则会报错
}

// 接收两个参数, 返回两个数值
func main2() {
	t1, t2 := test(10,20)
	println(t1)
	println(t2)
}

func test(a int, b int)(r1 int, r2 int)  {
	return a + b, a - b
}

// 转义字符练习
func main3() {
	fmt.Println("test\n")    //换行符
	fmt.Println("test\\n")   //使用转义字符，打印\n；
	fmt.Println("test\rttt") //遇到\r的时候，换行打印后面内容
	fmt.Println("test\ttt")  //空一格tab键的长度，打印后面内容
	fmt.Println("test\vtt")  //空格长度是\t的一半
}

// 利用go求幂的几种方法
// 二分幂法, 求x^n
func main4() {
	var x float64
	var n int
	fmt.Scanf("%f%d", &x, &n)
	fmt.Println(powerf(x, n))
	fmt.Println(powerf2(x, n))
	fmt.Println(powerf3(x, n))
	fmt.Println(math.Pow(x, float64(n)))
}

func main5() {
	//n := powerfInt64(10,16)
	n := powerfInt64(2,10)
	fmt.Println(n)
}

func main6() {
	// n := powerInt64(2,10)
	n := powerInt64(10,16)
	fmt.Println(n)
}

// 二分幂法: 求x^n
func powerf(x float64, n int) float64 {
	ans := 1.0

	for n != 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n /= 2
	}
	return ans
}

// 二分幂法: 求x^n
func powerInt64(x int64, n int) int64 {
	var ans int64 = 1
	var m	 int

	for n != 0 {
		m = n % 2

		if m == 1 { //如果n是奇数, 就要多乘一次
			ans *= x
		}

		x *= x
		n /= 2		// 二分
	}

	return ans
}

// 递归法:		求x^n
func powerf2(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else {
		return x * powerf2(x, n-1)
	}
}

func powerfInt64(x int64, n int) int64 {
	if n == 0 {
		return 1
	} else {
		return x * powerfInt64(x, n-1)
	}
}

// 循环法 求x^n
func powerf3(x float64, n int) float64 {
	ans := 1.0

	for n != 0 {
		ans *= x
		n--
	}
	return ans
}

/*
	报名数=0，显示0；
		  0<报名数<=9，显示数字，如6；
		  10直接返回10;
		  10<报名数<=99，显示十位数+，如63、67都显示60+；
		  100<=报名数<=999，显示百位数+，如635、675都显示600+；
		  1000<=报名数<=9999，显示千位数+，如6355、6755都显示6000+；
		  10000<=报名数<=99999，显示万位数+，如63555、67555都显示60000+；
		  100000<=报名数，显示100000+；
*/
func NumberTakeInt(number int) int {
	if (0 < number) && (number <= 9) {
		return 6
	} else if 10 == number {
		return 10
	} else if (10 < number) && (number <= 99) {
		return (number - (number % 10))
	} else if (100 <= number) && (number <= 999) {
		return (number - (number % 100))
	} else if (1000 <= number) && (number <= 9999) {
		return (number - (number % 1000))
	} else if (10000 <= number) && (number <= 99999) {
		return (number - (number % 10000))
	} else {
		return 100000
	}
}

func main7() {
	// var inputNum int = 1
	// var inputNum int = 9
	// var inputNum int = 10

	// var inputNum int = 11
	// var inputNum int = 19
	// var inputNum int = 21
	// var inputNum int = 51
	// var inputNum int = 98
	// var inputNum int = 99

	// var inputNum int = 100
	// var inputNum int = 199
	// var inputNum int = 200
	// var inputNum int = 299
	// var inputNum int = 998
	// var inputNum int = 999

	// var inputNum int = 1000
	// var inputNum int = 7859
	// var inputNum int = 9999

	// var inputNum int = 10000
	// var inputNum int = 78983
	// var inputNum int = 99999

	// var inputNum int = 100000
	var inputNum int = 100345

	ret := NumberTakeInt(inputNum)
	fmt.Println(ret)
}