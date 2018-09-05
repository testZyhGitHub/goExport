package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
	"sort"
	"fmt"
	"os"
	"log"
	"strconv"
	"time"
	"net/smtp"
	"bufio"
	"io"
	"encoding/gob"
	"errors"
)

func main() {
	// main1()
	// main2()
	// main3()
	// main4()
	// main5()
	// main6()
	// main7()
	main8()
}

func main1() {
	var inTE, outTE *walk.TextEdit
	
	MainWindow {
		Title:		"SCREAMO",
		MinSize:	Size{600, 400},
		Layout:		VBox{},
		
		Children:	[]Widget {
			HSplitter {
				Children: []Widget {
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton {
				Text:		"SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
        },
    }.Run()
}

func main2() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "xiaochuan测试",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE, MaxLength: 10},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}

// 这里用到了LineEdit、LineEdit控件
func main3() {
	var usernameTE, passwordTE *walk.LineEdit
	MainWindow{
		Title:   "登录",
		MinSize: Size{270, 290},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2, Spacing: 10},
				Children: []Widget{
					VSplitter{
						Children: []Widget{
							Label{
								Text: "用户名",
							},
						},
					},
					VSplitter{
						Children: []Widget{
							LineEdit{
								MinSize:  Size{160, 0},
								AssignTo: &usernameTE,
							},
						},
					},
					VSplitter{
						Children: []Widget{
							Label{MaxSize: Size{160, 40},
								Text: "密码",
							},
						},
					},
					VSplitter{
						Children: []Widget{
							LineEdit{
								MinSize:  Size{160, 0},
								AssignTo: &passwordTE,
							},
						},
					},
				},
			},

			PushButton{
				Text:    "登录",
				MinSize: Size{120, 50},
				OnClicked: func() {
					if usernameTE.Text() == "" {
						var tmp walk.Form
						walk.MsgBox(tmp, "用户名为空", "", walk.MsgBoxIconInformation)
						return
					}
					if passwordTE.Text() == "" {
						var tmp walk.Form
						walk.MsgBox(tmp, "密码为空", "", walk.MsgBoxIconInformation)
						return
					}
				},
			},
		},
	}.Run()
}

// TableView的使用
// 这里主要使用的是TableView控件，代码参考github：
type Condom struct {
	Index   int
	Name    string
	Price   int
	checked bool
}

type CondomModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Condom
}

func (m *CondomModel) RowCount() int {
	return len(m.items)
}

func (m *CondomModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index
	case 1:
		return item.Name
	case 2:
		return item.Price
	}
	panic("unexpected col")
}

func (m *CondomModel) Checked(row int) bool {
	return m.items[row].checked
}

func (m *CondomModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked
	return nil
}

func (m *CondomModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *CondomModel) Len() int {
	return len(m.items)
}

func (m *CondomModel) Less(i, j int) bool {
	a, b := m.items[i], m.items[j]

	c := func(ls bool) bool {
		if m.sortOrder == walk.SortAscending {
			return ls
		}

		return !ls
	}

	switch m.sortColumn {
	case 0:
		return c(a.Index < b.Index)
	case 1:
		return c(a.Name < b.Name)
	case 2:
		return c(a.Price < b.Price)
	}

	panic("unreachable")
}

