package main

import (
	"strings"
	"fmt"
	"unicode/utf8"
	"encoding/json"
	"reflect"
	"strconv"
)

func main() {
	// main1()
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
	// main15()
	// main16()
	// main19()
	// main20()
	main21()
}

// LastIndex 返回子串 sep 在字符串 s 中最后一次出现的位置
// 如果找不到，则返回 -1，如果 sep 为空，则返回字符串的长度
// 使用朴素字符串比较算法实现
// func LastIndex(s, sep string) int
func main1() {
	s := "Hello,世界! Hello!"
	i := strings.LastIndex(s, "h")
	fmt.Println(i) // -1
	i = strings.LastIndex(s, "H")
	fmt.Println(i) // 14
	i = strings.LastIndex(s, "")
	fmt.Println(i) // 20
}

// golang截取字符串, 对于字符串操作, 截取字符串是一个常用的, 而当你需要截取字符串中的一部分时, 可以使用像截取数组某部分那样来操作, 示例代码如下:
func main2() {
	str := "XBodyContentX"
	content := str[1 : len(str)-1]
	fmt.Println(content)
}

// SplitAfter(s, sep string) []string
// 说明:		将字符串s中的字符串以字符sep为分隔符拆分成若干个字符串切片并且保留原字符串中的分隔符号, 并返回字符串切片!
func main3() {
	//str := "one,two,three"
	str := "one,two,three,four,five"
	var dstStr string = ""
	var i int = 0

	for _, v := range strings.SplitAfter(str, ",") {
		fmt.Println(v)
		dstStr += v
		i++
		if 3 == i {
			break;
		}
	}

	fmt.Println(dstStr)
	fmt.Println(i)
}

// Split(s, sep string) []string
// 说明:		将字符串s中的字符串以字符sep为分隔符拆分成若干个元素的字符串切片, 并返回字符串切片
func main4() {
	// str := "one,two,three"
	//str := "one,two,three,four,five"
	str := "239,256,268,238,244"
	var dstStr string = ""
	var i int = 0

	for _, v := range strings.Split(str, ",") {
		fmt.Println(v)
		// strings.Join(dstStr, ",")
		i++
		if 3 == i {
			dstStr += v
			break;
		}
		dstStr += v + ","
	}

	fmt.Println(dstStr)
	fmt.Println("here!!!")
}

// Replace(s, old, new string, n int) string
// 将字符串 s 中出现字符 old 的前 n 个替换成 new字符，并返回替换后的字符串，如果要替换全部则 n 为 -1 即可
func main5() {
	//str := "hi hi hi are you ok"
	str := "hi hi hi are you ok"
	fmt.Println(strings.Replace(str, "hi", "ok", 3))
	fmt.Println(str)
	fmt.Println(strings.Replace(str, "hi", "ok", 2))
	fmt.Println(strings.Replace(str, "hi", "ok", -1))
}

// Repeat(s string, count int) string
// 说明:	将 count 个字符串 s 合并成一个字符串并返回
func main6() {
	str := "Hello "
	fmt.Println(strings.Repeat(str, 5))
}

// LastIndex(s, sep string) int
// 说明: 判断字符 sep 在字符串 s 中最后一次出现的位置，如果成功返回 sep 位置的索引，如果字符 sep 不在字符串 s 中则返回 -1
func main7() {
	str := "Hello World"
	fmt.Println(strings.LastIndex(str, "l"))
}

// Join(a []string, sep string) string
// 说明:   将一个字符串切片中的元素以字符 sep 进行分隔然后合并成一个字符串并返回
func main8() {
	str := []string{"Hello", "World", "Good"}
	// fmt.Println(strings.Join(str, " "))
	fmt.Println(strings.Join(str, ","))
}

// 测试下split函数处理空字符串会怎么样子
func main9() {
	// long_lat := r.GetString("long_lat")
	long_lat := ""
	// long_lat := "121.382158,31.391967"
	tempStr := strings.Split(long_lat, ",")

	if 2 == len(tempStr) {
		Longitude := tempStr[0]										// 经度
		Latitude  := tempStr[1]										// 纬度

		fmt.Println(Longitude)
		fmt.Println(Latitude)
	} else {
		fmt.Printf("tempStr is empty!!!")
	}
}

