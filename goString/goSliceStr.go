package main

import (
	"fmt"
	"strings"
	"strconv"
	"reflect"
	"unicode/utf8"
	"bytes"
)

func main() {
	// main2()
	// main3()
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

// func SplitAfterN(s, sep string, n int) []string该函数s根据sep分割，返回分割之后子字符串的slice,和split一样，只是返回的子字符串保留sep，如果sep为空，那么每一个字符都分割
func main1() {
	fmt.Println(strings.SplitAfterN("a,b,c,d,r", ",", 4)) //["a," "b," "c," "d,r"]
	fmt.Println(strings.SplitAfterN("a,b,c,d,r", ",", 5)) //["a," "b," "c," "d," "r"]
	
	// fmt.Println(strings.SplitAfterN("118.812039,31.894425", ",", 2)) // ?
	
	var tempStr []string
	//tempStr = strings.SplitAfterN("118.812039,31.894425", ",", 2)
	// func Split(s, sep string) []string,有join就有Split这个就是把字符串按照指定的分隔符切割成slice
	tempStr = strings.Split("118.812039,31.894425", ",")
	s1 := tempStr[0]
	s2 := tempStr[1]	
	
	fmt.Println(s1)
	fmt.Println(s2)	
}

// GO数值和字符串的相互转换
// 在做项目的时候, 通常都会碰到字符串转换, 在这介绍一下字符串与整型的相互转换。
// 在golang中, 用字符串与整型有两种方法, 一种是使用rune(int32位的别名)来转换, 一种是golang中stroncv包的函数来转换, 下面的是第二种方法, 另外还介绍了如何获取接口类型所代表值的类型, 直接上代码:
func main2() {
	var a string
	a = "123456"
	
	b, error := strconv.Atoi(a)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
	}
	
	b = b + 1
	fmt.Println(b)
	
	var c int = 1234
	d := strconv.Itoa(c)	// 数字变成字符串
	d = d + "sdfs"
	fmt.Println(d)
	
	var f int64 = 1234
	g := strconv.FormatInt(f, 10)	// 数字变成字符串
	fmt.Println(g)
	
	var e interface {}
	//e = 10
	e = "10"
	
	switch v := e.(type) {
		case int:
			fmt.Println("整型",v)
			break;
		case string:
			fmt.Println("字符串",v)
		break;
	}
}

// go字符串的遍历输出
func main3() {
	// go中有两种方式对字符串进行遍历, 一种是utf-8遍历, 另一种是Unicode遍历。
    // utf-8和Unicode编码的区别详见http://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html
	str := "Hello,世界"
	
	// utf-8遍历
	fmt.Println("utf-8遍历")
    for i := 0; i < len(str); i++ {
		ch := str[i]
		fmt.Println(ch)
    }
	
	fmt.Println("=============>Unicode遍历")
	// Unicode遍历
	for _, ch1 := range str {
        fmt.Println(ch1)
    }

	/*
		上面代码执行后, 会打印一串数字而不是字符。
		这是由于go语言中的字符串实际上是类型为byte的只读切片。
		或者说一个字符串就是一堆字节。
		这意味着, 当我们将字符存储在字符串中时, 实际存储的是这个字符的字节。
		一个字符串包含了任意个byte, 它并不限定Unicode, UTF-8或者任何其他预定义的编码。
		那么go语言用什么来表示字符呢, 下面的例子可以验证一下:
	*/
	str1 := "Hello,世界"
	
	// utf-8遍历
	fmt.Println("\n=============>utf-8遍历")
	for i := 0; i < len(str1); i++ {
		ch := str1[i]
		ctype := reflect.TypeOf(ch)
		fmt.Printf("%s ", ctype)
    }
	
	fmt.Println("\n=============>Unicode遍历")
	// Unicode遍历
	for _, ch1 := range str1 {
		ctype := reflect.TypeOf(ch1)
		fmt.Printf("%s ", ctype)
    }

	/*
		代码运行后显示ch的类型为uint8, 也就是byte类型, 而ch1的类型为int32, 也就是rune类型。
		go语言中的源码定义为utf-8文本, 不允许其他的表示。
		但是也存在特殊处理, 那就是字符串上使用for…range循环。
		range循环迭代时, 就会解码一个utf-8编码的rune。
		现在既然已经知道上述不管哪种遍历方式, 其实质都是字节。
		所以在打印时, 只需要将这些结果转化为字符字面值或者转换其输出类型就可以了。
		下面是两种字符串遍历方式:
	*/
	str2 := "Hello,世界"
	
	fmt.Println("\n==========>方法一")
	// 方法一:	格式化打印
    for _, ch1 := range str2 {
        fmt.Printf("%q",ch1)			// 单引号围绕的字符字面值, 由go语法安全的转义
    }
    
	fmt.Println("\n==========>方法二")
	// 方法二:	转化输出格式
	for _, ch2 := range str2 {
		fmt.Println(string(ch2))
    }
}

