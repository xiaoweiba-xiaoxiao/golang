





## golang 学习笔记

### 一. go语言基础

#### 1.变量，常量， 基本数据类型

##### 1.1 变量

```go
变量的申明 不能以数字开头
var a int = 1 //可用用于申明全局变量 
a := 1 //不用用于申明全局变量

变量申明不能以关键命名
go 中的关键字
break default func interface select
case defer go map struct
chan else goto package switch
const fallthough if range type
continue for import return var

多变量申明
var (
    a int
    b bool
    c string
    d int = 8
    f float64 = 3.1415    
)
```

##### 1.2 常量

```go
常量声明
const a int = 16go
批量申明
const (
    a = 1 << iota
    b
    c
)
```

##### 1.3 数据类型:

###### 1.3.1 数据类型介绍

```go
整数类型：
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64
var a int
a = 16

bool类型
var b bool
b = true

浮点数类型
float32 float64
var f float64 = 3.1415

字符串类型
var str string = "hello"

整数型的初始化为 0, 浮点数初始化为 0.0, 字符串为 "", bool为 false
其他数据类型
(a) 指针类型（Pointer）
(b) 数组类型
(c) 结构化类型(struct)
(d) Channel 类型
(e) 函数类型
(f) 切片类型
(g) 接口类型（interface）
(h) Map 类型

```

###### 1.3.2 操作符

```go
bool 类型
var b bool = false
!b 取反
|| 或
&& 与
== 等等于
%t 格式化shc

int float数值类型
+ - * /
==
%d

string 类型
"" 可以包含控制字符
`` 所有字原样输出

常用操作
len(str) 获取长度
+,fmt.Sprintf 拼接字段
strings.Split(str,"") 切割
strings.Contains 包含
strings.HasPrefix, strings.HasSuffix 判断前缀和后缀
strings.Index(), strings.LastIndex() ⼦串出现的位置
strings.Join(a[]string, sep string) join操作

操作符
逻辑操作符
==,!=, >, <, >=, <=,
运算算
+ - * / %
```

##### 1.4 类型推导

```go
go 可以根据赋值类型进行变量推导
示例1
b := 1 由1可以推导 b 为int类型
c := false 由值推导出c是bool类型
```



##### 1.5 go 程序的基本结构

```go
1. 任何⼀个代码⽂件⾪属于⼀个包

2. import 关键字，引⽤其他包：

   import(“fmt”)

   import(“os”)

   通常习惯写成：

   import (

    “fmt”

    “os”

   )

   开发可执⾏程序，package **main**，

    并且有且只有⼀个main⼊⼝函数

   4. 包中函数调⽤：

   a. 同⼀个包中函数，直接⽤函数名调⽤

   b. 不同包中函数，通过包名+点+

   函数名进⾏调⽤

   5. 包访问控制规则：

   a. ⼤写意味着这个函数/变量是可导出的

   b. ⼩写意味着这个函数/变量是私有的，包外部不能访问
```



#### 2. 函数

函数是一块执行特定任务的代码。一个函数是在输入源基础上，通过执行一系列的算法，生成预期的输出。

##### 2.1 函数的形式

```go
1.函数的关键字 func functionname(args)(int) {}

package main

import (
    "fmt"
)

func main(){
    fmt.Println("hello world")
}
2.带参数的函数
func testFunc(a,b int){
    fmt.Printf("a + b =%d",a+b)
}
3.带不定长参数的函数
func testFunc(a ...interface{}){
    for i := 0;i < len(a);i++{
        fmt.Printf("args[%d] is %v",i,a[i])
    }
}
4.函数的返回值
func testFunc(a,b int)(int){
    return a * b
}
5.多返回值
func testFunc(a,b int)(sum int,mod int){
    sum = a+b
    mod = a%b
    return //当返回值被定义时，直接返回皆可
}
```

##### 2.2  函数的调用

```go
1. 函数调用示例
package main

import (
	"fmt"
)

func Add(a,b int)(int){
    return a + b
}

func main(){
    a,b := 1,3
    d := Add(a,b)//调用函数Add
    fmt.Printf(d)
}

```

##### 2.3 匿名函数和闭包

```go
1.没有函数名的函数被称为匿名函数

示例
package main

import (
	"fmt"
)

func main(){
    result := func (a int,b int) (int){
        retrun a + b
    }(1,2) //匿名函数以及传参形式
    fmt.Println(result)
}

2.函数也是值可以像其他的值一样传递,也可以将函数作为参数

2.1 函数作为值传递
package main

import (
    "fmt"
)

func main(){
 	result := func (a int,b int) (int){
        retrun a + b
    }//将函数赋值给result
    fmt.Println(result(1,2))   
}
2.2 函数作为参数
package main

import (
    "fmt"
)

func testFunc(fn func(int,int) int,a,b int) int{ //函数作为参数
    return fn(a,b)
}

func main(){
    result := func (a int,b int) int{
        return a + b
    }
    fmt.Println(testFunc(result,1,3))    
}
2.3 闭包
闭包就是返回一个函数
func add() func(int) int{
    sum := 0
    return funx(x int) int{ //返回一个函数
        sum += x
        return sum
    }
}

func main(){
    result := add()
    for i :=0;i < 10;i++ {
        fmt.Println(result(i))
    }
}
2.3.1 闭包实例，写一个Fibonacci
package main

import (
    "fmt"
)

func Fibonacci() func() int{
    x,y,res := 0,1,0
    return func () int{
        res = x
        x,y = y,x+y
        return x
    }
}

func main(){
    fibo := Fibonacci()
    for i := 0;i< 20;i++{
        fmt.Println(fibo())
    }
}
2.3.2 利用channal 实现(提前预习)
package main

import (
    "fmt"
)

func fibonacci(fobi chan int,count int){
    x,y,res := 0,1,0
    for i :=0;i<count;i++{
        res,x,y = x,y,x+y
        fibo <- res
    }
    close(fibo)
}

func main() {
    fobi := make(chan int,10)
    go fibonacci(fobi,10)
    for val := range fobi{
        fmt.Printf(" val")
    }
}


```

#### 3.  if , switch,for

##### 3.1 if 

if 是go语言的条件表达式

```go
3.1.1 if 基本用法
if 条件 { 条件成立执行
    ...
}else{ //否则执行
    ...
} 
3.1.2 多条件判断
if {
    ...
}else if{
    ...
}else if{
    ...
}else if{
    ...
}else{
    ...
}
3.1.3 if 可以有一个简单的表达式，然后再进行判断
func c(a int) int{
    return a * (-1)
}
func main(){
    if a := c(a); a < 0{ //可以添加简单表达式a := c(a)
        fmt.Println("a is uint")
    }
}

```

##### 3.2 for 循环

go语言没有while go 语言的循环是使用for实现

```go
3.2.1 for
package main

import (
    "fmt"
)

func main(){
    for i := 0;i < 10;i++{
        fmt.Println(i)
    }
}
3.2.2 for range 
for := range 需要两个接受值，第一值为下标，第二个值为下标所对应的值。 
示例：
package main

import (
	"fmt"
)

func main(){
    str := "hello chunyan"
    for index,val := range str{ 
        fmt.Println("str[%d] is letter %s",index,val)
    }
}

3.2.3 嵌套循环
嵌套循环实现99乘法表
for i := 1; i < 10;i ++ {
    for j := 1;j < i+1;i++ {
        fmt.Printf("%dx%d=%d\t",j,i,i*j)
    }
    fmt.Println()
}
3.2.3 死循环
死循环 格式for {} 
for {
    
}
3.2.4 continue break
continue 跳过之后的代码进行下一次循环
break 退出循环，常常和死循环搭配
示例:
利用死循环，从字符串中找到第一次出现的位置并打印。
package main

import (
    "fmt"
)

str = "hello world"
childstr = " "
index := 0
for {
    if str[index] == childstr {
        fmt.Printf("%s is the %d letter of %s",childstr,index,str)
        break
    }
}
3.2.5 特殊 for
for (退出条件) {
}


```

##### 3.3 switch