func (m *CondomModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func NewCondomModel() *CondomModel {
	m := new(CondomModel)
	m.items = make([]*Condom, 3)

	m.items[0] = &Condom{
		Index: 0,
		Name:  "杜蕾斯",
		Price: 20,
	}

	m.items[1] = &Condom{
		Index: 1,
		Name:  "杰士邦",
		Price: 18,
	}

	m.items[2] = &Condom{
		Index: 2,
		Name:  "冈本",
		Price: 19,
	}

	return m
}

type CondomMainWindow struct {
	*walk.MainWindow
	model *CondomModel
	tv    *walk.TableView
}

func main4() {
	mw := &CondomMainWindow{model: NewCondomModel()}

	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Condom展示",
		Size:     Size{800, 600},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text: "Add",
						OnClicked: func() {
							mw.model.items = append(mw.model.items, &Condom{
								Index: mw.model.Len() + 1,
								Name:  "第六感",
								Price: mw.model.Len() * 5,
							})
							mw.model.PublishRowsReset()
							mw.tv.SetSelectedIndexes([]int{})
						},
					},
					PushButton{
						Text: "Delete",
						OnClicked: func() {
							items := []*Condom{}
							remove := mw.tv.SelectedIndexes()
							for i, x := range mw.model.items {
								remove_ok := false
								for _, j := range remove {
									if i == j {
										remove_ok = true
									}
								}
								if !remove_ok {
									items = append(items, x)
								}
							}
							mw.model.items = items
							mw.model.PublishRowsReset()
							mw.tv.SetSelectedIndexes([]int{})
						},
					},
					PushButton{
						Text: "ExecChecked",
						OnClicked: func() {
							for _, x := range mw.model.items {
								if x.checked {
									fmt.Printf("checked: %v\n", x)
								}
							}
							fmt.Println()
						},
					},
					PushButton{
						Text: "AddPriceChecked",
						OnClicked: func() {
							for i, x := range mw.model.items {
								if x.checked {
									x.Price++
									mw.model.PublishRowChanged(i)
								}
							}
						},
					},
				},
			},
			Composite{
				Layout: VBox{},
				ContextMenuItems: []MenuItem{
					Action{
						Text:        "I&nfo",
						OnTriggered: mw.tv_ItemActivated,
					},
					Action{
						Text: "E&xit",
						OnTriggered: func() {
							mw.Close()
						},
					},
				},
				Children: []Widget{
					TableView{
						AssignTo:         &mw.tv,
						CheckBoxes:       true,
						ColumnsOrderable: true,
						MultiSelection:   true,
						Columns: []TableViewColumn{
							{Title: "编号"},
							{Title: "名称"},
							{Title: "价格"},
						},
						Model: mw.model,
						OnCurrentIndexChanged: func() {
							i := mw.tv.CurrentIndex()
							if 0 <= i {
								fmt.Printf("OnCurrentIndexChanged: %v\n", mw.model.items[i].Name)
							}
						},
						OnItemActivated: mw.tv_ItemActivated,
					},
				},
			},
		},
	}.Run()
}

func (mw *CondomMainWindow) tv_ItemActivated() {
	msg := ``
	for _, i := range mw.tv.SelectedIndexes() {
		msg = msg + "\n" + mw.model.items[i].Name
	}
	walk.MsgBox(mw, "title", msg, walk.MsgBoxIconInformation)
}

// 文件选择器(加入了icon) , 这里就是调用Windows的文件选择框 , 主要是使用：FileDialog
type MyMainWindow struct {
	*walk.MainWindow
	edit *walk.TextEdit

	path string
}

func main5() {
	mw := &MyMainWindow{}
	MW := MainWindow{
		AssignTo: &mw.MainWindow,
		Icon:     "test.ico",
		Title:    "文件选择对话框",
		MinSize:  Size{150, 200},
		Size:     Size{300, 400},
		Layout:   VBox{},
		Children: []Widget{
			TextEdit{
				AssignTo: &mw.edit,
			},
			PushButton{
				Text:      "打开",
				OnClicked: mw.pbClicked,
			},
		},
	}
	if _, err := MW.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func (mw *MyMainWindow) pbClicked() {

	dlg := new(walk.FileDialog)
	dlg.FilePath = mw.path
	dlg.Title = "Select File"
	dlg.Filter = "Exe files (*.exe)|*.exe|All files (*.*)|*.*"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		mw.edit.AppendText("Error : File Open\r\n")
		return
	} else if !ok {
		mw.edit.AppendText("Cancel\r\n")
		return
	}
	mw.path = dlg.FilePath
	s := fmt.Sprintf("Select : %s\r\n", mw.path)
	mw.edit.AppendText(s)
}

