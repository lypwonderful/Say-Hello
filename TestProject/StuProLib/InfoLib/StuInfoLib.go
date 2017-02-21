package StuInfoLib

import (
	/*	"fmt"
		"log"*/
	"strconv"
)

/*
函数：学生信息结构体
*/
type StuStructinfo struct {
	Name  string
	Num   string
	Phone string
}

/*
函数：处理学生信息，将学生信息转为字符串
参数：具体哪个序号；学生信息结构体
输出：字符串，错误
*/
func DealStuStructInfo(num int, Stc StuStructinfo) (retstr string, err error) {
	var AllStuSting string

	AllStuSting = IntoStr(num) + "\t" + Stc.Name + "\t" + Stc.Num + "\t" + Stc.Phone

	return AllStuSting, nil
}

/*
函数：字符串转整型
输出：字符串
*/
func IntoStr(num int) string {

	return strconv.Itoa(num)
}
