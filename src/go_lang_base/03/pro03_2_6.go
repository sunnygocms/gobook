package main

func main() {
	var str string                                     //声明一个string类型的变量
	str = `The quick brown fox jumps over a lazy dog．` //赋值
	r := str[0]                                        //获取第一个元素的值
	length := len(str)                                 //字符串的长度
	str[0] = 's'                                       //编译报错：cannot assign to str[0]
}
