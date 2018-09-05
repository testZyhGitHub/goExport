package main

import (
    "fmt"
    "reflect"
    "time"
    "strconv"
    "strings"
    "errors"
)

const (
    TimeFormatMS = "2006-01-02 15:04:05.000"
    ThreeDaysTotalSec  = 3*24*60*60
    SevenDaysTotalSec  = 7*24*60*60
)

func getCurrentTS() (string) {
    return time.Now().Format(TimeFormatMS)
}

func main() {
    // main2()
    // main3()
    // main4()
    // main5()
    main6()
    // main7()
}

// Golang语言中, time对象和string对象之间的转换
func main1() {
    fmt.Println("----------------当前时间/时间戳/字符串----------------")
    t := time.Now()
    timestamp := t.Unix()
    fmt.Println("当前本时区时间：", t)
    fmt.Println("当前本时区时间时间戳：", timestamp)

    t = time.Now().UTC()
    timestamp = t.Unix()
    fmt.Println("当前零时区时间：", t)
    fmt.Println("当前零时区时间时间戳：", timestamp)
    fmt.Println("当前时间对应字符串：", t.Format("2006-01-02 15:04:05"))

    fmt.Println("")

    fmt.Println("------指定字符串后，字符串和时间戳之间的相互转换------")
    // 字符串-->时间戳
    // 方法一
    the_time := time.Date(2017, 7, 7, 9, 0, 0, 0, time.Local)
    unix_time := the_time.Unix()
    fmt.Println("方法一 时间戳：", unix_time, reflect.TypeOf(unix_time))

    // 方法二
    the_time, err := time.ParseInLocation("2006-01-02 15:04:05", "2017-07-07 09:00:00", time.Local)
    if err == nil {
        unix_time = the_time.Unix()
        fmt.Println("方法二 时间戳：", unix_time, reflect.TypeOf(unix_time))
    }

    // 时间戳--> 字符串
    res := time.Unix(unix_time, 0).Format("2006-01-02 15:04:05")
    fmt.Println("时间戳对应字符串：", res, reflect.TypeOf(res))
}

func main2() {

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
    //timeStr := "2017-10-27 15:41:38"
    //timeStr := "2018-06-27 20:47:07.856301"
    timeStr := "2018-10-27 20:47:07.856301"
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

    res1 := time.Unix(ddTimestamp, 0).Format("2006-01-02 15:04:05")
    fmt.Println("时间戳对应字符串2：", res1, reflect.TypeOf(res1))
}

func getCurrentDate() (string) {
    return time.Now().Format("2006-01-02")
}

func main3() {
    currDateStr := getCurrentDate()
    fmt.Println(currDateStr)

    currTimeStr := getCurrentTS()
    fmt.Println(currTimeStr)

    year := time.Now().Year()
    fmt.Println(year)

    month := time.Now().Month()
    fmt.Println(month)

    day := time.Now().Day()
    fmt.Println(day)

    hour := time.Now().Hour()
    fmt.Println(hour)

    minute := time.Now().Minute()
    fmt.Println(minute)

    second := time.Now().Second()
    fmt.Println(second)
}

// GOLang time 时间戳、日期字符串相互转化
func main4() {
    const base_format = "2006-01-02 15:04:05"
    const base_format1 = "2006-01-02"

    //获取当前时间
    nt := time.Now()
    fmt.Printf("now datetime:%v\n", nt)

    //延时年月日
    adnt := nt.AddDate(1,2,3)
    fmt.Printf("now adddate:%v\n", adnt)

    //延时秒
    ant := nt.Add(3600*1e9) //延时1小时=60*60 秒
    fmt.Printf("now add:%v\n", ant)

    //转换为时间格式字符串
    str_time1 := nt.Format(base_format)
    fmt.Printf("now time string:%v\n", str_time1)

    //转换为时间格式字符串
    str_time2 := nt.Format(base_format1)
    fmt.Printf("now time string:%v\n", str_time2)

    //时间字符串转换为日期格式
    parse_str_time, err1 := time.Parse(base_format, str_time1)
    fmt.Printf("string to datetime :%v\n", parse_str_time)
    if err1 != nil {
        fmt.Printf("err1 != nil !!!")
    }

    //时间字符串转换为日期格式
    parse_str_time2, err2 := time.Parse(base_format1, str_time2)
    fmt.Printf("string to datetime :%v\n", parse_str_time2)
    if err2 != nil {
        fmt.Printf("err2 != nil !!!")
    }

    //时间戳 秒
    timestamp := time.Now().Unix()
    println("timestamp:", timestamp)
    //时间戳 毫秒
    msec := time.Now().UnixNano() / 1e6
    println("timestamp msec:", msec)

    float_ms := msec % timestamp
    v := fmt.Sprintf("%.3f\n", float64(float_ms)/1000.0)
    println("float msec:", v)

    //时间戳转日期格式
    date_time := time.Unix(timestamp, 0)
    fmt.Printf("timestamp to datetime:%v\n", date_time)

    //时间字符串转时间戳
    t, _ := time.Parse(base_format, "2018-01-18 01:01:01")
    datetime_str_to_timestamp := t.Unix()
    println("datetime_str_to_timestamp:", datetime_str_to_timestamp)
}

