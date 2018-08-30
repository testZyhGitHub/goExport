package main

import (
	"fmt"
	"sync"
	"runtime"
	"time"
)

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
	// main6()
	// main7()
	// main8()
	// main9()
	// main10()
	// main11()
	// main12()
	// main13()
	main14()
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

func f1(in chan int) {
	fmt.Println(<-in)
}

func tstZuS2() {
	out := make(chan int)		// 这是由于out <- 2之前不存在对out的接收, 所以, 对于out <- 2来说, 永远是阻塞的, 即一直会等下去!!!
	out <- 2
	go f1(out)
}

func tstZuS3() {
	out := make(chan int)		// 这是由于out <- 2前存在对管道的读操作, 所以out <- 2 是合法的, 就像前文说的, 发送操作在接收者准备好之前是阻塞的!!!
	go f1(out)
	out <- 2
}

func main7() {
	/*
			golang协程		----->		通道channel阻塞

			说到channel, 就一定要说一说线程了!!!
			任何实际项目, 无论大小, 并发是必然存在的!!!
			并发的存在, 就涉及到线程通信!!!

			在当下的开发语言中, 线程通讯主要有两种, 共享内存与消息传递!!!
					(1).共享内存一定都很熟悉, 通过共同操作同一对象, 实现线程间通讯;
					(2).消息传递即通过类似聊天的方式;
					(3).golang对并发的处理采用了协程的技术;
					(4).golang的goroutine就是协程的实现;
					(5).协程的概念很早就有, 简单的理解为"轻量级线程", goroutine就是为了解决"并发任务间"的"通信"而设计的;
					(6).golang解决通信的理念是:				不要通过共享内存来通信, 而应该通过通信来共享内存!!!					--------->		解决方案是?
					(7).golang解决方案是"消息传递机制", "消息的传递"就是通过channel来实现的!!!
					(8).如何使用channel?				谈一谈对channe阻塞的理解?
			①.发送者角度:
								对于同一个通道, 发送操作(协程或者函数中的), 在接收者准备好之前是阻塞的;
								如果chan中的数据无人接收, 就无法再给通道传入其他数据;
								因为新的输入无法在通道非空的情况下传入;
								所以发送操作会等待chan再次变为可用状态:			就是通道值被接收时(可以传入变量)。

			②.接收者角度:
								对于同一个通道, 接收操作是阻塞的(协程或函数中的), 直到发送者可用;
								如果通道中没有数据, 接收者就阻塞了;

					(9).
					(10).
	*/

	// tstZuS2()
	tstZuS3()
}

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main8() {
	/* GoLang之协程、channel、select、同步锁 */
	// WebServer几种主流的并发模型
	/*
			1.多线程, 每个线程一次处理一个请求, 在当前请求处理完成之前不会接收其它请求; 但在高并发环境下, 多线程的开销比较大;
			2.基于回调的异步IO, 如Nginx服务器使用的epoll模型, 这种模式通过事件驱动的方式使用异步IO, 使服务器持续运转, 但人的思维模式是串行的, 大量回调函数会把流程分割, 对于问题本身的反应不够自然;
			3.协程, 不需要抢占式调度, 可以有效提高线程的任务并发性, 而避免多线程的缺点; 但原生支持协程的语言还很少!!!
			4.协程(coroutine)是Go语言中的轻量级线程实现, 由Go运行时(runtime)管理;
			5.在一个函数调用前加上go关键字, 这次调用就会在一个新的goroutine中并发执行;
			6.当被调用的函数返回时, 这个goroutine也自动结束;
			7.需要注意的是, 如果这个函数有返回值, 那么这个返回值会被丢弃!
			8.
	*/

	/*
			执行上面的代码, 会发现屏幕什么也没打印出来, 程序就退出了!!!
			下面的例子, main()函数启动了10个goroutine, 然后返回, 这时程序就退出了, 而被启动的执行Add()的goroutine没来得及执行!!!
			我们想要让main()函数等待所有goroutine退出后再返回, 但如何知道goroutine都退出了呢?   这就引出了多个goroutine之间通信的问题!!!

			在工程上, 有两种最常见的并发通信模型:		共享内存和消息
	*/
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}

var counter int = 0

