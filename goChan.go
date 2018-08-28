package main

import "fmt"

/*
		Golang关于通道Chan详解
			1.线程: 在golang里面也叫goroutine
			2.并发与并行???
			3.golang的线程是一种并发机制, 而不是并行;

		并发和并行区别?
			1.并发的关键是你有处理多个任务的能力, 不一定要同时;
			2.并行的关键是你有同时处理多个任务的能力;

		并发和并行有什么区别?
			做并发编程之前, 必须首先理解什么是并发, 什么是并行, 什么是并发编程, 什么是并行编程。
			并发(concurrency)和并行(parallellism)是:
				1.并行是指两个或者多个事件在同一时刻发生; 而并发是指两个或多个事件在同一时间间隔发生;
				2.并行是在不同实体上的多个事件, 并发是在同一实体上的多个事件;
				3.在一台处理器上"同时"处理多个任务, 在多台处理器上同时处理多个任务, 如hadoop分布式集群;

		所以并发编程的目标是充分的利用处理器的每一个核, 以达到最高的处理性能!!!

		在golang里面, 使用go这个关键字, 后面再跟上一个函数就可以创建一个线程!!!
				后面的这个函数可以是已经写好的函数, 也可以是一个匿名函数!!!


*/

func main() {
	// main1()
	// main4()
	// main5()
	main6()
}

func main1() {
	go fmt.Println("1")
	fmt.Println("2")
}

func main2() {
	/*
			golang里channel的实现原理

			channel是消息传递的机制, 用于多线程环境下lock free synchronization.
			它同时具备2个特性:
					(1).消息传递
					(2).同步

			golang里的channel的性能???

			自带的runtime package里已经提供了benchmark代码, 可以运行下面的命令查看其性能
				go test -v -test.bench=".*" runtime

			channel的实现，都在$GOROOT/src/pkg/runtime/chan.c里

			它是通过共享内存实现的:
				struct Hchan {
				}

			具体的实现是chan.c里的 Hchan* runtime·makechan_c(ChanType *t, int64 hint)
			此时, hint=5, t=interface{}


	*/

	// ch := make(chan interface{}, 5)
}

func main3() {
	/*
				golang之Channel
				Channel是Go中的一个核心类型, 可以将其看成一个管道, 通过它并发单元就可以发送或者接收数据进行通信(communication)!!!

				Do not communicate by sharing memory; instead, share memory by communicating.

				channel基础知识
					(1).创建channel:		使用内建函数make创建channel
					(2).
					(3).

	*/

	// 使用内建函数make创建channel
	// 	unBufferChan	:= make(chan int)		// 1 无缓冲的channel
	// 	bufferChan		:= make(chan int, N)	// 2 带缓冲的channel

	//	无缓冲:			发送和接收动作是同时发生的, 如果goroutine读取channel(<-channel), 则发送者(channel<-)会一直阻塞;
	//  缓冲channel:	类似一个有容量的队列, 当队列满的时候发送者会阻塞; 当队列空的时候接收者会阻塞!!!
}

func main4() {
	/* 往一个nil channel中发送数据会一直被阻塞, 从一个nil channel中接收数据会一直被block, 所以才会产生如下死锁的现象!!! */

	// fatal error: all goroutines are asleep - deadlock!
	//
	var x chan int
	go func() {
		x <- 1
	}()
	<-x
}

// channel读写操作
func main5() {
	/*				channl使用要小心, 使用前切记要初始化, 初始化函数用make!!!		*/
	var x int

	ch := make(chan int, 10)				// 初始化函数用make

	// 读操作
	// x <- ch

	// 写操作
	ch <- x
}

func main6() {
	// golang 单向管道使用

	/*
			一直听说代码即注释的概念, 但是一直没有一个具体的概念, 看到golang中通过单向chan的来做代码即注释的例子!!!

			单向管道
						对于单向channel我们可以这样定义!!!
						发送值的通道类型
												chan<-T
						接收值的通道类型
												<-chan T

			中单向管道应用场景
						在os/signal中使用了如下定义Notify函数只会对该通道发送元素值, 而不会从该通道接收值!
							func Notify(c chan<- os.Signal, sig ...os.Signal)
			这里, 问题来了, 对于单向的通道如何来应用那?
					毕竟只向一个通道发送值, 而没有接收过程是没有意义的。

			单向管道应用
					在func Notify(c chan<- os.Signal, sig …os.Signal)中, chan<- 表达了该函数只会向通道发送数据。
					我们在调用此函数的时候, 从表面上看需要传入一个只能发送元素不能接收元素的通道, 但是传入这样的通过是错误的!!!

			函数的传入通道应该是一个双向通道
					调用过程中, 我们应向函数传入双向通道并自觉遵守这个隐性规定, 传入的双向通道会转为一个单向通道!!!

			

	*/
}