```go
switch 用于多分支判断可以优化if的代码，提供代码的可读性
swith {
case
}
示例：
判断某个变量的类型

package main

import (
"fmt"
)
func main(){
	var a int
    switch a.(type) {
    case int:
        fmt.Println("a is int")
    case string:
        fmt.Println("a is string")
    default:
        fmt.Println("can not reconized a's type")
    }    
}

```

#### 4. 数组，切片，map

##### 4.1 数组

数组是存储一组相同类型的数据，数组的长度不可变

```go
4.1.1 数组声明
var arrInt [8]int 声明一个长度为8,元素类型为int的数组
4.1.2 数组初始化
var arrInt [8]int = [8]int{1,2,3,4,5,6,7,8}
4.1.3 数组赋值
var arrInt [8]int = [8]int{}
arrInt[0]=1
fmt.Println(arrInt)
4.1.4 数组遍历
var arrInt [8]int = [8]int{1,2,3,4,5,6,7,8}
for index,val = range arrInt{
    fmt.Printf("arrInt[%d] is %d",index,val)
}
注：数组不能改变长度
```



#### 5.字符串

##### 5.1 字符串

由于和其他语言相比，字符串在 Go 语言中有着自己特殊的实现，因此在这里需要被特别提出来

Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，我们可以创建一个字符串。

```go
5.1.1 Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，我们可以创建一个字符串。
package main

import (
	"fmt"
)

func main() {
	str := "hello world!"
	fmt.Println(str)
}
输出hello world
5.1.2 字符串遍历
package main

import (
	"fmt"
)

func main() {
	str := "hello world!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
}
输出:68 65 6c 6c 6f 20 77 6f 72 6c 64 21 为hello world Unicode UTF-8 编码的结果

package main

import (
	"fmt"
)

func main() {
	str := "hello world!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()
	for _, val := range str {
		fmt.Printf("%c", val)
	}
}
输出结果：
68 65 6c 6c 6f 20 77 6f 72 6c 64 21 
hello world!
上面的程序获取字符串的每一个字符，虽然看起来是合法的，但却有一个严重的 bug。让我拆解这个代码来看看我们做错了什么。
package main

import (
	"fmt"
)

func printByte(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Printf("\n")
}

func printChar(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Printf("\n")
}

func main() {
	str := "hello world!"
	printByte(str)
	printChar(str)
	str = "Señor"
	printByte(str)
	printChar(str)
}
打印结果：
68 65 6c 6c 6f 20 77 6f 72 6c 64 21 
hello world!
53 65 c3 b1 6f 72 
SeÃ±or
在上面程序的第 28 行，我们尝试输出 Señor 的字符，但却输出了错误的 S e Ã ± o r。 为什么程序分割 Hello World 时表现完美，但分割 Señor 就出现了错误呢？这是因为 ñ 的 Unicode 代码点（Code Point）是 U+00F1。它的 UTF-8 编码占用了 c3 和 b1 两个字节。它的 UTF-8 编码占用了两个字节 c3 和 b1。而我们打印字符时，却假定每个字符的编码只会占用一个字节，这是错误的。在 UTF-8 编码中，一个代码点可能会占用超过一个字节的空间。那么我们该怎么办呢？rune 能帮我们解决这个难题。
```

##### 5.2 rune

rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示

```go
package main

import (
	"fmt"
)

func printByte(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Printf("\n")
}

func printChar(str string) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
	fmt.Printf("\n")
}

func main() {
	str := "hello world!"
	printByte(str)
	printChar(str)
	str = "Señor"
	printByte(str)
	printChar(str)
}
输出结果:
68 65 6c 6c 6f 20 77 6f 72 6c 64 21 
hello world!
53 65 c3 b1 6f 72 
Señor

```

##### 5.3 for range 循环

上述问题可以通过for rang循环解决

```go
package main

import (
	"fmt"
)

func printStr(str string) {
	for index, runes := range str {
		fmt.Printf("%c bytes is %d\n", runes, index)
	}
}

func main() {

	str := "Señor"

	printStr(str)
}
输出结果：
S bytes is 0
e bytes is 1
ñ bytes is 2
o bytes is 4
r bytes is 5
可以看到ñ 占了两个字节

用字节切片构造字符串
package main

import (  
    "fmt"
)

func main() {  
    byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice)
    fmt.Println(str)
}
打印结果：
Café

如果我们把 16 进制换成对应的 10 进制值会怎么样呢
package main

import (  
    "fmt"
)

func main() {  
    byteSlice := []byte{67, 97, 102, 195, 169} 
    str := string(byteSlice)
    fmt.Println(str)
}
打印结果：
Café

用 rune 切片构造字符串
package main

import (  
    "fmt"
)

func main() {  
    runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str := string(runeSlice)
    fmt.Println(str)
}
打印结果：
Señor

```

##### 5.4 字符串的长度

```go
utf8 package 包中的 func RuneCountInString(s string) (n int) 方法用来获取字符串的长度。这个方法传入一个字符串参数然后返回字符串中的 rune 的数量。
package main

import (
	"fmt"
	"unicode/utf8"
)

func lenth(str string){
	fmt.Printf("the lenth of %s is %d\n",str,utf8.RuneCountInString(str))
}

func main(){
	str := "Señor"
	lenth(str)
	str = "chun"
	lenth(str)
}

len(str) 也可以获取字符串的长度
```

##### 5.5  字符串不可变

字符串是不可变的,一旦被创建它将无法修改

```go
package main

import (
	"fmt"
)

func changString(str string) string{
	str[0] = 'a'
	return str
}

func main(){
	str := "hello"
	fmt.Println(changString(str))
}
试图修改字符串：
# xiao.com/golang/writefile/lesson1
./lesson1.go:8:9: cannot assign to str[0]
exit status 2
Process exiting with code: 1
```

可以把字符串转化为一个 rune 切片。然后这个切片可以进行任何想要的改变，然后再转化为一个字符串。

```go
package main

import (
	"fmt"
)

func changStr(s []rune) string{
	s[0] = 'a'
	return string(s)
}

func main(){
	str := "hello"
	fmt.Println(changStr([]rune(str)))
}
此例实现了对字符串的修改
终端输出：
aello
```

#### 6. 指针

指针是一种存储变量内存地址（Memory Address）的变量。

##### 6.1 指针的声明

```go
6.1.1
var t *T //声明一个指向T类型的指针变量

示例：
package main

import (
	"fmt"
)

func main(){
	var a *int
	var b int = 5
	a = &b
	fmt.Printf("the type of a is %T\n",a)
	fmt.Println("The adrres of a is", a)
}
终端输出结果:
Type of a is *int  
address of b is 0x1040a124

6.1.2 指针的零值：
指针的零值（Zero Value）为nil
package main

import (
	"fmt"
)

func main(){
	a := 15
	var b *int
	if b == nil {
		fmt.Println("b is",b)
		b = &a
		fmt.Println("b is",b)
	}
}
终端输出:
b is <nil>
b is 0xc0000ac010
```

##### 6.2 指针解引用

指针解引用是可以通过指针获取到指针指向变量的值

```go
6.2.1 示例
package main

import (
	"fmt"
)

func main(){
	var b int = 23
	var a *int = &b
	fmt.Println("the address of b is",a)
	fmt.Println("the value of b is",*a)
}

终端输出:
the address of b is 0xc000016090
the value of b is 23

6.2.2 利用指针解引用来修改变量的值
package main

import (
	"fmt"
)

func main(){
	var b int = 24
	var a *int = &b
	*a++
	fmt.Println("the value of b is",b)
}
终端输出：
the value of b is 25
```

##### 6.3 指针传参

指针可以作为参数传递给变量

```go
6.3.1 示例1
package main

import (
	"fmt"
)

func testPointer(t *int){
	*t *= 2 
}

func main(){
	b := 2
	testPointer(&b)
	fmt.Println(b)
}
终端输出：
API server listening at: 127.0.0.1:47030
4
6.3.2 示例2
这里使用了数组指针
package main

import (
	"fmt"
)

func testSplprt(t *[3]int){
	for i := 0;i<len(*t);i++ {
		t[i] +=1
	}
	
}

func main(){
	var arr [3]int=[3]int{90,91,92}
	testSplprt(&arr)
	fmt.Println(arr)
}
终端输出：
API server listening at: 127.0.0.1:3011
[91 92 93]
//这种方式向函数传递一个数组指针参数，并在函数内修改数组。尽管它是有效的，但却不是 Go 语言惯用的实现方式。我们最好使用切片来处理。

使用切片
package main

import (
	"fmt"
)

func testArr(t []int){
	for i :=0;i<len(t);i++{
		t[i] += 1 
	}
}

func main(){
	var arr [3]int=[3]int{90,91,92}
	testArr(arr[:])
	fmt.Println(arr)
}
终端输出：
API server listening at: 127.0.0.1:43329
[91 92 93]
```