func main5() {
    UpBeginDate := "2018-07-07"
    //UpEndDate   := "2018-07-05"
    //UpEndDate   := "2018-07-07"
    //UpEndDate   := "2018-07-09"
    UpEndDate   := "2018-01-09"

    _, bHisOrder1 := compDate(UpBeginDate, UpEndDate)
    if (true == bHisOrder1) {
        fmt.Println("error!!!")
    }


    ret, bHisOrder := compDate("2018-07-07", "2018-07-05")
    if ret != nil {
        fmt.Println(ret)
    } else {
        fmt.Println(bHisOrder)
    }
}

func dateToYMD(date string) (year int, month int, day int, errCode int) {
    if "" != date {
        cnt := strings.Count(date,"-")
        if 2 == cnt {
            tempStr	:= strings.Split(date, "-")

            if 3 == len(tempStr) {
                subStr1 := tempStr[0]
                subStr2 := tempStr[1]
                subStr3 := tempStr[2]

                year, errYear := strconv.Atoi(subStr1)
                if errYear != nil{
                    return 0, 0, 0, -1
                }

                month, errMonth := strconv.Atoi(subStr2)
                if errMonth != nil{
                    return 0, 0, 0, -1
                }

                day, errDay := strconv.Atoi(subStr3)
                if errDay != nil{
                    return 0, 0, 0, -1
                }

                return year, month, day,0
            } else {
                return 0, 0, 0, -1
            }
        } else {
            return 0, 0, 0, -1
        }
    } else {
        return 0, 0, 0, -1
    }
}

// 比较日期大小: true:第二个日期更早(小), false:第二个日期更新(大), 日期相等怎返回false
func compDate(DateFirst string, DateSecond string) (error, bool) {
    var firstYear	int
    var firstMonth	int
    var firstDay	int
    var secondYear	int
    var secondMonth	int
    var secondDay	int

    firstYear, firstMonth, firstDay, errRet1 := dateToYMD(DateFirst)
    //errRet1 = -1
    if (-1 == errRet1) {
        return errors.New("DateFirst Failed!!!"), false
    }

    secondYear, secondMonth, secondDay, errRet2 := dateToYMD(DateSecond)
    //errRet2 = -1
    if (-1 == errRet2) {
        return errors.New("DateSecond Failed!!!"), false
    }

    if secondYear < firstYear {
        return nil, true
    } else if secondYear == firstYear {
        if 	secondMonth < firstMonth {
            return nil, true
        } else if secondMonth == firstMonth {
            if secondDay < firstDay {
                return nil, true
            } else if secondDay == firstDay {
                return nil, false
            } else {
                return nil, false
            }
        } else {
            return nil, false
        }
    } else {
        return nil, false
    }
}

func main6() {
    // go中time比较时需要注意写法
    format := "2006-01-02 15:04:05"
    now := time.Now()
    //now, _ := time.Parse(format, time.Now().Format(format))
    a, _ := time.Parse(format, "2015-03-10 11:00:00")
    b, _ := time.Parse(format, "2015-03-10 16:00:00")

    fmt.Println(now.Format(format), a.Format(format), b.Format(format))
    fmt.Println(now.After(a))
    fmt.Println(now.Before(a))
    fmt.Println(now.After(b))
    fmt.Println(now.Before(b))
    fmt.Println(a.After(b))
    fmt.Println(a.Before(b))
    fmt.Println(now.Format(format), a.Format(format), b.Format(format))
    fmt.Println(now.Unix(), a.Unix(), b.Unix())

    // 先把当前时间格式化成相同格式的字符串,然后使用time的Before, After, Equal 方法即可.
    time1 := "2015-03-20 08:50:29"
    time2 := "2015-03-21 09:04:25"
    //先把时间字符串格式化成相同的时间类型
    t1, err := time.Parse("2006-01-02 15:04:05", time1)
    t2, err := time.Parse("2006-01-02 15:04:05", time2)
    if err == nil && t1.Before(t2) {
        //处理逻辑
        fmt.Println("true")
    }

    time3 := "2015-03-20 08:50:29"
    time4 := "2015-03-21 09:04:25"

    t5, err := time.Parse("2006-01-02 15:04:05.681824", time3)
    t6, err := time.Parse("2006-01-02 15:04:05.681824", time4)
    if err == nil && t5.Before(t6) {
        //处理逻辑
        fmt.Println("true")
    }

    commiTime1 := "2018-08-28 13:39:36"
    commiTime2 := "2018-08-28 13:40:50"

    var t3 time.Time
    var err1 error

    t3, err1 = time.Parse("2006-01-02 15:04:05.681824", commiTime1)
    if err1 == nil {
        return
    }

    var t4 time.Time
    var err2 error
    t4, err2 = time.Parse("2006-01-02 15:04:05.681824", commiTime2)
    if err2 == nil {
        return
    }

    if t3.Before(t4) {
        fmt.Println("aaaaaa")
        fmt.Println(t3)
        fmt.Println(t4)
    } else {
        fmt.Println("bbbbbb")
        fmt.Println(t3)
        fmt.Println(t4)
    }
}

