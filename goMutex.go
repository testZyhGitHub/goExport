package main

import (
	"sync"
	"fmt"
)

func main() {
	main1()
}

// 同步锁
// 另外, 为了更好的地控制并行中的原子操作, sync包还提供了一个atomic子包, 支持对于一些基础数据类型的原子操作函数, 比如经典的CAS函数:
//					func CompareAndSwapUnit64(val *uint64, old, new uint64) (swapped bool)

/*
		Go语言包中的sync包提供了两种锁类型:			sync.Mutex和sync.RWMutex, 前者是互斥锁, 后者是读写锁!!!

		使用锁的经典模式:
			var lck sync.Mutex

			func foo() {
				lck.Lock()
				defer lck.Unlock()
				// ...
			}
		lck.Lock()会阻塞直到获取锁, 然后利用defer语句在函数返回时自动释放锁!!!

		对于从全局角度只需要运行一次的代码, 比如全局初始化操作, Go语言提供了一个once类型来保证全局的唯一性操作, 如下:
*/
var flag int32
var once sync.Once

func initialize() {
	flag = 3
	fmt.Println(flag)
}

func setup() {
	once.Do(initialize)
}

func main1() {    // flag只别打印了一次!!!
	setup()
	setup()
}