##### 6.5  不支持指针运算

Go 并不支持其他语言（例如 C）中的指针运算。

```go
//这是一个错误案例

package main

import (
	"fmt"
)

func main(){
	b := [...]int{109,110,111}
	p = &b
	p++
}
```



##### 7.  面向对象

##### 7.1 结构体

###### 7.1.1 结构体声明

结构体是用户定义的类型，表示若干个字段（Field）的集合。有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。

例如，一个职员有 `firstName`、`lastName` 和 `age` 三个属性，而把这些属性组合在一个结构体 `employee` 中就很合理。

```go
7.1.1示例1
package main

import (
	"fmt"
)

type employee struct{
	fistname string
	lastname string
	age int
}
//可以简写 属于同一类型的字段
type Employee struct {
    firsname,lastename string
    age,sex int
}

```

###### 7.1.2  匿名结构体（Anonymous Structure）

声明结构体时也可以不用声明一个新类型，这样的结构体类型称为匿名结构体

```go 
package main

import (
	"fmt"
)

func main(){
	var employee struct {
		firstname,lastname string
		age,sex int
	}
	employee.firstname = "w"
	employee.lastname = "c"
	employee.age = 31
	employee.sex = 1
	fmt.Println(employee)
}
```

###### 7.1.3 命名的结构体 

```go
package main

import (
	"fmt"
)

type Employee struct{
	firstname,lastname string
	sex,age int
}
```

​          上面的结构体 Employee 称为 命名的结构体（Named Structure）。我们创建了名为 Employee 的新类型，而它可以用于创建 Employee 类型的结构体变量。

###### 7.1.4 结构体零值

当定义好的结构体并没有被显式地初始化时，该结构体的字段将默认赋为零值。

```go
package main

import (
	"fmt"
)

type Employee struct{
	firstname,lastname string
	sex,age int
}

func main(){
	var cai Employee
	fmt.Println(cai)
}
终端输出
API server listening at: 127.0.0.1:9978
{  0 0}
firstname lastname 为""
sex age 为 0
package main

import (
	"fmt"
)

type Employee struct{
	firstname,lastname string
	sex,age int
}

func main(){
	var cai Employee
	cai.firstname = "c"
	cai.lastname = "xx"
	fmt.Println(cai)
}
终端输出
API server listening at: 127.0.0.1:19042
{c xx 0 0}
由于sex age 没有赋值，所以就采用0
```

###### 7.1.5 结构体字段访问

结构体字段访问方式为结构体变量.字段名

```go
package main

import (
	"fmt"
)

type Employee struct{
	firstname,lastname string
	sex,age int
}

func main(){
	var cai Employee
	cai.firstname = "c"
	cai.lastname = "xx"
	fmt.Println(cai)
}
```

###### 7.1.6 结构体指针

结构体指针是指向结构体的指针，反问结构体指针指向变量的值的字段时类似于（*a)解引用 可以直接

```go
package main

import (
	"fmt"
)

type Employee struct{
	Name string
	Sex,Age int
}

func main(){
	var employee *Employee = &Employee{
		Name: "hello",
		Sex: 1,
		Age: 31,
	}
	fmt.Println(employee.Name)//指针解引用
}
终端输出：
API server listening at: 127.0.0.1:46671
hello
```

###### 7.1.7 匿名字段

当我们创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段（Anonymous Field）

```go
package main

import (
	"fmt"
)

type person struct{
	string
	int
}

func main(){
	var people person=person{
		"hello",
		11,
	}
	fmt.Println(people.string)
	fmt.Println(people.int)
}
终端输出:
API server listening at: 127.0.0.1:31998
hello
11
```

虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型。比如在上面的 `Person` 结构体里，虽说字段是匿名的，但 Go 默认这些字段名是它们各自的类型。所以 `Person` 结构体有两个名为 `string` 和 `int` 的字段。

```go
package main

import (
	"fmt"
)

type person struct{
	string
	int
}

func main(){
	var people person=person{
		"hello",
		11,
	}
	fmt.Println(people.string)
	fmt.Println(people.int)
}
终端输出：
API server listening at: 127.0.0.1:31998
hello
11
```

###### 7.1.8 嵌套结构体

结构体的字段有可能也是一个结构体。这样的结构体称为嵌套结构体

```go
package main

import  (
	"fmt"
)

type Address struct{
	City string
}

type Employee struct{
	Name string
	Age int
	Sex int
	Addr Address
}

func main(){
	var e Employee=Employee{
		Name: "Hello",
		Age: 31,
		Sex: 2,
		Addr: Address{
			City: "ChengDu",
		},
	}
	fmt.Println(e)
}
终端输出:
API server listening at: 127.0.0.1:39086
{Hello 31 2 {ChengDu}}
```

###### 7.1.9 提升字段（Promoted Fields）

如果是结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段。这是因为提升字段就像是属于外部结构体一样，可以用外部结构体直接访问

```go
package main

import  (
	"fmt"
)

type Address struct{
	City string
}

type Employee struct{
	Name string
	Age int
	Sex int
	Address
}

func main(){
	var e Employee=Employee{
		Name: "Hello",
		Age: 31,
		Sex: 2,
		Addr: Address{
			City: "ChengDu",
		},
	}
	fmt.Println(e.City) //用外层结构体直接访问City字段
}
终端输出：
ChengDu
```

但是应提升结构体字段冲突：

```go
package main

import  (
	"fmt"
)

type Address struct{
	City string
}

type School struct{
	City string
}

type Employee struct{
	Name string
	Age int
	Sex int
	School
	Address
}

func main(){
	var e Employee=Employee{
		Name: "Hello",
		Age: 31,
		Sex: 2,
		School: School{
			City: "Wuhan",
		},
		Address: Address{
			City: "ChengDu",
		},
	}
	fmt.Println(e.City) //用外层结构体直接访问City字段
}
终端输出:
# xiao.com/golang/writefile/lesson1
./lesson1.go:35:15: ambiguous selector e.City
exit status 2
Process exiting with code: 1

出现报错的原因：
School,Address里面都有一个City,此时go程序不知道需要去获取那个结构体的City,因此就发生了冲突
冲突的解决方案1
package main

import  (
	"fmt"
)

type Address struct{
	City string
}

type School struct{
	City string
}

type Employee struct{
	Name string
	Age int
	Sex int
	School
	Address
}

func main(){
	var e Employee=Employee{
		Name: "Hello",
		Age: 31,
		Sex: 2,
		School: School{
			City: "Wuhan",
		},
		Address: Address{
			City: "ChengDu",
		},
	}
	fmt.Println(e.School.City) //用内部的结构体加上字段名反问直接访问City字段
}
终端输出:
Wuhan
第二种
package main

import  (
	"fmt"
)

type Address struct{
	City string
}

type School struct{
	City string
}

type Employee struct{
	Name string
	Age int
	Sex int
	City string
	School
	Address
}

func main(){
	var e Employee=Employee{
		Name: "Hello",
		Age: 31,
		Sex: 2,
		City: "Beijing",
		School: School{
			City: "Wuhan",
		},
		Address: Address{
			City: "ChengDu",
		},
	}
	fmt.Println(e.City)
}
终端输出：
API server listening at: 127.0.0.1:39228
Beijing
也由此我们得出一个结论，当外层的结构体和提升字段中包含同样名称字段时，优先级时外层结构体优于提升字段
```

###### 7.1.10 导出结构体和字段

如果结构体名称以大写字母开头，则它是其他包可以访问的导出类型（Exported Type）。同样，如果结构体里的字段首字母大写，它也能被其他包访问到。

