package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"  //找ai学的正则
	"strconv" //学了下，用来转换string和int的
	"strings"
	//具体可以看另一个TXl（用c坤坤写的【用基础语法堆出来的史】
)

// 定义接口类型用来适配通讯录
type txler interface {
	choose() int
}

// 定义一个通讯录的结构体类型
type TXL struct {
	user []USER //一个切片，用于存储联系人的结构体
	num  int
}

// 用来判断输入的结构体
type input_it struct {
	number_tell regexp.Regexp
	age_tell    regexp.Regexp
}

// 定义联系人的结构体类型
type USER struct {
	name   string
	addr   string
	number string
	age    string
	sex    string
}

func write_text(txl_r *TXL) {
	// 覆盖写入没问题
	name, err := os.Create("txl_text")
	//似乎不得不用err
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	name.WriteString(strconv.Itoa(txl_r.num) + "\n")
	for i := 0; i < txl_r.num; i++ {
		name.WriteString(txl_r.user[i].name + "\n")
		name.WriteString(txl_r.user[i].addr + "\n")
		name.WriteString(txl_r.user[i].number + "\n")
		name.WriteString(txl_r.user[i].age + "\n")
		name.WriteString(txl_r.user[i].sex + "\n")
	}
	name.Close()
}
func read_text(txl *TXL) {
	f, err := os.Open("txl_text")
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		return
	}
	numStr := scanner.Text()
	if numStr == "" {
		return
	}
	txl.num, _ = strconv.Atoi(numStr)
	txl.user = make([]USER, txl.num)
	for i := 0; i < txl.num; i++ {
		if scanner.Scan() {
			txl.user[i].name = scanner.Text()
		}
		if scanner.Scan() {
			txl.user[i].addr = scanner.Text()
		}
		if scanner.Scan() {
			txl.user[i].number = scanner.Text()
		}
		if scanner.Scan() {
			txl.user[i].age = scanner.Text()
		}
		if scanner.Scan() {
			txl.user[i].sex = scanner.Text()
		}
	}
} /*本来是
scanner.Scan()
txl.user[i].name =scanner.Text()
但似乎会因为第一次运行导致错误
*/

func makeTell(it *input_it) {
	*it = input_it{
		number_tell: *regexp.MustCompile("^\\d{11}$"),
		age_tell:    *regexp.MustCompile("^\\d+$"),
	}
}
func tell_number(i *input_it, s string) bool {
	makeTell(i)
	return i.number_tell.MatchString(s)
}
func tell_age(i *input_it, s string) bool {
	makeTell(i)
	return i.age_tell.MatchString(s)
}
func soft_TXL(txl *TXL) {
	for i := 0; i < txl.num; i++ {
		for j := 0; j < txl.num-i-1; j++ {
			if txl.user[j].age > txl.user[j+1].age {
				var txl_max_temp USER
				txl_max_temp = txl.user[j+1]
				txl.user[j+1] = txl.user[j]
				txl.user[j] = txl_max_temp
			}
		}
	}
}

// 最麻烦的add，用Scanln出现缓存区错误，err又用不明白，只能用正则了
func add(txl *TXL) {
	var userTemp USER
	var inputTemp input_it
	fmt.Println("输入名字：")
	fmt.Scanf("%s", &userTemp.name)
	fmt.Println("请输入地址：")
	fmt.Scanf("%s", &userTemp.addr)
	fmt.Println("请输入电话号码：")
	for {
		fmt.Scanf("%s", &userTemp.number)
		if tell_number(&inputTemp, userTemp.number) == false {
			fmt.Println("so?what can i say,下一次让cyjj抓住你喵,请输入正确的电话号码：")
			continue
		} else {
			break
		}
	}
	fmt.Println("请输入年龄：")
	for {
		fmt.Scanf("%s", &userTemp.age)
		if tell_age(&inputTemp, userTemp.age) == false {
			fmt.Println("cyjj变成小猫咪来抓捕你了，请务必输入正确的年龄：")
			continue
		} else {
			if userTemp.age < "6" {
				fmt.Println("bro以为自己很年轻{}-_-{}")
				break
			} else if userTemp.age > "150" {
				fmt.Println("给老资历跪下了@_@")
				break
			} else {
				break
			}
		}
	}
	fmt.Println("请输入性别(man or woman)：")
	for {
		fmt.Scanf("%s", &userTemp.sex)
		if userTemp.sex != "man" && userTemp.sex != "woman" {
			fmt.Println("马萨卡，你是传说中的男娘！请输入正确的性别吧>_<:")
			continue
		} else {
			break
		}
	}
	txl.user = append(txl.user, userTemp)
	txl.num++
}
func delect(txl *TXL) {
	var del_temp int
	emerge(txl)
	fmt.Println("请输入要删去的联系人的序号:")
	fmt.Scanf("%d", &del_temp)
	txl.user = append(txl.user[:del_temp], txl.user[del_temp+1:]...)
	//源代码为txl.user[del_temp] = USER{}，问的ai告诉我：“你的写法只把元素置空，但不会从切片里删除”，但这语法还不是很懂，
	txl.num--
}
func search(txl *TXL) int {
	var name_temp string
	var win_3g bool = true
	var user_Temp_s int = -1
	fmt.Println("请输入查找人的名字")
	fmt.Scanf("%s", &name_temp)
	for i := 0; i < len(txl.user); i++ {
		if txl.user[i].name == name_temp {
			fmt.Println("\n", i, "...")
			fmt.Println("名字：", txl.user[i].name)
			fmt.Println("年龄：", txl.user[i].age)
			fmt.Println("电话号码：", txl.user[i].number)
			fmt.Println("性别：", txl.user[i].sex)
			fmt.Println("地址：", txl.user[i].addr)
			user_Temp_s = i
			win_3g = false
		} else if win_3g && i == len(txl.user)-1 {
			fmt.Println("没有该联系人喵！>-<")
		}
	}
	return user_Temp_s
}