// 因为10个goroutine是并发执行的, 所以我们还引入了锁, 也就是代码中的lock变量;
func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println("counter =", counter)
	lock.Unlock()
}

// 在工程上, 有两种最常见的并发通信模型:		共享内存和消息
// 下面的例子, 使用了锁变量(属于一种共享内存)来同步协程, 事实上Go语言主要使用消息机制(channel)来作为通信模型;
func main9() {
	/*
		来看下面的例子, 10个goroutine共享了变量counter, 每个goroutine执行完成后, 将counter值加1;
		因为10个goroutine是并发执行的, 所以我们还引入了锁, 也就是代码中的lock变量;
		在main()函数中, 使用for循环来不断检查counter值, 当其值达到10时, 说明所有goroutine都执行完毕了, 这时main()返回, 程序退出!!!
	*/

	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()

		runtime.Gosched()	// 出让时间片

		if c >= 10 {
			break
		}
	}

	fmt.Println("All Go routine run finished!!!")
}

// channel
/*
		消息机制认为每个并发单元是自包含的、独立的个体, 并且都有自己的变量, 但在不同并发单元间这些变量不共享;
		每个并发单元的输入和输出只有一种, 那就是消息;
		channel是Go语言, 在语言级别提供的goroutine间的通信方式, 我们可以使用channel在多个goroutine之间传递消息;
		channel是进程内的通信方式, 因此通过channel传递对象的过程和调用函数时的参数传递行为比较一致, 比如也可以传递指针等;
		channel是类型相关的, 一个channel只能传递一种类型的值, 这个类型需要在声明channel时指定;

		channel的声明形式为:
				var chanName chan ElementType

		举个例子, 声明一个传递int类型的channel:
				var ch chan int

		使用内置函数make()定义一个channel:
				ch := make(chan int)

		在channel的用法中, 最常见的包括写入和读出:
				将一个数据value写入至channel, 这会导致阻塞, 直到有其他goroutine从这个channel中读取数据;
				将一个数据value写入至channel，这会导致阻塞，直到有其他goroutine从这个channel中读取数据;
						ch <- value

		从channel中读取数据, 如果channel之前没有写入数据, 也会导致阻塞, 直到channel中被写入数据为止!!!
						value := <-ch

		可以关闭不再使用的channel:
						close(ch)

		我们还可以创建一个带缓冲的channel:
						c := make(chan int, 1024)		此时, 创建一个大小为1024的int类型的channel, 即使没有读取方, 写入方也可以一直往channel里写入, 在缓冲区被填完之前都不会阻塞!!!

		从带缓冲的channel中读数据
						for i:=range c {
						　　...
						}
*/

// 现在利用channel来重写上面的例子
/*
func Count1(ch chan int, goRoutineNo int) {
	ch <- 1									// 将一个数据value写入至channel, 这会导致阻塞, 直到有其他goroutine从这个channel中读取数据
	fmt.Println("Counting")

	//info := fmt.Sprintf("go rountine no: %d", goRoutineNo)
	//fmt.Println(info)
}
*/

func Count1(ch chan int, goRoutineNo int) {
	ch <- 1									// 将一个数据value写入至channel, 这会导致阻塞, 直到有其他goroutine从这个channel中读取数据
	//fmt.Println("Counting")

	info := fmt.Sprintf("%d\n", goRoutineNo)
	fmt.Print(info)
	// fmt.Println("Counting")
}

func main10() {
	chs := make([] chan int, 10)			// 一个channel只能传递一种类型的值, 创建一个带有10个缓冲的, 传递int类型的channel数组, 定义了一个包含10个channel的数组, 并把数组中的每个channel分配给10个不同的goroutine

	fmt.Println("Start to create channel......")

	// 在每个goroutine完成后, 向goroutine写入一个数据, 在这个channel被读取前, 这个操作是阻塞的!!!
	// 在所有的goroutine启动完成后, 依次从10个channel中读取数据, 在对应的channel写入数据前, 这个操作也是阻塞的!
	// 这样, 就用channel实现了类似锁的功能, 并保证了所有goroutine完成后main()才返回!!!
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count1(chs[i], i+1)
	}

	/*
	for j := 0; j < 100000; j++ {
		fmt.Println(j)
	}

	for k := 0 ; k < 10000; k++ {
		fmt.Println("\n")
	}
	*/

	for _, ch := range(chs) {
		<-ch
	}

	//time.Sleep(10 * time.Second)
	time.Sleep(1)							// 主线程还是需要休息一把!!!
	//time.Sleep(1 * time.Second)

	//fmt.Println("Print all huanhang......")
	fmt.Println("End to create channel......")
}