```go
package person

type Student struct {
	Name     string
	Age, Sex int
	score    float32
}
//在文件夹person的文件person.go 新建一个结构体



//在main.go中访问
package main

import (
	"fmt"
	"xiao.com/golang/writefile/person"
)


func main(){
	var xiaoxiao person.Student = person.Student{
		Name: "xiaoxiao",
		Age: 18,
		Sex: 2,
		score: 100.00,//注释掉这行程序就能正常运行
	}
	fmt.Println(xiaoxiao)
}
终端输出:
# xiao.com/golang/writefile/lesson1
./main.go:14:3: unknown field 'score' in struct literal of type person.Student
exit status 2
Process exiting with code: 1
因为score字段是小写开头，所以我们不能直接访问。

```

###### 7.1.11结构体相等性

结构体是值类型。如果它的每一个字段都是可比较的，则该结构体也是可比较的。如果两个结构体变量的对应字段相等，则这两个变量也是相等的。

```go
package main

import (
	"fmt"
	"xiao.com/golang/writefile/person"
)


func main(){
	var xiaoxiao person.Student = person.Student{
		Name: "xiaoxiao",
		Age: 18,
		Sex: 2,
	}
	var xiaochao person.Student
	xiaochao.Name = "xiaoxiao"
	xiaochao.Age = 18
	xiaochao.Sex = 2
	if xiaoxiao == xiaochao {
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}
}
终端输出:
API server listening at: 127.0.0.1:13069
true
```

###### 7.1.12 tag

tag 是用于序列化struct 字段的，解析成tag相映的名称

```go
package person

type Student struct {
	Name    string `ymal:"name" json:"name" ini:"name"`
	Age   int
	Sex  int
	score    float32
}

package main

import (
	"fmt"
	"xiao.com/golang/writefile/person"
	"encoding/json"
)


func main(){
	var xiaoxiao person.Student = person.Student{
		Name: "xiaoxiao",
		Age: 18,
		Sex: 2,
	}
	var xiaochao person.Student
	xiaochao.Name = "xiaoxiao"
	xiaochao.Age = 18
	xiaochao.Sex = 2
	if xiaoxiao == xiaochao {
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}
	data,_:= json.Marshal(xiaoxiao)
	fmt.Println(string(data))
}
输出：
API server listening at: 127.0.0.1:24945
true
{"name":"xiaoxiao","Age":18,"Sex":2}
//这是一个简单的应用。
```

##### 7.2 结构体方法

方法其实就是一个函数，在 `func` 这个关键字和方法名中间加入了一个特殊的接收器类型。接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。

```go
package main

import (
	"fmt"
)

type Employee struct{
	Name string
	Money float32
}

func (e Employee)Show(){
	fmt.Printf("%#v",e)
}

func main(){
	cai := Employee{
		Name: "cai",
		Money: 15000.00,
	}
	cai.Show() 
}

通过 变量.方法 的形式调用了方法
```

既然我们可以使用函数写出相同的程序，那么为什么我们需要方法？这有着几个原因，让我们一个个的看看。

