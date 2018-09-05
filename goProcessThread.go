package main

import (
	"fmt"
	"time"
)

/*
	1.线程, 在golang里面也叫goroutine;
	2.了解一下并发与并行	
	3.golang的线程是一种并发机制, 而不是并行;
	4.在golang里面, 使用go这个关键字, 后面再跟上一个函数就可以创建一个线程!!!
	5.后面的这个函数可以是已经写好的函数, 也可以是一个匿名函数;
	6.
	7.
*/

func main() {
	// main2()
	// main3()
	// main3_jjfa1()
	// main3_jjfa2()
	// main5()
	// main6()
	// main7()
	main8()
}

func main1() {
	go fmt.Println("1")
	fmt.Println("2")
}

func main2() {
	//var i=3
	var i=5
	
	// 上面的代码就创建了一个匿名函数, 并且还传入了一个参数i, 下面括号里的i是实参, a是形参!!!
	//		那么下面的代码能按照我们预想的打印1、2、3吗?  告诉你们吧, 不能, 程序只能打印出2!
	go func(a int) {
		fmt.Println(a)
		fmt.Println("1\n")
	}(i)
	
	// 让主线程休眠1秒的代码, 那为什么会这样呢?
	// 因为程序会优先执行主线程, 主线程执行完成后, 程序会立即退出, 没有多余的时间去执行子线程!
	// 如果在程序的最后让主线程休眠1秒钟, 那程序就会有足够的时间去执行子线程!!!
	time.Sleep(1 * time.Second) 
	//time.Sleep(3 * time.Second) // 让主线程休眠3秒的代码
	fmt.Println("3")
}

// 通道又叫channel, 顾名思义, channel的作用就是在多线程之间传递数据的。
// 创建无缓冲channel:		chreadandwrite :=make(chan int)
// 创建只读channel:			chonlyread := make(<-chan int)
// 创建只写channel:			chonlywrite := make(chan<- int)  
func main3() {
	ch :=make(chan int)			// 创建无缓冲channel
	ch <- 1
	
	// 这段代码执行时会出现一个错误:  fatal error: all goroutines are asleep - deadlock! 
	// 这个错误的意思是说线程陷入了死锁, 程序无法继续往下执行, 那么造成这种错误的原因是什么呢???
	// 我们创建了一个无缓冲的channel, 然后给这个channel赋值了, 程序就是在赋值完成后陷入了死锁!!!
	// 因为我们的channel是无缓冲的, 即同步的, 赋值完成后来不及读取channel, 程序就已经阻塞了。 
	// 这里介绍一个非常重要的概念:	channel的机制是先进先出, 如果你给channel赋值了, 那么必须要读取它的值, 不然就会造成阻塞, 当然这个只对无缓冲的channel有效!!!
	//									对于有缓冲的channel, 发送方会一直阻塞直到数据被拷贝到缓冲区, 如果缓冲区已满, 则发送方只能在接收方取走数据后才能从阻塞状态恢复!!!
	go func() {
		<- ch
		fmt.Println("1")
	}()
	
	fmt.Println("2")  	
}

// 对于上面的例子有两种解决方案
func main3_jjfa1() {
	// 给channel增加缓冲区, 然后在程序的最后让主线程休眠一秒, 代码如下:
	ch := make(chan int, 1)
	//ch <- 1
	ch <- 555
		
	go func() {
		v := <-ch
		fmt.Println(v)
	}()
	
	time.Sleep(1 * time.Second)
	fmt.Println("2")					// 这样的话程序就会依次打印出1、2 	
}

// 把ch<-1这一行代码放到子线程代码的后面, 代码如下:
func main3_jjfa2() {
	ch := make(chan int)
	
	go func() {
		v := <-ch
		fmt.Println(v)
	}()
	
	//ch <- 1
	//ch <- 1111111
	fmt.Println("2")		// 这里就不用让主线程休眠了, 因为channel在主线程中被赋值后, 主线程就会阻塞, 直到channel的值在子线程中被取出。
	ch <- 2222222
	//time.Sleep(10 * time.Second)
	time.Sleep(2 * time.Second)
}

// golang函数以及函数和方法的区别
//		在接触到go之前, 我认为函数和方法只是同一个东西的两个名字而已(在我熟悉的c/c++, python, java中没有明显的区别);
//		
func main4() {
}

