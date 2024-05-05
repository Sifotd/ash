# **Go day 01 笔记**

## 常见的问题

#### var就是定义，不能使用：=

#### expected error ,foundxx就是多了xx,需要改xx

## go语言环境搭建

### 下载

下载地址 [All releases - The Go Programming Language (google.cn)](https://golang.google.cn/dl/)

安装目录选择一个好找的

go version查看是否安装好（cmd查看）

### 配置gopath

1 在自己的带闹闹上新建一个目录（存放自己编写的go语言代码）

2 在环境变量里面新建一个gopath

3 在刚刚新建的文件夹下新建三个文件夹 bin src pkg

4 把 bin目录添加到path里面去

5 删掉默认值 （%USERPROFILE%\go\bin）自己新建



## 第一个程序

```go
package main

import "fmt"//导包 包名 双引号

//函数外只能放置标识符（变量/函数/常量/类型）的声明
func main() {//程序入口
	fmt.Println("hello world")
}
```



### 编译

使用go buid  在项目目录下面执行 go build再执行 ./文件名.exe

go build -o 希望生成的exe文件名.exe

#### go run 

像执行脚本文件一样执行go代码

#### go install 

分为2步 

1 先编译 得到一个可执行文件

2 将可执行文件拷贝到GOPATH/bin

### 交叉编译

go支持跨平台编译 例如在windows平台编译一个能在linux平台执行的可执行文件

```go
  SET CGO_ENABLED=0
  SET GOOS=darwin
  SET GOARCH=amd64
  go build main.go//mac
  
  SET CGO_ENABLED=0//禁用go
  SET GOOS=linux//目标平台是
  SET GOARCH=amd64//目标处理器架构是amd64
  go build main.go
```

## Go语言文件的基本结构

```go
package main

import "fmt"//导包 包名 双引号

//函数外只能放置标识符（变量/函数/常量/类型）的声明
func main() {//程序入口
	fmt.Println("hello world")
}
```

## 变量和常量

#### 标识符 

变量名 函数名 常量名等等 标识符由字母数字和_（下划线组成），并且只能以字母和-(下划线)开头

#### 关键字

编程语言中预先定义好的具有特殊含义的标识符 go语言中 有25个关键字

#### 保留字

保留字 (reserved word)，指在高级语言中已经定义过的字，程序员不能再将这些字作为变量名、过程名或函数名使用  go语言中 有37个保留字

#### 变量的来历

利用变量将数据的内存地址保存起来，通过变量找到对应数据

#### 变量的类型

变量的类型需要声明 声明才可以使用

#### 变量声明

变量声明之后才可以使用

##### 标准声明

Go语言中变量声明格式

```go
var 变量名 变量类型
```

行尾无需引号

##### 批量声明

```go
var(
age int//空值 0
name string //空值""
s1 bool//空值 falae
)
```

##### 实例

```go
package main

import "fmt"

var (
	age  int
	name string
)

func main() {
	age = 18
	name = "门前雪"
	fmt.Printf("name:%s",name)//%s 占位符，使用name这个变量的值去替换占位符
	fmt.Println("age:",age)//ln 换行
}

```

#### 变量赋值

##### 声明变量同时赋值

```go
var s1 string="眼中人"
```

##### 类型推导(根据值判断类型是什么变量)

```go
var s2="门前雪"
fmt.Println(s2)
```

##### 短变量声明

在函数内部，可以使用更简短的:=声明变量

```go
s3 :="test"
fmt.Println(s3)
```

##### 匿名变量（哑元变量）

在使用多重赋值是，如果想要忽略某个值，可以使用匿名变量，匿名变量用一个下划线_表示

匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明

#### 注意事项

函数外的每个语句都必须以关键字开始

：=不能存在函数外

_多用于占位，表示忽略值

#### 常量

常量是恒定不变的值 声明方式如下

```go
const i=13.5
```

##### 多个常量一起声明

```go
const(
i=5
c=6
d=8
)
```

##### const声明多个变量时，如果省略了，则表示和上一行的值相同

```go
const(
q=1
c
w    
)
```

#### iota

iota 是go语言的常量计数器，只能在常量的表达式中使用

iota在const关键字出现时将被重置为0，const语言中每新增一行常量声明，将使iota计数增加一次，使用iota能简化定义，在定义枚举时有用

```go
const(
n1=iota//0
n2//1
n3//2
n4//3    
)
```

##### 几个常见的iota示例

###### 使用_跳过某些值

```go
const(
a1=iota//0
a2=iota //1
_
a3=iota//3
)
```

###### iota实现中间插队

```go
const(
n1=iota//0
n2=100//100
n3=iota//2
n4=iota//3 
)
const n5=iota//0
```

###### 多个常量声明在一行

```go
const(
d1,d2=iota+1,iota+2
d3,d4=iota+3,iota+4 
    //d1 1 d2 2  d3 4 d4 5
)
```

###### 定义数量级

```go
const(
_=iota
    kb=1<<(10*iota)//二进制数，左移10位，结果为1024
    mb=1<<(10*iota)
    gb=1<<(10*iota)
    tb=1<<(10*iota)
    pb=1<<(10*iota)
)
```

## 数据类型

### 整形数据类型

![img](https://pic4.zhimg.com/80/v2-b8da9a4e850435bd7f308fd134d54f2b_1440w.webp)

### 特殊整形

#### uint

32位操作系统上就是uint32 64位操作系统就是uint64

#### int

32位操作系统上就是int32 64位操作系统就是int64

#### uintptr

无符号整型 用于存放一个指针

### 注意事项

在使用int和uint的时候，不能假定它是32位或64位的整形，而是考虑int和uint可能在不同平台上的差异



### 八进制与十六进制

go语言中无法直接定义二进制数

go语言中使用方法如下

```go
	//输出十进制
	var a int = 10
	fmt.Printf("%d \n", a)

	//输出八进制，以0开头
	var b int = 012
	fmt.Printf("%o \n", b)

	//输出十六进制,以0x开头
	var c int = 0x16
	fmt.Printf("%x\n", c)
//查看类型变量
fmt.Printf("%T",c)

//声明int8类型的变量
i4:=int8(9)
fmt.Printf("%T",i4)
```

### 浮点型

Go语言支持2种浮点型数,float32和float64,这2种浮点数数据格式都遵循ieee 754标准 。float的浮点数的最大范围约为3.4e38,可以使用常量定义 math.MaxFloat32。float64的浮点数的最大范围位1.8e308，可以使用一个常量定义math.MaxFloat64

打印浮点数的时候，可以使用fmt包配合动词%f，代码如下

```go
func main(){
    fmt.Printf("%f.2f" math.pi)
}
```

```go

	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认go语言中的小数都是float64类型

	f2 := float32(1.3456)
	fmt.Printf("%t\n", f2) //显示声明float32类型
```



float32类型的值不能直接赋值给float64类型的变量

### 复数

complex64和complex128

```go
var c1 complex64
c1=1+2i
```

复数有实部和虚部，complex64的实部和虚部为32位 ，complex128的实部和虚部为64位

### 布尔值

Go语言中的以bool类型进行声明布尔类型数据声明，布尔型数据只有true和false2个值

#### 注意事项

布尔类型的默认值为false

go语言中不允许将整形数据强制转换为布尔型

布尔型无法参与数值运算，也无法与其他类型进行转换

## fmt总结

```go
//fmt占位符
	fmt.Printf("%t\n", n)
	//查看变量的值
	fmt.Printf("%v\n", n)
	//查看二进制
	fmt.Printf("%d\n", n)
	//查看八进制
	fmt.Printf("%o\n", n)
	//查看十六进制
	fmt.Printf("%x\n", n)

	var s = "门前雪"
	//字符串
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)

```

## 字符串

go语言中的字符串以原生数据类型出现，可以在go语言的源码中直接添加非asci码 字符串必须以""包括，''单引号包裹的是字符

```go
s1="love"
s2='5'
s3='爱'
s4='e'
```

1字节-8bit(8个二进制)

一个字符'a'=1个字节

一个utf8编码的汉族'爱'=一般占3个字节

### 字符串转义符

go语言的字符串常见转义符包含会场 换行 单双引号 制表符等

| 转义符 | 含义                           |
| ------ | ------------------------------ |
| `\r`   | 转义回车，返回行首             |
| `\n`   | 转义换行，跳到下一行的同列位置 |
| `\t`   | 转义制表                       |
| `\'`   | 转义单引号                     |
| `\"`   | 转义双引号"                    |
| `\\`   | 转义反斜杠\                    |

```go
// \防止程序误会
	path := "\"C:\\Users\""
	fmt.Print(path)
```

```go
s = "i ' m ok"
	fmt.Print(s)
```

### 多行字符串

go语言中要定义一个多行字符串，救必须使用反引号字符

```go
package main

import "fmt"

func main() {
	//使用反引号来定义多行字符串
	const str = `
			asdasddw
   			asfeffeef
	 `
	fmt.Println(str)
}

```

```go
	s6 := `
	门前雪
	眼中人
	`
fmt.Print(s6)
```

### 字符串的常用操作

方法	   		介绍
len(str)	 	求长度
+或fmt.Sprintf		拼接字符串
strings.Split	分割
strings.contains	判断是否包含
strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
strings.Index(),strings.LastIndex()	子串出现的位置
strings.Join(a[]string, sep string)	join操作

```go

	//字符串相关操作
	fmt.Println(len((s6))) //长度

	//字符串拼接
	name1 := "门前雪"
	name2 := "意中人"
	yi := name1 + name2
	fmt.Println(name1 + name2)
	fmt.Println(yi)
	ss2 := fmt.Sprintf("%s%s", name1, name2)
	fmt.Printf(ss2)
	//分隔
	te := strings.Split(ss2, "\\")
	fmt.Println(te)
	//包含
	fmt.Println(strings.Contains(ss2,"前"))
	//前缀和后缀判断
	fmt.Println(strings.HasPrefix(ss2,"门"))//前缀
	fmt.Println(strings.HasSuffix(ss2,"雪"))//后缀
	//判断子串出现的位置
	s5:="qwrersstr"
	fmt.Println(strings.Index(s5,"a"))
	fmt.Println(strings.LastIndex(s5,"r"))//判断最后一次出现的位置
	//拼接
	fmt.Println(strings.Join(te,"+"))
//从字符串中拿出具体的字符
	for _, c := range s {
		fmt.Printf("%c\n", c)//%c:字符
	}
}
```

### byte和runne类型

字符用单引号（''）组成 

```go
var a:="色"
```

go语言的字符有以下2种

1 uint8类型 或者叫byte类型 代表了asci码的一个字符

2 rune类型 代表一个utf-8字符

当需要除了中文，日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32

go使用了特殊的rune类型来除了unicode,让基于unicode的文本除了更为方便，也可以使用默认的byte类型进行默认字符串处理，性能和拓展性都有照顾

### 字符串修改及类型转换

要修改字符串，需要先将其转换成[]rune或[]byte完成后再转换成string.。无论哪转换，都会重新分配内存，并复制字节数组

```go
	//字符串修改
	s9:="意中人" //意 中 人
	s10:=[]rune(s9)//把字符串强制转换成一个rune切片
	s10[0]='百'
	fmt.Printf(string(s10))
//1位的用byte，3位的用rune
}
```

### 类型转换

go语言中只有强制类型转换，没有隐式类型转换，该语法只在2个类型之间支持强制类型转换的时候使用

```go
T(表达式)
```

T表示要转换的类型，表达式包括变量，复杂算子和函数返回值等

### 查找字符串中的汉字数

```go
	 t := 0
	s8 := "门前雪 意中人 何时再相逢 年少的梦不再有"
	for _, d := range s8 {
		if d > 'z' {
			t++
		}
	}

	fmt.Printf("%d", t)
```



## 流程控制

go语言中最常用的流程控制 if和for 而switch和goto主要是为了简化代码

### if else(分支结构)

go语言中if条件判断的格式如下

```go
if 表达式1{
    分支1
}else if 表达式2{
    分支2
}else {
    分支3
}
```

```go
package main

import "fmt"

//if条件判断
func main() {
	age := 16
	if age == 14 {

		fmt.Printf("初中")
	} else if age > 14 && age < 18 {
		fmt.Println("高中")
	} else if age > 20 {
		fmt.Println("大学")
	} else {
		fmt.Printf("该当帕鲁了")

	}

}

```

### if条件判断特殊写法

if条件判断还有一种特殊的写法，可以在if表达式之前添加一个执行语句 再根据变量值进行判断

```go
func ifdemo2(){
	if score :=99 ;score>=90{
		fmt.Println("A")
	}else if  score<90{
		fmt.Println("B")
	}
}
```

### for循环结构

for循环的基本格式如下

```go
for 初始语句;条件表达式 ;结束语句{
    循环体语句
}
```

```go
	//基本格式
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	//变种1
	var i=5
	for;i<10;i++{
		fmt.Println(i)
	}
	//变种2
	
	for;i<10; {
		fmt.Println(i)
		i++
	}
}
```

```go
	//无限循环
	for{
		fmt.Println("123")
	}
```

```go
//流程控制之跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")
	//当i=5时，跳过此次for循环()
	for i=0;i<10;i++{
		if i==5{
			continue//继续下一次循环
		}
	}
```

for循环可以通过break goto return panic 语句强制退出循环

### for range(键值循环)

Go 语言可以使用 for range 遍历数组、切片、字符串、map 及通道（channel）。通过 for range 遍历的返回值有一定的规律：
数组、切片、字符串返回索引和值。
map 返回键和值。
通道（channel）只返回通道内的值。
遍历数组、切片——获得索引和元素

```go
// 	//for range循环
	s := "门前雪"
	for i, v := range s { //i是索引
		fmt.Printf("%d %c\n", i, v)
	}
}

```

### 打印九九乘法表

```go

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {

			fmt.Printf("%d*%d=%-2d ", j, i, i*j)
			if j == i {
				fmt.Println()
			}

		}
		//	fmt.Println()

	}
打印倒着的九九乘法表
for i := 9; i >0; i-- {
		for j := 1; j <= i; j++ {

			fmt.Printf("%d*%d=%-2d ", j, i, i*j)
			if j == i {
				fmt.Println()
			}

		}
		//	fmt.Println()

	}
```



### switch case

使用switch语句可方便地对大量的值进行判断

go语言规定每个switch语句只能有一个default

```go
package main

import "fmt"

func main(){
	var n int=5
	switch n{
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
	    fmt.Println("无效的数字")
	}
}
```

```go
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	}
```

```go
age:=40
	switch {
	case age>40:
		fmt.Println("猜错了")
	case age==40:
		fmt.Println("猜对了")
	}
```

#### fallthrough 

可以执行满足条件的case的下一个case 是为了兼容c语言中的case设计的

### goto（跳转到指定标签）

goto语句通过标签进行代码间的无条件跳转 goto语句可以在快手跳出循环 避免重复突出有一定的帮助 

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	breakTag:
	fmt.Println("结束for循环")
}
```

### continue 

继续下一次循环

### break

结束循环

## 运算符

### 算术运算符

加减乘除求余数

++和--在go语言中不是算术运算符

```go
	var (
		a = 10
		b = 15
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
a++//单独的语句，不能放在等号的右边赋值
b--//单独的语句
}
```

### 关系运算符

比较2个值 返回布尔值

```go
//关系运算符
fmt.Println(a==b)//go语言是强类型的，只有相同的变量才能比较
fmt.Println(a!=b)//不等于
fmt.Println(a>=b)//大于等于
fmt.Println(a>b)//大于
fmt.Println(a<=b)//小于等于
fmt.Println(a<b)//小于
}
```

### 逻辑运算符

&& and

|| or

! not

```go
//逻辑运算符
	age := 20
	if age > 18 && age < 60 {
		fmt.Println("苦逼上班的")
	} else {
		fmt.Println("不用上班")
	}
	if age<18||age>60{
		fmt.Println("不用上班")
	}else {
		fmt.Println("不用上班")
	}
	isMarried:=false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)
