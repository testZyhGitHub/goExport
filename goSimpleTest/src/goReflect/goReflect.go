package main

// Go Reflect
/*
		1.最近在看一些go语言标准库以及第三方库的源码时, 发现go的reflect被大量使用;
		2.虽然反射的机制大多数语言都支持, 但好像都没有go一样这么依赖反射的特性;
		3.个人觉得, reflect使用如此频繁的一个重要原因离不开go的另一个特性, 空接口interface{}, reflect配合空接口, 让原本是静态类型的go具备了很多动态类型语言的特征;
		4.另外, 虽然反射大大增加了go语言的灵活性, 但要完全掌握它的原理和使用也还是有一点难度的!!!
		5.go的reflect库有两个重要的类型:
					reflect.Type
					reflect.Value

		6.Type,Value分别对应对象的类型和值数据
				还有两个重要的函数:
						reflect.TypeOf(i interface{}) Type				reflect.TypeOf()返回值的类型就是reflect.Type
						reflect.ValueOf(i interface{}) Value			reflect.ValueIOf()返回值的类型就是reflect.Value

		7.reflect.Type
				reflect.TypeOf(i interface{}) Type
				因为reflect.Typeof的参数是空接口类型, 因此可以接收任意类型的数据。
				TypeOf()的返回值是这个接口类型对应的reflect.Type对象。


		8.

		9.
*/