// go rune 简要分析
func main4() {
	// 今天看golang代码看到一个单词rune, 熟悉而陌生
	// 之前学习go并没有过多注意这个“神秘符号”
	// rune在golang中是int32的别名, 在各个方面都与int32相同。 
	// 被用来区分字符值和整数值。
	s := "hello你好"
	fmt.Println(len(s))				// 输出长度为11
	fmt.Println(len([]rune(s)))		// 输出长度为7
	//fmt.Println("\n")
	
	s = "你好"
	fmt.Println(len(s))				// 输出长度为6
	fmt.Println(len([]rune(s)))		// 输出长度为2
	fmt.Println("\n")
	
	s = "你"
	fmt.Println([]byte(s))			// 输出长度为6
	fmt.Println(rune('你'))			// 输出20320
	fmt.Println("\n")
	
	// 通过上述代码可以将rune理解为一个可以表示unicode编码的值int的值, 称为码点(code point), 只不过go语言把这个码点抽象为rune。 
}

// go语言字符串处理
//		go语言处理字符串主要介绍下"strconv"和"strings", 直接看代码和注释, 如下:
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
    }
}

func main5() {
	fmt.Println(strings.Contains("abbccddd", "cd"))	// 字符串s中是否包含substr, 返回bool值, true
	fmt.Println(strings.Contains("aa", ""))			// true
	
	fmt.Println(strings.Join([]string{"sqn", "hello", "world"}, " - "))				// 字符串链接, 把slice a通过sep链接起来
	fmt.Println(strings.Index("aabcddae", "a"))										// 在字符串s中查找sep所在的位置, 返回位置值, 找不到返回-1
	fmt.Println(strings.Repeat("SQN", 5))											// 重复s字符串count次, 最后返回重复的字符串
	fmt.Println(strings.Replace("ababababababab", "ab", "sqn", 5))					// 在s字符串中, 把old字符串替换为new字符串, n表示替换的次数, 小于0表示全部替换
	fmt.Println(strings.Replace("ababababababab", "ab", "sqn", 35))		
	fmt.Println(strings.Split("sqn-sqn-sqn.sq-.sqn", "-"))							// 把s字符串按照sep分割, 返回slice
	fmt.Println(strings.Trim("sqnfewfewsqn", "sqn"))								// 在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Println(strings.Join(strings.Fields("a b sqn hello world sqn"), "***"))		// 去除s字符串的空格符, 并且按照空格分割返回slice, 并连接
    
	/* 字符串转化的函数在strconv中 */
	// Append系列函数将整数等转换为字符串后, 添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4444, 10)		// 10进制
	fmt.Println(str)
	fmt.Println(string(str))
	
	str = strconv.AppendInt(str, 55555, 2)		// 2进制
	fmt.Println(str)
	fmt.Println(string(str))
	 str = strconv.AppendBool(str, false)

    str = strconv.AppendQuote(str, "abcdefg")

    str = strconv.AppendQuoteRune(str, '单')

    str = strconv.AppendQuoteRune(str, '你')
	
	//Parse 系列函数把字符串转换为其他类型
    a1, err := strconv.ParseBool("false")
    checkError(err)
    b1, err1 := strconv.ParseFloat("123.23", 64)
    checkError(err1)
    c1, err2 := strconv.ParseInt("1234", 10, 64)
    checkError(err2)
    d1, err3 := strconv.ParseUint("12345", 10, 64)
    checkError(err3)
    e1, err4 := strconv.Atoi("1023")
    checkError(err4)
    fmt.Println(a1, b1, c1, d1, e1)
    b3 := "1000"
    fmt.Println(string(b3))
    m := 6
    fmt.Println(int(m))
    fmt.Println(float32(m))
    fmt.Println(float64(m))
}

// golang中获取字符串长度的几种方法
//	一、获取字符串长度的几种方法
/*
		使用bytes.Count()统计
		使用strings.Count()统计
		将字符串转换为[]rune后调用len函数进行统计
		使用utf8.RuneCountInString()统计  
*/
func main6() {
	str := "HelloWord"
	l1	:= len([]rune(str))
	l2 := bytes.Count([]byte(str), nil) - 1
	l3  := strings.Count(str, "") - 1
	l4	:= utf8.RuneCountInString(str)
	
	// 打印结果：都是 9  
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)  
}