```

### 位运算符

&  参与运算的两数对应的二级位相与（2位均为1才为1）

|   参与运算的两数对应的二进位相或 （2位有一个为1就为1）

^   参与运算的两数对应的二进位异或，当两对应的二进位相异时，   结果为1（2位不一样就为1）

<< 左移n位就是乘以2的n次方  “a<<b”就是把a的各二次方全部左移b位，高位丢弃，低位补0

2个>  右移n位 就是除以2的n次方  “a>>b”就是把a的各二次方全部右移b位

```go
package main

import "fmt"

func main() {
	//位运算：针对的是二进制
	//5的二进制 101
	//2的二进制 010

	//& 按位与
	fmt.Println(5&2)//结果000
	//!：按位或
	fmt.Println(5|2)//结果111
	//^ 按位异或
	fmt.Println(5^2)//结果111
	//<< 将二进制位左移指定位数
	fmt.Println(5<<1)//结果1010 结果就是10
	//>> 将二进制位右移指定位数
	fmt.Println(5>>1)//结果 10 就是2

	// var m=int8(1) 只能存八位
	// fmt.Println(m<<10) 超过8位 

	//位运算的作用
	//ip地址 权限
}
```

### 赋值运算符

```go
	//赋值运算符
	// 用来给变量赋值
	var x int
	x = 10
	x += 1 //x=x+1
	fmt.Println(x)
	x -= 1  //x=x-1
	x /= 2  //x=x/2
	x %= 2  //x=x%2
	x <<= 2 //x=x<<2
	x >>= 2 //x=x>>2
	x &= 2  //x=x&2
	x |= 3  //x=x|3
	x ^= 2  //x=x^2
}

