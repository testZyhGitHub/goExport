package main

import (
	"fmt"
	"reflect"
)

// go语言map的复杂用法
// 对于简的map:				例如, map[string]string还是很好掌握的!!!

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

// 企业素材信息 zyh
type MDetailInfo struct {
	EntMaterialDetailID int64  // 企业素材明细ID
	PicPriority         int    // 图片优先级标记
	PicTip              string // 图片标题
	PicUrl              string // 图片地址
	OperType            int    // 1.新增 2 删除
}

func main() {
	// main1()
	// main2()
	// main4()
	// main5()
	// main6()
	// main7()
	main8()
}

func main1 () {
	//EntMaterialList map[int][]MDetailInfo	// 企业素材
	//EntMaterialList := make(map[int][]MDetailInfo)
	
	personDB := make(map[string][2]PersonInfo)
	
	// 初始化, 注意对数组的初始化
	personDB["test1"] = [2]PersonInfo{{"12345", "Tom", "aaa"}, {"12346", "Xym", "bbb"}}
	
	// 元素查找, 这是通用的使用方法
	v, ok := personDB["test1"]
	if !ok {
		fmt.Println(" 没有找到信息")
		return
	}
	
	// 打印出全部值和各个分值
	fmt.Printf("v=%v v[0]=%v v[1]=%v \n", v, v[0], v[1])
	
	fmt.Printf("\n")
	
	// 用range方便的得到其中的值
	/*
		for i, v := range v {
			fmt.Println(i, v, v.ID, v.Address, v.Name)
			// 可以做其他的处理
		}
	*/
	
	fmt.Printf("print map info......\n\n")
	for i, row := range v {
		fmt.Println(i, row, row.ID, row.Address, row.Name)
		// 可以做其他的处理
    }
}

func main2() {
	EntMaterialList := make(map[int][]MDetailInfo)
	
	// "EntMaterialList":{"1":[],"2":[{"PicUrl":"/web/ent/material/5541","EntMaterialDetailID":0,"OperType":1}],"3":[],"4":[]},
	EntMaterialList[0] = []MDetailInfo{{1, 0, "test1", "/web/ent/material/5541", 1}, {2, 0, "test2", "/web/ent/material/5542", 1}, {3, 0, "test3", "/web/ent/material/5541", 1}}	// banner新增3张图片
	EntMaterialList[1] = []MDetailInfo{{4, 0, "test4", "/web/ent/material/5545", 2}}					// 工资条删除1张图片
	EntMaterialList[2] = []MDetailInfo{{5, 0, "test5", "/web/ent/material/5546", 1}}					// 企业环境新增1张图片
	EntMaterialList[3] = []MDetailInfo{{6, 0, "test6", "/web/ent/material/5547", 1},{7, 0, "test7", "/web/ent/material/5548", 2}}
	EntMaterialList[4] = []MDetailInfo{}	// 食堂无变动
		
	// 元素查找, 这是通用的使用方法
	v, ok := EntMaterialList[0]	
	if !ok {
		fmt.Println(" 没有找到信息")
		return
	}
	
	// 打印出全部值和各个分值
	fmt.Printf("v=%v v[0]=%v v[1]=%v v[2]=%v \n", v, v[0], v[1], v[2])
	fmt.Printf("\n\n\n")
	
	for i, row := range v {
		fmt.Println(i, row, row.EntMaterialDetailID, row.PicPriority, row.PicTip, row.PicUrl, row.OperType)
		// 可以做其他的处理
    }
	
	fmt.Printf("\n\n\nprint map info......\n")
	for in, row := range EntMaterialList[0] {
		fmt.Println(in, row, row.EntMaterialDetailID, row.PicPriority, row.PicTip, row.PicUrl, row.OperType)
		// 可以做其他的处理
    }
	
	fmt.Printf("\n\n\nPrint All Map Info......\n")
	for in, row := range EntMaterialList {
		for _, col := range row {
			fmt.Println(in, col, col.EntMaterialDetailID, col.PicPriority, col.PicTip, col.PicUrl, col.OperType)
		}
		fmt.Printf("\n")
	}
}

// golang中map声明及初始化
func main3() {
	// map的声明
	// var m1 map[string]int		// map[key]value, key必须支持==(避免使用浮点型), value不做规范
    
	// map的初始化
	
	var m1 map[string]int=map[string]int{"key":0} 	// 方式1
	m1["key"] = 2
	
	m2:=make(map[string]int)						// 方式2
	
	m2["hello"] = 22
	
	// map声明后初始化前, 可进行查找、删除、len和range操作, 并不会报错
	// map声明后不能进行赋值, 只有初始化后才能进行赋值操作
}