- [Go 不是纯粹的面向对象编程语言](https://golang.org/doc/faq#Is_Go_an_object-oriented_language)，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
- 相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。假设我们有一个 `Square` 和 `Circle` 结构体。可以在 `Square` 和 `Circle` 上分别定义一个 `Area` 方法。见下面的程序。

```go
package main

import (
	"fmt"
	"math"
)

type Rectangle struct{
	length int
	width int
}

type Circle struct{
	radius float64
}

func (r Rectangle)Area() int{
	return r.length * r.width
}

func (c Circle)Area() float64{
	return math.Pi * c.radius * c.radius
}

func main(){
	r := Rectangle{
		length: 10,
		width: 5,
	}
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of rectangle %d\n",r.Area())
	fmt.Printf("Area of circle %f",c.Area())
}
终端输出：
API server listening at: 127.0.0.1:16689
Area of rectangle 50
Area of circle 452.389342

从上可以看出每个结构体都可以有自己的area计算方法。
```

###### 7.2.1 指针接收器与值接收器

到目前为止，我们只看到了使用值接收器的方法。还可以创建使用指针接收器的方法。值接收器和指针接收器之间的区别在于，在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的。让我们用下面的程序来帮助理解这一点。

```go
package main

import (
	"fmt"
	"math"
)

type Rectangle struct{
	length int
	width int
}

type Circle struct{
	radius float64
}

func (r *Rectangle)Area() int{
	return r.length * r.width
}

func (c *Circle)Area() float64{
	return math.Pi * c.radius * c.radius
}

func main(){
	r := Rectangle{
		length: 10,
		width: 5,
	}
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of rectangle %d\n",r.Area())
    fmt.Printf("Area of circle %f",(&c).Area())
}
终端输出:
API server listening at: 127.0.0.1:38075
Area of rectangle 50
Area of circle 452.389342

可以看到(&c).Area() 和 c.Area()等价
```

什么时候使用指针接收器，什么时候使用值接收器？

一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。

指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。

在其他的所有情况，值接收器都可以被使用。

建议使用指针接收器，除非有特殊的要求

###### 7.2.2 匿名字段的方法

属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段的结构体一样。

```go
package main

import (
	"fmt"
)

type address struct{
	city string
	state string
}

func (a *address) fullAdress(){
	fmt.Printf("Full address: %s，%s",a.city,a.state)
}

type person struct{
	firsrName string
	lastName string
	address
}

func main(){
	p := person{
		firsrName: "hello",
		lastName: "world",
		address: address{
			city: "chengdu",
			state: "gaoxin",
		},
	}
	p.fullAdress()
}

终端输出：
API server listening at: 127.0.0.1:18025
Full address: chengdu，gaoxin
```

在方法中使用值接收器 与 在函数中使用值参数

当一个函数有一个值参数，它只能接受一个值参数。

当一个方法有一个值接收器，它可以接受值接收器和指针接收器。

```go
package main

import (
	"fmt"
)

type rectangle struct{
	length int
	width int
}

func area(r rectangle){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func (r rectangle)area(){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func main(){
	r := &rectangle{
		length:  10,
		width: 5,
	}
	area(r)
	r.area()
}
终端输出:
# xiao.com/golang/writefile/lesson1
./main.go:25:6: cannot use r (type *rectangle) as type rectangle in argument to area
exit status 2
Process exiting with code: 1

原因是area函数只能接收一个值类型。
正确的写法：
package main

import (
	"fmt"
)

type rectangle struct{
	length int
	width int
}

func area(r rectangle){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func (r rectangle)area(){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func main(){
	r := &rectangle{
		length:  10,
		width: 5,
	}
	//area(r)
	area(*r)
	r.area()
}
终端输出：
API server listening at: 127.0.0.1:15773
Area Function result: 50
Area Function result: 50
反之亦然
函数使用指针参数只接受指针，而使用指针接收器的方法可以使用值接收器和指针接收器。
package main

import (
	"fmt"
)

type rectangle struct{
	length int
	width int
}

func area(r *rectangle){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func (r *rectangle)area(){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func main(){
	r := &rectangle{
		length:  10,
		width: 5,
	}
	area(r)
	// area(*r) 不能使用值类型
	r.area()
}
终端输出：
API server listening at: 127.0.0.1:23433
Area Function result: 50
Area Function result: 50

```

###### 7.2.3在非结构体上的方法

到目前为止，我们只在结构体类型上定义方法。也可以在非结构体类型上定义方法，但是有一个问题。为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中。到目前为止，我们定义的所有结构体和结构体上的方法都是在同一个 `main` 包中，因此它们是可以运行的。

```go
package main

import (
	"fmt"
)

type Myint int

func (a Myint)add(b Myint) Myint{
	return a + b
}

func main() {
	num1 := Myint(5)
	num2 := Myint(10)
	sum := num1.add(num2)
	fmt.Println("sum is",sum)
}

终端输出:
API server listening at: 127.0.0.1:38977
sum is 15


示例2
package main

func (a int) add(b int) {
}

func main() {

}
在上面程序的第 3 行，我们尝试把一个 add 方法添加到内置的类型 int。这是不允许的，因为 add 方法的定义和 int 类型的定义不在同一个包中。该程序会抛出编译错误 cannot define new methods on non-local type int
```

##### 7.3 接口

接口就是一组方法的签名

任何实现了该接口方法的结构体，就实现了该接口

###### 7.3.1 接口的声明与实现

```go
package main

import (
	"fmt"
)

type Animal interface{
	Eat()
	Talk()
}

type Dog struct{
	name string
	kind string
}

func (d Dog)Eat(){
	fmt.Printf("I am eat\n")
}

func (d Dog)Talk(){
	fmt.Printf("I am %s,I am %s",d.name,d.kind)
}

func main(){
	var a Animal
	d := Dog{
		name: "maomao",
		kind: "samoye",
	}
	a = d
	a.Eat()
	a.Talk()
}
终端输出:
API server listening at: 127.0.0.1:36648
I am eat
I am maomao,I am samoye
```

###### 7.3.2 接口的实际用途

前面的例子教我们创建并实现了接口，但还没有告诉我们接口的实际用途。在上面的程序里，如果我们使用 `name.FindVowels()`，而不是 `v.FindVowels()`，程序依然能够照常运行，但接口并没有体现出实际价值。

因此，我们现在讨论一下接口的实际应用场景。

我们编写一个简单程序，根据公司员工的个人薪资，计算公司的总支出。为了简单起见，我们假定支出的单位都是美元。

```go
package main

import (
	"fmt"
)

type SalaryCalculator interface{
	CalculateSalary() int
}

type Permanent struct {
	empId int
	basicpay int
	pf int
}

type Contract struct {
	empId int
	basicpay int
}

func (p Permanent) CalculateSalary() int{
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int{
	return c.basicpay
}

func totalExpense(s []SalaryCalculator){
	expense := 0
	for _,v := range s{
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total EXpense per Month $%d",expense)
}

func main(){
	pemp1 := Permanent{1,5000,20}
	pemp2 := Permanent{2,6000,30}
	cemp1 := Contract{3,3000}
	empolyees := []SalaryCalculator{pemp1,pemp2,cemp1}
	totalExpense(empolyees)
}
终端输出:
API server listening at: 127.0.0.1:29236
Total EXpense per Month $14050
```

###### 7.3.3 接口的内部表示

我们可以把接口看作内部的一个元组 `(type, value)`。 `type` 是接口底层的具体类型（Concrete Type），而 `value` 是具体类型的值

```go
package main

import (
	"fmt"
)

type Test interface{
	Tester()
}

type myFloat float64

func (m myFloat)Tester(){
	fmt.Println(m)
}

func describe(t Test){
	fmt.Printf("Interface type %T value %v\n",t,t)
}

func main(){
	var t Test
	f := myFloat(89.7)
	t = f
	describe(t)
	t.Tester()
}
Test 接口只有一个方法 Tester()，而 MyFloat 类型实现了该接口。在第 24 行，我们把变量 f（MyFloat 类型）赋值给了 t（Test 类型）。现在 t 的具体类型为 MyFloat，而 t 的值为 89.7。第 17 行的 describe 函数打印出了接口的具体类型和值。该程序输出：
API server listening at: 127.0.0.1:22370
Interface type main.myFloat value 89.7
89.7
```

###### 7.3.4 空接口

没有包含方法的接口称为空接口。空接口表示为 `interface{}`。由于空接口没有方法，因此所有类型都实现了空接口。

空接口的应用非常广泛

```go
var a []interface{} // 这个切片可以存储任意类型的值
示例1：特殊的切片和map
package main

import (
	"fmt"
)

func main(){
	var arr []interface{}=[]interface{}{"hello",1,3.14}
	fmt.Println(arr)
}
终端输出:
API server listening at: 127.0.0.1:17679
[hello 1 3.14]

package main

import (
	"fmt"
)

func main(){
	var arr []interface{}=[]interface{}{"hello",1,3.14}
	fmt.Println(arr)
	var maps map[interface{}]interface{} = map[interface{}]interface{}{
		1:"hello",
		"hello":1,
		3.14: "Pi",
		"Pi":3.14,
	}
	fmt.Println(maps)
}
终端输出:
API server listening at: 127.0.0.1:43578
[hello 1 3.14]
map[3.14:Pi 1:hello Pi:3.14 hello:1]

示例2:实现接收任意类型参数的函数
package main

import (
	"fmt"
)

func test(i interface{}){
	fmt.Println(i)
}

func main(){
	var arr []interface{}=[]interface{}{"hello",1,3.14}
	var maps map[interface{}]interface{} = map[interface{}]interface{}{
		1:"hello",
		"hello":1,
		3.14: "Pi",
		"Pi":3.14,
	}
	test(arr)
	test(maps)
}
终端输出:
API server listening at: 127.0.0.1:39922
[hello 1 3.14]
map[3.14:Pi 1:hello Pi:3.14 hello:1]
```

###### 7.3.5类型断言

类型断言用于提取接口的底层值（Underlying Value）。

在语法 `i.(T)` 中，接口 `i` 的具体类型是 `T`，该语法用于获得接口的底层值。

一段代码胜过千言。下面编写个关于类型断言的程序。

```go
package main

import (
	"fmt"
)

func assert(i interface{}){
	s := i.(int)
	fmt.Println(s)
}

func main(){
	var s interface{} = 56
	assert(s)
}
终端输出:
API server listening at: 127.0.0.1:3834
56

断言类型错误是，package main

import (
	"fmt"
)

func assert(i interface{}){
	s,ok := i.(int)
	fmt.Println(s，ok)
}

func main(){
	var s interface{} = "hello"
	assert(s)
}

推荐的类型断言实现:
API server listening at: 127.0.0.1:8129
0 false
```

###### 7.3.6 类型选择

类型选择用于将接口的具体类型与很多 case 语句所指定的类型进行比较。它与一般的 switch 语句类似。唯一的区别在于类型选择指定的是类型，而一般的 switch 指定的是值。

类型选择的语法类似于类型断言。类型断言的语法是 `i.(T)`，而对于类型选择，类型 `T` 由关键字 `type` 代替。下面看看程序是如何工作的。

```go
package main

import (
	"fmt"
)

func findType(i interface{}){
	switch i.(type) {
	case int:
		fmt.Printf("I am a int and my value is %d\n",i.(int))
	case string:
		fmt.Printf("I am an string and my value %s\n",i.(string))
	default:
		fmt.Printf("Unknow type\n")
	}
}

func main(){
	findType("Naveen")
	findType(77)
	findType(89.98)
}
终端输出:
API server listening at: 127.0.0.1:30192
I am an string and my value Naveen
I am a int and my value is 77
Unknow type
```

还可以将一个类型和接口相比较。如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较。

```go
package main

import (
	"fmt"
)

type person struct{

}
type Des interface{
	describe()
}

func (p person)describe(){

}

func findType(i interface{}){
	switch i.(type) {
	case int:
		fmt.Printf("I am a int and my value is %d\n",i.(int))
	case string:
		fmt.Printf("I am an string and my value %s\n",i.(string))
	case Des:
		fmt.Printf("I am an Des\n")
	default:
		fmt.Printf("Unknow type\n")
	}
}

func main(){
	findType("Naveen")
	findType(77)
	findType(89.98)
	var p person
	findType(p)
}
```

###### 7.3.7 指针接收者和值接受者

指针接收者是指针实现了这个接口，值接受者是值实现的这个接口

```go
package main

import (
	"fmt"
)

type Animal interface{
	Eat()
	Talk()
}

type Dog struct{
	name string
}

type Cat struct{
	name string
}

func (d *Dog)Eat(){
	fmt.Println("dog is eatting")
}

func (d *Dog)Talk(){
	fmt.Printf("I am dog,my name is %s\n",d.name)
}

func (c Cat)Eat(){
	fmt.Println("cat is eatting")
}

func (c Cat)Talk(){
	fmt.Printf("I am cat,my name is %s\n",c.name)
}

func main(){
	d := Dog{
		name: "Lily",
	}
	c := Cat{
		name: "Lucy",
	}

	var a1,a2 Animal
	a1 = &d //是d的指针实现了这个接口所以需要传入一个接口，如果a1=d就是错误的
	a1.Eat()
	a1.Talk()
	a2 = c
	a2.Eat()
	a2.Talk()
}

终端输出:
dog is eatting
I am dog,my name is Lily
cat is eatting
I am cat,my name is Lucy
```

###### 7.3.8 多接口

```go
package main

import (
	"fmt"
)

type Animal interface{
	Eat()
}

type Pet interface{
	peiban()
}

type Dog struct{
	Name string
}


func (d *Dog)Eat(){
	fmt.Println("Dog is eatting")
}

func (d *Dog)peiban(){
	fmt.Println("Dog peiban people")
}

func main(){
	var a Animal
	var p Pet
	var d *Dog = &Dog{
		Name: "maomao",
	}
	a = d
	p = d
	a.Eat()
	p.peiban()
}
终端输出:
API server listening at: 127.0.0.1:24037
Dog is eatting
Dog peiban people
```

###### 7.3.9 接口嵌套

尽管 Go 语言没有提供继承机制，但可以通过嵌套其他的接口，创建一个新接口。

```go
package main

import (
	"fmt"
)

type Animal interface{
	Eat()
}

type Pet interface{
	peiban()
}

type Lovely interface{
	Animal
	Pet
} 

type Dog struct{
	Name string
}


func (d *Dog)Eat(){
	fmt.Println("Dog is eatting")
}

func (d *Dog)peiban(){
	fmt.Println("Dog peiban people")
}

func main(){
	var l Lovely
	var d *Dog = &Dog{
		Name: "maomao",
	}
	l = d
	l.Eat()
	l.peiban()
}
终端输出:
API server listening at: 127.0.0.1:19039
Dog is eatting
Dog peiban people

Lovely就是一个新接口，包含了Animal和Pet的接口签名方法
Dog的指针实现了Animal和Pet接口，所以他就实现了Lovely接口
```

###### 7.3.10 接口零值

```go
package main

import (
	"fmt"
)

type Test interface{
	Test()
}

type A struct{

}

func (a A)Test(){
	fmt.Println(true)
}

func main(){
	var t Test
	if t == nil {
		a := A{}
		t = a
		a.Test()
	}
}

终端输出:
API server listening at: 127.0.0.1:21973
true
```

##### 8. IO

###### 8.1 格式化输⼊

 从终端获取⽤户的输⼊

fmt.Scanf(format string, a…interface{}): 格式化输⼊，空格作为分隔符，占位符和

 格式化输出⼀致

fmt.Scan(a …interface{}): 从终端获取⽤户输⼊，存储在**Scanln**中的参数⾥，空格和换⾏符

 作为分隔符

fmt.Scanln(a …interface{}): 从终端获取⽤户输⼊，存储在**Scanln**中的参数⾥，空格作为分隔符，

遇到换⾏符结束

```go
import (
	"fmt"
)

var str1,str2 string

func testInput(){
	fmt.Scanf("%s%s",&str1,&str2)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
}
//遇到换行符事结束用户输入

package main

import (
	"fmt"
)

var str1,str2 string

func testInput(){
	fmt.Scan(&str1,&str2)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
}
//输入完成后结束输出
package main

import (
	"fmt"
)

var str1,str2 string

func testInput(){
	fmt.Scanln(&str1,&str2)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
}
//遇到换行符就结束用户输入

import (
	"fmt"
)

var str1,str2 string

func testInput(){
	fmt.Scanf("%s\n%s",&str1,&str2)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
}
//\n可以占位\n所以可以隔行输入
```

###### 8.2 从字符串格式化输入

fmt.Sscanf(str, format string, a…interface{}): 格式化输⼊，空格作为分隔符，占位符和

 格式化输出⼀致

fmt.Sscan(str string, a …interface{}): 从终端获取⽤户输⼊，存储在**Scanln**中的参数⾥，

空格和换⾏符作为分隔符

fmt.Sscanln(str string, a …interface{}): 从终端获取⽤户输⼊，存储在**Scanln**中的参数⾥，

空格作为分隔符，遇到换⾏符结束

```go
package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello world !"

func testInput(){
	fmt.Sscanf(str,"%s%s%s",&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出：
API server listening at: 127.0.0.1:20797
hello
world
!


package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello\nworld\n!"

func testInput(){
	fmt.Sscanf(str,"%s%s%s",&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出:
hello

package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello\nworld\n!"

func testInput(){
	fmt.Sscanf(str,"%s\n%s\n%s",&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出:
API server listening at: 127.0.0.1:18615
hello
world
!

package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello\nworld\n!"

func testInput(){
	fmt.Sscan(str,&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出:
API server listening at: 127.0.0.1:34727
hello
world
!

package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello world !"

func testInput(){
	fmt.Sscan(str,&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}

终端输出:
API server listening at: 127.0.0.1:10148
hello
world
!

package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello\tworld\t!"

func testInput(){
	fmt.Sscan(str,&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}


package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello world !"

func testInput(){
	fmt.Sscanln(str,&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出:
API server listening at: 127.0.0.1:48954
hello
world
!

package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello\nworld\n!"

func testInput(){
	fmt.Sscanln(str,&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出:
API server listening at: 127.0.0.1:42289
hello

```

###### 8.3 格式化输出

fmt.Sprintf(format string, a…interface{}): 格式化并返回字符串

fmt.Sprintln(a …interface{}): 把零个或多个变量按空格进⾏格式化并换⾏，返回字符串

fmt.Sprint(a …interface{}): 把零个或多个变量按空格进⾏格式化，返回字符串

```go
package main

import (
	"fmt"
)

func testSprint(){
	str := fmt.Sprintln("hello world!")
	fmt.Println(str)
	str = fmt.Sprintf("%s hello world!","hehe")
	fmt.Println(str)
	str = fmt.Sprint("hello world!")
	fmt.Println(str)
}

func main(){
	testSprint()
}
终端输出:
API server listening at: 127.0.0.1:9921
hello world!

hehe hello world!
hello world!


```

###### 8.4 终端输⼊输出

终端其实是⼀个⽂件 

终端相关⽂件的实例

os.Stdin：标准输⼊的⽂件实例，类型为*File

os.Stdout：标准输出的⽂件实例，类型为*File

os.Stderr：标准错误输出的⽂件实例，类型为*File

```go
package main

import (
	"fmt"
	"os"
)

var str1,str2 string

func testOs(){
	fmt.Fscanf(os.Stdin,"%s%s",&str1,&str2)
	fmt.Fprintf(os.Stdout,"%s %s\n",str1,str2)
}

func main(){
	testOs()
}


```

###### 8.5 格式化输出

fmt.Fprintf(fifile, format string, a…interface{}): 格式化输出，并写⼊到⽂件中

fmt.Fprintln(fifile, a …interface{}): 把零个或多个变量写⼊到⽂件中， 并换⾏

fmt.Fprint(fifile, a …interface{}): 把零个或多个变量写⼊到⽂件

```go
import (
	"fmt"
	"os"
)

func testPrint(){
	var str1,str2 string = "hello world","hehe"
	fmt.Fprint(os.Stdout,str1," ",str2,"\n")//不能换行所以要加换行符
	fmt.Fprintf(os.Stdout,"%s %s\n",str1,str2) //格式化输出到文件
	fmt.Fprintln(os.Stdout,str1,str2)//格式化输出并换行
}

func main(){
	testPrint()
}

终端输出:
API server listening at: 127.0.0.1:40053
hello world hehe
hello world hehe
hello world hehe
```

###### 8.7 带缓冲区的读写

```go
package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	var inputReader *bufio.Reader = bufio.NewReader(os.Stdin)
	fmt.Println("please input some things:")
	str,err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Fprintf(os.Stdout,"you input is %s\n",str)
	}
}

终端输出:
please input some things:
hello world
you input is hello world
```

###### 8.8 命令行参数获取

```go
package main

import (
	"fmt"
	"os"
)

func testArgs(){
	if len(os.Args) >1 {
		fmt.Println(len(os.Args))
		for i := 1;i < len(os.Args);i++{
			fmt.Println(os.Args[i])
		}
	}
}

func main(){
	testArgs()
}
终端输出:
hello
-p
11
```

使⽤flag包获取命令⾏参数

```go
package main

import (
	"fmt"
	"flag"
)

var length int
var charset string

func ParseArgs(){
	flag.IntVar(&length,"l",16,"l 生成密码的长度")
	flag.StringVar(&charset,"s","num",`s 制定密码⽣成的字符集, 
	num:只使⽤数字[0-9], 
	char:只使⽤英⽂字母[a-zA-Z], 
	mix: 使⽤数字和字母，
	advance:使⽤数字、字母以及特殊字符`)
	flag.Parse()
}


func main(){
	ParseArgs()
	fmt.Println(length)
	fmt.Println(charset)
}


```

urfave/cli包的使⽤

```go
package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
)

func main(){
	app := cli.NewApp()
	app.Name = "Test"
	app.Action = func(c *cli.Context) error{
		fmt.Println("Hello friend!") 
 		return nil
	}
	app.Run(os.Args)
}
```

获取命令⾏参数

```go
package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
)

func main(){
	app := cli.NewApp()
	app.Name = "Test"
	app.Action = func(c *cli.Context) error{
		fmt.Printf("Hello friend! %q\n",c.Args().Get(0)) 
 		return nil
	}
	app.Run(os.Args)
}

打包：
./main 222
终端输出:
Hello friend! "222"
```

获取选项参数

```go
package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
)

func main(){
	var langueage string
	var recusive bool
	app := cli.NewApp()
	app.Name = "Test"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "language,l",
			Value: "Enlish",
			Usage: "language for the greeting",
			Destination: &langueage,
		},
		cli.BoolFlag{
			Name: "recusive,r",
			Usage: "recusive for the greeting",
			Destination: &recusive,
		},
	}
	app.Action = func(c *cli.Context) error{
		var cmd string
		fmt.Println(c.NArg())
		if c.NArg() > 0{
			cmd = c.Args()[0]
			fmt.Printf("cmd is %s\n",cmd)
		}
		fmt.Println("recusive is ", recusive) 
 		fmt.Println("language is ", langueage) 
 		return nil
	}
	app.Run(os.Args)
}
编译：
./main -l chinese -r
终端输出:
0
recusive is  true
language is  chinses
```

##### 9. 文件读写

###### 9.1 文件打开和读写

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file := "/home/golang/log/test.log"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(2)
	}
	defer f.Close()
	var b [1024]byte
	var str []byte
	for {
		n, err := f.Read(b[:])
		if err != nil {
			fmt.Println(err)
		}
		if err == io.EOF {
			break
		}
		str = append(str, b[:n]...)
	}
	fmt.Println(string(str))
}