```

## 数组

#### 数组定义

```go
package main

import "fmt"

//数组
//存放元素的容器
//必须指定存放的元素的类型和容量（长度）
//数组的类型和长度是数组类型的一部分
func main() {
	var v1 [3]bool
	var v2 [4]bool
	fmt.Printf("v1:%T v2:%T\n", v1, v2)

}

```

```go
var a[3]int
```

#### 数组初始化

初始化数组时可以使用初始化列表来设置数组元素的值

```go
	//初始化
	//如果不初始化，默认元素都是零值（布尔值：false 整形和浮点型都是0 字符串为""）
	var testArray [3]int                   // 数组会初始化为int类型的0值
	var numArray [3]int = [3]int{1, 22, 0} //使用指定的初始值进行初始化
	fmt.Println(testArray, numArray)
	//初始化方式1

	a1:=[2]bool{true,false}
	fmt.Println(a1)
	//初始化方式2
	a2:=[...]int{1,2,3,4,5}//根据初始值自动推断数组的长度是多少
	fmt.Println(a2)
	

	//初始化方式3
	a3:=[5]int{0:1,4:2} //根据索引进行初始化
	fmt.Println(a3)
```

#### 数组的遍历

```go
//数组的遍历
	
	cities:=[...]string{"北京","上海","深圳","广州"}
	//方法1 for索引
	for i:=0;i<len(cities);i++{
		fmt.Println(cities[i])
	}
	//方法2 for range
	for index,value:=range cities{ //index 索引 value 值
		fmt.Println(index,value)
	}