// 文本检索器
func main6() {
	mw := &MyMainWindow1{}

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "SearchBox",
		Icon:     "test.ico",
		MinSize:  Size{300, 400},
		Layout:   VBox{},
		Children: []Widget{
			GroupBox{
				Layout: HBox{},
				Children: []Widget{
					LineEdit{
						AssignTo: &mw.searchBox,
					},
					PushButton{
						Text:      "检索",
						OnClicked: mw.clicked,
					},
				},
			},
			TextEdit{
				AssignTo: &mw.textArea,
			},
			ListBox{
				AssignTo: &mw.results,
				Row:      5,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}

}

type MyMainWindow1 struct {
	*walk.MainWindow
	searchBox *walk.LineEdit
	textArea  *walk.TextEdit
	results   *walk.ListBox
}

func (mw *MyMainWindow1) clicked() {
	word := mw.searchBox.Text()
	text := mw.textArea.Text()
	model := []string{}
	for _, i := range search(text, word) {
		model = append(model, fmt.Sprintf("%d检索成功", i))
	}
	log.Print(model)
	mw.results.SetModel(model)
}

func search(text, word string) (result []int) {
	result = []int{}
	i := 0
	for j, _ := range text {
		if strings.HasPrefix(text[j:], word) {
			log.Print(i)
			result = append(result, i)
		}
		i += 1
	}
	return
}

// 邮件群发器, 别人写的邮件群发器, 出自: https://studygolang.com/articles/2078, 关于golang中stmp的使用, 请看博客: 		Go实战–通过net/smtp发送邮件(The way to go)
type ShuJu struct {
	Name    string
	Pwd     string
	Host    string
	Subject string
	Body    string
	Send    string
}

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	fmt.Println("Send to " + to)
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/html;charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain;charset=UTF-8"
	}
	body = strings.TrimSpace(body)
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func readLine2Array(filename string) ([]string, error) {
	result := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		return result, errors.New("Open file failed.")
	}
	defer file.Close()
	bf := bufio.NewReader(file)
	for {
		line, isPrefix, err1 := bf.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				return result, errors.New("ReadLine no finish")
			}
			break
		}
		if isPrefix {
			return result, errors.New("Line is too long")
		}
		str := string(line)
		result = append(result, str)
	}
	return result, nil
}

func DelArrayVar(arr []string, str string) []string {
	str = strings.TrimSpace(str)
	for i, v := range arr {
		v = strings.TrimSpace(v)
		if v == str {
			if i == len(arr) {
				return arr[0 : i-1]
			}
			if i == 0 {
				return arr[1:len(arr)]
			}
			a1 := arr[0:i]
			a2 := arr[i+1 : len(arr)]
			return append(a1, a2...)
		}
	}
	return arr
}

func LoadData() {
	fmt.Println("LoadData")
	file, err := os.Open("data.dat")
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		SJ.Name = "用户名"
		SJ.Pwd = "用户密码"
		SJ.Host = "SMTP服务器:端口"
		SJ.Subject = "邮件主题"
		SJ.Body = "邮件内容"
		SJ.Send = "要发送的邮箱，每行一个"
		return
	}
	dec := gob.NewDecoder(file)
	err2 := dec.Decode(&SJ)
	if err2 != nil {
		fmt.Println(err2.Error())
		SJ.Name = "用户名"
		SJ.Pwd = "用户密码"
		SJ.Host = "SMTP服务器:端口"
		SJ.Subject = "邮件主题"
		SJ.Body = "邮件内容"
		SJ.Send = "要发送的邮箱，每行一个"
	}
}

func SaveData() {
	fmt.Println("SaveData")
	file, err := os.Create("data.dat")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	enc := gob.NewEncoder(file)
	err2 := enc.Encode(SJ)
	if err2 != nil {
		fmt.Println(err2)
	}
}

var SJ ShuJu
var runing bool
var chEnd chan bool

func main7() {
	LoadData()
	chEnd = make(chan bool)
	var emails, body, msgbox *walk.TextEdit
	var user, password, host, subject *walk.LineEdit
	var startBtn *walk.PushButton
	MainWindow{
		Title:   "邮件发送器",
		MinSize: Size{800, 600},
		Layout:  HBox{},
		Children: []Widget{
			TextEdit{AssignTo: &emails, Text: SJ.Send, ToolTipText: "待发送邮件列表，每行一个"},
			VSplitter{
				Children: []Widget{
					LineEdit{AssignTo: &user, Text: SJ.Name, CueBanner: "请输入邮箱用户名"},
					LineEdit{AssignTo: &password, Text: SJ.Pwd, PasswordMode: true, CueBanner: "请输入邮箱登录密码"},
					LineEdit{AssignTo: &host, Text: SJ.Host, CueBanner: "SMTP服务器:端口"},
					LineEdit{AssignTo: &subject, Text: SJ.Subject, CueBanner: "请输入邮件主题……"},
					TextEdit{AssignTo: &body, Text: SJ.Body, ToolTipText: "请输入邮件内容", ColumnSpan: 2},
					TextEdit{AssignTo: &msgbox, ReadOnly: true},
					PushButton{
						AssignTo: &startBtn,
						Text:     "开始群发",
						OnClicked: func() {
							SJ.Name = user.Text()
							SJ.Pwd = password.Text()
							SJ.Host = host.Text()
							SJ.Subject = subject.Text()
							SJ.Body = body.Text()
							SJ.Send = emails.Text()
							SaveData()

							if runing == false {
								runing = true
								startBtn.SetText("停止发送")
								go sendThread(msgbox, emails)
							} else {
								runing = false
								startBtn.SetText("开始群发")
							}
						},
					},
				},
			},
		},
	}.Run()
}

