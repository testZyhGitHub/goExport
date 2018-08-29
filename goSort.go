package main

/*
			https://blog.csdn.net/u010983881/article/details/52460998
			https://blog.csdn.net/huhudeni/article/details/80255688
			https://blog.csdn.net/books1958/article/details/46896449
			https://blog.csdn.net/books1958/article/details/23773371
			https://blog.csdn.net/fyxichen/article/details/46595965
*/

/*
		基本类型排序和slice排序

		Go是通过sort包提供排序和搜索, 因为Go暂时不支持泛型(将来也不好说支不支持), 所以, Go的sort和search使用起来跟类型是有关的, 或是需要像c一样写比较函数等, 稍微显得也不是很方便!!!
		Go的排序思路和C和C++有些差别。
		C默认是对数组进行排序;
		C++是对一个序列进行排序;
		Go则更宽泛一些, 待排序的可以是任何对象, 虽然很多情况下是一个slice(分片, 类似于数组), 或是包含slice的一个对象!

*/