// golang里的字符串替换函数使用方法
func main10 () {
	// golang里可以通过strings.Replace来进行字符串的替换, 下面是一段strings.Replace使用的演示代码, 这段golang代码用于将字符串中的空格替换成逗号。
	str := "welcome to sharejs.com"
	str = strings.Replace(str, " ", ",", -1)
	fmt.Println(str)

	idNumStr1 := "11010519491231002X"
	str = strings.Replace(idNumStr1, "X", "0", -1)
	fmt.Println(str)

	idNumStr2 := "110105194912a1002x"
	str = strings.Replace(idNumStr2, "x", "0", -1)
	fmt.Println(str)

	idNumStr3 := "11010519491261002x"
	str = strings.Replace(idNumStr3, "x", "0", -1)
	fmt.Println(str)

	idNumStr4 := "1XXXX5XXXXXX3X002X"
	str = strings.Replace(idNumStr4, "X", "0", -1)
	fmt.Println(str)

	idNumStr5 := "1xx1xx1xxx1xx1002x"
	str = strings.Replace(idNumStr5, "x", "0", -1)
	fmt.Println(str)
}

// Go语言中字符串的查找方法小结
func main11() {
	// func Contains(s, substr string) bool这个函数是查找某个字符是否在这个字符串中存在, 存在返回true
	fmt.Println(strings.Contains("widuu", "wi")) //true
	fmt.Println(strings.Contains("wi", "widuu")) //false

	// func ContainsAny(s, chars string) bool这个是查询字符串中是否包含多个字符
	fmt.Println(strings.ContainsAny("widuu", "w&d")) //true

	// func ContainsRune(s string, r rune) bool,这里边当然是字符串中是否包含rune类型，其中rune类型是utf8.RUneCountString可以完整表示全部Unicode字符的类型
	fmt.Println(strings.ContainsRune("widuu", rune('w'))) //true
	fmt.Println(strings.ContainsRune("widuu", 20))        //fasle

	// func Count(s, sep string) int这个的作用就是输出，在一段字符串中有多少匹配到的字符
	fmt.Println(strings.Count("widuu", "uu")) //1
	fmt.Println(strings.Count("widuu", "u"))  //2

	// func Index(s, sep string) int 这个函数是查找字符串，然后返回当前的位置，输入的都是string类型，然后int的位置信息
	fmt.Println(strings.Index("widuu", "i")) //1
	fmt.Println(strings.Index("widuu", "u")) //3

	// func IndexAny(s, chars string) int 这个函数是一样的查找，字符串第一次出现的位置，如果不存在就返回-1
	fmt.Println(strings.IndexAny("widuu", "u")) //3

	// func IndexByte(s string, c byte) int,这个函数功能还是查找第一次粗线的位置，只不过这次C是byte类型的，查找到返回位置，找不到返回-1
	fmt.Println(strings.IndexByte("hello xiaowei", 'x')) //6

	// func IndexRune(s string, r rune) int，还是查找位置，只不过这次是rune类型的
	fmt.Println(strings.IndexRune("widuu", rune('w'))) //0

	// func IndexFunc(s string, f func(rune) bool) int这个函数大家一看就知道了，是通过类型的转换来用函数查找位置，我们来代码看下哈
	// fmt.Println(strings.IndexFunc("nihaoma", split)) //3

	// https://yq.aliyun.com/ziliao/99569?spm=a2c4e.11155472.blogcont.16.14a68deaOAZtte
}

// golang截取字符串
func main12() {
	BlueCollarIdCardNo := "432831196411150810"

	str := Substr2(BlueCollarIdCardNo, 0, 2)
	fmt.Printf(str)
}

func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// 截取字符串 start 起点下标 end 终点下标(不包括)
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

