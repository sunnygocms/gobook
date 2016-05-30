package main //package名称，主运行的package名规定为main
//包引入区
import (
	"fmt"
)

//主函数
func main() {
	消息 := "Mr. Watson, Come Here, I Want You! \r\n 沃特森先生,过来，我想见你！"
	fmt.Println(消息) //个人不建议这种形式的变量名
	string := 1     //个人不建议使用预定义标识符作为变量名称
	fmt.Println(string)
}