func main7() {
    // 计算两个时间戳是否相隔3天

    // 获取当前时间戳
    t := time.Now()
    fmt.Println(t)
    fmt.Println(t.Unix())

    // 获取时间戳
    timestamp := time.Now().Unix()
    fmt.Println(timestamp)

    //时间字符串转时间戳
    const base_format = "2006-01-02 15:04:05"
    t1, _ := time.Parse(base_format, "2018-01-18 01:01:01")
    datetime_str_to_timestamp := t1.Unix()
    println("datetime_str_to_timestamp:", datetime_str_to_timestamp)

    const base_format1 = "2006-01-02"
    t2, _ := time.Parse(base_format1, "2018-08-27")
    datetime_str_to_timestamp1 := t2.Unix()
    println("datetime_str_to_timestamp1:", datetime_str_to_timestamp1)

    t3, _ := time.Parse(base_format1, "2018-08-24")
    datetime_str_to_timestamp2 := t3.Unix()
    println("datetime_str_to_timestamp2:", datetime_str_to_timestamp2)

    tmp := datetime_str_to_timestamp1 - datetime_str_to_timestamp2      // 精度不够的
    if tmp <= 24*60*60 {
        println("tmp: ", tmp)
    }

    // 必须要精确到时分秒
    //timeMove,_  := time.Parse(base_format, "2018-08-23 23:59:59")
    //timeMove,_  := time.Parse(base_format, "2018-08-24 00:00:00")
    //timeMove,_  := time.Parse(base_format, "2018-08-26 23:59:59")
    //timeEnd,_   := time.Parse(base_format, "2018-08-26 23:59:59")
    timeEnd,_   := time.Parse(base_format, "2018-08-27 00:00:00")
    timeMove1,_  := time.Parse(base_format, "2018-08-28 00:00:00")      //1 day     ----->      86400
    timeMove2,_  := time.Parse(base_format, "2018-08-29 00:00:00")      //2 days    ----->      172800
    timeMove3,_  := time.Parse(base_format, "2018-08-30 00:00:00")      //3 days    ----->      259200
    timeMove4,_  := time.Parse(base_format, "2018-08-31 00:00:00")      //4 days    ----->      345600
    timeMove5,_  := time.Parse(base_format, "2018-09-01 00:00:00")      //5 days    ----->      432000
    timeMove6,_  := time.Parse(base_format, "2018-09-02 00:00:00")      //6 days    ----->      518400
    timeMove7,_  := time.Parse(base_format, "2018-09-03 00:00:00")      //7 days    ----->      604800

    timeMove8,_  := time.Parse(base_format, "2018-09-03 00:00:01")

    // timeMove,_  := time.Parse(base_format, "2018-09-04 00:00:00")   // 7未来7天内
    // timeMove,_  := time.Parse(base_format, "2018-09-05 00:00:00")

    //totalSecsMove   := timeMove.Unix()

    totalSecsEnd    := timeEnd.Unix()

    totalSecsMove1 := timeMove1.Unix()
    totalSecsMove2 := timeMove2.Unix()
    totalSecsMove3 := timeMove3.Unix()
    totalSecsMove4 := timeMove4.Unix()
    totalSecsMove5 := timeMove5.Unix()
    totalSecsMove6 := timeMove6.Unix()
    totalSecsMove7 := timeMove7.Unix()

    totalSecsMove8 := timeMove8.Unix()

    tmp1 := totalSecsEnd - totalSecsMove1
    tmp2 := totalSecsEnd - totalSecsMove2
    tmp3 := totalSecsEnd - totalSecsMove3
    tmp4 := totalSecsEnd - totalSecsMove4
    tmp5 := totalSecsEnd - totalSecsMove5
    tmp6 := totalSecsEnd - totalSecsMove6
    tmp7 := totalSecsEnd - totalSecsMove7
    tmp8 := totalSecsEnd - totalSecsMove8

    println("tmp1: ", tmp1)
    println("tmp2: ", tmp2)
    println("tmp3: ", tmp3)
    println("tmp4: ", tmp4)
    println("tmp5: ", tmp5)
    println("tmp6: ", tmp6)
    println("tmp7: ", tmp7)
    println("tmp8: ", tmp8)

    // tmp = tmp7

    tmp = tmp8

    if tmp <= ThreeDaysTotalSec && tmp > 0 {    // 3天时间的总秒数
        println("tmp: ", tmp)
    } else if tmp <= 0 && tmp >= -SevenDaysTotalSec {
        println("tmp: ", tmp)
    } else {
        ;
    }
}