终端输出:
file is open 32222file is open 32222file is open 32222file is open 32222file is open 32222
file is open 32222
file is open 32222

file is open 32222

2020-07-22 10:27:44.506file is open 322222020-07-22 10:29:15.998file is open 32222
2020-07-22 10:29:23.008file is open 32222
2020-07-22 10:30:10.033: file is open 32222
2020-07-22 10:49:09.488: DEBUG: file is open 32222
2020-07-22 10:58:47.505: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991)file is open 32222
2020-07-22 10:59:49.187: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991) file is open 32222
2020-07-22 11:12:02.319: DEBUG: (/home/golang/lesson/logger/util.go:xiao.com/golang/lesson/logger.GetLineInfo:8) file is open 32222
```

文件读取, file.Read和file.ReadAt。读到文件末尾返回:io.EOF

###### 9.2 bufio读取文件

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file := "/home/golang/log/test.log"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(2)
	}
	defer f.Close()
	bufReader := bufio.NewReader(f)
	for {
		context, err := bufReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("the input is %s", context)
	}

}
终端输出:
the input is file is open 32222file is open 32222file is open 32222file is open 32222file is open 32222
the input is file is open 32222
the input is file is open 32222
the input is 
the input is file is open 32222
the input is 
the input is 2020-07-22 10:27:44.506file is open 322222020-07-22 10:29:15.998file is open 32222
the input is 2020-07-22 10:29:23.008file is open 32222
the input is 2020-07-22 10:30:10.033: file is open 32222
the input is 2020-07-22 10:49:09.488: DEBUG: file is open 32222
the input is 2020-07-22 10:58:47.505: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991)file is open 32222
the input is 2020-07-22 10:59:49.187: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991) file is open 32222
the input is 2020-07-22 11:12:02.319: DEBUG: (/home/golang/lesson/logger/util.go:xiao.com/golang/lesson/logger.GetLineInfo:8) file is open 32222
the input is 2020-07-22 11:12:31.926: DEBUG: (/home/golang/lesson/logger/file_logger.go:xiao.com/golang/lesson/logger.(*Filelog).DeBug:72) file is open 32222
the input is 2020-07-22 11:13:02.336: DEBUG: (/home/golang/lesson/logger/file_test.go:xiao.com/golang/lesson/logger.TestFileLog:14) file is open 32222
the input is 2020-07-22 11:17:18.379: DEBUG: (file_test.go:xiao.com/golang/lesson/logger.TestFileLog:14) file is open 32222
the input is 2020-07-22 11:19:48.674: DEBUG: (file_test.go:logger.TestFileLog:14) file is open 32222
the input is 2020-07-22 11:22:01.664: DEBUG: (file_test.go: logger.TestFileLog: 14) file is open 32222
the input is 2020-07-22 11:23:02.104: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) file is open 32222
the input is 2020-07-22 14:24:33.765: DEBUG: (filefile_logger.go:  function:logger.(*Filelog).DeBug line:82) file is open 32222
the input is 2020-07-22 14:26:50.36: DEBUG: (filefile_logger.go:  function:logger.(*Filelog).DeBug line:82) ERROR: code 32222
the input is 2020-07-22 14:27:40.152: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
the input is 2020-07-22 15:40:46.841: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
the input is 2020-07-22 16:02:14.879: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
the input is 2020-07-23 14:41:41.522: DEBUG: (filelogs.go:  function:logger.Debug line:20) init logger success
the input is 2020-07-23 14:43:44.555: DEBUG: (filemain.go:  function:main.main line:33) init logger success
the input is 2020-07-23 18:10:53.816: DEBUG: (filemain.go:  function:main.main line:35) init logger success
```