// strings.Count函数和bytes.Count函数, 这两个函数的用法是相同, 只是一个作用在字符串上, 一个作用在字节上
func main7() {
	// strings中的Count方法
	// func Count(s, sep string) int{}		// 判断字符sep在字符串s中出现的次数, 没有找到则返回-1, 如果为空字符串("")则返回字符串的长度+1
	str := "HelloWorld"
	fmt.Println(strings.Count(str, "o"))	// 打印o出现的次数, 打印结果为2
	
	/*
		注:
			在Golang中, 如果字符串中出现中文字符不能直接调用len函数来统计字符串字符长度, 这是因为在Go中, 字符串是以UTF-8为格式进行存储的,
							在字符串上调用len函数, 取得的是字符串包含的byte的个数。
	*/
	str1	:= "HelloWorld"
	str2	:= "Hello, 世界"
	
	fmt.Println(len(str2))		// 打印结果:	13
	fmt.Println(len(str1))		// 打印结果:	9		(如果是纯英文字符的字符串, 可以使用来判断字符串的长度)
}

// Golang中获取中文字符串的子串字符位置及截取子串
func main8() {
	/*
			昨天准备用golang做一个简单的文本分析, 需要简单的对字符串进行一些操作, 在查看了strings和strconv库时,
				我没找到截取字符串的函数, 同时strings.Index返回的是子串的字节位置,
				例如这个例子:
			strings.Index("早上好, 张先生!","好")的返回值是6, 而不是2(从0开始算)。	
	*/
	
	// 于是我自己写了一个处理中文的返回字符串子串位置的函数, 思想其实很简单, 
	//		首先通过strings库中的Index函数获得子串的字节位置, 再通过这个位置获得子串之前的字节数组pre, 再将pre转换成[]rune, 获得[]rune的长度,
	//			便是子串之前字符串的长度, 也就是子串在字符串中的字符位置, 具体代码如下:
	
}

// 注意, 这里用的是string.Index函数, 类似的, 也可以写中文字符串的类似strings中的IndexAny,LastIndex等函数
func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str,substr)
	
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}
	
	return result  
}  

// 同样的思想, 我也写了一个截取中文字符串的函数, 如下:
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs	:= []rune(str)
	lth	:= len(rs)
	
	// 简单的越界判断  
	if begin < 0 {
		begin = 0  
	}
	
	if begin >= lth {
		begin = lth
	}
	
	end := begin + length
	
	if end > lth {
		end = lth
	}
	
	// 返回子串
	return string(rs[begin:end])  
}

// Go字符串统计
func WordCount(s string) map[string]int {
	var word string
	
	m := make(map[string]int)				// map是以string字符串stringsg中的字符为key的!		------>			最后map中有6个元素
	
	for i := 0; i < len(s); {
		word = s[i : i+1]
		fmt.Println(word)
		v, ok := m[word]					// map指定key取对应的value时, 可以指定返回两个值, 第一个是对应的value, 第二个是一个bool, 表示是否有值!
		if ok != false {
			m[word] = v + 1
		} else {
			m[word] = 1
		}
		i += 1
    }
	
	return m
}

func main9() {
	/*
			s
			t
			r
			i
			n
			g
			s
			g
			map[i:1 n:1 g:2 s:2 t:1 r:1]
	*/
    var str = "stringsg"
    fmt.Println(WordCount(str))
} 

// go字符串的遍历输出
//		go中有两种方式对字符串进行遍历, 一种是utf-8遍历, 另一种是Unicode遍历。
//		utf-8和Unicode编码的区别详见http://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html
func main10() {
	str := "Hello, 世界"
	
	// utf-8遍历
	fmt.Println("===========================>utf-8遍历")
	for i := 0; i < len(str); i++ {
        ch := str[i]
        fmt.Println(ch)					// 打印13行
    }
	
	// Unicode遍历
	fmt.Println("===========================>Unicode遍历")
	for _, ch1 := range str {
		fmt.Println(ch1)				// 打印9行
	}
	
	/*
		上面代码执行后, 会打印一串数字而不是字符!!!
		
		这是由于go语言中的字符串实际上是类型为byte的只读切片!!!
		或者说一个字符串就是一堆字节!!!
        这意味着, 当我们将字符存储在字符串中时, 实际存储的是这个字符的字节。
		一个字符串包含了任意个byte, 它并不限定Unicode, UTF-8或者任何其他预定义的编码。
		那么go语言用什么来表示字符呢, 下面的例子可以验证一下!!!
	*/
}

