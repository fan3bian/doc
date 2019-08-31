//当前程序包名
package main
//导入其他的包
import "fmt"
//使用包别名
import std "fmt"
import . "fmt"
// import "os"
// import (
//   "strings"
// )
//常量的定义
const PI =3.14
//全局变量的定义与赋值
var name = "gopher"
//一般类型的声明
type newType int
//结构的声明
type gopher struct{}
//接口的声明
type golang interface{}
//由main函数作为程序入口启动
func main() {
	fmt.Println("Hello world!你好，世界")
	//别名调用
	std.Println("Hello world!你好，世界")
	//省略调用
	Println("Hello world!你好，世界")
}
//导入包后未调用的函数或者类型会报编译错误

//常量 变量 接口 类型 结构 函数 首字母大写为public,小写为private