整个文件读取

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "/home/golang/log/test.log"
	bufReader, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", string(bufReader))
}
终端输出:

```

读取压缩文件

```go
package main

import (
	"fmt"
	"compress/gzip"
	"bufio"
	"os"
)

func main(){
	file := "/home/golang/log.tar.gz"
	var r *bufio.Reader
	fi,err := os.Open(file)
	if err != nil {
		fmt.Printf("Error: %s",err)
		return
	}
	defer fi.Close()
	fz,err := gzip.NewReader(fi)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open gzip failed, err: %v\n", err)
		return
	}
	r = bufio.NewReader(fz)
	for {
		line,err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(line)
	}

}
终端输出:

```

###### 9.3 文件写入

```go
package main

import (
	"fmt"
	"os"
)

func main(){
	file := "/home/golang/log/test.txt"
	var str string = "hello xiaoxiao" 
	f,err := os.OpenFile(file,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0644)
	if err != nil {
		fmt.Fprintf(os.Stderr,"error:%v",err)
		return
	}
	defer f.Close()
	n,err := f.Write([]byte(str))
	if err != nil {
		fmt.Fprintf(os.Stderr,"error:%v",err)
	}
	fmt.Println(n)
}
终端输出：
14
```

```go
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	file := "/home/golang/log/test.txt"
	f,err := os.OpenFile(file,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("ERROR:%s",err)
		return
	}
	defer f.Close()
	var str string = "hello world\n"
	neWrite := bufio.NewWriter(f)
	neWrite.WriteString(str)
	neWrite.Flush()
}
终端输出:
cat file
hello xiaoxiao
hello world
```

写入整个文件的示例

```go
import (
	"fmt"
	"os"

	"io/ioutil"
)

func main(){
	file := "/home/golang/log/test.txt"
	var str string = "hello world\n" 
	err := ioutil.WriteFile(file,[]byte(str),0644)
	if err != nil {
		panic(err.Error())
	}
}
终端输出：
cat file
hello world
由此可见会覆盖文件的内容
```

文件拷贝

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile()(writing int64,err error){
	src := "/home/golang/log/test.txt"
	des := "/home/golang/log/test1.txt" 
	sr,ok := os.Open(src)
	if ok != nil {
		err = ok
		return
	}
	defer sr.Close()
	dest,ok := os.OpenFile(des,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0600)
	if ok != nil{		
		err = ok
		return
	}
	defer dest.Close()
	return io.Copy(dest,sr)
}

func main(){
	fmt.Println(copyFile())
}
终端输出:
12 nil
```

