package main

import (
	"fmt"
)

func main() {
	// main1()
	// main2()
	// main3()
	main4()
}

func main1() {
	a := make([]int, 10, 20)
	fmt.Printf("%d, %d\n", len(a), cap(a))
	fmt.Println(a)
	
	b := a[:cap(a)]
	fmt.Println(b)
}

// golang分配内存有一个make函数
//		该函数第一个参数是类型, 第二个参数是分配的空间, 第三个参数是预留分配空间
//		例如 a:=make([]int, 5, 10), len(a)输出结果是5, cap(a)输出结果是10, 然后我对a[4]进行赋值发现是可以得, 
//				但对a[5]进行赋值发现报错了, 于是郁闷这个预留分配的空间要怎么使用呢?
//		于是google了一下发现原来预留的空间需要重新切片才可以使用!!!

func main2() {
	//a := make([]int, 5, 10)		// 浪费5个内存!!!
	a := make([]int, 0, 10)
	// a := make([]int, 0)
	
	// a[4] = 5
	// fmt.Println(a)
	
	// 但对a[5]进行赋值发现报错了，于是郁闷这个预留分配的空间要怎么使用呢，于是google了一下发现原来预留的空间需要重新切片才可以使用，于是做一下记录
	// a[5] = 6
	// b := 6

	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)
	a = append(a, 5)
	a = append(a, 6)			// append函数会在数组容量不足的时候自动扩容一倍!
	a = append(a, 7)
	a = append(a, 8)
	a = append(a, 9)
	a = append(a, 10)
	a = append(a, 11)
	fmt.Println(a)
	fmt.Printf("%d, %d\n", len(a), cap(a))
}

type SPEntNameCount struct {
	Count	int
}

func main3() {
	retList1 := make([]*SPEntNameCount, 0, 3)
	for a := 0; a < 7; a++ {
		record := new(SPEntNameCount)
		record.Count = a
		retList1 = append(retList1, record)
	}

	retList2 := make([]SPEntNameCount, 0, 3)		// 没有空间报错!!!
	//retList2 := make([]SPEntNameCount, 3, 3)
	// for b := 0; b < 3; b++ {
	for b := 0; b < 4; b++ {			// 超出内存范围报错!!!
		retList2[b].Count = b
	}
}

func main4() {
	retMap := map[int64]int64{}

	var b int64
	var c int64 = 10
	for b = 0; b < 4; b++ {			// 超出内存范围报错!!!

		if b > 0 {
			//retMap[b] += c
			retMap[b] = c
		}

		c++
	}

}