func main11() {
	str := "Hello,世界"
	
    // utf-8遍历
	fmt.Println("=============>utf-8遍历")
	for i := 0; i < len(str); i++ {
		ch		:= str[i]
		ctype	:= reflect.TypeOf(ch)
		fmt.Printf("%s ", ctype)					// 打印12行
    }
	
	// Unicode遍历
	fmt.Println("\n=============>Unicode遍历")
	for _, ch1 := range str {
		ctype := reflect.TypeOf(ch1)
		fmt.Printf("%s ",ctype)						// 打印8行 
	}
	
	/*
		代码运行后显示ch的类型为uint8, 也就是byte类型, 而ch1的类型为int32, 也就是rune类型。
		go语言中的源码定义为utf-8文本, 不允许其他的表示。
		但是也存在特殊处理, 那就是字符串上使用for…range循环。
		range循环迭代时, 就会解码一个utf-8编码的rune。
		现在既然已经知道上述不管哪种遍历方式, 其实质都是字节。
		所以在打印时, 只需要将这些结果转化为字符字面值或者转换其输出类型就可以了。
	*/
}

// 下面是两种字符串遍历方式
func main12() {
	str := "Hello,世界"
	
	// 方法一:  格式化打印
	fmt.Println("==========>方法一")
	for _, ch1 := range str {
		fmt.Printf("%q", ch1)				// 单引号围绕的字符字面值, 由go语法安全的转义, 打印出来共8个字符(1行以为)!
    }
	
	// 方法二:	转化输出格式
	fmt.Println("\n==========>方法二")
	for _, ch2 := range str {	
		fmt.Println(string(ch2))			// 打印出来共8行字符串!	
    }
}

// go语言学习----字符串、数组和切片的应用
func main13() {
	// 字符串、数组和切片的应用
	// 从字符串生成字节切片
	// 假设s是一个字符串(本质上是一个字节数组)
	// 那么就可以直接通过 c := []bytes(s) 来获取一个字节的切片c。
	// 另外, 您还可以通过 copy 函数来达到相同的目的：copy(dst []byte, src string)。
	// 同样的, 还可以使用 for-range 来获得每个元素(Listing 7.13—for_string.go):
	s := "\u00ff\u754c"
    for i, c := range s {
        fmt.Printf("%d:%c ", i, c)		// 输出:	0:ÿ 2:界
    }
	
	/*
			我们知道, Unicode字符会占用2个字节, 有些甚至需要3个或者4个字节来进行表示。
			如果发现错误的UTF8字符, 则该字符会被设置为U+FFFD并且索引向前移动一个字节。
			和字符串转换一样, 您同样可以使用c := []int(s)语法, 这样切片中的每个int都会包含对应的Unicode代码, 因为字符串中的每次字符都会对应一个整数。
			类似的, 您也可以将字符串转换为元素类型为rune的切片:				r := []rune(s)。
			可以通过代码len([]int(s))来获得字符串中字符的数量, 但使用utf8.RuneCountInString(s)效率会更高一点。(参考count_characters.go)
			
			您还可以将一个字符串追加到某一个字符数组的尾部:
					var b []byte
					var s string
					b = append(b, s...)
					
			
	*/

	var b []byte
	var s1 string = "jack"
	b = append(b, s1...)
	//b = append(b, s1)
	fmt.Println("\n")
	fmt.Println(s1)
	
	// 获取字符串的某一部分
	// 使用 substr := str[start:end] 可以从字符串str获取到从索引start开始到end-1位置的子字符串。
	// 同样的, str[start:] 则表示获取从 start 开始到 len(str)-1 位置的子字符串。
	// 而str[:end] 表示获取从 0 开始到 end-1的子字符串。
	
	// 字符串和切片的内存结构
	// 在内存中, 一个字符串实际上是一个双字结构, 即一个指向实际数据的指针和记录字符串长度的整数(见图 7.4)
	// 因为指针对用户来说是完全不可见, 因此我们可以依旧把字符串看做是一个值类型, 也就是一个字符数组。
	// 字符串 string s = "hello" 和子字符串 t = s[2:3] 在内存中的结构可以用下图表示：
	
	// 修改字符串中的某个字符
	// Go语言中的字符串是不可变的, 也就是说 str[index] 这样的表达式是不可以被放在等号左侧的。
	// 如果尝试运行str[i] = 'D'会得到错误：cannot assign to str[i]。
	// 因此, 您必须先将字符串转换成字节数组, 然后再通过修改数组中的元素值来达到修改字符串的目的, 最后将字节数组转换回字符串格式。
	// 例如, 将字符串 "hello" 转换为 "cello"：
	fmt.Println("\n")
	s2 := "hello"
	fmt.Println(s2)
	c := []byte(s2)
	fmt.Println(c)
	
	c[0] = 'c'
	s3 := string(c)		// s2 == "cello"
	fmt.Println(s3)		// 所以, 您可以通过操作切片来完成对字符串的操作!
	
    // 字节数组对比函数
	// 下面的Compare函数会返回两个字节数组字典顺序的整数对比结果, 即 0 if a == b, -1 if a < b, 1 if a > b。
		
}