```

#### 多维数组

```go
	//多维数组
	//[[1,2][3,4][5,6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)
}
```

#### 多层数组的遍历

```go
	//多维数组的遍历
	for _, v := range a11 {
		//fmt.Println(v)
		for j, k := range v {
			fmt.Println(j, k)
		}
	}
```

#### 数组是值类型

```go
	//数组是值类型
	b1:=[3]int{1,2,3}
	b2:=b1//b2={1,2,3}
	b2[0]=100 //b2:={100,2,3}
	fmt.Println(b1,b2)
}
```

数组支持"!=""=="操作符

[n]*T是指针数组  *[n]T是数组指针 

#### 练习题

```go

	//求数组内元素的和
	s := 0
	a1 := [5]int{1, 5, 7, 9, 4}
	for _, v := range a1 {
		s += v
	}
	fmt.Println(s)

	//找出数组中和为指定值两个元素的下标
	add := 16
	a2 := [8]int{1, 2, 4, 8, 8, 12, 14, 15}
	for i, v := range a2 {
		for j, k := range a2 {
			if i != j && v+k == add { //i!=j 是为了满足题意的2个元素
				fmt.Println(i, j)
			}
		}
	}
}s
```

## 切片（slice）

切片是一个拥有相同元素类型的可变长度的序列，它是基于数组类型做的一层封装，它非常灵活，支持自动扩容

切片是一个引用类型，它的内部结构包含 地址 长度和容量。切片一般用于快速地操作一块数据集合

### 声明

```go
var name[]T//name表示变量名 T表示切片中的元素类型
```

#### 实例

```go
//初始化
	num = []int{1, 2, 3, 4}
	s = []string{"hello", "world"}
	fmt.Println(num, s) //输出切片中的元素
```

```go
	fmt.Println(num, s) //输出切片中的元素
		fmt.Println(num == nil, s == nil) //true true nil未分配内存地址
	//切片是引用类型，不支持直接比较，只能和nil比较
```

### 切片的长度和容量

切片拥有自己的长度和容量，我们可以通过内置的len()函数来求长度 使用内置的cap()函数来求切片的容量

```go
	//长度和容量
	fmt.Printf("len=%d cap=%d slice=%v\n", len(num), cap(num), num)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(s), cap(s), s)

```

### 由数组得到切片

```go
//由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //数组
	s1 := a1[0:4] //基于一个数组切割左包含，右不包含 1,2,3,4
	s2 := a1[2:7]
	fmt.Println(s1, s2)
