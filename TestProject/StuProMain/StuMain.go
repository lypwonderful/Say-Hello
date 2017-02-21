/*
	1.简单的学生信息管理
	2.AUTHER：lypwonderful
	3.Date:2017/2/21
************************************
	1.写入文件默认为根文件下的	stufile.txt

*/
package main

import (
	"TestProject/StuProLib/BaseLib"
	"TestProject/StuProLib/Infolib"
	"bufio"
	"fmt"
	//"log"
	"os"
	"strings"
)

/*提示输入*/
func PrintPromptInfo() {
	fmt.Println("\t=================================================================")
	StuProLib.PrintInputInfo() //打印提示信息，在文件里面添加头信息（NUM NAME ....）
	fmt.Println("\t\tWhat You Want To Do,Input As The Follow Prompt.")
	fmt.Println("\t\t1.Add The Stu Item.")
	fmt.Println("\t\t2.Find a Stu Item.")
	fmt.Println("\t\t3.Del a Stu Item.")
	fmt.Println("\t\t4.Update a Stu Item.")
	fmt.Println("\t\t5.Show All Stu Info.")
	fmt.Println("\t\t6.Quit EXE.")
	fmt.Println("\t=================================================================")
}
func main() {
	PrintPromptInfo()
	var Num int
	for {
		fmt.Print("Chose:")
		fmt.Scanf("%d", &Num)
		switch Num {
		case 1:
			{
				fmt.Println("\t\t1.Add The Stu Item.The Input Format is: NAME NUM PHONE")
				AddStuInem()
			}
		case 2:
			{
				var num = 1

				fmt.Println("\t\t2.Find a Stu Item.")
				fmt.Print("Find Item Num:")
				fmt.Scanf("%d", &num)
				sli, _ := StuProLib.FindStuInfo(num)
				fmt.Println(sli)
			}
		case 3:
			{
				var num = 1
				fmt.Println("\t\t3.Del a Stu Item.")
				fmt.Print("Del Item Num:")
				fmt.Scanf("%d", &num)
				StuProLib.DeletStuinfo(num)
			}
		case 4:
			{
				var num = 1
				fmt.Println("\t\t4.Update a Stu Item.")
				fmt.Print("Update Item Num:")
				fmt.Scanf("%d", &num)
				fmt.Print("Update Item Info:")
				UpdateStuInem(num)

			}
		case 5:
			{
				fmt.Println("\t\t5.Show All Stu Info.")
				StuProLib.ReadAllStuinfo()
			}
		case 6:
			{
				fmt.Println("\t\t6.Quit EXE.")
				//log.Fatal("Quit EXE Success.")
				//exit()
				os.Exit(0)
			}
		default:
			{
				PrintPromptInfo()
			}
		}
	}
}

/*添加元素*/
func AddStuInem() {
	i, _ := StuProLib.StuInfoCount()
	var AddStuStr string
	var StuStruct = StuInfoLib.StuStructinfo{
		Name:  "nil",
		Num:   "nil",
		Phone: "nil",
	}

	for {
		InfoRead := bufio.NewReader(os.Stdin)
		AddStuStr, _ = InfoRead.ReadString('\n')
		AddStuStr = strings.Trim(AddStuStr, "\n")
		AddStuSlic := strings.Split(AddStuStr, " ")
		/*q 退出输入*/
		if strings.Compare("q", AddStuSlic[0]) == 0 {
			fmt.Println("Add Ok!")
			break
		}

		/*简单判断输入错误，不能满足NAME NUM PHONE*/
		if len(AddStuSlic) != 3 {
			fmt.Println("Add:Input Error.")
		} else {
			StuStruct.Name = AddStuSlic[0]
			StuStruct.Num = AddStuSlic[1]
			StuStruct.Phone = AddStuSlic[2]
			Str5, _ := StuInfoLib.DealStuStructInfo(i, StuStruct)
			StuProLib.WriteStuInfo(Str5)
			i++
		}

	}

}

/*更新数据*/
func UpdateStuInem(pos int) {
	//i, _ := StuProLib.StuInfoCount()
	var AddStuStr string
	var StuStruct = StuInfoLib.StuStructinfo{
		Name:  "nil",
		Num:   "nil",
		Phone: "nil",
	}

	InfoRead := bufio.NewReader(os.Stdin)
	AddStuStr, _ = InfoRead.ReadString('\n')
	AddStuStr = strings.Trim(AddStuStr, "\n")
	AddStuSlic := strings.Split(AddStuStr, " ")

	if len(AddStuSlic) != 3 {
		fmt.Println("Add:Input Error.")
	} else {
		StuStruct.Name = AddStuSlic[0]
		StuStruct.Num = AddStuSlic[1]
		StuStruct.Phone = AddStuSlic[2]
		StuProLib.UpdateStuinfo(pos, StuStruct)
	}

}
