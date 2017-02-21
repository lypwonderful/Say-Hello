package StuProLib

import (
	"TestProject/StuProLib/Infolib"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	//"strconv"
	"strings"
)

/*
函数：打印一些简单的提示信息
参数：无
输出：无
*/
func PrintInputInfo() {

	SFile := OpenStuFile("stufile.txt")
	defer SFile.Close()
	/*读取文件，行读取*/
	LineFile := bufio.NewReader(SFile)
	_, err := LineFile.ReadString('\n')
	if err != nil {
		AllStuSting := "NUM" + "\t" + "NAME" + "\t" + "STUNUM" + "\t" + "PHONE" + "\t\t" + "TIPS" + "\r\n"
		InputStuInfo(SFile, AllStuSting)
	}
	fmt.Println("\t\t\tThis is a Project Test.")
	//WriteStuInfo(AllStuSting)
}

/*
函数：打开文件1.返回打开的文件
参数：文件名称
输出：文件句柄
*/
func OpenStuFile(FilePathName string) *os.File {

	f, err := os.OpenFile(FilePathName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)

	if err != nil {
		/*输出错误，并退出程序*/
		log.Fatal(err)
	}
	/*提示输出成功*/
	//fmt.Println("Open Success!")
	return f
}

/*
函数：输入信息   用于初始输入时的信息输入（需要打开文件）
输出：需要输入的信息
*/
func WriteStuInfo(InputStr string) {

	/*末尾添加换行*/
	ChStr := "\r\n"
	InputStr = InputStr + ChStr

	SFile := OpenStuFile("stufile.txt")

	SFile.WriteString(InputStr)

	defer SFile.Close()

}

/*
函数：输入信息（不需要打开文件）用于删除文件函数里面
参数：文件名
参数：字符串
*/
func InputStuInfo(SFile *os.File, InputStr string) {

	/*末尾添加换行*/
	ChStr := ""
	InputStr = InputStr + ChStr

	SFile.WriteString(InputStr)

}

/*
函数：读取某行信息
参数：具体行数
输出：字符串信息
*/
func ReadStuinfo(pos int) (readstring string, err error) {

	num := 0

	SFile := OpenStuFile("stufile.txt")
	defer SFile.Close()
	LineFile := bufio.NewReader(SFile)
	/*TXT行读取*/
	for {

		Line, err := LineFile.ReadString('\n')
		if num == pos {
			//fmt.Println(Line)
			return Line, nil
		}
		if err != nil {
			if err == io.EOF {
				//fmt.Println("Pos is outof range...")
				return " ", nil
			}
			return " ", err
		}
		num++
	}
	return " ", nil
}

/*
函数：读取所有学生信息
参数：无
输出：错误
*/
func ReadAllStuinfo() (err error) {

	SFile := OpenStuFile("stufile.txt")
	defer SFile.Close()
	LineFile := bufio.NewReader(SFile)
	/*TXT行读取*/
	for {

		Line, err := LineFile.ReadString('\n')
		Line = strings.Trim(Line, "\n")
		fmt.Println(Line)
		if err != nil {
			if err == io.EOF {

				return nil
			}
			return err
		}

	}
	return nil
}

/*
函数：删除某一行数据
参数：删除数据的位置
输出：错误
*/
func DeletStuinfo(pos int) error {
	num := 1
	n := 0        //编号
	pos = pos + 1 //原始编号初始为0，需要先加1进行修正
	count, _ := StuInfoCount()
	//fmt.Println(count)
	SFile := OpenStuFile("stufile.txt")
	LineFile := bufio.NewReader(SFile)
	/*TXT行读取*/
	SFileNew := OpenStuFile("stufilenew.txt")

	for {
		Line, err := LineFile.ReadString('\n')
		if num != pos {
			/*写入数据，除了需要删除的数据*/

			if num < pos {
				InputStuInfo(SFileNew, Line)
				n++
			} else if num > pos && num <= count {
				NewStc := StuInfoLib.StuStructinfo{}
				Slic := strings.Split(Line, "\t")
				NewStc.Name = Slic[1]
				NewStc.Num = Slic[2]
				NewStc.Phone = Slic[3]
				newdatast, _ := StuInfoLib.DealStuStructInfo(n, NewStc)
				InputStuInfo(SFileNew, newdatast)
				n++
			}
		}

		if err != nil {
			if err == io.EOF {
				//fmt.Println("DeletStuinfo Pos is outof range...")
				/*重写文件函数*/
				defer func() {
					/*关闭打开的文件*/
					SFileNew.Close()
					SFile.Close()

					/*移除文件然后重命名为原来的文件名*/
					err := os.Remove("stufile.txt")
					if err != nil {
						log.Print(err)
					}
					os.Rename("stufilenew.txt", "stufile.txt")
				}()
				return nil
			}
			return err
		}
		num++
	}
	return nil
}

/*
函数：修改某一行数据
参数：删除数据的位置
输出：错误
*/
func UpdateStuinfo(pos int, newdatastc StuInfoLib.StuStructinfo) error {
	num := 1
	pos = pos + 1 //原始编号初始为0，需要先加1进行修正
	/*判断文件里面是否包含需要修改的目录*/
	count, _ := StuInfoCount()
	if count <= 1 {
		log.Fatal("UpdateStuinfo:There is No item...")
	} else if count < pos {
		log.Fatal("UpdateStuinfo:Outof Range item...")
	}

	SFile := OpenStuFile("stufile.txt")
	LineFile := bufio.NewReader(SFile)
	/*TXT行读取*/
	SFileNew := OpenStuFile("stufilenew.txt")
	for {
		Line, err := LineFile.ReadString('\n')

		if num == pos {
			/*写入数据，除了需要修改的数据*/
			newdatainfo, _ := StuInfoLib.DealStuStructInfo(num-1, newdatastc)
			newdatainfo = newdatainfo + "\r\n"
			InputStuInfo(SFileNew, newdatainfo)
		} else {
			InputStuInfo(SFileNew, Line)
		}
		if err != nil {
			if err == io.EOF {
				//fmt.Println("UpdateStuinfo:Pos is outof range...")
				/*重写文件函数*/
				defer func() {
					/*关闭打开的文件*/
					SFileNew.Close()
					SFile.Close()

					/*移除文件然后重命名为原来的文件名*/
					err := os.Remove("stufile.txt")
					if err != nil {
						log.Print(err)
					}
					os.Rename("stufilenew.txt", "stufile.txt")
				}()
				return nil
			}
			return err
		}
		num++
	}
	return nil
}

/*
函数：读取里面文件个数
参数：无
输出：返回学生信息的个数，错误
*/
func StuInfoCount() (n int, err error) {
	num := 0
	SFile := OpenStuFile("stufile.txt")
	LineFile := bufio.NewReader(SFile)
	/*TXT行读取*/
	for {
		_, err := LineFile.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				//fmt.Println("StuInfoCount:Pos is outof range...")
				defer SFile.Close()
				return num, nil
			}
			return 0, err
		}
		num++
	}
	return 0, nil
}

/*
函数：找出某一行数据
参数：数据的位置
输出：信息切片，错误
*/
func FindStuInfo(pos int) (FindSlic []string, err error) {
	FindString, _ := ReadStuinfo(pos)
	//FindLine := strings.Split(FindString, "\r\n")
	FindLine := strings.Trim(FindString, "\r\n")
	FindStrSli := strings.Split(FindLine, "\t")
	return FindStrSli, nil
}