```

### 切片再切片

切片指向了一个底层数组

切片的长度就是它元素的个数

切片的容量是底层数组从切片的第一个元素到最后一个元素的数量

```go
	// 切片再切片
	//  底层数组从切片的第一个元素到最后的元素就是切片的容量
	var s []string
	fmt.Println(s, len(s), cap(s)) //[] 0 0
	s = []string{"hello", "world"}
	fmt.Println(s, len(s), cap(s)) //[] 0 0
	b := s[0:1]
	fmt.Println(s, len(b), cap(b)) //[] 0 0
	fmt.Println(s, b)
}
```

### 使用make()函数构造切片

```go
package main

import "fmt"

func main() {
	s1 := make([]int, 6, 10)//类型 长度 容量 如果没学容量 ，默认与长度一致
	fmt.Printf("len=%d cap=%d slice=%v\n", len(s1), cap(s1), s1)
}
```

### 切片的本质

就是一个框，框住了一块连续的内存

切片属于引用类型，真正的数据都是保存在底层数组里的

### 切片不能直接比较

切片只能和nil比较

一个nil值的切片没有底层数组

一个nil值的切片的长度和容量都是0，但我们不能说一个长度和容量都是0的切片一定是nil

#### 如果要判断一个切片是否为空，要用len(s)==0来判断，不能用s==nil来判断

### 切片的赋值拷贝

```go
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //切片
	s4 := s3                                   //s3和s4都指向同一个底层数组
	fmt.Println(s3, s4)
	s3[0] = 100
	fmt.Println(s3, s4) //s3，s4值相同
}
```

### 切片的遍历

切片的遍历与数组的遍历一样

支持索引遍历和range遍历

range遍历

```go
package main

import "fmt"

func main() {
	s := []int{12, 23, 34, 45, 656}
	for _, v := range s {
		fmt.Println(v)
	}
}

```

for遍历

```go
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
```

### append()方法为切片添加元素

append()函数可以为切片动态添加元素，每一个切片会指向一个底层数组

```go
package main

import "fmt"

//append()为切片追加元素
func main() {
	s1 := []string{"北京", "上海", "贵州"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//调用append函数必须用原来的切片变量接收返回值
	s1 = append(s1, "遵义") //append追加元素，原来的底层数组放不下的时候，go底层就会把底层数组换一个，必须用变量接收append的返回值
	fmt.Println(s1)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}

```

扩容策略

如果需要的容量 `cap` 超过原切片容量的两倍 `doublecap`，会直接使用需要的容量作为新容量`newcap`。否则，当原切片长度小于1024时，新切片的容量会直接翻倍。而当原切片的容量大于等于1024时，会反复地增加25%，直到新容量超过所需要的容量

```go
s2 := []string{"成都", "深圳", "烟台", "重庆", "苏州"}
	s1 = append(s1, s2...) //append可以追加一个切片，...表示拆开追加
```

### 使用copy()函数复制切片

#### 使用copy函数一定要算好长度

```go
func main() {
	a1 := []int{1, 2, 3}
	a2 := a1 //赋值
	// var a3 []int//nil  无内存 不可拷贝
	var a3 = make([]int, 3, 3) //make 创建切片
	copy(a3, a1)               //copy 前目标 后原
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3) //a3 是 1 2 3

}

```

### 从切片中删除元素

go语言中没有删除切片元素的专用方法，我们可以使用切片的特性来删除元素

```go
a1 := []int{1, 2, 3}
	//将a1中索引为1的元素删掉
	a1 = append(a1[:1], a1[2:]...) //a1[:1] 1 2  a1[2:] 3
	fmt.Println(a1)

	a11 := [...]int{1, 3, 5}
	s1 := a11[:]
	fmt.Println(s1, len(s1), cap(s1))
	s1 = append(s1[:1], s1[2:]...) //修改了底层数组为[1 5 5]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(a11) //[1 5 5]
```

```go
//关于append删除切片中的某个元素
	a1 := [...]int{1, 2, 3, 85, 4, 56, 5, 6, 7, 78, 8, 9, 10}
	s1 := a1[:]

	//删除索引为8的元素
	s1 = append(s1[:8], s1[9:]...)
	fmt.Println(s1)
```

## 指针

go语言中不存在指针操作 只需要记住2个符号

1 &：取地址

2 *：根据地址取值 

```go
package main

import "fmt"

//pointer
func main() {
	// 取地址
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Println(p) // 输出内存地址 0xc00000a0d8
	// 根据地址取值
	m := *p
	fmt.Println(m)
}
```

### new

用于基本数据类型创建指针，返回的是对应类型的指针

```go
	//new函数申请一个内存地址

	var a = new(int) //new函数申请一个内存地址
	fmt.Println(a)
	*a =
```

### make

给map ,slice,chan申请内存 返回的是这三个类型本身

## map

Go语言中提供映射关系容器为map 其内部使用散列表(hash)实现

map是一种无序的基于key-value的数据结构，go语言中的map是引用类型，必须初始化才能使用

### map定义

```go
	var m1 map[string]int         //还没有初始化
	m1 = make(map[string]int, 10) //初始化,要估算好map容量，避免在程序运行过程中再动态扩容
	m1["门前雪"] = 1000
	m1["眼中人"] = 2000
	fmt.Println(m1)
}
```

### map的遍历

```go
//map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
```

如果想遍历key的时候 

```go
	//只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
//只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}
```

### 使用delete函数删除键值对

```go
	//删除
	delete(m1, "门前雪")
	fmt.Println(m1)
	delete(m1, "大沙比") //删除不存在的key，不会报错