// go 迭代string数组,直接拷贝去用即可
func main13() {
	subsCodes := []string{"aaaa", "vvvvv", "dddd", "eeeee", "gfgggg"}
	for _, s := range subsCodes {
		fmt.Println(s)
	}

	arrVarifyCode := []string{"1", "0", "x", "9", "8", "7", "6", "5", "4", "3", "2"}
	for _, s := range arrVarifyCode {
		fmt.Println(s)
	}
}

// golang 一行代码 把array/slice转成逗号分隔的字符串
func main14() {
	// [a] -> a -> a
	// [a b c] -> a b c -> a,b,c
	// strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)
}

func main15() {
	// string 不能直接和byte数组转换
	// string可以和byte的切片转换

	// 1,string 转为[]byte
	//var str string = "test"
	var str string = "11010519491231002"
	var data []byte = []byte(str)

	fmt.Println(data)

	res := strings.Replace(strings.Trim(fmt.Sprint(data), "[]"), " ", ",", -1) // golang 一行代码 把array/slice转成逗号分隔的字符串
	fmt.Printf(res)


	// 2,byte转为string
	var data1 [10]byte
	data1[0] = 'T'
	data1[1] = 'E'

	var str1 string = string(data1[:])
	fmt.Printf(str1)

	// golang的字符称为rune, 等价于C中的char, 可直接与整数转换
	var c rune='a'
	var i int =98
	i1:=int(c)
	fmt.Println("'a' convert to",i1)
	c1:=rune(i)
	fmt.Println("98 convert to",string(c1))

	//string to rune
	for _, char := range []rune("世界你好") {
		fmt.Println(string(char))
	}

	// rune实际是整型，必需先将其转换为string才能打印出来，否则打印出来的是一个整数
	c2 := 'a'
	fmt.Println(c2)
	fmt.Println(string(c2))
	fmt.Println(string(97))

	// rune在golang中是int32的别名，在各个方面都与int32相同, 被用来区分字符值和整数值。
	s:="hello你好"
	fmt.Println(len(s))//输出长度为11
	fmt.Println(len([]rune(s)))//输出长度为7
	s="你好"
	fmt.Println(len(s))//输出长度为6
	fmt.Println(len([]rune(s)))//输出长度为2
	s="你"
	fmt.Println([]byte(s))//输出长度为6
	fmt.Println(rune('你'))//输出20320

	// 通过上述代码可以将rune理解为 一个 可以表示unicode 编码的值int 的值，称为码点（code point）。只不过go语言把这个码点抽象为rune。

	// 浅析rune，byte
	// golang内置类型有rune类型和byte类型。
	// 需要知晓的是rune类型的底层类型是int32类型, 而byte类型的底层类型是int8类型, 这决定了rune能比byte表达更多的数。

	// 在unicode中, 一个中文占两个字节, utf-8中一个中文占三个字节, golang默认的编码是utf-8编码, 因此默认一个中文占三个字节, 但是golang中的字符串底层实际上是一个byte数组。
	// 因此可能会出现下面这种奇怪的情况
	str10 := "hello 世界"
	fmt.Println(len(str10)) //12

	// 我们期望得到的结果应该是8, 原因是golang中的string底层是由一个byte数组实现的, 而golang默认的编码是utf-8, 因此在这里一个中文字符占3个字节, 所以获得的长度是12.
	// 想要获得我们想要的结果也很简单, golang中的unicode/utf8包提供了用utf-8获取长度的方法
	str11 := "hello 世界"
	fmt.Println(utf8.RuneCountInString(str11)) //8

	// 上面说了byte类型实际上是一个int8类型, int8适合表达ascii编码的字符, 而int32可以表达更多的数, 可以更容易的处理unicode字符, 因此, 我们可以通过rune类型来处理unicode字符!!!
	str12 := "hello 世界"
	str13 := []rune(str12)
	fmt.Println(len(str13)) //8

	/*
			这里会将申请一块内存, 然后将str12的内容复制到这块内存, 实际上这块内存是一个rune类型的切片, 而str13拿到的是一个rune类型的切片的引用, 我们可以很容易的证明这是一个引用
	*/
	str14 := "hello 世界"
	str15 := []rune(str14)
	t := str15
	t[0] = 'w'
	fmt.Println(string(str15))		// "wello世界", 通过把str15赋值给t, t上改变的数据, 实际上是改变的是t指向的rune切片, 因此, str14也会跟着改变

	// 字符串的遍历
	// 对于字符串, 看一下如何遍历吧, 也许你会觉得遍历轻而易举, 然而刚接触golang的时候, 如果这样遍历字符串, 那么将是非常糟糕的
	str16 := "hello 世界"
	for i := 0;i < len(str16);i++ {
		fmt.Println(string(str16[i]))
	}

	// 如何解决这个问题呢？
	// 第一个解决方法是用range循环, 原因是range会隐式的unicode解码, 第二个方法是将str转换为rune类型的切片, 这个方法上面已经说过了, 这里就不再赘述了!!!  当然还有很多方法, 其本质都是将byte向rune上靠!!!
	str17 := "hello 世界"
	for _,v := range str17 {
		fmt.Println(string(v))
	}

    // rune和byte的区别:    除开rune和byte底层的类型的区别, 在使用上, rune能处理一切的字符, 而byte仅仅局限在ascii

	// 数组和切片
	/*
			在golang中，有一个“罕见”的复合类型，叫切片，切片是基于数组的，golang和其它语言不一样，在golang中，数组是不可变的，对数组进行类型转换等操作都会导致golang隐式的申请一块内存，然后将原数组的内容复制到这块内存。
			数组是不可变的，这就决定了字符串也是不可变的，因为字符串底层就是一个byte数组实现的。
			在实际的开发当中，我们经常使用的是切片，而不是数组。
			https://blog.csdn.net/haodawang/article/details/79988189
			https://blog.csdn.net/haodawang/article/details/80005072
			https://blog.csdn.net/haodawang/article/details/80006059

	*/

	// go 数组(array)、切片(slice)、map、结构体(struct)
	/*
				https://www.cnblogs.com/jackylee92/p/6171897.html
				http://blog.haohtml.com/archives/14239
				https://blog.csdn.net/YanJiangbo/article/details/43602055
				http://blog.51cto.com/guanyu/1895203
	*/
}

