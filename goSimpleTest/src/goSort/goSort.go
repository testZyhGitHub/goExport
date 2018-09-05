package main

import (
	"fmt"
	"sort"
	"reflect"
	"time"
	"strconv"
	"strings"
)

func main () {
	// main1()
	// main2()
	// main3()
	// main4()
	// main5()
	// main6()
	// main7()
	main8()
}

/*
			type Interface interface {
				Len() int				// Len 为集合内元素的总数
				Less(i, j int) bool		// 如果index为i的元素小于index为j的元素, 则返回true, 否则返回false
				Swap(i, j int)			// Swap交换索引为i和j的元素
			}
*/

// golang中sort包用法???    				golang中也实现了排序算法的包sort包
// sort包中实现了３种基本的排序算法:	插入排序, 快排和堆排序;  和其他语言中一样, 这三种方式都是不公开的, 他们只在sort包内部使用．
//												所以用户在使用sort包进行排序时无需考虑使用那种排序方式, sort.Interface定义的三个方法:
//			(1).获取数据集合长度的Len()方法
//			(2).比较两个元素大小的Less()方法
//			(3).交换两个元素位置的Swap()方法
//		就可以顺利对数据集合进行排序!!!
//
//			sort包会根据实际数据自动选择高效的排序算法!!!
//			任何实现了sort.Interface的类型(一般为集合), 均可使用该包中的方法进行排序!!!
//			这些方法要求集合内列出元素的索引为整数!!!
//
//			func Float64s(a []float64)					// Float64s将类型为float64的slice a以升序方式进行排序
//			func Float64sAreSorted(a []float64) bool	// 判定是否已经进行排序func Ints(a []int)
//			func Ints(a []int)							// Ints以升序排列int切片。
//			func IntsAreSorted(a []int) bool			// IntsAreSorted判断int切片是否已经按升序排列!
//			func IsSorted(data Interface) bool			// IsSorted判断数据是否已经排序, 包括各种可sort的数据类型的判断!
//			func Strings(a []string)					// Strings以升序排列 string 切片。
//			func StringsAreSorted(a []string) bool		// StringsAreSorted判断string切片是否已经按升序排列。
//

// 定义interface{}, 并实现sort.Interface接口的三个方法
type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}

func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}

// https://studygolang.com/articles/3360
// https://www.jianshu.com/p/6e52bad56e06
func main1() {
	a := IntSlice{1, 3, 5, 7, 2}
	b := []float64{1.1, 2.3, 5.3, 3.4}
	c := []int{1, 3, 5, 4, 2}

	fmt.Println(sort.IsSorted(a))		// false, IsSorted判断数据是否已经排序, 包括各种可sort的数据类型的判断!

	if !sort.IsSorted(a) {
		sort.Sort(a)
	}

	if !sort.Float64sAreSorted(b) {		// 判定是否已经进行排序
		sort.Float64s(b)
	}

	if !sort.IntsAreSorted(c) {			// 判断int切片是否已经按升序排列
		sort.Ints(c)
	}

	fmt.Println(a)//[1 2 3 5 7]
	fmt.Println(b)//[1.1 2.3 3.4 5.3]
	fmt.Println(c)// [1 2 3 4 5]
}

func main2() {
	// entPi := (100 / 500 / 5) * 4

	entPi := (5369 / 500 / 5) * 4

	if entPi >= 5 {
		entPi = 5
	}

	//var entScale int = 0
	var entScale int = 2300

	var scale float64
	if entScale >= 7000 {
		scale = 5
	} else if entScale >= 2000 && entScale < 7000 {
		scale = 3
	} else {
		scale = 1
	}
	var piAllEnt float64
	piAllEnt = (float64(entPi) + scale) / 15

	fmt.Println("type:", reflect.TypeOf(entPi))				// golang打印变量类型
	fmt.Println(entPi)
	fmt.Println(piAllEnt)
}

// 权重是什么???
func main3() {
	// https://baike.baidu.com/item/%E6%9D%83%E9%87%8D/10245966?fr=aladdin
	// https://baike.baidu.com/item/%E6%9D%83%E9%87%8D%E5%80%BC/10602484
	// https://zhidao.baidu.com/question/1987826676297167387.html
	// http://www.ccutu.com/wenwen/answer7201.html
}

type SpAddrInfo struct {
	SpAddrId		int64		// 服务商地址主键，自增长
	SpAddr			string 		// 服务商地址
	LongLat			string		// 经纬度
}

// 同一个sp多个地址数组
type SpAddrInfoArr []SpAddrInfo