```

### 按照指定顺序遍历map

```go
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//去吃map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
    }
```

### 元素为map类型的切片

```go
//map和slice组合
	//元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10) //定义一个长度为10的切片，元素类型为map[int]string
	//没有对内部的map进行初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "a"
	fmt.Println(s1
```

### 值为切片类型的map

```go
	//值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["遵义"] = []int{1, 2, 3}
```

## 函数

函数定义

```go
func 函数名（参数）（返回值）{
    函数体
}
```

函数名第一个字母不能是数字

```go
package main

import "fmt"

func add(x int, y int) (sum int) {

	return x + y
}

//没有返回值
func sum(x int, y int) {
	fmt.Println(x + y)
}

//没有参数
func sayHello() {
	fmt.Printf("say gua")
}

//无参数但有返回值
func getSum() int {
	return 3
}

//参数可以命名，也可以不命名
//命名返回值就相当于再函数中声明一个变量
func add2(x, y int) (ret int) {
	ret = x + y
	return //使用命名返回值可以return省略
}

//多个返回值 ，当参数中联系连续多个参数的类型一致时，可以将最后一个非参数的类型简写
func swap(x, y string) (string, string) {
	return y, x
}

//可变长参数
//必须放在函数参数的最后
func sumAll(nums ...int) {
	fmt.Println(nums) //切片
}

func main() {
	sum := add(100, 200)
	fmt.Print(sum)
	m, n := swap("hello", "world")
	fmt.Println(m, n)
	sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

//go语言中没有默认参数这个概念

```

## 函数定义和defer

在一个命名的函数中不能够再声明命名函数

### defer语句

defer语句会将其后面跟随的语句进行延迟处理，在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行 

```
defer语句执行  返回值x
               |
               defer语句执行
               |
               返回值return
```

```
package main

import "fmt"
//go语言的rentrun不是原子操作,分为多步
//第一步 返回值赋值
//第二步  真正的ret返回

func deferdemo() {
	fmt.Println("门前雪")
	defer fmt.Println("嘿嘿嘿")
	defer fmt.Println("love") //一个函数中可以有多个defer
	fmt.Println("眼中人")        //多个defer语句按照先进后出的顺序执行
}

func main() {
	deferdemo()
}
```

### 函数进阶

#### 变量的作用域

函数中查找变量的顺序

1 先在函数内部查找

2 找不到就往函数外面查找，一直查找到全局

##### 全局变量

全局都可以使用

##### 局部变量

不能在函数外面使用

### 函数类型和变量

#### 定义函数类型

```
type calculation func (int int )int
```

定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值

函数也可以作为参数的类型

#### 函数也可以作为参数的类型

```
func f3(x func() int){
代码段
}
```

#### 函数还可以作为返回值

```
func f5(x func()int) func(int, int) int{
代码段
ret :=func (a,b int)int
}
```

#### 匿名函数

```
var f1=func(x,y int){
fmt.Println(x+y)
}
```

一般用在函数内部

#### 立即执行函数

```
func(x,y int){
fmt.Println(x+y)
fmt.Println("hello world")
}(100,200)
```



### 闭包

闭包指的是一个函数和与其相关的引用环境（引用另一个函数外部变量）而组合而成的实体，简单来说，闭包=函数+引用环境

```
package main

//闭包
import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//定义一个函数对f2进行闭包
func f3(f func(x, y int), m, n int) func() {
	temp := func() {
		f(m, n)
	}
	return temp
}
func main() {
	f4 := func() {
		fmt.Println("this is f4")
	}
	f1(f4)
	temp := f3(f2, 100, 200)
	temp()
}
```

### defer再讲解

释放资源 解锁

#### defer执行原理

1 返回值赋值 2 defer 3 真的返回

### 内置函数

#### panic和recover

panic可以在任何地方引发，recover只有在defer调用的函数中有用

defer语句一定要在可能引发panic之前定义

## 递归

函数自己调用自己

递归的本质就是逐级拆分，按照提示拆分

```go
package main

import "fmt"

//递归
//递归一定要有明确的退出条件
//递归适合处理 问题相同 ，问题规模越来越小的场景
//函数自我调用

//例子1
//实现阶乘
func jc(n int) (ret int) {
	//对n进行判断 ，n>0进行递归运算
	if n <= 1 {
		return 1
	}
	ret = n * jc(n-1)
	return ret
}
func main() {
	var n int
	fmt.Scan(&n)
	sum := jc(n)
	fmt.Println("阶乘是", sum)
}
```

## 结构体

go语言是强类型的，对变量有个确切的类型

### 自定义类型和类型别名

```go
//type后面跟的是类型
type myInt int //自定义类型
type yourInt = int //类型别名
func main(){
    var n myInt
    n=100
    fmt.Println(n)
    fmt.Printf("%T\n",n)
    var m yourInt
    m=100
    fmt.Println(n)
    fmt.Printf("%T\n",n)
}
```

### 结构体定义

使用type和struct关键字来定义结构体

```go
type 类型名 struct{
    字段名 字段类型
    字段名 字段类型
}
```

```go
package main

import "fmt"

//type后面跟的是类型
// type myInt int     //自定义类型
// type yourInt = int //类型别名

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// var n myInt = 100

	// fmt.Println(n)
	// fmt.Printf("%T\n", n)

	// var m yourInt = 100

	// fmt.Println(m)
	// fmt.Printf("%T\n", m)

	// var c rune = '雪'

	// fmt.Println(c)
	// fmt.Printf("%T\n", c)
	var p person
	p.name = "门前雪"
	p.age = 18
	p.gender = "男"
	p.hobby = []string{"游戏", "骑车", "美食"}
	fmt.Println(p)
	//访问变量p的字段
	fmt.Println(p.age)
}

```



### 匿名结构体

```go
	var s struct {
		name string
		x    int
		y    int
	}
	s.x = 18
	s.name = "love"
	fmt.Println(s)
```

### 结构体指针及其初始化

```
b:=a
b存的是a的内存地址
```

```go
package main

import "fmt"

//结构体是值类型
type person struct {
	name, gender string
}

func f(x person) { //go语言中函数传参永远是副本，传的是拷贝
	x.gender = "女"
}

func f2(x *person) { //go语言中函数传参永远是副本，修改原变量需要指针
	// (*x).gender = "女"
	x.gender = "女" //语法糖，自动根据指针找到对应的变量
}

//创建指针类型结构体

func main() {
	var p person
	p.gender = "男"
	p.name = "松上霜"
	f(p)
	fmt.Println(p.gender)
	fmt.Println(p.name)
	f2(&p)
	fmt.Println(p.gender) //女 f2改的是原本
	//结构体指针1
	var p2 = new(person)
	fmt.Println(p2)
	p2.name = "门前雪"
	fmt.Printf("%p\n", p2)

	//结构体指针2
	var p3 = &person{
		name:   "元帅",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)
	//使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序的一样
	p4 := &person{
		"门前雪",
		"男",
	}
	fmt.Printf("%#v\n", p4)
}

```

#### 初始化

构造函数：返回一个结构体变量的函数

```go
package main

import "fmt"

//构造函数
type person struct {
	name string
	age  int
}

//构造函数 ：约定成俗 用new 开头
//返回的式结构体还是结构体指针
//当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}
func main() {
	p1 := newPerson("门前雪", 18)
	p2 := newPerson("松上霜", 18)
	fmt.Println(p1, p2)

}
```

#### 方法和接收者

go语言中的方法是一种作用于特定变量的函数，这种特定类型变量叫做接收者

定义方式如下

```
func (接收者变量 接收者类型)方法名(参数列表)(返回参数){
函数体
}
```



```go
package main

import "fmt"

//方法
type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//方法是作用于特定函数的函数
//接收者表示的是具体类型方法变量，多用类型名字名首字母小写表示
func (d dog) wang() { //(d dog) 指定函数调用
	fmt.Printf("%s:汪汪汪~", d.name)
}
func main() {
	var name []string = []string{"肖健", "周文洪", "靳俭凯", "肖登宇"}
	var n int
	fmt.Println("请输入0~3的数")
	fmt.Scan(&n)
	d1 := newDog(name[n])
	d1.wang()
}

```

##### 值接收者和指针接收者

```go
package main

import "fmt"

//方法
// 标识符 变量名 函数名 类型名 方法名
//go 语言中如果有标识符首字母大写的，就表示对外部可见（暴露的，公有的）
type dog struct {
	name string
}
type person struct {
	name string
	age  int
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}

}