func Compare(a, b[]byte) int {
    for i:=0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }
	
    // 数组的长度可能不同
    switch {
    case len(a) < len(b):
        return -1
    case len(a) > len(b):
        return 1
    }
    
	return 0 // 数组相等
}

// 搜索及排序切片和数组
/*
		标准库提供了sort包来实现常见的搜索和排序操作。
		您可以使用sort包中的函数func Ints(a []int)来实现对int类型的切片排序。
		例如sort.Ints(arri), 其中变量arri就是需要被升序排序的数组或切片。
		为了检查某个数组是否已经被排序, 可以通过函数IntsAreSorted(a []int) bool 来检查, 如果返回true则表示已经被排序。
		类似的, 可以使用函数func Float64s(a []float64)来排序float64的元素, 或使用函数func Strings(a []string)排序字符串元素。
		想要在数组或切片中搜索一个元素, 该数组或切片必须先被排序(因为标准库的搜索算法使用的是二分法)。
		然后, 您就可以使用函数func SearchInts(a []int, n int) int进行搜索, 并返回对应结果的索引值。
		当然, 还可以搜索 float64 和字符串:
				func SearchFloat64s(a []float64, x float64) int
				func SearchStrings(a []string, x string) int
		您可以通过查看 官方文档 来获取更详细的信息。
		这就是如何使用 sort 包的方法，我们会在第 11.6 节对它的细节进行深入, 并实现一个属于我们自己的版本。
*/

// append函数常见操作
/*
		我们在第 7.5 节提到的 append 非常有用, 它能够用于各种方面的操作:
			1.将切片 b 的元素追加到切片 a 之后：a = append(a, b...)
			2.复制切片 a 的元素到新的切片 b 上：
							b = make([]T, len(a))
							copy(b, a)
			3.删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)
			4.切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)
			5.为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)
			6.在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)
			7.在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)
			8.在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)
			9.取出位于切片 a 最末尾的元素 x：x, a = a[len(a)-1], a[:len(a)-1]
			10.将元素 x 追加到切片 a：a = append(a, x)
			
	因此，您可以使用切片和 append 操作来表示任意可变长度的序列。
	从数学的角度来看，切片相当于向量，如果需要的话可以定义一个向量作为切片的别名来进行操作。
	如果您需要更加完整的方案，可以学习一下 Eleanor McHugh 编写的几个包：slices、chain 和 lists。
*/

// 切片和垃圾回收
/*
		切片的底层指向一个数组, 该数组的实际体积可能要大于切片所定义的体积。
		只有在没有任何切片指向的时候, 底层的数组内层才会被释放, 这种特性有时会导致程序占用多余的内存。
		
		示例 
				函数 FindDigits 将一个文件加载到内存，然后搜索其中所有的数字并返回一个切片。
		
		var digitRegexp = regexp.MustCompile("[0-9]+")

		func FindDigits(filename string) []byte {
			b, _ := ioutil.ReadFile(filename)
			return digitRegexp.Find(b)
		}
		
		这段代码可以顺利运行, 但返回的 []byte指向的底层是整个文件的数据。
		只要该返回的切片不被释放, 垃圾回收器就不能释放整个文件所占用的内存。
		换句话说, 一点点有用的数据却占用了整个文件的内存。
		
		想要避免这个问题, 可以通过拷贝我们需要的部分到一个新的切片中:
		func FindDigits(filename string) []byte {
			b, _ := ioutil.ReadFile(filename)
			b = digitRegexp.Find(b)
			c := make([]byte, len(b))
			copy(c, b)
			return c
		}		
*/

func main14() {
	// golang substring方法的几种实现
	// 方法1:  原生方法, 直接使用slice切片实现, 但此方法对于包括中文字符就截取错误
	s := "abcde"
	fmt.Println(s[0:2]);			// 输出ab
	
	//包含中文的字符串
	s2 := "我是中国人"
	fmt.Println(s2[0:2])			// 输出有乱码呀
	
	// 第二种方法呢, 自己实现, 如下是我实现的几种方法, 最主要还是通过rune来实现的
	
}

// 获取source的子串, 如果start小于0或者end大于source长度则返回
// 
func substring(source string, start int, end int) string {
	
}