func main4() {
	var OrderDt string = time.Now().Format("2006-01-02")
	//var spEntIdStr string = "1000, 1001, 1002, 1003, 1004, 1005, 1006, 1007, 1008"
	var spEntIdStr string = "1000, 1071, 1088, 1326, 1327, 1087"

	sql := "select * from rcrt_b_order a "
	sql = sql + "where exists ("
	sql = sql + "select 1 from "
	sql = sql + "(SELECT sp_id,sp_ent_id,max(case when updated_tm is null then created_tm else updated_tm end) as updated_tm FROM rcrt_b_order where order_dt = '%s' group by sp_id,sp_ent_id) b "
	sql = fmt.Sprintf(sql, OrderDt)
	sql = sql + "where a.sp_id = b.sp_id and a.sp_ent_id = b.sp_ent_id and (a.updated_tm = b.updated_tm or a.created_tm = b.updated_tm) and sp_ent_id in (%s)"
	sql = fmt.Sprintf(sql, spEntIdStr)
	sql = sql + ")"
	sql = sql + "and a.order_dt = '%s' "
	sql = fmt.Sprintf(sql, OrderDt)
	sql = sql + "and is_deleted = 0 and is_enabled = 1 order by updated_tm desc"

	fmt.Println(sql)
}

func main5() {
	//spEntIdNoSignArr := make([]interface{}, 0)
	spEntIdNoSignArr := make([]int64, 10, 100)

	spEntIdNoSignArr[0] = 1000
	spEntIdNoSignArr[1] = 1071
	spEntIdNoSignArr[2] = 1088
	spEntIdNoSignArr[3] = 1326
	spEntIdNoSignArr[4] = 1327
	spEntIdNoSignArr[5] = 1087

	var addtoStr string = ""
	var itemStr string = ""

	for _, value := range spEntIdNoSignArr{
		if 0 != value {
			itemStr = strconv.FormatInt(value,10)
			//addtoStr = fmt.Sprintf(addtoStr,itemStr)
			addtoStr += itemStr + ","
		}
	}

	fmt.Println(addtoStr)

	i := strings.LastIndex(addtoStr, ",")
	addtoStr = addtoStr[: i]

	fmt.Println(addtoStr)
}

// 发单信息结构体
type RcrtOrderInfo struct {
	SpentID			int
	PrcRemark		string
	AgeRemark		string
}

type RcrtOrderRsp struct {
	RcrtOrderTotalNum	int	// 获取到的发单总个数
	RcrtOrderList		[]*RcrtOrderInfo
}

// 发单信息结构体
type RcrtOrderInfo1 struct {
	// SpentID			int64
	PrcRemark		string
	AgeRemark		string
}



func GetSpentBillInfoFrom_RcrtbOrder1() (*RcrtOrderRsp, error, int) {
	rcrtOrderInfos := new(RcrtOrderRsp)

	arr := [...]int{6,2,4,9,8,3}

	// 1.遍历方式一
	for i:= 0;i<len(arr);i++{
		rcrtOrderInfo := new(RcrtOrderInfo)
		rcrtOrderInfo.SpentID	= i
		rcrtOrderInfo.PrcRemark	= "Price: " + strconv.Itoa(i)
		rcrtOrderInfo.AgeRemark	= "Age: " + strconv.Itoa(i+10)

		rcrtOrderInfos.RcrtOrderList = append(rcrtOrderInfos.RcrtOrderList, rcrtOrderInfo)
	}

	rcrtOrderInfos.RcrtOrderTotalNum = len(arr)

	return rcrtOrderInfos, nil, 2
}

func GetSpentBillInfoFrom_RcrtbOrder2() (map[int]RcrtOrderInfo1, error, int) {
	//rcrtOrderMapInfos := new(RcrtOrderRsp)
	SpentIdRcrtOrderMapInfos := make(map[int]RcrtOrderInfo1)		// key: SpentID, value: RcrtOrderInfo

	arr := [...]int{6,2,4,9,8,3}

	var temp RcrtOrderInfo1
	// 1.遍历方式一
	for i:= 0;i<len(arr);i++{
		spentID := i
		temp.PrcRemark	= "Price: " + strconv.Itoa(i)
		temp.AgeRemark	= "Age: " + strconv.Itoa(i+10)

		SpentIdRcrtOrderMapInfos[spentID] = temp
		//rcrtOrderInfos.RcrtOrderList = append(rcrtOrderInfos.RcrtOrderList, rcrtOrderInfo)
	}

	return SpentIdRcrtOrderMapInfos, nil, 2
}

func main6() {
	rcrtOrderRsp,err,ret := GetSpentBillInfoFrom_RcrtbOrder1()

	goto RetCode
	if (err != nil) {
		fmt.Println("err != nil!!!")
	} else {
			if 0 == ret {
				fmt.Println("0 == ret")
			} else {
				len := rcrtOrderRsp.RcrtOrderTotalNum
				fmt.Println(len)

				for _, row := range rcrtOrderRsp.RcrtOrderList {
					fmt.Println(row.SpentID)
					fmt.Println(row.PrcRemark)
					fmt.Println(row.AgeRemark)
				}
			}
	}

RetCode:
	fmt.Println("test use goto loop!!!")
}

