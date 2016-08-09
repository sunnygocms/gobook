package main //package名称，主运行的package名规定为main
//包引入区
import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

func Sunnysplit(s rune) bool {
	if s == '$' {
		return true
	}
	return false
}

func JudgeType(e interface{}) (result string) {
	switch e.(type) {
	case int:
		result = "整型"
		break
	case string:
		result = "字符串"
		break
	case []byte:
		result = "[]bytes"
		break
	}
	return
}

func Byte2Str(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}

func Str2Byte(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(s))))
}

//主函数
func main() {
	yourStr := "Mr. Watson, Come Here, I Want You!"
	fmt.Println(len(yourStr))
	fmt.Println(yourStr[4:16])
	fmt.Println(strings.Index(yourStr, "Come"))
	fmt.Println(strings.Index(yourStr, "ff"))
	yourStr = "123456789"
	yourNumber, error := strconv.Atoi(yourStr)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
	} else {
		fmt.Println(JudgeType(yourNumber), yourNumber)
	}
	yourStr = strconv.Itoa(yourNumber)
	fmt.Println(JudgeType(yourStr), yourStr)

	yourNumberFloat := "1234.56789"
	yourStr = fmt.Sprintf("%f", yourNumberFloat)
	fmt.Println(JudgeType(yourStr), yourStr)

	yourStr = "Mr.$Watson,$Come$Here,$I$Want$You!"
	fmt.Println(strings.Replace(yourStr, "$", " ", -1)) //out  [Mr. Watson, Come Here, I Want You!]

	yourStr = "Mr. Watson, Come Here, I Want You!"
	yourByte := Str2Byte(&yourStr)
	fmt.Println(JudgeType(yourByte), yourByte)

	yourStr = Byte2Str(yourByte)
	fmt.Println(JudgeType(yourStr), yourStr)

	yourStr = "Mr.              Watson, Come       Here, I          Want    You!"
	fmt.Println(strings.Fields(yourStr)) //out  [Mr. Watson, Come Here, I Want You!]

	yourStr = "Mr.$Watson,$Come$Here,$I$Want$You!"
	fmt.Println(strings.FieldsFunc(yourStr, Sunnysplit)) //out  [Mr. Watson, Come Here, I Want You!]

	yourStr = "Mr.$Watson,$Come$Here,$I$Want$You!"
	fmt.Println(strings.Split(yourStr, "$")) //out  [Mr. Watson, Come Here, I Want You!]

	yourStr = "Mr.$Watson,$Come$Here,$I$Want$You!"
	fmt.Println(strings.SplitAfter(yourStr, "$")) //out  [Mr.$ Watson,$ Come$ Here,$ I$ Want$ You!]

	yourStr = "   Mr. Watson, Come Here, I Want You!  "
	fmt.Println(strings.Trim(yourStr, " "))

	yourStr = "\r\n\t   Mr. Watson, Come Here, I Want You!  \r\n\t"
	fmt.Println(strings.TrimSpace(yourStr))

}
