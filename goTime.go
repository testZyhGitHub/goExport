package main

import (
	"fmt"
	"time"
	"strconv"
)

const (
	RegMobile = "^1(3|4|5|6|7|8|9)\\d{9}$"
	RegVerifyCode = "^\\d{4}$"
	//TimeFormatMS = "2006-01-02 15:04:05.000"
	//TimeFormatMS = "2006-01-02 15:04:05"
	TimeFormatMS = "2006-01-02"
)

func main() {
	// main1()
	// main2()
	// main3()
	main4()
	
	//timeStamp := getCurrentTS()
	//fmt.Printf(timeStamp)
}

func main1() {
	fmt.Printf("Welcomt to use go time package")
	fmt.Printf("\n")
	currDayTime := time.Now()
	fmt.Print(currDayTime)
	fmt.Printf("\n")
	const DateFormat = "2006-01-02"
	//timeNow := currDayTime.Format("2006-01-02")
	timeNow := currDayTime.Format(DateFormat)
	fmt.Print(timeNow)
}

func main2() {
	//delay_wh, Due_wh := DelayOrDue()
	//fmt.Printf(delay_wh)
	//fmt.Printf("\n")
	//fmt.Printf(Due_wh)
	fmt.Printf("\n")
}

/*
func DelayOrDue() (string, string) {
	//delay_wh := fmt.Sprintf("(DATE_ADD(dtPromiseSettleDate,INTERVAL 4 DAY)<='%s')", CurrTime("2006-01-02"))
	//Due_wh	 := fmt.Sprintf("(DATE_SUB(dtPromiseSettleDate,INTERVAL 2 DAY)<='%s' AND '%s'<=DATE_ADD(dtPromiseSettleDate,INTERVAL 3 DAY))", CurrTime("2006-01-02"), CurrTime("2006-01-02"))
	//return delay_wh, Due_wh
}
*/

func main3() {
// https://blog.csdn.net/skh2015java/article/details/70051512
	t := time.Now()
    fmt.Println(t)
	fmt.Println(t.UTC().Format(time.UnixDate))
	fmt.Println(t.Unix())
	
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
    fmt.Println(timestamp)
	
	timestamp = timestamp[:10]
    fmt.Println(timestamp)
	
	fmt.Println("\n\n\n")
	// 获取当前时间戳
	timeUnix := time.Now().Unix()			// 单位s, 打印结果:		1491888244
	timeUnixNano:=time.Now().UnixNano()  	// 单位纳秒, 打印结果:	1491888244752784461
	fmt.Println(timeUnix)
	fmt.Println(timeUnixNano)
}

func getCurrentTS() (string) {
	return time.Now().Format(TimeFormatMS)
}

func main4() {
	// 时间字符串解析成时间格式
	//timeStr := "2018-01-01"
	timeStr := "2017-10-27 15:41:38"
    fmt.Println("timeStr:", timeStr)
    t, _ := time.Parse("2006-01-02", timeStr)
    fmt.Println(t.Format(time.UnixDate))
}

func main5() {

    t := time.Now()
    timestamp := t.Unix()
    fmt.Println("time type t is：", t)
    fmt.Println("timestamp is: ", timestamp)

    t = time.Now().UTC()
    timestamp = t.Unix()
    fmt.Println("time type t is：", t)
    fmt.Println("timestamp is: ", timestamp)

    // 字符串-->时间戳
    // 方法一
    the_time := time.Date(2017, 10, 27, 15, 41, 38, 856301, time.Local)
    unix_time := the_time.Unix()
    fmt.Println("方法一 时间戳：", unix_time, reflect.TypeOf(unix_time))

    // 方法二
    //the_time, err := time.ParseInLocation("2006-01-02 15:04:05", "2017-10-27 15:41:38", time.Local)
    timeStr := "2017-10-27 15:41:38"
    the_time, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
    if err == nil {
        unix_time = the_time.Unix()
        fmt.Println("方法二 时间戳：", unix_time, reflect.TypeOf(unix_time))
    }

    // 时间戳--> 字符串
    res := time.Unix(unix_time, 0).Format("2006-01-02 15:04:05")
    fmt.Println("时间戳对应字符串1：", res, reflect.TypeOf(res))

    //cntHours := 30*24
    cntHours := 30*24
    cntHoursStr := strconv.Itoa(cntHours) + "h"

    //now := time.Now()
    //dd, _ := time.ParseDuration("24h")
    dd, _ := time.ParseDuration(cntHoursStr)
    dd1 := the_time.Add(dd)
    ddTimestamp := dd1.Unix()
    fmt.Println(dd1)

    res1 := t