// struct --> map --> string
type demo struct {
	Id   string
	Name string
}

// go语言之结构体数组转为string字符串, 转换顺序:  先将struct结构体转为map,  再将map转为string字符串
func main16() {
	demos := []demo{demo{Id: "1", Name: "zs"}, demo{Id: "2", Name: "ls"}, demo{Id: "3", Name: "ww"}}

	for _, v := range demos {
		tmpdata := Struct2Map(v)
		str, err := json.Marshal(tmpdata)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(str))
	}
}

// 结构体转为map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// go语言初始化内部结构体3中方式
func main17() {
	type User struct {
		Id   int
		Name string
		Age  int
	}

	type Manger struct {
		User
		title string
	}

	m := Manger{User:User{1, "ok", 12}, title:"123"}		// 可以
	m2 := Manger{User{1, "ok", 12}, "123"}		// 可以
	m3 := Manger{User:User{Id:1, Name:"ok", Age:12}, title:"123"}				// 可以

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)

	// golang中结构体的初始化方法(new方法)

	// 自定义一个结构体
	type Rect struct {
		x		float64
		y		float64
		width	float64
		height	float64
	}

	/*
	// 初始化方法
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width:100, height:200}

	// 注意这几个变量全部为指向Rect结构的指针(指针变量)，因为使用了new()函数和&操作符．而如果使用方法
	a := Rect{}  // 则表示这个是一个Rect{}类型．两者是不一样的．参考代码：
	*/

	rect1 := &Rect{0, 0, 100, 200}
	rect1.x = 10

	a := Rect{}
	a.x = 15

	fmt.Printf("%v\n%T\n", a, a)
	fmt.Printf("%v\n%T\n", rect1, rect1)

	/*
			在Go语言中, 未进行初始化的变量都会被初始化为该类型的零值, 例如bool类型的零值为false, int类型的零值为0, string类型的零值为空字符串;
			在Go语言中没有构造函数的概念, 对象的创建通常交由一个全局的创建函数来完成, 以NewXXX来命令, 表示"构造函数":
			func NewRect(x ,y ,width, height float64) {
					return &Rect{x, y, width, height}
			}

			用 new 分配内存 内建函数 new 本质上说跟其他语言中的同名函数功能一样：new(T) 分配了零值填充的 T 类型的内存空间，并且返回其地址，一个 *T 类型的值。
			用 Go 的术语说，它返回了一个指针，指向新分配的类型 T 的零值。
			记住这点非常重要。
			这意味着使用者可以用 new 创建一个数据结构的实例并且可以直接工作。
			如 bytes.Buffer的文档所述 “Buffer 的零值是一个准备好了的空缓冲。”
			类似的，sync.Mutex 也没有明确的构造函数或 Init 方法。
			取而代之，sync.Mutex 的零值被定义为非锁定的互斥量。
			零值是非常有用的。
			例如这样的类型定义，56 页的”定义自己的类型” 内容。

			务必记得make仅适用于map, slice和channel, 并且返回的不是指针, 应当用new获得特定的指针!!!
	*/
}