// Go语言嵌套Map类型，类似PHP的二维或者多维数组
func main4() {
	/*
		关键点:			interface{}可以代表任意类型
		
		原理知识点:		interface{} 就是一个空接口, 所有类型都实现了这个接口, 所以它可以代表所有类型
	*/
	
	var tags map[string]interface{}
	tags = make(map[string]interface{})
	
	var tagsLocal map[string]string
	tagsLocal = make(map[string]string)
	
	var tagsTest map[string]string
	tagsTest = make(map[string]string)
	
	var tagsProduction map[string]string
	tagsProduction = make(map[string]string)
	
	tagsLocal["dev.www.9178.us"] = "dev.www.9178.us"
	tagsLocal["dev.static.9178.us"] = "dev.static.9178.us"

	tagsTest["dev.www.9178.us"] = "www.ninja911.com"
	tagsTest["dev.static.9178.us"] = "static.ninja911.com"
	
	tagsProduction["dev.www.9178.us"] = "ipx.www.ninja911.com"
	tagsProduction["dev.static.9178.us"] = "ipx.static.ninja911.com"
	
	tags["local"] = tagsLocal
	//tags["test"] = tagsLocal
	tags["test"] = tagsTest
	//tags["production"] = tagsLocal
	tags["production"] = tagsProduction
	
	fmt.Println(tags)
}

// go语言之结构体数组转为string字符串
// 转换顺序：先将struct结构体转为map,再将map转为string字符串:			struct --> map --> string
type demo struct {  
    Id                    string  
    Name                  string  
}  

func main5() {
	/*
	demos := [{"Id":"1","Name":"zs"},{"Id":"2","Name":"ls"},{"Id":"3","Name":"ww"}]  
  
	for _, v := range demos {  
        tmpdata := Struct2Map(v)  
        str, err := json.Marshal(tmpdata)  
        if err != nil {  
            fmt.Println(err)  
        }         
		fmt.println(string(str))  
	} */ 
}

// https://www.cnblogs.com/liang1101/p/6741262.html

//结构体转为map  
func Struct2Map(obj interface{}) map[string]interface{} {  
    t := reflect.TypeOf(obj)  
    v := reflect.ValueOf(obj)  
  
    var data = make(map[string]interface{})  
    for i := 0; i < t.NumField(); i++ {  
        data[t.Field(i).Name] = v.Field(i).Interface()  
    }  
    return data  
}  

// GO语言映射(Map)用法分析
// 本文实例讲述了GO语言映射(Map)用法, 分享给大家供大家参考, 具体如下:
// 映射是一种内置的数据结构, 用来保存键值对的无序集合。

// 1.映射的创建
func main6() {

	/*
		make(map[KeyType]ValueType, initialCapacity)
		make(map[KeyType]ValueType)
		map[KeyType]ValueType{}
		map [KeyType ] ValueType { key1 : value1, key2: value2, ... , keyN : valueN}	
	*/
	
	/*
		如下, 用4种方式分别创建数组, 其中第一种和第二种的区别在于, 有没有指定初始容量,
				不过使用的时候则无需在意这些,
				因为map的本质决定了, 一旦容量不够, 它会自动扩容:
	*/
	map1 := make(map[string]string, 5)
    map2 := make(map[string]string)
    map3 := map[string]string{}
    map4 := map[string]string{"a": "1", "b": "2", "c": "3"}
    fmt.Println(map1, map2, map3, map4)	
}

// 2.映射的填充和遍历
func main7() {
	/*
		数组的填充使用map[key] = value的方式, 遍历映射的时候, 每一项都返回2个值, 键和值。
		结果如下:
						a->1    b->2    c->3
	*/
	map1 := make(map[string]string)
	map1["a"] = "1"
	map1["b"] = "2"
	map1["c"] = "3"
		
	for key, value := range map1 {
		fmt.Printf("%s->%-10s", key, value)
	}
}

// 3.映射的查找、修改和删除
func main8() {
map4 := map[string]string{"a": "1", "b": "2", "c": "3"}
	val, exist := map4["a"]
	val2, exist2 := map4["d"]
	fmt.Printf("%v,%v\n", exist, val)
	fmt.Printf("%v,%v\n", exist2, val2)

	map4["a"] = "8" //修改映射和添加映射没什么区别
	fmt.Printf("%v\n", map4)
	
	/*
			map指定key取对应的value时, 可以指定返回两个值,
					第一个是对应的value, 第二个是一个bool, 表示是否有值!
			如上, "a"肯定有值, "b"肯定没值。
			
			修改映射和添加映射的操作没什么区别, 若指定的键不存在则创建, 否则, 修改之。
			删除则是使用go的内置函数delete, 输出如下:
					true,1
					false,
					map[a:8 b:2 c:3]
					删除b：
					map[a:8 c:3]
	*/
	
	fmt.Println("删除b：")
	delete(map4, "b")
	fmt.Printf("%v", map4)
}

// 4.