func main7() {
	rcrtOrderRsp,err,ret := GetSpentBillInfoFrom_RcrtbOrder2()

	if (err != nil) {
		fmt.Println("err != nil!!!")
	} else {
		if 0 == ret {
			fmt.Println("0 == ret")
		} else {

			for rol, row := range rcrtOrderRsp {
				fmt.Println(rol)
				fmt.Println(row.PrcRemark)
				fmt.Println(row.AgeRemark)
			}
		}
	}
}

// go语言十大排序算法总结
//		https://blog.csdn.net/guoer9973/article/details/51924715
//		https://blog.csdn.net/guoer9973/article/details/51933468
func main8() {

	// 快速排序实现
	/*
			快速排序由C. A. R. Hoare在1962年提出。
			它的基本思想是:
							通过一趟排序将要排序的数据分割成独立的两部分, 其中一部分的所有数据都比另外一部分的所有数据都要小,
									然后再按此方法对这两部分数据分别进行快速排序,
										整个排序过程可以递归进行,
										以此达到整个数据变成有序序列。

			算法步骤：
设要排序的数组是A[0]……A[N-1]，首先任意选取一个数据（通常选用数组的第一个数）作为关键数据，然后将所有比它小的数都放到它前面，所有比它大的数都放到它后面，这个过程称为一趟快速排序。值得注意的是，快速排序不是一种稳定的排序算法，也就是说，多个相同的值的相对位置也许会在算法结束时产生变动。
一趟快速排序的算法是：
1）设置两个变量i、j，排序开始的时候：i=0，j=N-1；
2）以第一个数组元素作为关键数据，赋值给key，即key=A[0]；
3）从j开始向前搜索，即由后开始向前搜索(j–)，找到第一个小于key的值A[j]，将A[j]和A[i]互换；
4）从i开始向后搜索，即由前开始向后搜索(i++)，找到第一个大于key的A[i]，将A[i]和A[j]互换；
5）重复第3、4步，直到i=j； (3,4步中，没找到符合条件的值，即3中A[j]不小于key,4中A[i]不大于key的时候改变j、i的值，使得j=j-1，i=i+1，直至找到为止。找到符合条件的值，进行交换的时候i， j指针位置不变。另外，i==j这一过程一定正好是i+或j-完成的时候，此时令循环结束）。

快速排序是不稳定的。最理想情况算法时间复杂度O(nlog2n)，最坏O(n ^2)。
个人总结：
快速排序其实是以数组第一数字作为中间值，开始排序，当这个值不是最佳中间值的时候，就会出现最坏的情况，当一次排序完成后，准备进入递归，递归传入的是slice，递归的退出条件是，当这个slice已经只能和自己比较了，也就是变为了中间值，slice为1。
	*/

	arry := []int{6, 1, 3, 5, 8, 4, 2, 0, 9, 7}
	learnsort := Sortor{name: "快速排序--从小到大--不稳定--nlog2n最坏n＊n---"}
	//learnsort.sort(arry)
	learnsort.sortDesc(arry)
	fmt.Println(learnsort.name, arry)
}

type Sortor struct {
	name string
}

type SortInterface interface {
	sort()
}

func (sorter Sortor) sort(arry []int) {
	if len(arry) <= 1 {
		return //递归终止条件，slice变为0为止。
	}
	mid := arry[0]
	i := 1 //arry[0]为中间值mid，所以要从1开始比较
	head, tail := 0, len(arry)-1
	for head < tail {
		if arry[i] > mid {
			arry[i], arry[tail] = arry[tail], arry[i] //go中快速交换变量值，不需要中间变量temp
			tail--
		} else {
			arry[i], arry[head] = arry[head], arry[i]
			head++
			i++
		}
	}
	arry[head] = mid
	sorter.sort(arry[:head]) //这里的head就是中间值。左边是比它小的，右边是比它大的，开始递归。
	sorter.sort(arry[head+1:])
}

func (sorter Sortor) sortDesc(arry []int) {
	if len(arry) <= 1 {
		return //递归终止条件，slice变为0为止。
	}
	mid := arry[0]
	i := 1 //arry[0]为中间值mid，所以要从1开始比较
	head, tail := 0, len(arry)-1
	for head < tail {
		if arry[i] > mid {
			arry[i], arry[tail] = arry[tail], arry[i] //go中快速交换变量值，不需要中间变量temp
			tail--
		} else {
			arry[i], arry[head] = arry[head], arry[i]
			head++
			i++
		}
	}
	arry[head] = mid
	sorter.sort(arry[:head]) //这里的head就是中间值。左边是比它小的，右边是比它大的，开始递归。
	sorter.sort(arry[head+1:])
}