// golang初始化结构体数组
func main18() {
	// 最近组里新项目要求用go来写, 没办法只能边看文档边写代码, 今天遇到郁闷的问题, 查了好久最终发现居然是一个标点符号的导致的, 遂纪录之!!!
	// 刚刚给一个接口写单元测试时想初始化一个结构体数组, 然后遍历该数组并建立http.Request进行测试, 结果一直报错, 最后。。。才发现golang结构体初始化的正确姿势
	// 问题就出在大括号中最后一个元素的后面必须要加逗号，golang对语法的严格要求算是领教了。
	/*
	requests := []AuthRequest {
		AuthRequest {
			"plain",
			"xl_test@xunlei.net",
			"123456",
			"smtp",
			3,
			"192.168.34.104",
			"client.example.com",
		},
		AuthRequest{
			"plain",
			"xl_test@xunlei.net",
			"123456",
			"pop3",
			3,
			"192.168.34.104",
			"client.example.com",
		},
	}
	*/
}

func main19() {
	orderDt := "2016-04-27"

	cnt := strings.Count(orderDt,"-")
	if "" != orderDt && 2 == cnt {
		tempStr := strings.Split(orderDt, "-")

		if 3 == len(tempStr) {
			subStr1 := tempStr[0]
			subStr2 := tempStr[1]
			subStr3 := tempStr[2]

			year, errYear := strconv.Atoi(subStr1)
			if errYear != nil{

			}

			month, errMonth := strconv.Atoi(subStr2)
			if errMonth != nil{
			}

			day, errDay := strconv.Atoi(subStr3)
			if errDay != nil{
			}

			fmt.Println(subStr1)
			fmt.Println(subStr2)
			fmt.Println(subStr3)

			fmt.Println(year)
			fmt.Println(month)
			fmt.Println(day)
		} else {
			fmt.Printf("tempStr is empty!!!")
		}
	}
}

// GO数值和字符串的相互转换
func main20() {
	a := "123456"
	b, error := strconv.Atoi(a)

	if error != nil{
		fmt.Println("字符串转换成整数失败")
	}

	b = b + 1
	fmt.Println(b)

	var c int = 1234
	d := strconv.Itoa(c)   //数字变成字符串
	d = d + "sdfs"
	fmt.Println(d)
	var e interface {}
	e = 10
	switch v := e.(type){
	case int:
		fmt.Println("整型",v)
		break;
	case string:
		fmt.Println("字符串",v)
		break;

	}
}

func main21() {
	upTimeTmp := ""
	creTimeStr := strings.Split(upTimeTmp,".")

	createdTime		:= creTimeStr[0]
	//println(createdTime)
	fmt.Println(createdTime)
}