package main

import (
	"fmt"
	"strconv"
	"sort"
	"reflect"
)

func main() {
	// main1()
	// main2()
	// main3()
	// main4()
	// main5()
	// main6()
	// main7()
	// main8()
	// main9()
	// main10()
	// main11()		// 专门用于go代码中的int元素类型的数组去重!!!
	main12()
}

// go语言基础 遍历数组 range			---->		我们可以通过关键字range来遍历数组中的值
func main1() {
	/*
		遍历数组：依次获取数组中的数据
		range 数组名:
			index,  value
	*/
	arr := [...]int{6,2,4,9,8,3}

	// 1.遍历方式一
	for i:= 0;i<len(arr);i++{
		fmt.Print(arr[i],"\t")
	}
	fmt.Println()

	// 2.range遍历数组
	sum := 0
	for _, value := range arr{
		//fmt.Println(index, "\t",value)
		fmt.Println(value)
		sum += value
	}
	fmt.Println(sum)
}

func main2() {
	spEntIdArr := [...]int{6,2,4,9,8,3}
	spEntSignupIdArr := [...]int{6,9,8}
	spEntIdPpArr := make([]interface{}, 0)			// 2, 4, 3

	var bInSignupIdArr int = 0
	for _, value1 := range spEntIdArr {
		for _, value2 := range spEntSignupIdArr {
			if value1 == value2 {
				bInSignupIdArr = 1
				break;
			}
		}

		if 1 != bInSignupIdArr {
			spEntIdPpArr = append(spEntIdPpArr, value1)
		} else {
			bInSignupIdArr = 0
		}
	}

}

type stru1 struct {
	id int
	ageBig string
	prcBig string
}

type stru2 struct {
	age string
	prc string
}

func main3() {
	bigArr		:= [...]int{6,2,4,9,8,3}
	smallIdArr	:= [...]int{6,9,8}

	tstStru1 := make([]stru1, 0, 100)
	tstStru2 := make(map[int]stru2)

	for _, value1 := range bigArr{
		tstStru1 = append(tstStru1, stru1{
			id:	value1,
			ageBig: "",
			prcBig: "",
		})
	}

	var tempStru2 stru2
	for _, value2 := range smallIdArr {
		tmpStr1 := "age"+strconv.Itoa(value2)
		tmpStr2 := "prc"+strconv.Itoa(value2)

		tempStru2.age = tmpStr1
		tempStru2.prc = tmpStr2

		tstStru2[value2] =tempStru2
	}

	// 采用 range 获取数组项不能修改数组中结构体的值：
	for _, dayRecrEntInfo := range tstStru1{
		dayRecrEntInfo.ageBig = tstStru2[dayRecrEntInfo.id].age
		dayRecrEntInfo.prcBig = tstStru2[dayRecrEntInfo.id].prc
	}

	// 采用 range 获取的下标值，然后用下标方式引用的数组项也可以直接修改：
	for idx, _ := range tstStru1 {
		idTemp := tstStru1[idx].id

		tstStru1[idx].ageBig = tstStru2[idTemp].age
		tstStru1[idx].prcBig = tstStru2[idTemp].prc
	}

	fmt.Println(tstStru1)
}

// golang数组去重, 去空
func main4() {
	// removeDuplicatesAndEmpty
}

// golang数组去重, 去空
func removeDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)

	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}

		ret = append(ret, a[i])
	}

	return
}

// golang移除数组中重复的元素
// 方式一:			这种发放适用于string, int, float等切片, 会对切片中的元素进行排序
func SliceRemoveDuplicates(slice []string) []string {

	sort.Strings(slice)
	i := 0
	var j int

	for {
		if i >= len(slice)-1 {
			break
		}

		for j = i + 1; j < len(slice) && slice[i] == slice[j]; j++ {
		}

		slice= append(slice[:i+1], slice[j:]...)
		i++
	}

	return  slice
}

// 方法二
func RemoveDuplicate(list *[]int) []int {
	var x []int = []int{}

	for _, i := range *list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}

				if k == len(x)-1 {
					x = append(x,i)
				}
			}
		}
	}

	return x
}