func sendThread(msgbox, es *walk.TextEdit) {
	sendTo := strings.Split(SJ.Send, "\r\n")
	susscess := 0
	count := len(sendTo)
	for index, to := range sendTo {
		if runing == false {
			break
		}
		msgbox.SetText("发送到" + to + "..." + strconv.Itoa((index/count)*100) + "%")
		err := SendMail(SJ.Name, SJ.Pwd, SJ.Host, to, SJ.Subject, SJ.Body, "html")
		if err != nil {
			msgbox.AppendText("\r\n失败:" + err.Error() + "\r\n")
			if err.Error() == "550 Mailbox not found or access denied" {
				SJ.Send = strings.Join(DelArrayVar(strings.Split(SJ.Send, "\r\n"), to), "\r\n")
				es.SetText(SJ.Send)
			}
			time.Sleep(1 * time.Second)
			continue
		} else {
			susscess++
			msgbox.AppendText("\r\n发送成功!")
			SJ.Send = strings.Join(DelArrayVar(strings.Split(SJ.Send, "\r\n"), to), "\r\n")
			es.SetText(SJ.Send)
		}
		time.Sleep(1 * time.Second)
	}
	SaveData()
	msgbox.AppendText("停止发送!成功 " + strconv.Itoa(susscess) + " 条\r\n")
}

// 身份证号码校验
func main8() {
	var BlueCollarIdCardNo string		// 蓝领会员身份证号

	// BlueCollarIdCardNo = "11010519491231002X"
	// BlueCollarIdCardNo = 440524188001010014
	// BlueCollarIdCardNo = "432831196411150810"
	// BlueCollarIdCardNo = "432831196411150817"

	// BlueCollarIdCardNo = "320311770706001"   // (男)
	BlueCollarIdCardNo = "320311770706003"   // (男)
	// BlueCollarIdCardNo = "320311770706002"	//  (女)

	if (len(BlueCollarIdCardNo) == 18) {
		check := CheckIDCard18(BlueCollarIdCardNo)
		if (false == check) {
			fmt.Println("身份证号18位但是非法身份证!!!")
		} else {
			fmt.Println("合法18位身份证号!!!")
		}
	} else if (len(BlueCollarIdCardNo) == 15) {
		check := CheckIDCard15(BlueCollarIdCardNo)
		if (false == check) {
			fmt.Println("身份证号15位但是非法身份证!!!")
		} else {
			fmt.Println("合法15位身份证号!!!")
		}
	} else {
		fmt.Println("身份证号不是15位或者18位")
	}
}