//方法是作用于特定函数的函数
//接收者表示的是具体类型方法变量，多用类型名字名首字母小写表示
func (d dog) wang() { //(d dog) 指定函数调用
	fmt.Printf("%s:汪汪汪~", d.name)
}

//使用值接收者 传拷贝值
func (p person) guonian() {

	p.age++
	fmt.Println(p.age)
}

//使用指针接收者 传内存地址
func (p *person) rgn() {
	p.age++

}

func main() {
	d1 := newDog("ssh")
	d1.wang()

	p1 := newPerson("门前雪", 18)
	fmt.Println(p1.age)
	p1.guonian()
	p1.rgn()
	fmt.Println(p1.age)
}

```

### 函数版学生管理系统

```go
package main

import (
	"fmt"
	"os"
)

//学生管理系统（函数版）
/*
写一个系统，能够查看/新增/编辑/删除学生
*/
var (
	allStudents = make(map[int64]*student, 50) //学号不重复，名字能重复，学号就是key

)

type student struct {
	id   int64
	name string
}

// nerStudents 是student类型的构造函数
func nerStudents(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudents() {
	//把所有的学生都打印出来
	for k, v := range allStudents {
		fmt.Printf("学号：%d  姓名: %s", k, v.name)
	}
}

func addNewStudents() {
	//向all中添加学生
	//创建一个新学生
	//获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Printf("请输入学生的学号")
	fmt.Scan(&id)
	fmt.Printf("请输入学生的姓名")
	fmt.Scan(&name)
	//创造一个一个学生的变量
	//然后追加到all中，调用构造函数
	newStu := nerStudents(id, name)
	allStudents[id] = newStu //追加，赋值就行
}
func deleteStudents() {
	//请用户输入要删除的学生学号
	var (
		deleteID int64
	)
	fmt.Printf("请输入你要删除学生的学号:")
	fmt.Scan(&deleteID)

	//根据输入的键值对删除allStudents中的学生
	delete(allStudents, deleteID)

}
func main() {
	//allStudents = make(map[int64]*student, 50)
	// 打印菜单
	for {
		fmt.Println("欢迎使用学生管理系统")
		fmt.Println(`
	1 查看所有学生
	2 新增学生
	3 删除学生
	4 退出
	`)

		// 等待用户选择操作
		fmt.Print("请输入对应序号以执行相应操作 :")
		var choice int
		fmt.Scan(&choice)
		//反馈用户，输出用户的选择
		fmt.Printf("您选择了%d选项\n", choice)

		// 执行对应的函数
		//先判断choice
		switch choice {
		case 1:
			showAllStudents()
		case 2:
			addNewStudents()
		case 3:
			deleteStudents()
		case 4:
			os.Exit(1) //1表示退出的状态码
		default:
			fmt.Println("输入无效，请重新输入")
		}
	}
}
```

### 结构体的匿名字段与字段结构体嵌套

### 匿名字段

结构体运行其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段

```go
package main

import "fmt"

//匿名字段

type person struct {
	string
	int
}

func main() {
	p1 := person{
		"门前雪",
		18,
	}
	fmt.Println(p1.int)
	fmt.Println(p1.string)
}

```

### 嵌套结构体和匿名结构体嵌套

```go
package main

import "fmt"

//嵌套结构体  把一个结构体嵌套在另一个结构体里面
type address struct {
	address string
}
type person struct {
	name    string
	age     int
	address //匿名嵌套结构体
}
type company struct {
	name    string
	address address
}

func main() {
	p1 := person{
		name: "门前雪",
		age:  18,
		address: address{
			address: "贵州省遵义市",
		},
	}
	fmt.Println(p1.address.address)
	fmt.Println(p1.address) //现在自己结构体寻找，找不到就在匿名结构体查找，匿名结构体有重复命名的，要写全

}
```

### 结构体模拟实现继承

go语言中使用结构体也可以实现其他编程语言中面向对象的继承

```go
package main

import "fmt"

//结构体模拟实现其他语言中的继承

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Println("%s会动", a.name)
}