cat 命令实现

```go
package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func cat(){
	if len(os.Args) < 2{
		fmt.Printf("can not find file\n")
		return
	}
	for i := 1;i <len(os.Args);i++{
		file := os.Args[i]
		data,err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("cat %s ERROR: %s",file,err)
			continue
		}
		fmt.Println(file)
		fmt.Println(string(data))
	}
}

func main(){
	cat()
}
执行
cat 
终端输出:
/home/golang/log/test.txt
hello world

/home/golang/log/test.log
file is open 32222file is open 32222file is open 32222file is open 32222file is open 32222
file is open 32222
file is open 32222

file is open 32222

2020-07-22 10:27:44.506file is open 322222020-07-22 10:29:15.998file is open 32222
2020-07-22 10:29:23.008file is open 32222
2020-07-22 10:30:10.033: file is open 32222
2020-07-22 10:49:09.488: DEBUG: file is open 32222
2020-07-22 10:58:47.505: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991)file is open 32222
2020-07-22 10:59:49.187: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991) file is open 32222
2020-07-22 11:12:02.319: DEBUG: (/home/golang/lesson/logger/util.go:xiao.com/golang/lesson/logger.GetLineInfo:8) file is open 32222
2020-07-22 11:12:31.926: DEBUG: (/home/golang/lesson/logger/file_logger.go:xiao.com/golang/lesson/logger.(*Filelog).DeBug:72) file is open 32222
2020-07-22 11:13:02.336: DEBUG: (/home/golang/lesson/logger/file_test.go:xiao.com/golang/lesson/logger.TestFileLog:14) file is open 32222
2020-07-22 11:17:18.379: DEBUG: (file_test.go:xiao.com/golang/lesson/logger.TestFileLog:14) file is open 32222
2020-07-22 11:19:48.674: DEBUG: (file_test.go:logger.TestFileLog:14) file is open 32222
2020-07-22 11:22:01.664: DEBUG: (file_test.go: logger.TestFileLog: 14) file is open 32222
2020-07-22 11:23:02.104: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) file is open 32222
2020-07-22 14:24:33.765: DEBUG: (filefile_logger.go:  function:logger.(*Filelog).DeBug line:82) file is open 32222
2020-07-22 14:26:50.36: DEBUG: (filefile_logger.go:  function:logger.(*Filelog).DeBug line:82) ERROR: code 32222
2020-07-22 14:27:40.152: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
2020-07-22 15:40:46.841: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
2020-07-22 16:02:14.879: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
2020-07-23 14:41:41.522: DEBUG: (filelogs.go:  function:logger.Debug line:20) init logger success
2020-07-23 14:43:44.555: DEBUG: (filemain.go:  function:main.main line:33) init logger success
2020-07-23 18:10:53.816: DEBUG: (filemain.go:  function:main.main line:35) init logger success
```

```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)


func cat(buf *bufio.Reader){
	for {
		line,err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
    
}

func Parse(){
	flag.Parse()
	if flag.NArg() < 2{
		fmt.Printf("can not find file\n")
		return
	}
	for i := 1;i <len(os.Args);i++{
		file := os.Args[i]
		f,err := os.Open(file)
		if err != nil {
			fmt.Printf("cat %s ERROR: %s",file,err)
			return
		}
		defer f.Close()	    	
		cat(bufio.NewReader(f))		
	}
}

func main(){
	Parse()
	
}
```

###### 9.4 defer解析

defer原理分析

​			return x

​			返回值= x

​			RET指令

defer原理

​		返回值= x

​		RET指令

​		运行defer

```go
package main

import (
	"fmt"
)


func testDefer1()(x int){
	x = 5
	defer func(){
		x += 1
	}() 
	return x
}

func main(){
	fmt.Println(testDefer1())
}

终端输出:
API server listening at: 127.0.0.1:33190
6
//已经设置了返回值为x,defer的时候x的值改变

package main

import (
	"fmt"
)


func testDefer2()(int){
	x := 5
	defer func(){
		x += 1
	}() 
	return x
}

func main(){
	fmt.Println(testDefer2())
}
终端输出:
API server listening at: 127.0.0.1:13897
5

package main

import (
	"fmt"
)


func testDefer2()(y int){
	x := 5
	defer func(){
		x += 1
	}() 
	return x
}

func main(){
	fmt.Println(testDefer2())
}
终端输出:
API server listening at: 127.0.0.1:30377
5
```

##### 10.反射

反射就是程序能够在运行时检查变量和值，求出它们的类型。

```go
package main

import (
	"fmt"
	"os"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefect(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	k := t.Kind()
	switch k {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			fmt.Fprintf(os.Stdout, "%s=%v\n", field.Tag.Get("ini"), v.Field(i))
		}
	default:
		fmt.Printf("unkonwn type")
	}
}

func main() {
	var student Student = Student{
		Name: "chun",
		Sex:  2,
	}
	testRefect(student)
}
中断输出：
API server listening at: 127.0.0.1:33535
name=chun
sex=2
```

###### 10.1reflect包

reflect 包提供反射机制

`reflect.Type` 表示 `interface{}` 的具体类型，而 `reflect.Value` 表示它的具体值。`reflect.TypeOf()` 和 `reflect.ValueOf()` 两个函数可以分别返回 `reflect.Type` 和 `reflect.Value`。这两种类型是我们创建查询生成器的基础

```go
1.reflect.ValueOf()

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefectValue(a interface{}) {
	v := reflect.ValueOf(a) //获取value
	t := v.Type() //通过v.Type()来获取Type
	fmt.Printf("the type is %s\n", t)
	fmt.Printf("the value is %v\n", v)
}

func main() {
	var student Student = Student{
		Name: "chun",
		Sex:  2,
	}
	testRefectValue(student)
}
终端输出:
API server listening at: 127.0.0.1:33121
the type is main.Student
the value is {chun 2}

2.reflect.TypeOf()
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefectType(a interface{}) {
	t := reflect.TypeOf(a) //获取type
	fmt.Printf("the type is %s\n", t)
}

func main() {
	var student Student = Student{
		Name: "chun",
		Sex:  2,
	}
	testRefectType(student)
}
终端输出:
API server listening at: 127.0.0.1:34516
the type is main.Student
```

reflect.Kind

`reflect` 包中还有一个重要的类型：[`Kind`](https://golang.org/pkg/reflect/#Kind)。

在反射包中，`Kind` 和 `Type` 的类型可能看起来很相似，但在下面程序中，可以很清楚地看出它们的不同之处。

```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefectKind(a interface{}) {
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	k := v.Kind() //获取type
	fmt.Printf("the type is %s\n", t)
	fmt.Printf("the kind is %s\n", k)
}

func main() {
	var student Student = Student{
		Name: "chun",
		Sex:  2,
	}
	testRefectKind(student)
}
终端输出:
API server listening at: 127.0.0.1:24274
the type is main.Student
the kind is struct

```

我想你应该很清楚两者的区别了。`Type` 表示 `interface{}` 的实际类型（在这里是 **`main.Order`**)，而 `Kind` 表示该类型的特定类别（在这里是 **`struct`**）。

```go
const (
    Invalid Kind = iota
    Bool
    Int
    Int8
    Int16
    Int32
    Int64
    Uint
    Uint8
    Uint16
    Uint32
    Uint64
    Uintptr
    Float32
    Float64
    Complex64
    Complex128
    Array
    Chan
    Func
    Interface
    Map
    Ptr
    Slice
    String
    Struct
    UnsafePointer
)
```

通过反射设置变量值

```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefectSetValue(a interface{}) {
	t := reflect.ValueOf(a)
	t.Elem().FieldByName("Name").SetString("chunyan") //修改字段的值
}

func main() {
	var student *Student = &Student{
		Name: "chun",
		Sex:  2,
	}
	testRefectSetValue(student)
	fmt.Println(*student)
}
终端输出:
API server listening at: 127.0.0.1:35319
{chunyan 2}
```

