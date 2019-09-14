#### 变量的声明与赋值
- 变量的声明格式: var <变量名称> <变量类型>
- 变量的赋值格式: <变量名称> = <表达式>
- 变量的同时赋值:var <变量名称> [变量类型]=<表达式>

```golang
	var a int
	var a int = 1
	var a  = 1 //类型推断
	a := 1;

	var (
		aa = "hello"
		sss,bb = 1,2
	)
```
全局变量声明不可以省略var,局部变量不可以使用变量组

###### 常量

- 常量的值在编译时已经确定
- 常量的右侧必须是常量或常量表达式
- 常量表达式中的函数必须是内置函数

```golang
const a int = 1
const b = 'A'
const (
	c = 1
	d = 2
)
const e,d = 1,"c"

const(
	a = 'A'
	b
	c = iota
	d
)
// 65 65 2 3
```
- iota: 常量计数器，常量组中每定义一个常量，递增


#### func
```goLang
func name() (...int){
	
}
```
- 变长参数，多返回值
- defer，类似java中finally
- 匿名函数与闭包

#### reflect
反射可以提高程序的灵活性