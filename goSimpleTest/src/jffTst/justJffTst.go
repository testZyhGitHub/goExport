package main

import (
	"runtime"
	"strconv"
	"fmt"
	log "github.com/xiaomi-tc/log15"
	"regexp"
)

func main() {
	main1()
	main2()
}

func main1() {
	runTimeStr := GetRuntimeInfo()
	fmt.Println(runTimeStr)

	//var inMap map[string]interface{}
	inMap := make(map[string]interface{})

	inMap["SPCoopName"]	= "testCoopSendMsg7"
	inMap["SendToA"]		= 1
	inMap["SendToB"]		= 2
	inMap["ContactName"]	= "HelloOk2"
	inMap["UserID"]		= 101856

	LogParam("NewSPCoop", inMap)
}

func GetRuntimeInfo() (string) {
	pc, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc)
	return " ( " + file + " " + strconv.Itoa(line) + " " + funcName.Name() + " ) "
}

// [2018/08/21 10:18:12] [info] [justJffTst.go:37] msg="NewSPCoop输入参数:map[UserID:101856 SPCoopName:testCoopSendMsg7 SendToA:1 SendToB:2 ContactName:HelloOk2]"
func LogParam(inName string, inParam interface{}) {
	log.Info(inName + "输入参数:" + fmt.Sprintf("%v", inParam))
}

const RegMobile = "^1(3|4|5|6|7|8|9)\\d{9}$"
const PhoneNumberReg = "^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$"     // 最新、最全、最准确的手机号正则表达式

//检查手机号码合法性
func isMobile(mobile string) bool {
	reg := regexp.MustCompile(RegMobile)
	return reg.MatchString(mobile)
}

// 检查手机号码格式是否正确!!!
func CheckPhoneNumStr(dateStr string) bool {
	reg := regexp.MustCompile(PhoneNumberReg)
	return reg.MatchString(dateStr)
}

// 手机号码校验
func main2() {

}