type dog struct {
	feet   uint8
	animal //animal拥有的方法，dog此时也有了
}

func (d dog) wang() {
	fmt.Println(" %s 汪汪汪", d.name)
}
func main() {
	d1 := dog{
		animal: animal{name: "小黑"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
```

### 结构体与json

```go
package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json

// 1 序列化 把go语言中的结构体变量--->json格式的字符串
// 2 反序列化 把 json格式的字符串 -->go语言能够识别的结构体变量
type person struct {
	Name string `json:"name" db:"name" ini:"name"` //首字母大写，才能被json解析
	Age  int    `json:"age"`
}

func main() {

	//序列化
	p1 := person{
		Name: "门前雪",
		Age:  18,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshale failed ,err :%v", err)
		return
	}
	fmt.Printf("%v\n", string(b))

	//反序列化
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了能在函数内部修改p2的值
	fmt.Printf("%#v\n", p2)
}

```

## 接口（interface）

接口是一种类型，是一种特殊的类型，它规定了变量有哪些方法

在编程中会遇到以下场景

我不关心一个变量是什么类型，我只关心能调用它的什么方法

```go
package main

import (
	"fmt"
)

//引出接口示例

// 定义一个能叫的类型
type speaker interface {
	speak() //只要实现了speak方法，都能当做speak变量

}

// 定义一个car类型
// 不管是什么结构体，只要有run方法都是car类型
type car interface {
	run()
}
type falali struct {
	brand string
}
type baoshijie struct {
	brand string
}

type cat struct{}

type dog struct{}
type person struct{}

// 接收car类型
func run(c car) {
	c.run()
}
func (b baoshijie) run() {
	fmt.Println("fafa")
}
func (f falali) run() {
	fmt.Println("bbbb")
}
func (c cat) speak() {
	fmt.Println("喵喵喵~")

}
func (d dog) speak() {
	fmt.Println("汪汪汪~")
}
func (p person) speak() {
	fmt.Println("啊啊啊~")
}
func da(x speaker) {
	//接收一个参数，传进来什么，我就打什么
	x.speak() //挨打了就要叫
}
func main() {
	var c1 cat
	var c2 dog
	var p1 person

	da(c1)
	da(c2)
	da(p1)

	f1 := falali{
		brand: "法拉利",
	}
	b1 := baoshijie{
		brand: "保时捷",
	}
	run(f1)
	run(b1)

}

```



### 接口的定义与实现

```go
type 接口名 interface {
    方法名1 （参数1，参数2）（返回值1，返回值2）
    方法名2 （参数1，参数2）（返回值1，返回值2）
}
```

用来给变量\参数\返回值等设置类型

一个变量如果实现了接口中规定的所有方法，那么这个变量就实现了这个接口，可以称为这个接口类型的变量

接口分为接口类型和接口值

接口保存的分为两部分：值类型和值

动态类型 动态值

main.cat  cat{

}

这样就实现了接口变量能够存储不同的值

### 接口类型变量

```go
package main

import "fmt"

type animal interface {
	move()
	eat(string)
}
type cat struct {
	naem string
	feet int8
}
type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡你太美")
}
func (c chicken) eat(food string) {
	fmt.Printf("吃 %s", food)
}
func (c cat) move() {
	fmt.Println("走猫步")
}
func (c cat) eat(s string) {
	fmt.Printf("吃 %s", s)
	fmt.Println()
}
func main() {
	var a1 animal //定义一个接口类型的变量
	var b1 animal
	mimi := cat{
		naem: "咪咪",
		feet: 4,
	}
	kfc := chicken{
		feet: 4,
	}
	b1 = kfc
	b1.eat("鸡饲料")
	fmt.Println(b1)
	a1 = mimi
	a1.eat("肉")
	fmt.Println(a1)

}

```



#### 值接收者 指针接收者

使用值接收者与使用值接收者的区别

使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存

指针接收者实现接口智能存结构体指针类型的变量

## 包



## 文件操作