// 在一个函数中使用单向读channel
func Parse(ch <-chan int) {
	for value := range ch {
		fmt.Println("Parsing value", value)
	}
}

// 我们在将一个channel变量传递到一个函数时, 可以通过将其指定为单向channel变量, 从而限制该函数中可以对此channel的操作!!!
// 单向channel的作用有点类似于c++中的const关键字, 用于遵循代码"最小权限原则"!!!
//
func main11() {
	// 单向channel变量的声明
	// var ch1 chan int			// 普通channel
	// var ch2 chan <- int		// 只用于写int数据
	// var ch3 <-chan int		// 只用于读int数据

	// 可以通过类型转换, 将一个channel转换为单向的:
	// ch4 := make(chan int)
	// ch5 := <-chan int(ch4)	// 单向读
	// ch6 := chan<- int(ch4)	// 单向写
}

/*
		channel作为一种原生类型, 本身也可以通过channel进行传递, 例如下面这个流式处理结构:
*/
type PipeData struct {
	value		int
	handler		func(int) int
	next		chan int
}

func handle(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

// select
/*
		在UNIX中, select()函数用来监控一组描述符, 该机制常被用于实现高并发的socket服务器程序;
		Go语言直接在语言级别支持select关键字, 用于处理异步IO问题, 大致结构如下:

		select {
			case <- chan1:			// 如果chan1成功读到数据
			case chan2 <- 1:		// 如果成功向chan2写入数据
			default:				// 默认分支
		}

		Go语言没有对channel提供直接的超时处理机制, 但我们可以利用select来间接实现, 例如:
*/
func main12() {
	timeout := make(chan bool, 1)
	ch := make(chan bool, 1)

	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	switch {
		// 这样使用select就可以避免永久等待的问题, 因为程序会在timeout中获取到一个数据后继续执行, 而无论对ch的读取是否还处于等待状态!!!
		case <- ch:			// 从ch中读取到数据
		case <- timeout: 	// 没有从ch中读取到数据, 但从timeout中读取到了数据
	}
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

// GO Channel并发、死锁问题
func main13() {
	// 如果不是我对真正并行的线程的追求, 就不会认识到Go有多么的迷人!!!
	// Go语言从语言层面上就支持了并发, 这与其他语言大不一样, 不像以前我们要用Thread库, 来新建线程, 还要用线程安全的队列库来共享数据!!!

	// Go语言的goroutines、信道和死锁
	/*
			Go语言中有个概念叫做goroutine, 这类似我们熟知的线程, 但是更轻!!!

	*/

	loop()
	loop()
}

//
// 连表查询
// https://blog.csdn.net/smilesundream/article/details/80209026
// https://blog.csdn.net/kjfcpua/article/details/17710331
// https://blog.csdn.net/phantom_111/article/details/79489313
// https://studygolang.com/articles/10172
// https://blog.csdn.net/netdxy/article/details/54564436
// https://www.cnblogs.com/suoning/p/7259106.html
// https://www.cnblogs.com/suoning/p/7237444.html
//
func main14() {
	// 把一个loop放在一个goroutine里跑，我们可以使用关键字go来定义并启动一个goroutine
	go loop()				// 启动一个goroutine
	loop()					// 可是为什么只输出了一趟呢? 明明我们主线跑了一趟, 也开了一个goroutine来跑一趟啊!!! 原来, 在goroutine还没来得及跑loop的时候, 主函数已经退出了!!!

	// main函数退出地太快了, 我们要想办法阻止它过早地退出, 一个办法是让main等待一下:
	time.Sleep(time.Second) // 停顿一秒

	/*
			可是采用等待的办法并不好, 如果goroutine在结束的时候, 告诉下主线说"Hey, 我要跑完了!!!", 就好了, 即所谓阻塞主线的办法, 回忆下我们Python里面等待所有线程执行完毕的写法:
			for thread in threads:
					thread.join()

			是的, 我们也需要一个类似join的东西来阻塞住主线, 那就是信道!!!

	*/
}

