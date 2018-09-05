package main

import (
	"fmt"
	"encoding/json"
	"time"
)

func main() {
	// main1()
	// main2()
	// main3()
	// main4()
	main5()
}

// golang字符串json格式解析
// 最近在用golang写关于微信方面的东西, 首先遇到的就是将字符串转换成golang的json格式, 利用corpid和corpsecret返回的是一个json格式的字符串, 其格式如下:
// {"access_token":"uAUS6o5g-9rFWjYt39LYa7TKqiMVsIfCGPEN4IZzdAk5-T-ryVhL7xb8kYciuU_m","expires_in":7200}

// 我们可以把它转换成一个map[string]interface{}类型的数据, 相关代码如下:

var dat map[string]interface{}

// 注意相应变量首字母的大小写(首字母小写不可见, 大写可见, 具体查看golang的变量相关的内容), 将JSON绑定到结构体, 结构体的字段一定要大写, 否则不能绑定数据。
// 我们还可以定义一个结构体, 将数据转换成对应的结构体对象, 再获取相应的数据, 定义一个weixintoken结构体:
type weixintoken struct {
    Tokens string `json:"access_token"`
    Expires int `json:"expires_in"`
}

func main1() {
	str := "{\"access_token\":\"uAUS6o5g-9rFWjYt39LYa7TKqiMVsIfCGPEN4IZzdAk5-T-ryVhL7xb8kYciuU_m\",\"expires_in\":7200}"
	
	if err := json.Unmarshal([]byte(str), &dat); err == nil {
		fmt.Println(dat)
		fmt.Println(dat["expires_in"])
	} else {
		fmt.Println(err)
	}
	
	ret:="{\"access_token\":\"uAUS6o5g-9rFWjYt39LYa7TKqiMVsIfCGPEN4IZzdAk5-T-ryVhL7xb8kYciuU_m\",\"expires_in\":7200}"
	var config weixintoken
    if err := json.Unmarshal([]byte(ret), &config); err == nil {
        fmt.Println(config)
        fmt.Println(config.Tokens)
    }else {
        fmt.Println(err)
    }
}

// 如何使用Go语言自带的库把对象转换为JSON格式, 并在channel中进行传输后, 并把JSON格式的信息转换回对象?
// Go语言的JSON库			----->					Go语言自带的JSON转换库为 encoding/json
// 			(1).其中把对象转换为JSON的方法(函数)为json.Marshal(), 其函数原型:					func Marshal(v  interface{}) ([]byte, error)
//							也就是说, 这个函数接收任意类型的数据v, 并转换为字节数组类型, 返回值就是我们想要的JSON数据和一个错误代码, 当转换成功的时候, 这个错误代码为nil;
//					在进行对象转换为JSON的过程中, 会遵循如下几条规则:
//							a.布尔型转换为JSON后仍是布尔型, 如true	->	true
//							b.浮点型和整数型转换后为JSON里面的常规数字, 如 1.23 -> 1.23
//							c.字符串将以UTF-8编码转化输出为Unicode字符集的字符串, 特殊字符比如<将会被转义为\u003c
//							d.数组和切片被转换为JSON 里面的数组，[]byte类会被转换为base64编码后的字符串，slice的零值被转换为null
//							e.结构体会转化为JSON对象，并且只有结构体里边以大写字母开头的可被导出的字段才会被转化输出，而这些可导出的字段会作为JSON对象的字符串索引
//							f.转化一个map 类型的数据结构时，该数据的类型必须是 map[string]T（T 可以是encoding/json 包支持的任意数据类型)
//
//			(2).把JSON转换回对象的方法(函数)为json.Unmarshal(), 其函数原型为:		func Unmarshal(data [] byte, v interface{}) error
//							a.这个函数会把传入的 data 作为一个JSON来进行解析, 解析后的数据存储在参数 v 中
//							b.这个参数 v 也是任意类型的参数(但一定是一个类型的指针), 原因是我们在是以此函数进行JSON 解析的时候，这个函数不知道这个传入参数的具体类型, 所以它需要接收所有的类型。
//							c.那么，在进行解析的时候，如果JSON 和 对象的结构不对口会发生什么呢，这就需要解析函数json.Unmarshal()遵循以下规则
//										json.Unmarshal()函数会根据一个约定的顺序查找目标结构中的字段, 如果找到一个即发生匹配。
//					那什么是找到了呢?
//							关于“找到了”又有如下的规则:		假设一个JSON对象有个名为"Foo"的索引, 要将"Foo"所对应的值填充到目标结构体的目标字段上, json.Unmarshal()将会遵循如下顺序进行查找匹配
//																	一个包含Foo 标签的字段, 一个名为Foo 的字段, 一个名为Foo 或者Foo 或者除了首字母其他字母不区分大小写的名为Foo 的字段。 这些字段在类型声明中必须都是以大写字母开头、可被导出的字段。
//										注意：如果JSON中的字段在Go目标类型中不存在，json.Unmarshal() 函数在解码过程中会丢弃该字段。
//					当JSON 的结构是未知的时候，会遵循如下规则：
//								JSON中的布尔值将会转换为Go中的bool类型
//										数值会被转换为Go中的float64类型
//										字符串转换后还是string类型
//										JSON数组会转换为[]interface{} 类型
//										JSON对象会转换为map[string]interface{}类型
//										null值会转换为nil
//
//		在Go的标准库encoding/json包中，允许使用map[string]interface{}和[]interface{} 类型的值来分别存放未知结构的JSON对象或数组

// 假设我们有如下一个类(结构体)student及其一个实例对象st
type Student struct {
	Name		string
	Age			int
	Guake		bool
	Classes		[]string
	Price		float32
}

func main2() {
/*
	st := &Student {
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}
*/
	
	// 现在我们需要把这个类的一个对象转换为JSON格式, 并且传输给远方的朋友, 那么我们就可以这么做:
	//b, err := json.Marshal(st)
	
	// 这样就转换好了, 是不是很简单! 转换回来就更简单了, 比如我们有一个新的student对象, 就叫stb, 那么我们可以这样转换回来:
	// stb := &Student{}
	//err = json.Unmarshal([]byte(strData), &stb)
}

// golang中struct转json后键名首字母大小写问题解决:    主要介绍一下struct转json后键名首字母大小写的问题
//  结构体里的字段首字母必须大写, 否则无法正常解析
type Person struct {
	Name	string		// Name字段首字母大写
	age		int			// age字段首字母小写
}

// 从上面代码可以看出如果结构体的字段首字母小写, 该字段将无法正常解析
func main3() {
	person := Person{"小黑",18}
	if result,err := json.Marshal(&person);err==nil {
		//fmt.Println(string(result))
		fmt.Println(string(result))
	}
}

// 如果我们想让struct转json后的首字母小写，我们可以通过字段的tag指定

// 未指定tag
type Person1 struct {
	Name	string
	Age		int               
}

func main4() {
	person:=Person1{"小黑",18}
	if result,err:=json.Marshal(&person);err==nil{
		fmt.Println(string(result))
	}
}

// 指定字段的tag, 实现json字符串的首字母小写
type Person2 struct {
	Name	string	`json:"name"`
	Age		int		`json:"age"`
	Time	int64	`json:"-"`			// 直接忽略字段
}

func main5(){
	person:=Person2{"小x",18, time.Now().Unix()}
	if result,err:=json.Marshal(&person);err==nil{
		fmt.Println(string(result))
	}
}