// Golang去除slice和list的重复元素, 非常有用的功能
// Golang中, 利用反射和interface就可以做到, 不废话看代码
func main5() {
	b := []string{"a", "b", "c", "c", "e", "f", "a", "g", "b", "b", "c"}
	sort.Strings(b)
	fmt.Println(Duplicate(b))


	c := []int{1, 1, 2, 4, 6, 7, 8, 4, 3, 2, 5, 6, 6, 8, 9, 9, 9, 9, 9, 1,2,3,4}
	sort.Ints(c)
	//fmt.Println(Duplicate(c))
	d := Duplicate(c)
	fmt.Println(d)

	//spIdArrTmp := make([]interface{}, 0)
	//spIdArrTmp := make([]int64, 0)
	spIdArrTmp := []int64{1003, 10085, 1010, 1008, 10142, 10173, 10113, 9999, 9999, 10218, 9999, 10690, 10693, 9999, 10208, 10218}
	//sort.Ints(spIdArrTmp)
	e := Duplicate(spIdArrTmp)
	fmt.Println(e)
}

func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

type SpListInfo1 struct {
	SpId               int64  //服务商ID
	Distance           float64 //位置数据
}

func main6() {
	SpDistanceMap	:= make(map[int64]*SpListInfo1, 0)			// key: SpID, value: 最短距离

	//
	for _, v := range SpDistanceMap {
		v.Distance = 400000
	}
}

func main7() {
	var tmpSpIdArr SpIdArr =  	[]int64{1003, 10085, 1010, 1008, 10142, 10173, 10113, 9999, 9999, 10218, 9999, 10690, 10693, 9999, 10208, 10218}

	fmt.Println(tmpSpIdArr)
	sort.Sort(tmpSpIdArr)
	fmt.Println(tmpSpIdArr)

	ret := Duplicate(tmpSpIdArr)
	fmt.Println(ret)
}

type SpIdArr  []int64
func (p SpIdArr) Len() int { return len(p) }
func (p SpIdArr) Less(i, j int) bool { return p[i]<p[j]}
func (p SpIdArr) Swap(i, j int)      { p[i], p[j] = p[j], p[i]}

func findInInt64Array(array []interface{}, i int64) bool {
	for _, v := range array {
		if v.(int64) == i {
			return true
		}
	}

	return false
}

func main8() {
	// type StdEntIdArr []int64
	// var stdEntIdArr StdEntIdArr =  	[]int64{1003, 10085, 1010, 1008, 10142, 10173, 10113, 9999, 9999, 10218, 9999, 10690, 10693, 9999, 10208, 10218}
	entIdArr := [...]int64{1003, 10085, 1010, 1008, 10142, 10173, 10113, 9999, 9999, 10218, 9999, 10690, 10693, 9999, 10208, 10218}

	var entId int64 = 1008

	fmt.Println(entIdArr)

	stdEntIdArr := make([]interface{}, 0)

	for _, row := range entIdArr {
		stdEntID	:= row
		stdEntIdArr	= append(stdEntIdArr, stdEntID)
	}

	ret := findInInt64Array(stdEntIdArr, entId)

	fmt.Println(ret)
}

type TestArr struct {
	tmep		int
	even		int		// 1:是偶数   0: 不是偶数
	remark		string
}

func main9() {
	//arr := [...]int{6,2,4,9,8,3}
	arr := [...]int{6,2,4,9,8,3}
	fmt.Println(arr)

	// index := -1
	index := 0
	arrLen := len(arr)

	brr := append(arr[:index], arr[index+arrLen:]...)		// 切片删除元素
	fmt.Println(brr)

	// arr = append(arr[:index], arr[index+arrLen:]...)		// 切片删除元素
	// fmt.Println(arr)

	var dataArr []TestArr

	var tmpData TestArr
	for i := 0; i < 6; i++ {
		tmpData.tmep = 10+i

		if 0 == (i % 2) {
			tmpData.even = 1
		} else {
			tmpData.even = 0
		}

		dataArr = append(dataArr, tmpData)
	}

	fmt.Println(dataArr)

	index1 := 0
	dataArrLen := len(dataArr)

	brr1 := append(dataArr[:index1], dataArr[index1+dataArrLen:]...)		// 切片删除元素
	fmt.Println(brr1)

	// dataArr = append(dataArr[:index1], dataArr[index1+dataArrLen:]...)		// 切片删除元素
	fmt.Println(dataArr)

	// 删除为奇数的切片
	index2 := 0
	for _, value := range dataArr {
		if 1 == value.even {
			dataArr = append(dataArr[:index2], dataArr[index2+1:]...)		// 切片删除元素
		}
	}
	fmt.Println(dataArr)

	//for idx, _	:=  range dataArr{
	//	if 1 == value.even {
	//		append(s[:i], s[i+1:]...)
	//	}
	//}
}