// 最后我们看一个生产者和消费者的例子:		通道不带缓冲!
func main5() {
	/*
		在这段代码中, 因为channel是没有缓冲的, 所以当生产者给channel赋值后, 生产者这个线程会阻塞, 直到消费者线程将channel中的数据取出。
		消费者第一次将数据取出后, 进行下一次循环时, 消费者的线程也会阻塞, 因为生产者还没有将数据存入, 这时程序会去执行生产者的线程。
		程序就这样在消费者和生产者两个线程间不断切换, 直到循环结束!!!
	*/
	ch := make(chan int)				// 创建一个无缓冲的channel ch
	go produce(ch)						// 生产者给通道ch赋值!				<-------		生产者线程会阻塞,  直到消费者线程将channel中的数据取出!
	go consumer(ch)						// 消费者从通道ch取值!				<-------		
	//time.Sleep(1 * time.Second)
	time.Sleep(10 * time.Second)
	fmt.Println("main thread over!!!")
}

func produce(p chan<- int) {  			// p是一个只写的channel
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}

func consumer(c <-chan int) {			// c是一个只读的channel
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}

// 下面我们再看一个带缓冲的例子
func main6() {
	/* 在这个程序中, 缓冲区可以存储10个int类型的整数, 在执行生产者线程的时候, 线程就不会阻塞, 一次性将10个整数存入channel, 在读取的时候, 也是一次性读取!!! */
	ch := make(chan int, 10)
	go produce1(ch)
	go consumer1(ch)
	time.Sleep(1 * time.Second)
}

func produce1(p chan<- int) {
    for i := 0; i < 10; i++ {
        p <- i
        fmt.Println("send:", i)
    }
}

/*
func FUNCTION(msgs Msgs) {
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
}
*/

func consumer1(c <-chan int) {
    for i := 0; i < 10; i++ {
        v := <-c
        fmt.Println("receive:", v)
    }
}

/*
		go语言的go func(){}()是什么意思???
		
		http.HandleFunc("/next", handler)
		go func() {
			for i := 0; ; i++ {
				nextID <- i
			}
		}()
		
		http.ListenAndServe("localhost:8080", nil)
		
		go func表示func这个函数会是以协程的方式运行, 这样就可以提供程序的并发处理能力!!! 定义并调用函数func, 以并发的方式调用匿名函数func!!!
*/

/*
func main7() {
	//forever := make(chan bool)
	forever := make(chan bool) //此处叫channel

	// 此处提取出匿名函数, 便于理解
	
	// "<-forerver" 这是什么语法?
	// 这句语法的意思等待其他routine向forever这个channel中传递值,
	//			通过这种方式来阻塞<-forerver语句所在的routine的退出,
	//			一旦有其他routine向forerver中传入值,
	//		<-forerver语句所在的routine就会解除阻塞, 继续向下执行代码
			
	// go func(){}() 这又是什么语法？引用	------->  这条语句的意思是新生成一个routine去执行func(){}这个函数的逻辑
	go func() {															// go func(){}()就是一个立即执行的匿名函数, 是启动一个goroutine, go的协程, 来执行后边的func
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	
	// go FUNCTION(msg)		// go关键字启动一个携程(可以简单理解为轻量级的线程)
		
	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// <-forever				// 可以看成从channel中pop一个值, 如果没有, 则阻塞!!!   	<-forerver就是从channel里取数据, 这里由于channel是空的, 所以会产生阻塞!!!
	
	// 另外, 你要是真的想实现CTRL+C执行优雅退出的话, 这样写:
	forever1 := make(chan os.Signal)
	
	// 监听这么多信号, 具体的忘完了, 自己查吧!!!
	signal.Notify(forever1, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP, syscall.SIGTERM)
	
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever1
	
	// doSomeCleanHere()
	
	// 因为和题目的几个关键词比较像:	协程, 协程+匿名函数, channel
}
*/

// golang函数加go关键字后怎样返回值 --------> 可以用channel
func main8() {
	ch := make(chan int)
	
	//go sum(1, 2, ch)
	
	go sum(11, 22, ch)
	
	// 获取结果
	fmt.Println(<-ch)
}

func sum(a, b int, ch chan int) {
	// 将结果发送到channel
	ch <- a + b
}