func CheckIDCard18(idNumber string) bool {
	// 第1步:	数字验证
	mi := powerInt64(10, 17)

	bHas_x := strings.IndexByte(idNumber, 'x')
	bHas_X := strings.IndexByte(idNumber, 'X')

	var strTmp string
	if (-1 != bHas_x) && (1 == strings.Count(idNumber, "x")) {
		strTmp = strings.Replace(idNumber, "x", "0", -1)
	} else if (-1 != bHas_X) && (1 == strings.Count(idNumber, "X")) {
		strTmp = strings.Replace(idNumber, "X", "0", -1)
	} else if (-1 == bHas_x) && (-1 == bHas_X) {
		strTmp = idNumber
	} else {
		return false
	}

	n, err := strconv.ParseInt(strTmp, 10, 64)   // string到int64
	if err != nil {
		return false
	}

	if n < mi {
		return false
	}

	// 第2步:	省份验证
	var address string = "11x22x35x44x53x12x23x36x45x54x13x31x37x46x61x14x32x41x50x62x15x33x42x51x63x21x34x43x52x64x65x71x81x82x91"
	startTwoNum := getSubstr(idNumber, 0, 2)
	if true != strings.Contains(address, startTwoNum) {
		return false
	}

	// 第3步:	生日验证
	tmpStr1 := getSubstr(idNumber, 6, 10)
	tmpStr2 := getSubstr(idNumber, 10, 12)
	tmpStr3 := getSubstr(idNumber, 12, 14)
	birthDayStr :=  tmpStr1 + "-" + tmpStr2 + "-" + tmpStr3

	const base_format = "2006-01-02"
	_, err = time.Parse(base_format, birthDayStr) // 时间字符串转换为日期格式
	if err != nil {
		return false
	}

	// 第4步:	校验码验证, 校检码可以是0—9的数字, 有时也用x表示, 由于10被替换表达为x
	arrVarifyCode	:= []string{"1", "0", "x", "9", "8", "7", "6", "5", "4", "3", "2"}											// 根据身份证1~17位上的字符乘以,
	Wi				:= []string{"7", "9", "10", "5", "8", "4", "2", "1", "6", "3", "7", "9", "10", "5", "8", "4", "2"}		// 身份证号码每位对应的乘数因子数组

	ontologyCode	:= getSubstr(idNumber, 0, 17)
	var Ai []byte = []byte(ontologyCode)		// string 转为[]byte,

	var sum int

	/*
			整数乘法运算法则:			被乘数X乘数 = 积

			整数乘法法则:
										从右边起, 依次用第二个因数每位上的数去乘第一个因数, 乘到哪一位, 得数的末尾就和第二个因数的哪一位对齐;

然后把几次乘得的数加起来.
（整数末尾有0的乘法：可以先把0前面的数相乘,然后看各因数的末尾一共有几个0,就在乘得的数的末尾添写几个0.）
整数除法运算法则：
积/一个因数=另一个因数
被除数/除数=商
被除数/商=除数
除数X商=被除数
整数的除法法则：
1）从被除数的高位起,先看除数有几位,再用除数试除被除数的前几位,如果它比除数小,再试除多一位数；
2）除到被除数的哪一位,就在那一位上面写上商；
3）每次除后余下的数必须比除数小.
整数乘除法运算法则：
先乘除,后加减,有括号的先算括号里的.

	*/
	var i int
	var wiElementInt int
	var aiElementInt int
	for i = 0; i < 17; i++ {
		wiElementInt,_ = strconv.Atoi(Wi[i])
		aiElementInt,_ = strconv.Atoi(string(Ai[i]))
		sum += wiElementInt * aiElementInt
	}

	/*
		   余数只可能有:
					 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 总共有11个数字.
		   其分别对应的最后一位身份证的号码为:
					 1, 0, X, 9, 8, 7, 6, 5, 4, 3, 2
		   即余数0对应1, 余数1对应0, 余数2对应X......
	*/
	_, remainder := divide(sum, 11)
	lastBitStr := getSubstr(idNumber, 17, 18)
	sLastBitStr := strings.ToLower(lastBitStr)  // ToLower, ToUpper
	if (arrVarifyCode[remainder] != sLastBitStr) {
		return false;       // 校验码验证
	}

	return true
}

func CheckIDCard15(idNumber string) bool {
	// 第1步:	数字验证
	mi := powerInt64(10, 14)

	n, err := strconv.ParseInt(idNumber, 10, 64)   // string到int64
	if err != nil {
		return false
	}

	if n < mi {
		return false
	}

	// 第2步:	省份验证
	var address string = "11x22x35x44x53x12x23x36x45x54x13x31x37x46x61x14x32x41x50x62x15x33x42x51x63x21x34x43x52x64x65x71x81x82x91"
	startTwoNum := getSubstr(idNumber, 0, 2)
	if true != strings.Contains(address, startTwoNum) {
		return false
	}

	// 第3步:	生日验证
	tmpStr1 := getSubstr(idNumber, 6, 8)
	tmpStr2 := getSubstr(idNumber, 8, 10)
	tmpStr3 := getSubstr(idNumber, 10, 12)
	birthDayStr :=  "19" + tmpStr1 + "-" + tmpStr2 + "-" + tmpStr3  // 15位身份证号码在我们的系统中很少会用到!!!

	const base_format = "2006-01-02"
	_, err = time.Parse(base_format, birthDayStr) // 时间字符串转换为日期格式
	if err != nil {
		return false
	}

	return true
}

func divide(a, b int) (int, int) {
	quotient	:= a / b
	remainder	:= a % b
	return quotient, remainder		// 商, 余数
}

// 二分幂法: 求x^n
func powerInt64(x int64, n int) int64 {
	var ans int64 = 1
	var m	 int

	for n != 0 {
		m = n % 2

		if m == 1 { //如果n是奇数, 就要多乘一次
			ans *= x
		}

		x *= x
		n /= 2		// 二分
	}

	return ans
}

// 截取字符串 start 起点下标 end 终点下标(不包括)
func getSubstr(str string, start int, end int) string {
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