// 使用for循环对golang中结构体数组取值进行修改时, 需要注意的问题
func main10() {
	type a1 struct {
		key1 string
		key2 string
		key3 string
	}

	testData := []a1 {
		a1{"1","2", "3"},
		a1{"4","5", "6"},
	}

	// 上面的代码定义了一个结构体，声明了一个数组。

	// 采用循环变量可以修改数组中结构体的取值:		输出：[{1 2 999} {4 5 999}]
	for i := 0; i < len(testData); i++ {
		testData[i].key3 = "999"
	}
	fmt.Printf("%v", testData)

	// 采用range获取的下标值, 然后用下标方式引用的数组项也可以直接修改：		输出：[{1 2 999} {4 5 999}]
	/*
	for idx, _ := range testData {
		testData[idx].key3 = "999"
	}
	fmt.Printf("%v", testData)
	*/

	// 采用 range 获取数组项不能修改数组中结构体的值：							输出：[{1 2 3} {4 5 6}]
	/*
	for _, item := range testData {
		item.key3 = "999"
	}
	fmt.Printf("%v", testData)
	*/
}

// sp_ent获取列表数据的时候, updateId arr需要去重!!!
func main11() {
	var spIdArrTmp SpIdArrType
	var idArrTmp IdArrType
	SpIdArr1 := make([]interface{}, 0)
	SpIdArr2 := make([]interface{}, 0)
	IdArr3 := make([]interface{}, 0)

	//updateIdArr := make([]interface{}, 0)
	updateIdArr := make([]int64, 0)
	//updateIdArr = append(updateIdArr, 102190,102190,102190,0,0)
	// updateIdArr := append(updateIdArr, updatedBy)
	// stReq.NameIDs = append(stReq.NameIDs, 100000, 100001)

	updateIdArr = append(updateIdArr, 102120,101197,102199,475,88,88,475,0,3232,101197)


	spIdArrTmp = updateIdArr
	idArrTmp = updateIdArr
	sort.Sort(spIdArrTmp)
	sort.Sort(idArrTmp)

	SpIdArr1 = Duplicate(updateIdArr)						// 对sp_id数组除重一把!!!
	SpIdArr2 = Duplicate(spIdArrTmp)
	IdArr3 = Duplicate(idArrTmp)

	fmt.Print(SpIdArr1)
	fmt.Printf("\n")
	fmt.Print(SpIdArr2)
	fmt.Printf("\n")
	fmt.Print(IdArr3)
}

// Go语言string，int，int64 ,float之间类型转换方法, 绍的Go语言string，int，int64 ,float之间类型转换方法
func main12() {
	// 1.int转string
	//s := strconv.Itoa(i)  // 等价于s := strconv.FormatInt(int64(i), 10)

	// 2.int64转string
	//i := int64(123)		 					// 第二个参数为基数，可选2~36
	//s := strconv.FormatInt(i, 10)		// 注：对于无符号整形，可以使用FormatUint(i uint64, base int)

	// 3.string转int
	//i, err := strconv.Atoi(s)

	// 4.string转int64, 第二个参数为基数（2~36），第三个参数位大小表示期望转换的结果类型，其值可以为0, 8, 16, 32和64，分别对应 int, int8, int16, int32和int64
	//i, err := strconv.ParseInt(s, 10, 64)

	// 5.float相关
	// float转string：
	//v := 3.1415926535
	//s1 := strconv.FormatFloat(v, 'E', -1, 32)//float32s2 := strconv.FormatFloat(v, 'E', -1, 64)// 函数原型及参数含义具体可查看：https://golang.org/pkg/strconv/#FormatFloat

	// string转float：
	//s := "3.1415926535"
	//v1, err := strconv.ParseFloat(v, 32)
	//v2, err := strconv.ParseFloat(v, 64)

	// go语言string、int、int64互相转换
	//string到int
	//int,err:=strconv.Atoi(string)
	//string到int64
	//int64, err := strconv.ParseInt(string, 10, 64)
	//int到string
	//string:=strconv.Itoa(int)
	//int64到string
	//string:=strconv.FormatInt(int64,10)
	//string到float32(float64)
	//float,err := strconv.ParseFloat(string,32/64)
	//float到string
	//string := strconv.FormatFloat(float32, 'E', -1, 32)
	//string := strconv.FormatFloat(float64, 'E', -1, 64)
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
}


type SpIdArrType []int64

func (p SpIdArrType) Len() int           { return len(p) }
func (p SpIdArrType) Less(i, j int) bool { return p[i] < p[j] }
func (p SpIdArrType) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type IdArrType []int64

func (p IdArrType) Len() int           { return len(p) }
func (p IdArrType) Less(i, j int) bool { return p[i] < p[j] }
func (p IdArrType) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
