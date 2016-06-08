package sunny

import (
	"bytes"
	"errors"
	"fmt"
)

func Fun1() { //无参数，无返回值
	fmt.Println("This is function Fun1")
}

func Fun2(s1 string, s2 string) { //多个参数，没有返回值
	var buf bytes.Buffer
	buf.WriteString("This is ")
	buf.WriteString(s1)
	buf.WriteString(" ")
	buf.WriteString(s2)
	buf.WriteString(". ")
	fmt.Println(buf.String())
}

func Fun3(s1 string, s2 string) string { //多个参数，一个返回值，不指定返回值名称
	result := "This is "
	result += s1
	result += " "
	result += s2
	result += ". "
	return result
}
func Fun4(s1 string, s2 string) (result string) { //多个参数，一个返回值，指定返回值名称
	result = "This is "
	result += s1
	result += " "
	result += s2
	result += ". "
	return //也可以写成return result，不过简写更符合Go的习惯
}

func Fun5(s1 string, s2 string, s ...string) (string, error) {
	//多个参数,还包括可变长参数；多个返回值
	result := ""
	if len(s) == 0 {
		return "", errors.New("没有传递可变长变量") //生成一个简单的 error 类型
	}
	result = "This is "
	result += s1
	result += " "
	result += s2
	result += ". "
	for _, tmp := range s[:len(s)-1] {
		result += tmp
		result += "--"
	}
	for _, tmp := range s[len(s)-1:] {
		result += tmp
	}

	return result, nil
}