// 显示联系人
func emerge(txl *TXL) {
	if txl.num > 0 {
		fmt.Printf("当前通讯录共有%d人", txl.num)
		for i := 0; i < txl.num; i++ {
			fmt.Println("\n", i, ".")
			fmt.Println("名字：", txl.user[i].name)
			fmt.Println("年龄：", txl.user[i].age)
			fmt.Println("电话号码：", txl.user[i].number)
			fmt.Println("性别：", txl.user[i].sex)
			fmt.Println("地址：", txl.user[i].addr)
		}
	} else {
		fmt.Println("当前通讯录没人")
	}

}
func remix(txl *TXL) {
	var choose_del int
	var user_Temp_r int = search(txl)
	fmt.Println("是否修改？（0-yes，1-no")
	fmt.Scanf("%d", &choose_del)
	if choose_del == 0 {
		var userTemp USER
		var inputTemp input_it
		fmt.Println("输入名字：")
		fmt.Scanf("%s", &userTemp.name)
		fmt.Println("请输入地址：")
		fmt.Scanf("%s", &userTemp.addr)
		fmt.Println("请输入电话号码：")
		for {
			fmt.Scanf("%s", &userTemp.number)
			if tell_number(&inputTemp, userTemp.number) == false {
				fmt.Println("so?what can i say,下一次让cyjj抓住你喵,请输入正确的电话号码：")
				continue
			} else {
				break
			}
		}
		fmt.Println("请输入年龄：")
		for {
			fmt.Scanf("%s", &userTemp.age)
			if tell_age(&inputTemp, userTemp.age) == false {
				fmt.Println("cyjj变成小猫咪来抓捕你了，请务必输入正确的年龄：")
				continue
			} else {
				break
			}
		}
		fmt.Println("请输入性别(man or woman)：")
		for {
			fmt.Scanf("%s", &userTemp.sex)
			if userTemp.sex != "man" && userTemp.sex != "woman" {
				fmt.Println("马萨卡，你是传说中的男娘！请输入正确的性别吧>_<:")
				continue
			} else {
				break
			}
		}
		txl.user[user_Temp_r] = userTemp
	}
}
func remove(txl *TXL) {
	txl.user = []USER{}
	txl.num = 0
	fmt.Println("已清空")
}

// Scanln一点都不好用！！！sb缓存区
/*啊！用半天发现根本不需要这个，留个纪念
	func clean() {
	var fuck string
	fmt.Scanln(&fuck)
}*/
func (txl *TXL) choose() int {
	var chooseNum int = 0
	fmt.Println("选择你的功能喵：\n" + strings.Repeat("-", 10))
	fmt.Println("1.添加联系人\n2.删除联系人\n3.查找联系人\n4.显示联系人\n5.修改联系人\n6.清空联系人\n0.退出通讯录")
	fmt.Println(strings.Repeat("-", 10))
	fmt.Scanln(&chooseNum)
	switch chooseNum {
	case 1:
		add(txl)
		return 1
	case 2:
		delect(txl)
		return 2
	case 3:
		search(txl)
		return 3
	case 4:
		emerge(txl)
		return 4
	case 5:
		remix(txl)
		return 5
	case 6:
		remove(txl)
		return 6
	case 0:
		return 0
	default:
		fmt.Println("无效选择，请重新输入喵~")
		return -1
	}
}

// 强行使用接口，hhh，感觉目前还用不上的
func use_TXL(T txler) int {
	usefacility := T.choose()
	return usefacility
}
func main() {
	var t = &TXL{
		user: []USER{},
		num:  0,
	}
	read_text(t)
	for {
		if use_TXL(t) == 0 {
			break
		} else {
			soft_TXL(t)
			continue
		}
	}
	write_text(t)
}

//cyjj卡哇伊kqgg卡哇伊rtgg卡哇伊
